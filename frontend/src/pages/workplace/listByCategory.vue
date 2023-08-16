<script setup lang="ts">
const keyword = ref<string>('')
</script>

<template>
  <div class="container">
    <div class="title">
      <a-typography-title :heading="5" style="margin-top: 0">
        分类查询
      </a-typography-title>
      <a-input-search v-model="keyword" :style="{ width: '30%' }" placeholder="请输入节目分类" search-button />
    </div>
    <Suspense>
      <SubscribeTimelineByCategory :keyword="keyword" />
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
name: ListByCategory
meta:
  requiresAuth: true
  layout: workplace-layout
</route>
