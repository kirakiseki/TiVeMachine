<template>
  <div class="container">
    <a-layout style="height: 100vh;">
      <a-layout-header>
        <Navbar />
      </a-layout-header>
      <a-layout>
        <a-layout-sider
          :collapsible="true"
        >
          <WorkplaceMenu />
        </a-layout-sider>
        <a-layout-content>
          <RouterView />
        </a-layout-content>
      </a-layout>
      <a-layout-footer>
        <Footer />
      </a-layout-footer>
    </a-layout>
  </div>
</template>

<style scoped>
.container > .arco-layout-header,
.container > .arco-layout-footer,
.container > .arco-layout-content {
  display: flex;
  flex-direction: column;
  justify-content: center;
  font-stretch: condensed;
  text-align: center;
}
.container {
  background-color: var(--color-fill-2);
}
</style>
