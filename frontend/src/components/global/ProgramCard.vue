<template>
  <a-card :style="{ width: '100%' }" class="programCard">
    <template #cover>
      <div
        :style="{
          overflow: 'hidden',
        }"
      >
        <img
          :style="{ width: '100%' }"
          alt="dessert"
          src="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/a20012a2d4d5b9db43dfc6a01fe508c0.png~tplv-uwbnlip3yd-webp.webp"
        >
      </div>
    </template>
    <a-card-meta title="走近科学">
      <template #description>
        周六 10:20-10:45
      </template>
    </a-card-meta>
    <a-button type="primary" size="small" style="margin-top: 10px;">
      预约
    </a-button>
  </a-card>
</template>

<style scoped>
.programCard {
  transition-property: all;
}

.programCard:hover {
  transform: scale(1.01) translateY(-4px);
}
</style>
