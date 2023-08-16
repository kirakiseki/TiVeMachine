<script lang="ts" setup>
import { useUserStore } from '~/store'
import useUser from '~/hooks/user'
import logo from '~/assets/images/logo.svg'
import { isLogin } from '~/utils/auth'

defineProps<{
  off?: boolean
}>()
const userStore = useUserStore()
const router = useRouter()
const { logout } = useUser()
const avatar = computed(() => {
  return userStore.Avatar
})
function handleLogout() {
  logout()
}
</script>

<template>
  <div class="navbar">
    <div class="left-side">
      <a-space>
        <img
          v-if="!off"
          alt="logo"
          :src="logo"
        >
        <a-typography-title
          v-if="!off"
          :style="{ margin: 0, fontSize: '18px' }"
          :heading="5"
        >
          TiVeMachine
        </a-typography-title>
      </a-space>
    </div>
    <ul class="right-side">
      <li v-if="isLogin()">
        <a-dropdown trigger="click">
          <a-avatar
            :size="32"
            :style="{ marginRight: '8px', cursor: 'pointer' }"
          >
            <img alt="avatar" :src="avatar">
          </a-avatar>
          <template #content>
            <a-doption>
              <a-space @click="router.push({ name: 'Workplace' })">
                <icon-user />
                <span>
                  节目中心
                </span>
              </a-space>
            </a-doption>
            <a-doption>
              <a-space @click="router.push({ name: 'User' })">
                <icon-settings />
                <span>
                  用户设置
                </span>
              </a-space>
            </a-doption>
            <a-doption>
              <a-space @click="handleLogout">
                <icon-export />
                <span>
                  退出登录
                </span>
              </a-space>
            </a-doption>
          </template>
        </a-dropdown>
      </li>
      <li v-else>
        <a-button type="primary" @click="() => { router.push('/login') }">
          登录
        </a-button>
      </li>
    </ul>
  </div>
</template>

<style scoped lang="less">
  .navbar {
    display: flex;
    justify-content: space-between;
    height: 100%;
    background-color: var(--color-bg-2);
    border-bottom: 1px solid var(--color-border);
  }

  .left-side {
    display: flex;
    align-items: center;
    padding-left: 20px;
  }

  .center-side {
    flex: 1;
  }

  .right-side {
    display: flex;
    padding-right: 20px;
    list-style: none;
    :deep(.locale-select) {
      border-radius: 20px;
    }
    li {
      display: flex;
      align-items: center;
      padding: 0 10px;
    }

    a {
      color: var(--color-text-1);
      text-decoration: none;
    }
    .nav-btn {
      border-color: rgb(var(--gray-2));
      color: rgb(var(--gray-8));
      font-size: 16px;
    }
    .trigger-btn,
    .ref-btn {
      position: absolute;
      bottom: 14px;
    }
    .trigger-btn {
      margin-left: 14px;
    }
  }
</style>

<style lang="less">
  .message-popover {
    .arco-popover-content {
      margin-top: 0;
    }
  }
</style>
