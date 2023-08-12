<script>
import { IconExclamationCircleFill } from '@arco-design/web-vue/es/icon'

export default {
  components: { IconExclamationCircleFill },
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
    <a-timeline>
      <a-timeline-item label="2017-03-10" dot-color="#00B42A">
        The first milestone
      </a-timeline-item>
      <a-timeline-item label="2018-05-22">
        The second milestone
      </a-timeline-item>
      <a-timeline-item label="2020-06-22" dot-color="#F53F3F">
        The third milestone
        <IconExclamationCircleFill
          :style="{ color: 'F53F3F', fontSize: '12px', marginLeft: '4px' }"
        />
      </a-timeline-item>
      <a-timeline-item label="2020-09-30" dot-color="#C9CDD4">
        The fourth milestone
      </a-timeline-item>
    </a-timeline>
  </div>
</template>
