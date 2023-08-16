import { ref } from 'vue'
import type { Schedule } from './subscribe'
import { AddProgram, ChannelInfo, ChannelList, DeleteProgram, ListByCategory, ListByChannel, ListByPeriod, ProgramInfo, ProgramList, Search } from '~/api/modules/program'
import type { Channel, Program } from '~/api/modules/program'
import { ScheduleInfo, ScheduleList } from '~/api/modules/schedule'

export default function useProgram() {
  const programList = ref<Program[]>([])
  const channelList = ref<Channel[]>([])
  const scheduleList = ref<Schedule[]>([])
  const fetchProgramList = async () => {
    const list = (await ProgramList()).data
    programList.value = []
    list.list.forEach(async (element) => {
      programList.value.push((await ProgramInfo({ program_id: element })).data.info)
    })
    return programList.value
  }
  const fetchChannelList = async () => {
    const list = (await ChannelList()).data
    channelList.value = []
    list.list.forEach(async (element) => {
      channelList.value.push((await ChannelInfo({ channel_id: element })).data.info)
    })
    return channelList.value
  }
  const fetchScheduleList = async () => {
    const list = (await ScheduleList()).data
    scheduleList.value = []
    list.list.forEach(async (element) => {
      const info = (await ScheduleInfo({ schedule_id: element })).data.info
      scheduleList.value.push({
        program: (await ProgramInfo({ program_id: info.ProgramID })).data.info,
        channel: (await ChannelInfo({ channel_id: info.ChannelID })).data.info,
        start_time: info.StartTime,
        end_time: info.EndTime,
        id: info.ID,
      })
    })
    return scheduleList.value
  }
  const listByPeriod = async (start_time: number, end_time: number) => {
    const list = (await ListByPeriod({ start_time, end_time })).data
    programList.value = []
    list.list.forEach(async (element) => {
      programList.value.push((await ProgramInfo({ program_id: element })).data.info)
    })
    return programList.value
  }
  const listByName = async (name: string) => {
    const list = (await Search({ keyword: name })).data
    console.log(list)
    return list
  }

  const listByCategory = async (category: string) => {
    const list = (await ListByCategory({ category })).data
    programList.value = []
    list.list.forEach(async (element) => {
      programList.value.push((await ProgramInfo({ program_id: element })).data.info)
    })
    return programList.value
  }
  const listByChannel = async (channel_id: number) => {
    const list = (await ListByChannel({ channel_id })).data
    programList.value = []
    list.list.forEach(async (element) => {
      programList.value.push((await ProgramInfo({ program_id: element })).data.info)
    })
    return programList.value
  }
  const search = async (keyword: string) => {
    const list = (await Search({ keyword })).data
    programList.value = []
    list.list.forEach(async (element) => {
      programList.value.push((await ProgramInfo({ program_id: element })).data.info)
    })
    return programList.value
  }
  const addProgram = async (data: Program) => {
    await AddProgram(data)
  }
  const deleteProgram = async (data: number) => {
    await DeleteProgram({ program_id: data })
  }
  const scheduleCount = async () => {
    const list = (await ScheduleList()).data
    return list.list.length
  }
  return {
    programList,
    fetchProgramList,
    channelList,
    fetchChannelList,
    scheduleList,
    fetchScheduleList,
    listByPeriod,
    listByCategory,
    listByChannel,
    listByName,
    search,
    addProgram,
    deleteProgram,
    scheduleCount,
  }
}
