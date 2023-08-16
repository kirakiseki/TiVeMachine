<script setup lang="ts">
import type { Channel } from '~/api/modules/program'
import useProgram from '~/hooks/program'

const value = ref<number>()
const program = useProgram()
const chanList = ref<Channel[]>([])
onBeforeMount(async () => {
  chanList.value = await program.fetchChannelList()
})
</script>

<template>
  <div class="container">
    <div class="title">
      <a-typography-title :heading="5" style="margin-top: 0">
        频道查询
      </a-typography-title>
      <a-select v-model="value" :style="{ width: '30%' }" placeholder="请选择频道">
        <a-option v-for="item of chanList" :key="item.ID" :value="item.ID" :label="item.Name" />
      </a-select>
    </div>
    <Suspense>
      <SubscribeTimelineByChannel :channel="value" />
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
    padding: 20px 20px 20px 20px;
    background-color: var(--color-bg-2);
    border-radius: 4px 4px 0 0;
  }
</style>

<route lang="yaml">
name: ListByChannel
meta:
  requiresAuth: true
  layout: workplace-layout
</route>
