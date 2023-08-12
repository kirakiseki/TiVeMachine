<template>
  <div
    :style="{
      boxSizing: 'border-box',
      width: '100%',
      padding: '40px',
      backgroundColor: 'var(--color-fill-2)',
    }"
  >
    <a-row :gutter="30" :style="{ marginBottom: '30px' }">
      <a-col :span="6">
        <ProgramCard />
      </a-col>
      <a-col :span="6">
        <ProgramCard />
      </a-col>
      <a-col :span="6">
        <ProgramCard />
      </a-col>
      <a-col :span="6">
        <ProgramCard />
      </a-col>
    </a-row>
    <a-row :gutter="30" :style="{ marginBottom: '20px' }">
      <a-col :span="6">
        <ProgramCard />
      </a-col>
      <a-col :span="6">
        <ProgramCard />
      </a-col>
      <a-col :span="6">
        <ProgramCard />
      </a-col>
      <a-col :span="6">
        <ProgramCard />
      </a-col>
    </a-row>
    <a-row :gutter="30" :style="{ marginBottom: '20px' }">
      <a-col :span="6">
        <ProgramCard />
      </a-col>
      <a-col :span="6">
        <ProgramCard />
      </a-col>
      <a-col :span="6">
        <ProgramCard />
      </a-col>
      <a-col :span="6">
        <ProgramCard />
      </a-col>
    </a-row>
  </div>
</template>
