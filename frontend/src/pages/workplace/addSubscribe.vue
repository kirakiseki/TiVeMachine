<template>
  <div class="container">
    <div class="title">
      <a-typography-title :heading="5" style="margin-top: 0">
        添加订阅
      </a-typography-title>
    </div>
    <Suspense>
      <SubscribeTimeline />
    </Suspense>
  </div>
</template>

<style scoped>
  .container {
    background-color: var(--color-bg-2);
    margin: 16px 20px;
    margin-bottom: 0;
    display: flex;
    flex-direction: column;
  }

  .title {
    width: 100%;
    padding: 20px 20px 0 20px;
    background-color: var(--color-bg-2);
    border-radius: 4px 4px 0 0;
  }

  .timeline {
    width: 100%;
    padding: 20px 20px 0 20px;
    background-color: var(--color-bg-2);
  }
</style>

<route lang="yaml">
name: AddSubscribe
meta:
  requiresAuth: true
  layout: workplace-layout
</route>
