import { ref } from 'vue'
import { Notification as ArcoNotification } from '@arco-design/web-vue'
import { type Channel, ChannelInfo, type Program, ProgramInfo } from '~/api/modules/program'
import { ScheduleInfo, ScheduleList } from '~/api/modules/schedule'
import { AddSubscribe as AddSubscribeFunc, SubscriptionInfo, SubscriptionList, Unsubscribe } from '~/api/modules/subscribe'

export interface Schedule {
  id: number
  program: Program
  channel: Channel
  start_time: number
  end_time: number
}

export interface Subscribe {
  id?: number
  program: string
  channel: string
  alarm_time: number
}

export default function useSubscribe() {
  const subscribeList = reactive<Ref<number[]>>(ref<number[]>([]))

  const fetchSubscribeList = async () => {
    subscribeList.value = (await SubscriptionList()).data.subscription_list
  }

  const scheduleList = async () => {
    return (await ScheduleList()).data
  }

  const scheduleInfo = async (schedule_id: number) => {
    const info = (await ScheduleInfo({ schedule_id })).data.info
    const subscribeListInfo = ref<Schedule[]>([])
    const schedule: Schedule = {
      program: (await ProgramInfo({ program_id: info.ProgramID })).data.info,
      channel: (await ChannelInfo({ channel_id: info.ChannelID })).data.info,
      start_time: info.StartTime,
      end_time: info.EndTime,
      id: schedule_id,
    }
    subscribeListInfo.value.push(schedule)
    return schedule
  }

  const scheduleListInfo = async () => {
    const list = await scheduleList()
    const listArr = []
    for (const schedule of list)
      listArr.push(await scheduleInfo(schedule))
    return listArr
  }

  const subscriptionInfo = async (data: number) => {
    return await SubscriptionInfo({ subscribe_id: data })
  }

  const addSubscribe = async (data: number, alarm_time: number) => {
    await AddSubscribeFunc({ schedule_id: data, alarm_time })
    const scheduleInfo = (await ScheduleInfo({ schedule_id: data })).data.info
    const programName = (await ProgramInfo({ program_id: scheduleInfo.ProgramID })).data.info.Name
    const channelName = (await ChannelInfo({ channel_id: scheduleInfo.ChannelID })).data.info.Name
    setTimeout(() => {
      ArcoNotification.info({
        title: '订阅提醒',
        content: `您订阅的 ${programName} 即将在 ${channelName} 播出，播出时间为 ${new Date(scheduleInfo.StartTime * 1000).toString()}，请及时收看！`,
        duration: 10 * 1000,
      })
      if (window.Notification && Notification.permission !== 'denied') {
        Notification.requestPermission(() => {
          const n = new Notification('订阅提醒', { body: `您订阅的 ${programName} 即将在 ${channelName} 播出，播出时间为 ${new Date(scheduleInfo.StartTime * 1000).toString()}，请及时收看！` })
          n.onshow = function () {
            setTimeout(n.close.bind(n), 10 * 1000)
          }
        })
      }
    }, alarm_time - new Date().getTime())
  }

  const unsubscribe = async (data: number) => {
    await Unsubscribe({ schedule_id: data })
  }

  const fetchSubscribeListInfo = async () => {
    const subIDList = (await SubscriptionList()).data.subscription_list
    const subListInfo = [] as Subscribe[]
    subIDList.forEach(async (e) => {
      const subInfo = await subscriptionInfo(e)
      const scheduleInfo = (await ScheduleInfo({ schedule_id: subInfo.data.info.ScheduleID })).data.info
      subListInfo.push({
        id: scheduleInfo.ID,
        program: (await ProgramInfo({ program_id: scheduleInfo.ProgramID })).data.info.Name,
        channel: ((await ChannelInfo({ channel_id: scheduleInfo.ChannelID })).data.info.Name),
        alarm_time: subInfo.data.info.AlertTime,
      } as Subscribe)
    })
    return subListInfo
  }

  return {
    subscribeList,
    fetchSubscribeList,
    scheduleList,
    scheduleInfo,
    scheduleListInfo,
    addSubscribe,
    unsubscribe,
    subscriptionInfo,
    fetchSubscribeListInfo,
  }
}
