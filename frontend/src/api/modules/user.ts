import axios from 'axios'
import type { RouteRecordNormalized } from 'vue-router'
import type { UserState } from '~/store/modules/user/types'

export interface LoginData {
  username: string
  password: string
}

export interface LoginRes {
  token: string
}
export function login(data: LoginData) {
  return axios.post<LoginRes>('/api/user/login', data)
}

export function logout() {
  return axios.post<LoginRes>('/api/user/logout')
}

export function getUserInfo() {
  return axios.post<UserState>('/api/user/info')
}

export function getMenuList() {
  return axios.post<RouteRecordNormalized[]>('/api/user/menu')
}

export interface SetSexReq {
  sex: number
}

export function setSex(data: SetSexReq) {
  return axios.post('/api/user/setSex', data)
}

export interface SetAvatarReq {
  avatar: string
}

export function setAvatar(data: SetAvatarReq) {
  return axios.post('/api/user/setAvatar', data)
}

export interface SetDescriptionReq {
  description: string
}

export function setDescription(data: SetDescriptionReq) {
  return axios.post('/api/user/setDescription', data)
}
