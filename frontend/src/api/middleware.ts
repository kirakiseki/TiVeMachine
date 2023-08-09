import axios from 'axios'
import type { AxiosRequestHeaders, AxiosResponse, InternalAxiosRequestConfig } from 'axios'
import { Message, Modal } from '@arco-design/web-vue'
import { useUserStore } from '~/store'
import { getToken } from '~/utils/auth'

export interface HttpResponse<T = unknown> {
  status: number
  msg: string
  code: number
  data: T
}

if (import.meta.env.VITE_API_BASE_URL)
  axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL

axios.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // TODO: Authorization
    const token = getToken()
    if (token) {
      if (!config.headers)
        config.headers = {} as AxiosRequestHeaders

      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    // TODO: Error handling
    return Promise.reject(error)
  },
)

axios.interceptors.response.use(
  (response: AxiosResponse<HttpResponse>) => {
    const res = response
    // if the custom code is not 20000, it is judged as an error.
    if (res.data.code !== 20000) {
      Message.error({
        content: res.data.msg || 'Error',
        duration: 5 * 1000,
      })
      // 50008: Illegal token; 50012: Other clients logged in; 50014: Token expired;
      if (
        [50008, 50012, 50014].includes(res.data.code)
                && response.config.url !== '/api/user/info'
      ) {
        Modal.error({
          title: 'Confirm logout',
          content:
                        'You have been logged out, you can cancel to stay on this page, or log in again',
          okText: 'Re-Login',
          async onOk() {
            const userStore = useUserStore()

            await userStore.logout()
            window.location.reload()
          },
        })
      }
      return Promise.reject(new Error(res.data.msg || 'Error'))
    }
    return res
  },
  (error) => {
    Message.error({
      content: error.msg || '请求错误',
      duration: 5 * 1000,
    })
    return Promise.reject(error)
  },
)
