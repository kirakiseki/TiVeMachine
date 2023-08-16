import axios from 'axios'

export interface SubscriptionListRes {
  subscription_list: number[]
}

export function SubscriptionList() {
  return axios.get<SubscriptionListRes>('/api/subscribe/subscriptionList')
}

export interface SubscribeData {
  schedule_id: number
  alarm_time: number
}

export interface UnsubscribeData {
  schedule_id: number
}

export interface SubscriptionInfoData {
  subscribe_id: number
}

export interface Subscribe {
  ID: number
  UserID: number
  ScheduleID: number
  AlertTime: number
}

export interface SubscriptionInfoRes {
  info: Subscribe
}

export function AddSubscribe(data: SubscribeData) {
  return axios.post('/api/subscribe/subscribe', data)
}

export function Unsubscribe(data: UnsubscribeData) {
  return axios.post('/api/subscribe/unsubscribe', data)
}

export function SubscriptionInfo(data: SubscriptionInfoData) {
  return axios.post<SubscriptionInfoRes>('/api/subscribe/subscriptionInfo', data)
}
