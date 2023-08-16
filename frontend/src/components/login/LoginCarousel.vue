<script lang="ts" setup>
const carouselItem = computed(() => [
  {
    slogan: '频道覆盖更广',
    subSlogan: '',
    image: 'http://s3.tivemachine.com/avatar/channel.png',
  },
  {
    slogan: '节目种类更多',
    subSlogan: '',
    image: 'http://s3.tivemachine.com/avatar/program.png',
  },
  {
    slogan: '播出时刻更准',
    subSlogan: '',
    image: 'http://s3.tivemachine.com/avatar/schedule.png',
  },
])
</script>

<template>
  <div class="banner">
    <div class="banner-inner">
      <a-carousel class="carousel" animation-name="fade" :auto-play="true">
        <a-carousel-item v-for="item in carouselItem" :key="item.slogan">
          <div :key="item.slogan" class="carousel-item">
            <div class="carousel-title">
              {{ item.slogan }}
            </div>
            <div class="carousel-sub-title">
              {{ item.subSlogan }}
            </div>
            <img class="carousel-image" :src="item.image">
          </div>
        </a-carousel-item>
      </a-carousel>
    </div>
  </div>
</template>

<style lang="less" scoped>
.banner {
    display: flex;
    align-items: center;
    justify-content: center;

    &-inner {
        flex: 1;
        height: 100%;
    }
}

.carousel {
    height: 100%;

    &-item {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100%;
    }

    &-title {
        color: var(--color-fill-1);
        font-weight: 500;
        font-size: 20px;
        line-height: 28px;
    }

    &-sub-title {
        margin-top: 8px;
        color: var(--color-text-3);
        font-size: 14px;
        line-height: 22px;
    }

    &-image {
        width: 320px;
        margin-top: 30px;
    }
}
</style>
