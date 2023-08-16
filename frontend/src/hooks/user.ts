import { useRouter } from 'vue-router'
import { Message } from '@arco-design/web-vue'

import { useUserStore } from '~/store'
import { setAvatar, setDescription, setSex } from '~/api/modules/user'

export default function useUser() {
  const router = useRouter()
  const userStore = useUserStore()
  const logout = async (logoutTo?: string) => {
    await userStore.logout()
    const currentRoute = router.currentRoute.value
    Message.success('登出成功')
    router.push({
      name: logoutTo && typeof logoutTo === 'string' ? logoutTo : 'login',
      query: {
        ...router.currentRoute.value.query,
        redirect: currentRoute.name as string,
      },
    })
  }
  const updateSex = async (data: number) => {
    await setSex({ sex: data })
    await userStore.info()
  }
  const updateAvatar = async (data: string) => {
    console.log(data)
    await setAvatar({ avatar: data })
    await userStore.info()
  }
  const updateDescription = async (data: string) => {
    await setDescription({ description: data })
    await userStore.info()
  }
  return {
    logout,
    updateSex,
    updateAvatar,
    updateDescription,
  }
}
