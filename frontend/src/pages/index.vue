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
        <IndexContentSwitch />
        <ProgramTimeline />
        <ProgramListCard />
      </a-layout-content>
      <RouterLink to="/login">
        Login
      </RouterLink>
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
</style>

<route lang="yaml">
name: index
meta:
  requiresAuth: false
</route>
