import { defineStore } from 'pinia'
import useAppStore from '../view'
import type { UserState } from './types'
import type { LoginData } from '~/api/modules/user'
import {
  getUserInfo,
  login as userLogin,
  logout as userLogout,
} from '~/api/modules/user'
import { clearToken, setToken } from '~/utils/auth'

const useUserStore = defineStore('user', {
  state: (): UserState => ({
    name: undefined,
    avatar: undefined,
    email: undefined,
    phone: undefined,
    registrationDate: undefined,
    accountId: undefined,
    role: '',
  }),

  getters: {
    userInfo(state: UserState): UserState {
      return { ...state }
    },
  },

  actions: {
    switchRoles() {
      return new Promise((resolve) => {
        this.role = this.role === 'user' ? 'admin' : 'user'
        resolve(this.role)
      })
    },
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
        await userLogout()
      }
      finally {
        this.logoutCallBack()
      }
    },
  },
})

export default useUserStore
