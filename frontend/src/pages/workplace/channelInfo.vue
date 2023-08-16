<template>
  <div class="container">
    <div class="title">
      <a-typography-title :heading="5" style="margin-top: 0">
        频道信息
      </a-typography-title>
    </div>
    <Suspense>
      <ChannelInfoCardGroup />
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
  }
</style>

<route lang="yaml">
name: ChannelInfo
meta:
  requiresAuth: true
  layout: workplace-layout
</route>
