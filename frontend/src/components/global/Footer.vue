<template>
    <a-layout-footer class="footer">
        @2023 TiVeMachine, GitHub Repo: <a href="https://github.com/kirakiseki/TiVeMachine">kirakiseki/TiVeMachine</a>, Powered by Vue3, Vite, Arco Design Pro.
    </a-layout-footer>
</template>

<script lang="ts" setup></script>

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
