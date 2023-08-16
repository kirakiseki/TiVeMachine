<script setup lang="ts">
defineProps<{
  channel: {
    name: string
    description: string
    cover: string
  }
}>()
</script>

<template>
  <a-card :style="{ width: '100%' }" :title="channel.name">
    <a-descriptions
      :data="[{
        label: '频道名',
        value: channel.name,
      }, {
        label: '频道介绍',
        value: channel.description,
      }]" bordered :column="1"
    />
  </a-card>
</template>

<style scoped>
.channelCard {
  transition-property: all;
}

.channelCard:hover {
  transform: scale(1.01) translateY(-4px);
}

.description {
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}
</style>
