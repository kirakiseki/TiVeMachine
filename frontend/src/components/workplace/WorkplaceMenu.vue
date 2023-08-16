<script setup lang="ts">
const router = useRouter()
</script>

<template>
  <a-menu
    :style="{ width: '200px', height: '100%' }"
    breakpoint="xl"
  >
    <a-sub-menu key="0">
      <template #icon>
        <icon-heart-fill />
      </template>
      <template #title>
        订阅管理
      </template>
      <a-menu-item key="0_0" :onclick="() => { router.push({ name: 'MySubscribe' }) }">
        我的订阅
      </a-menu-item>
      <a-menu-item key="0_1" :onclick="() => { router.push({ name: 'AddSubscribe' }) }">
        添加订阅
      </a-menu-item>
    </a-sub-menu>
    <a-sub-menu key="1">
      <template #icon>
        <icon-ordered-list />
      </template>
      <template #title>
        节目管理
      </template>
      <a-menu-item key="1_0" :onclick="() => { router.push({ name: 'ExistProgram' }) }">
        已有节目
      </a-menu-item>
      <a-menu-item key="1_1" :onclick="() => { router.push({ name: 'AddProgram' }) }">
        添加节目
      </a-menu-item>
    </a-sub-menu>
    <a-sub-menu key="2">
      <template #icon>
        <icon-search />
      </template>
      <template #title>
        信息查询
      </template>
      <a-menu-item key="2_0" :onclick="() => { router.push({ name: 'ChannelInfo' }) }">
        频道信息
      </a-menu-item>
      <a-menu-item key="2_1" :onclick="() => { router.push({ name: 'Workplace' }) }">
        系统信息
      </a-menu-item>
      <a-sub-menu key="2_2">
        <template #title>
          节目信息
        </template>
        <a-menu-item key="2_2_0" :onclick="() => { router.push({ name: 'ListByName' }) }">
          名称查询
        </a-menu-item>
        <a-menu-item key="2_2_1" :onclick="() => { router.push({ name: 'ListByPeriod' }) }">
          时段查询
        </a-menu-item>
        <a-menu-item key="2_2_2" :onclick="() => { router.push({ name: 'ListByCategory' }) }">
          分类查询
        </a-menu-item>
        <a-menu-item key="2_2_3" :onclick="() => { router.push({ name: 'ListByChannel' }) }">
          频道查询
        </a-menu-item>
      </a-sub-menu>
    </a-sub-menu>
  </a-menu>
</template>

<style scoped>
.menu-demo {
  box-sizing: border-box;
  width: 100%;
  height: 600px;
  padding: 40px;
  background-color: var(--color-neutral-2);
}
</style>
