<script setup lang="ts">
import type { Channel } from '~/api/modules/program'
import useChannel from '~/hooks/program'

const channel = useChannel()
const channelList = await channel.fetchChannelList()

function channelRows(channelList: Array<Channel>) {
  const rows = []
  for (let i = 0; i < channelList.length; i += 2)
    rows.push(channelList.slice(i, i + 2))

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
    <a-row v-for="(row, i) in channelRows(channelList)" :key="i" :gutter="30" :style="{ marginBottom: '20px' }">
      <a-col v-for="p in row" :key="p.Name" :span="12">
        <ChannelCard :channel="{ name: p.Name, description: p.Description, cover: p.Cover }" />
      </a-col>
    </a-row>
  </div>
</template>
