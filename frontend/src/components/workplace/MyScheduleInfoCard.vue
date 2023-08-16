<script setup lang="ts">
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

const subscribe = useSubscribe()
await subscribe.fetchSubscribeList()
const subList = reactive(subscribe.subscribeList)
const nowSeconds = computed(() => Math.floor(new Date().getTime() / 1000))

async function handleUnsubscribeBtn(scheduleID: number) {
  subscribe.unsubscribe(scheduleID)
  const ind = subList.value.indexOf(scheduleID)
  if (ind > -1)
    subList.value.splice(ind, 1)
}
</script>

<template>
  <div class="scheduleCard">
    <a-card :style="{ width: '30%' }" :title="schedule.program">
      <template #extra>
        <a-button v-if="schedule.start_time < nowSeconds" type="outline" disabled>
          已过期
        </a-button>
        <a-button v-if="subList.includes(schedule.id) && schedule.start_time >= nowSeconds" type="outline" class="pickerbtn" @click="async () => { await handleUnsubscribeBtn(schedule.id) }">
          取消订阅
        </a-button>
        <a-button v-else type="outline" disabled class="pickerbtn">
          已取消
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
