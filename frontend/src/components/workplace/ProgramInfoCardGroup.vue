<script setup lang="ts">
import type { Program } from '~/api/modules/program'
import useProgram from '~/hooks/program'

const program = useProgram()
const programList = await program.fetchProgramList()

function programRows(programList: Array<Program>) {
  const rows = []
  for (let i = 0; i < programList.length; i += 1)
    rows.push(programList.slice(i, i + 1))

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
      <a-col v-for="p in row" :key="p.Name" :span="24">
        <ProgramInfoCard :program="{ id: p.id as number, name: p.Name, category: p.Category, description: p.Description }" />
      </a-col>
    </a-row>
  </div>
</template>
