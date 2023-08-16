<script setup>
const images = [
  'https://3vimg.hitv.com/100/2308/1116/4719/AKwnSPMooEJ/lwUr4B.jpg?x-oss-process=image/format,webp',
  'https://0vimg.hitv.com/100/2308/1215/4348/8CFWiYD/eqUJwo.jpg?x-oss-process=image/format,webp',
  'https://1vimg.hitv.com/100/2308/1219/4518/hOMokbmm4FTC/mkqWti.jpg?x-oss-process=image/format,webp',
]
</script>

<template>
  <a-carousel
    :style="{
      width: '100%',
      height: '380px',
    }"
    :auto-play="true"
  >
    <a-carousel-item v-for="image in images" :key="image">
      <img
        :src="image"
        :style="{
          width: '100%',
        }"
      >
    </a-carousel-item>
  </a-carousel>
</template>
