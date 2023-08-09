<script lang="ts" setup></script>

<template>
  <a-layout-footer class="footer">
    @2023 TiVeMachine, <icon-github />: <a href="https://github.com/kirakiseki/TiVeMachine">kirakiseki/TiVeMachine</a>, Powered by Vue3, Vite, Arco Design.
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
}

.footer a {
    color: var(--color-text-2);
}
</style>
