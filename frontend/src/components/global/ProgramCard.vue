<script setup lang="ts">
defineProps<{
  program: {
    name: string
    description: string
    category: string
    cover: string
  }
}>()
</script>

<template>
  <a-card :style="{ width: '100%' }" class="programCard">
    <template #cover>
      <div
        :style="{
          overflow: 'hidden',
        }"
      >
        <div
          :style="{ width: '100%', height: '0', paddingBottom: '50%', overflow: 'hidden', backgroundImage: `url(${program.cover})`, backgroundSize: 'cover', backgroundPosition: 'center' }"
        />
      </div>
    </template>
    <a-card-meta :title="program.name">
      <template #description>
        <p class="description">
          {{ program.description }}
        </p>
      </template>
    </a-card-meta>
    <a-card-meta>
      <template #description>
        {{ program.category }}
      </template>
    </a-card-meta>
  </a-card>
</template>

<style scoped>
.programCard {
  transition-property: all;
}

.programCard:hover {
  transform: scale(1.01) translateY(-4px);
}

.description {
  white-space: nowrap;
  text-overflow: ellipsis;
  overflow: hidden;
}
</style>
