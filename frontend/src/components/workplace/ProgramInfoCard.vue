<script setup lang="ts">
import useProgram from '~/hooks/program'

defineProps<{
  program: {
    id: number
    name: string
    description: string
    category: string
  }
}>()

const p = useProgram()
const deleted = ref<boolean>(false)

function handleDeleteBtn(id: number) {
  p.deleteProgram(id)
  deleted.value = true
}
</script>

<template>
  <a-card :style="{ width: '100%' }" :title="program.name">
    <template #extra>
      <a-button v-if="!deleted" type="outline" class="pickerbtn" @click="() => { handleDeleteBtn(program.id) }">
        删除节目
      </a-button>
      <a-button v-else type="outline" disabled class="pickerbtn">
        已删除
      </a-button>
    </template>
    <a-descriptions
      :data="[{
        label: '节目名',
        value: program.name,
      }, {
        label: '节目介绍',
        value: program.description,
      }, {
        label: '分类',
        value: program.category,
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
