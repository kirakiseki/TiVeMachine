<script lang="ts">
import useProgram from '~/hooks/program'

const program = useProgram()
const allList = await program.fetchScheduleList()
const scheduleList = reactive({ list: allList })
const scheduleListSorted = computed(() => scheduleList.list.sort((a, b) => a.start_time - b.start_time))
export default {
  props: {
    channel: {
      type: Number,
      default: -1,
    },
  },
  data() {
    return {
      scheduleListSorted,
    }
  },
  watch: {
    channel: {
      handler() {
        scheduleList.list = allList.filter((s) => {
          return s.channel.ID === this.channel
        })
      },
      immediate: true,
    },
  },
}
</script>

<template>
  <div class="timeline">
    <a-timeline>
      <a-timeline-item v-for="s in scheduleListSorted" :key="s.id" dot-color="#00B42A">
        <ScheduleInfoCardNoFunc
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
