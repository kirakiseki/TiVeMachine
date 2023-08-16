<script setup lang="ts">
import type { CalendarValue } from '@arco-design/web-vue/es/date-picker/interface'

const startTimeLeft = ref<Date>()
const startTimeRight = ref<Date>()
const endTimeLeft = ref<Date>()
const endTimeRight = ref<Date>()

function onChangeStart(dateString: (CalendarValue | undefined)[] | undefined, date: (Date | undefined)[] | undefined) {
  if (date) {
    startTimeLeft.value = date[0] as Date
    startTimeRight.value = date[1] as Date
  }
}

function onChangeEnd(dateString: (CalendarValue | undefined)[] | undefined, date: (Date | undefined)[] | undefined) {
  if (date) {
    endTimeLeft.value = date[0] as Date
    endTimeRight.value = date[1] as Date
  }
}
</script>

<template>
  <div class="container">
    <div class="title">
      <a-typography-title :heading="5" style="margin-top: 0">
        时段查询
      </a-typography-title>
      <div class="picker">
        开始时间范围
        <a-range-picker
          show-time
          :default-value="['2023-08-10 00:00:00', '2023-12-31 00:00:00']"
          @change="onChangeStart"
        />
      </div>
      <div class="picker">
        结束时间范围
        <a-range-picker
          show-time
          :default-value="['2023-08-10 00:00:00', '2023-12-31 00:00:00']"
          @change="onChangeEnd"
        />
      </div>
    </div>
    <Suspense>
      <SubscribeTimelineByPeriod :start-time-left="startTimeLeft" :end-time-left="endTimeLeft" :start-time-right="startTimeRight" :end-time-right="endTimeRight" />
      <template #fallback>
        <a-spin spinning />
      </template>
    </Suspense>
  </div>
</template>

<style scoped>
  .container {
    background-color: var(--color-fill-2);
    padding: 16px 20px;
    padding-bottom: 0;
    display: flex;
    flex-direction: column;
  }

  .title {
    width: 100%;
    padding: 20px 20px 0 20px;
    background-color: var(--color-bg-2);
    border-radius: 4px 4px 0 0;
    display: flex;
    flex-direction: column;
  }

  .picker {
    margin-bottom: 20px;
  }
</style>

<route lang="yaml">
name: ListByPeriod
meta:
  requiresAuth: true
  layout: workplace-layout
</route>
