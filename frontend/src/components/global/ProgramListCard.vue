<script setup lang="ts">
import type { Program } from '~/api/modules/program'
import useProgram from '~/hooks/program'

const program = useProgram()
const programList = await program.fetchProgramList()

function programRows(programList: Array<Program>) {
  const rows = []
  for (let i = 0; i < programList.length; i += 4)
    rows.push(programList.slice(i, i + 4))

  return rows
}
</script>

<template>
  <div
    :style="{
      boxSizing: 'border-box',
      width: '100%',
      padding: '40px',
      backgroundColor: 'var(--color-fill-2)',
    }"
  >
    <a-row v-for="(row, i) in programRows(programList)" :key="i" :gutter="30" :style="{ marginBottom: '20px' }">
      <a-col v-for="p in row" :key="p.id" :span="6">
        <ProgramCard :program="{ name: p.Name, description: p.Description, category: p.Category, cover: p.Cover }" />
      </a-col>
    </a-row>
  </div>
</template>
