import axios from 'axios'

export interface Schedule {
  ID: number
  ProgramID: number
  ChannelID: number
  StartTime: number
  EndTime: number
}

export interface ScheduleInfoData {
  schedule_id: number
}

export interface ScheduleInfoRes {
  info: Schedule
}

export interface ScheduleListRes {
  list: number[]
}

export function ScheduleInfo(data: ScheduleInfoData) {
  return axios.post<ScheduleInfoRes>('/api/program/scheduleInfo', data)
}

export function ScheduleList() {
  return axios.get<ScheduleListRes>('/api/program/scheduleList')
}
