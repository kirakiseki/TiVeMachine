<script setup lang="ts">
import useProgram from '~/hooks/program'
import useSubscribe from '~/hooks/subscribe'

const program = useProgram()
const subscribe = useSubscribe()
const scheduleList = await program.fetchScheduleList()
const scheduleListSorted = computed(() => scheduleList.sort((a, b) => a.start_time - b.start_time))
await subscribe.fetchSubscribeList()
const subList = subscribe.subscribeList.value
const displayList = computed(() => {
  return scheduleListSorted.value.filter((s) => {
    return subList.includes(s.id)
  })
})
</script>

<template>
  <div class="timeline">
    <a-timeline>
      <a-timeline-item v-for="s in displayList" :key="s.id" dot-color="#00B42A">
        <MyScheduleInfoCard
          :schedule="{
            id: s.id,
            program: s.program.Name,
            channel: s.channel.Name,
            category: s.program.Category,
            start_time: s.start_time,
            end_time: s.end_time,
          }"
        />
      </a-timeline-item>
    </a-timeline>
  </div>
</template>

<style scoped>
  .timeline {
    width: 100%;
    padding: 20px 20px 0 20px;
    background-color: var(--color-bg-2);
  }
</style>
