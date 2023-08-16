<script lang="ts" setup>
import { Message } from '@arco-design/web-vue'
import { API_BASE } from '~/api/middleware'
import useProgram from '~/hooks/program'

const program = useProgram()

const cover = ref<string>()
const name = ref<string>()
const description = ref<string>()
const category = ref<string>()

function onNameChange(e: string) {
  name.value = e
}

function onDescriptionChange(e: string) {
  description.value = e
}

function onCategoryChange(e: string) {
  category.value = e
}

function onUpload(e: any) {
  cover.value = e.response.data.url
}

async function onAddBtn() {
  if (!cover.value)
    return Message.error('请上传封面')
  if (!name.value)
    return Message.error('请输入节目名称')
  if (!description.value)
    return Message.error('请输入节目简介')
  if (!category.value)
    return Message.error('请输入节目分类')

  program.addProgram({
    Cover: cover.value,
    Name: name.value,
    Description: description.value,
    Category: category.value,
  })

  Message.success('添加成功')
}
</script>

<template>
  <div class="container">
    <a-col class="banner">
      <a-row>
        <a-col :span="2">
          <a-typography-title :heading="5" style="margin-top: 0">
            封面
          </a-typography-title>
        </a-col>
        <a-col :span="8">
          <a-upload :action="`${API_BASE}/api/user/uploadCover`" style="margin-top: 20px;" list-type="picture" :limit="1" @success="onUpload" />
        </a-col>
      </a-row>
      <a-row style="margin-top: 40px;">
        <a-col :span="2">
          <a-typography-title :heading="5" style="margin-top: 0">
            名称
          </a-typography-title>
        </a-col>
        <a-col :span="8">
          <a-input :style="{ width: '320px' }" placeholder="请输入节目名称" allow-clear @input="onNameChange" />
        </a-col>
      </a-row>
      <a-row style="margin-top: 40px;">
        <a-col :span="2">
          <a-typography-title :heading="5" style="margin-top: 0">
            简介
          </a-typography-title>
        </a-col>
        <a-col :span="8">
          <a-input :style="{ width: '320px' }" placeholder="请输入节目简介" allow-clear @input="onDescriptionChange" />
        </a-col>
      </a-row>
      <a-row style="margin-top: 40px;padding-bottom: 20px;">
        <a-col :span="2">
          <a-typography-title :heading="5" style="margin-top: 0">
            分类
          </a-typography-title>
        </a-col>
        <a-col :span="8">
          <a-input :style="{ width: '320px' }" placeholder="请输入节目分类" allow-clear @input="onCategoryChange" />
        </a-col>
      </a-row>
    </a-col>
    <a-button type="primary" style="margin-top: 20px;" @click="async () => { await onAddBtn() }">
      保存
    </a-button>
  </div>
</template>

<style lang="less" scoped>
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
