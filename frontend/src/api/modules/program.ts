import axios from 'axios'

export interface ListRes {
  list: number[]
}

export interface ProgramInfoData {
  program_id: number
}

export interface Program {
  id?: number
  Name: string
  Description: string
  Cover: string
  Category: string
}

export interface DeleteProgramData {
  program_id: number
}

export interface ProgramInfoRes {
  info: Program
}

export interface ChannelInfoData {
  channel_id: number
}

export interface Channel {
  ID: number
  Name: string
  Description: string
  Cover: string
}

export interface ChannelInfoRes {
  info: Channel
}

export interface ListByPeriodData {
  start_time: number
  end_time: number
}

export interface ListByCategoryData {
  category: string
}

export interface ListByChannelData {
  channel_id: number
}

export interface SearchData {
  keyword: string
}

export function ProgramList() {
  return axios.get<ListRes>('/api/program/list')
}

export function ChannelList() {
  return axios.get<ListRes>('/api/program/channelList')
}

export function ProgramInfo(data: ProgramInfoData) {
  return axios.post<ProgramInfoRes>('/api/program/info', data)
}

export function ChannelInfo(data: ChannelInfoData) {
  return axios.post<ChannelInfoRes>('/api/program/channelInfo', data)
}

export function ListByPeriod(data: ListByPeriodData) {
  return axios.post<ListRes>('/api/program/listByPeriod', data)
}

export function ListByCategory(data: ListByCategoryData) {
  return axios.post<ListRes>('/api/program/listByCategory', data)
}

export function ListByChannel(data: ListByChannelData) {
  return axios.post<ListRes>('/api/program/listByChannel', data)
}

export function Search(data: SearchData) {
  return axios.post<ListRes>('/api/program/search', data)
}

export function AddProgram(data: Program) {
  return axios.post('/api/program/addProgram', data)
}

export function DeleteProgram(data: DeleteProgramData) {
  return axios.post('/api/program/deleteProgram', data)
}
