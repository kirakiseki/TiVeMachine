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
</script>

<template>
  <div class="scheduleCard">
    <a-card :style="{ width: '30%' }" :title="schedule.program">
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
