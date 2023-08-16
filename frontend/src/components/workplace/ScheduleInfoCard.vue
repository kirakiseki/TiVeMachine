<script setup lang="ts">
import { Message } from '@arco-design/web-vue'
import useSubscribe from '~/hooks/subscribe'

defineProps<{
  schedule: {
    id: number
    program: string
    category: string
    channel: string
    start_time: number
    end_time: number
  }
}>()

function timeParse(time: number) {
  return new Date(time * 1000).toLocaleString()
}

function parseTimeStr(time: string) {
  const [h, m, s] = time.split(':')
  return new Date(new Date().getFullYear(), new Date().getMonth(), new Date().getDate(), Number(h), Number(m), Number(s))
}

function dateToStr(time: Date) {
  return `${time.getHours()}:${time.getMinutes()}:${time.getSeconds()}`
}

const subscribe = useSubscribe()
await subscribe.fetchSubscribeList()
const subList = reactive(subscribe.subscribeList)
const alarmTime = ref<string>('')
const nowSeconds = computed(() => Math.floor(new Date().getTime() / 1000))

async function handleSubscribeBtn(scheduleID: number) {
  if (alarmTime.value === '') {
    Message.error('请选择提醒时间')
    return
  }

  if (parseTimeStr(alarmTime.value).getTime() < new Date().getTime()) {
    Message.error('提醒时间不能早于当前时间')
    return
  }

  console.log(parseTimeStr(alarmTime.value))
  Notification.requestPermission()
  subscribe.addSubscribe(scheduleID, parseTimeStr(alarmTime.value).getTime())
  subList.value.push(scheduleID)
}
</script>

<template>
  <div class="scheduleCard">
    <a-card :style="{ width: '70%' }" :title="schedule.program">
      <template #extra>
        <a-button v-if="schedule.start_time < nowSeconds" type="outline" disabled>
          已过期
        </a-button>
        <div v-if="!subList.includes(schedule.id) && schedule.start_time >= nowSeconds">
          <a-button type="outline" @click="async () => { await handleSubscribeBtn(schedule.id) }">
            订阅
          </a-button>
          <a-time-picker
            v-model:model-value="alarmTime"
            class="picker"
            placeholder="选择提醒时间"
            style="margin-left: 20px;"
          >
            <template #extra>
              <a-button class="pickerbtn" @click="() => { alarmTime = dateToStr(new Date()) }">
                现在
              </a-button>
              <a-button class="pickerbtn" @click="() => { alarmTime = dateToStr(new Date(parseTimeStr(alarmTime).getTime() - 1800 * 1000)) }">
                -0.5h
              </a-button>
              <a-button class="pickerbtn" @click="() => { alarmTime = dateToStr(new Date(schedule.start_time * 1000)) }">
                播出时
              </a-button>
            </template>
          </a-time-picker>
        </div>
        <a-button v-if="subList.includes(schedule.id) && schedule.start_time >= nowSeconds" type="outline" disabled class="pickerbtn">
          已订阅
        </a-button>
      </template>
      <a-descriptions
        :data="[{
          label: '分类',
          value: schedule.category,
        }, {
          label: '频道',
          value: schedule.channel,
        }, {
          label: '开始时间',
          value: timeParse(schedule.start_time),
        }, {
          label: '结束时间',
          value: timeParse(schedule.end_time),
        }]" bordered :column="1"
      />
    </a-card>
  </div>
</template>

<style scoped>
.scheduleCard {
  margin-bottom: 30px;
}

.pickerbtn {
  margin-left: 10px;
}
</style>
