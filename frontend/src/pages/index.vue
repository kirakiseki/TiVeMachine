<script setup lang="ts">
defineOptions({
  name: 'IndexPage',
})
</script>

<template>
  <div>
    Index Page
  </div>
  <Footer />
</template>

<route lang="yaml">
name: index
meta:
  requiresAuth: false
</route>
