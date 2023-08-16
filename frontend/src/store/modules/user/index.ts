import { defineStore } from 'pinia'
import useAppStore from '../view'
import type { UserState } from './types'
import type { LoginData } from '~/api/modules/user'
import {
  getUserInfo,
  login as userLogin,
} from '~/api/modules/user'
import { clearToken, setToken } from '~/utils/auth'

const useUserStore = defineStore('user', {
  state: (): UserState => ({
    ID: 0,
    Username: '',
    Avatar: '',
    Description: '',
    Sex: 0,
  }),

  getters: {
    userInfo(state: UserState): UserState {
      return { ...state }
    },
  },

  actions: {
    // Set user's information
    setInfo(partial: Partial<UserState>) {
      this.$patch(partial)
    },

    // Reset user's information
    resetInfo() {
      this.$reset()
    },

    // Get user's information
    async info() {
      const res = await getUserInfo()
      if (res.data.Avatar === '')
        res.data.Avatar = 'http://s3.tivemachine.com/avatar/avatar.jpg'
      if (res.data.Description === '')
        res.data.Description = '这个人很懒，什么都没有留下'

      this.setInfo(res.data)
    },

    // Login
    async login(loginForm: LoginData) {
      try {
        const res = await userLogin(loginForm)
        setToken(res.data.token)
      }
      catch (err) {
        clearToken()
        throw err
      }
    },
    logoutCallBack() {
      const appStore = useAppStore()
      this.resetInfo()
      clearToken()
      appStore.clearServerMenu()
    },
    // Logout
    async logout() {
      try {
        // await userLogout()
      }
      finally {
        this.logoutCallBack()
      }
    },
  },
})

export default useUserStore
