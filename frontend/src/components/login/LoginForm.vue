<template>
    <div class="login-form-wrapper">
        <div class="login-form-title">登录 TiVeMachine</div>
        <div class="login-form-sub-title">登录 TiVeMachine</div>
        <div class="login-form-error-msg">错误信息Placeholder</div>
        <a-form ref="loginForm" :model="userInfo" class="login-form" layout="vertical" @submit="handleSubmit">
            <a-form-item field="username" hide-label>
                <a-input>
                    <template #prefix>
                        <icon-user />
                    </template>
                </a-input>
            </a-form-item>
            <a-form-item field="password" hide-label>
                <a-input-password allow-clear>
                    <template #prefix>
                        <icon-lock />
                    </template>
                </a-input-password>
            </a-form-item>
            <a-space direction="vertical">
                <div class="login-form-password-actions">
                    <a-checkbox checked="rememberPassword" :model-value="loginConfig.rememberPassword"
                        @change="setRememberPassword as any">
                        记住密码
                    </a-checkbox>
                    <a-link>忘记密码</a-link>
                </div>
                <a-button type="primary" html-type="submit" long>
                    登录
                </a-button>
                <a-button type="text" long class="login-form-register-btn">
                    注册账号
                </a-button>
            </a-space>
        </a-form>
    </div>
</template>

<script setup lang="ts">

const userInfo = reactive({
    username: "sa",
    password: "asf",
});

const handleSubmit = () => {
    console.log(userInfo);
}

const loginConfig = useStorage('login-config', {
    rememberPassword: true,
    username: 'admin',
    password: 'admin',
});

const setRememberPassword = (value: boolean) => {
    loginConfig.value.rememberPassword = value;
};
</script>

<style lang="less" scoped>
.login-form {
    &-wrapper {
        width: 320px;
    }

    &-title {
        color: var(--color-text-1);
        font-weight: 500;
        font-size: 24px;
        line-height: 32px;
    }

    &-sub-title {
        color: var(--color-text-3);
        font-size: 16px;
        line-height: 24px;
    }

    &-error-msg {
        height: 32px;
        color: rgb(var(--red-6));
        line-height: 32px;
    }

    &-password-actions {
        display: flex;
        justify-content: space-between;
    }

    &-register-btn {
        color: var(--color-text-3) !important;
    }
}
</style>
