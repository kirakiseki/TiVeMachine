<script lang="ts">
import useProgram from '~/hooks/program'

const program = useProgram()
const allList = await program.fetchScheduleList()
const scheduleList = reactive({ list: allList })
const scheduleListSorted = computed(() => scheduleList.list.sort((a, b) => a.start_time - b.start_time))
export default {
  props: {
    startTimeLeft: {
      type: Date,
      default: new Date(1691596800000),
    },
    startTimeRight: {
      type: Date,
      default: new Date(1703952000000),
    },
    endTimeLeft: {
      type: Date,
      default: new Date(1691596800000),
    },
    endTimeRight: {
      type: Date,
      default: new Date(1703952000000),
    },
  },
  data() {
    return {
      scheduleListSorted,
    }
  },
  watch: {
    startTimeLeft: {
      handler: 'update',
      immediate: true,
    },
    startTimeRight: {
      handler: 'update',
      immediate: true,
    },
    endTimeLeft: {
      handler: 'update',
      immediate: true,
    },
    endTimeRight: {
      handler: 'update',
      immediate: true,
    },
  },
  methods: {
    update() {
      scheduleList.list = allList.filter((s) => {
        return s.start_time * 1000 >= this.startTimeLeft.getTime() && s.start_time * 1000 <= this.startTimeRight.getTime() && s.end_time * 1000 >= this.endTimeLeft.getTime() && s.end_time * 1000 <= this.endTimeRight.getTime()
      })
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
