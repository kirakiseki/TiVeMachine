<script setup lang="ts">
import useUser from '~/hooks/user'
import { useUserStore } from '~/store'
import { API_BASE } from '~/api/middleware'

const user = useUserStore()
const userMani = useUser()
const value = ref<number>(user.Sex as number)

const onSexChange = useDebounceFn((e) => {
  console.log(e)
  userMani.updateSex(e)
}, 1000)
const onDescriptionChange = useDebounceFn((e) => {
  console.log(e)
  userMani.updateDescription(e)
}, 1000)

function onSuc(e: any) {
  console.log(e)
  userMani.updateAvatar(e.response.data.url)
}
</script>

<template>
  <div class="container">
    <a-col class="banner">
      <a-col :span="8">
        <a-typography-title :heading="5" style="margin-top: 0">
          个人信息
        </a-typography-title>
      </a-col>
      <a-divider class="panel-border" />
    </a-col>
    <a-col class="banner">
      <a-row>
        <a-col :span="2">
          <a-typography-title :heading="5" style="margin-top: 0">
            头像
          </a-typography-title>
        </a-col>
        <a-col :span="8">
          <a-avatar :size="64" :image-url="user.Avatar" />
          <a-upload :action="`${API_BASE}/api/user/uploadAvatar`" style="margin-top: 20px;" list-type="picture" :limit="1" @success="onSuc" />
        </a-col>
        <a-col :span="2">
          <a-typography-title :heading="5" style="margin-top: 0">
            性别
          </a-typography-title>
        </a-col>
        <a-col :span="8">
          <a-select v-model="value" :style="{ width: '30%' }" :placeholder="value === 0 ? '男' : '女'" @change="onSexChange">
            <a-option :value="0">
              男
            </a-option>
            <a-option :value="1">
              女
            </a-option>
          </a-select>
        </a-col>
      </a-row>
      <a-divider class="panel-border" />
      <a-row style="margin-top: 40px;">
        <a-col :span="2">
          <a-typography-title :heading="5" style="margin-top: 0">
            简介
          </a-typography-title>
        </a-col>
        <a-col :span="8">
          <a-input :style="{ width: '320px' }" :placeholder="user.Description" allow-clear @input="onDescriptionChange" />
        </a-col>
        <a-col :span="2">
          <a-typography-title :heading="5" style="margin-top: 0">
            UID
          </a-typography-title>
        </a-col>
        <a-col :span="8">
          <a-typography-title :heading="3" style="margin-top: 0">
            {{ user.ID }}
          </a-typography-title>
        </a-col>
      </a-row>
      <a-divider class="panel-border" />
    </a-col>
  </div>
</template>

<style scoped lang="less">
  .container {
    background-color: var(--color-fill-2);
    padding: 16px 20px;
    padding-bottom: 20px;
    display: flex;
    flex-direction: column;
  }

  .title {
    width: 100%;
    padding: 0 20px 0 20px;
    background-color: var(--color-bg-2);
    border-radius: 4px 4px 0 0;
  }

  .banner {
    width: 100%;
    padding: 20px 20px 0 20px;
    background-color: var(--color-bg-2);
    border-radius: 4px 4px 0 0;
  }

  :deep(.arco-icon-home) {
    margin-right: 6px;
  }
</style>

<route lang="yaml">
name: User
meta:
  requiresAuth: true
  layout: user-layout
</route>
