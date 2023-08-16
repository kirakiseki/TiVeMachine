<script lang="ts" setup>
import useProgram from '~/hooks/program'

const program = useProgram()
const programList = await program.fetchProgramList()
const programCount = computed(() => {
  return programList.length
})
const scheduleCount = await program.scheduleCount()
const channelList = await program.fetchChannelList()
const channelCount = computed(() => {
  return channelList.length
})
</script>

<template>
  <a-grid :cols="24" :row-gap="16" class="panel">
    <a-grid-item
      class="panel-col"
      :span="{ xs: 12, sm: 12, md: 12, lg: 12, xl: 12, xxl: 8 }"
    >
      <a-space>
        <a-avatar :size="54" class="col-avatar">
          <img
            alt="avatar"
            src="http://s3.tivemachine.com/avatar/channel.png"
          >
        </a-avatar>
        <a-statistic
          title="频道"
          :value="channelCount"
          :value-from="0"
          animation
          show-group-separator
        >
          <template #suffix>
            <span class="unit">个</span>
          </template>
        </a-statistic>
      </a-space>
    </a-grid-item>
    <a-grid-item
      class="panel-col"
      :span="{ xs: 12, sm: 12, md: 12, lg: 12, xl: 12, xxl: 8 }"
    >
      <a-space>
        <a-avatar :size="54" class="col-avatar">
          <img
            alt="avatar"
            src="http://s3.tivemachine.com/avatar/program.png"
          >
        </a-avatar>
        <a-statistic
          title="节目"
          :value="programCount"
          :value-from="0"
          animation
          show-group-separator
        >
          <template #suffix>
            <span class="unit">个</span>
          </template>
        </a-statistic>
      </a-space>
    </a-grid-item>
    <a-grid-item
      class="panel-col"
      :span="{ xs: 12, sm: 12, md: 12, lg: 12, xl: 12, xxl: 8 }"
    >
      <a-space>
        <a-avatar :size="54" class="col-avatar">
          <img
            alt="avatar"
            src="http://s3.tivemachine.com/avatar/schedule.png"
          >
        </a-avatar>
        <a-statistic
          title="播出记录"
          :value="scheduleCount"
          :value-from="0"
          animation
          show-group-separator
        >
          <template #suffix>
            <span class="unit">个</span>
          </template>
        </a-statistic>
      </a-space>
    </a-grid-item>
    <a-grid-item :span="24">
      <a-divider class="panel-border" />
    </a-grid-item>
  </a-grid>
</template>

<style lang="less" scoped>
  .panel {
    background-color: var(--color-bg-2);
    border-radius: 4px;
    overflow: auto;
  }
  :deep(.panel-border) {
    margin-bottom: 0;
    border-bottom: 1px solid rgb(var(--gray-2));
  }
  .arco-grid.panel {
    margin-bottom: 0;
    padding: 16px 20px 0 20px;
  }
  .panel-col {
    padding-left: 43px;
    border-right: 1px solid rgb(var(--gray-2));
  }
  .col-avatar {
    margin-right: 12px;
    background-color: var(--color-fill-2);
  }
  .up-icon {
    color: rgb(var(--red-6));
  }
  .unit {
    margin-left: 8px;
    color: rgb(var(--gray-8));
    font-size: 12px;
  }
  :deep(.panel-border) {
    margin: 4px 0 0 0;
  }
</style>
