<script setup lang="ts">
defineOptions({
  name: 'IndexPage',
})
</script>

<template>
  <div class="container">
    <a-layout>
      <a-layout-header>Header</a-layout-header>
      <a-layout-content>Content</a-layout-content>
      <RouterLink to="/login">
        Login
      </RouterLink>
      <Footer />
    </a-layout>
  </div>
</template>

<style scoped>
.container :deep(.arco-layout-header),
.container :deep(.arco-layout-footer),
.container :deep(.arco-layout-content) {
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
