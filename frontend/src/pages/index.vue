<script setup lang="ts">
defineOptions({
  name: 'IndexPage',
})
</script>

<template>
  <div class="container">
    <a-layout>
      <a-layout-header>
        <Navbar />
      </a-layout-header>
      <a-layout-content>
        <IndexCarousel />
        <a-typography-title :heading="3" class="title">
          频道列表
        </a-typography-title>
        <a-divider :size="12" style="border-bottom-style: dotted" />
        <Suspense>
          <ChannelListCard />
          <template #fallback>
            <a-spin spinning />
          </template>
        </Suspense>
        <a-typography-title :heading="3" class="title">
          节目列表
        </a-typography-title>
        <a-divider :size="12" style="border-bottom-style: dotted" />
        <Suspense>
          <ProgramListCard />
          <template #fallback>
            <a-spin spinning />
          </template>
        </Suspense>
      </a-layout-content>
      <Footer />
    </a-layout>
  </div>
</template>

<style scoped>
.container > .arco-layout-header,
.container > .arco-layout-footer,
.container > .arco-layout-content {
  display: flex;
  flex-direction: column;
  justify-content: center;
  font-stretch: condensed;
  text-align: center;
}

.title {
  margin-left: 20px;
}
.container {
  background-color: var(--color-fill-2);
}
</style>

<route lang="yaml">
name: index
meta:
  requiresAuth: false
</route>
