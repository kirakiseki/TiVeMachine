import axios from 'axios'
import type { AxiosRequestHeaders, AxiosResponse, InternalAxiosRequestConfig } from 'axios'
import { Message } from '@arco-design/web-vue'
import { getToken } from '~/utils/auth'

export interface HttpResponse<T = unknown> {
  status: number
  msg: string
  code: number
  data: T
}

export const API_BASE = import.meta.env.VITE_API_BASE_URL

if (import.meta.env.VITE_API_BASE_URL)
  axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL

axios.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = getToken()
    if (token) {
      if (!config.headers)
        config.headers = {} as AxiosRequestHeaders

      config.headers.Authorization = token
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

axios.interceptors.response.use(
  (response: AxiosResponse<HttpResponse>): any => {
    const res = response
    if (res.data.code !== 200) {
      Message.error({
        content: res.data.msg || 'Error',
        duration: 5 * 1000,
      })
      return Promise.reject(new Error(res.data.msg || 'Error'))
    }
    return res.data
  },
  (error) => {
    Message.error({
      content: error.msg || '请求错误',
      duration: 5 * 1000,
    })
    return Promise.reject(error)
  },
)
