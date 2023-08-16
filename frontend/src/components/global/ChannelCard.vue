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
  <a-card :style="{ width: '100%' }" class="channelCard">
    <template #cover>
      <div
        :style="{
          overflow: 'hidden',
        }"
      >
        <img
          :style="{ width: '100%' }"
          :src="channel.cover"
        >
      </div>
    </template>
    <a-card-meta :title="channel.name">
      <template #description>
        <p class="description">
          {{ channel.description }}
        </p>
      </template>
    </a-card-meta>
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
