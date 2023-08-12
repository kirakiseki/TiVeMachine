<script lang="ts" setup></script>

<template>
  <a-layout-footer class="footer">
    <span>@2023 TiVeMachine, <icon-github />: <a href="https://github.com/kirakiseki/TiVeMachine">kirakiseki/TiVeMachine</a>, Powered by Vue3, Vite, Arco Design.</span>
  </a-layout-footer>
</template>

<style lang="less" scoped>
.footer {
    display: flex;
    align-items: center;
    justify-content: center;
    height: 40px;
    color: var(--color-text-2);
    text-align: center;
    background-color: var(--color-bg-2);
}

.footer a {
    color: var(--color-text-2);
}
</style>
