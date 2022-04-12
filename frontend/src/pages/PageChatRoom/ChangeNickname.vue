<template>
  <el-dialog
    :modelValue="modelValue"
    :width="dialogWidth()"
    title="更改昵称"
    @update:modelValue="(value) => emit('update:modelValue', value)"
  >
    <el-form label-width="40px">
      <el-form-item label="昵称">
        <el-input v-model="inputNickName" @keydown.enter.native="confirmEditNickname"/>
      </el-form-item>
      <el-form-item label="">
        <el-button :icon="CheckIcon" round size="large" type="primary" @click="confirmEditNickname">确定</el-button>
        <el-button round size="large" @click="cancelEditNickname">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import {ref, watch} from 'vue'
import {Check as CheckIcon} from '@element-plus/icons-vue'

function dialogWidth() {
  return (window.innerWidth <= 600) ? '80%' : '480px'
}

const props = defineProps({
  nickname: String,
  modelValue: Boolean,
})
const emit = defineEmits(['update:nickname', 'update:modelValue'])

const inputNickName = ref(props.nickname)

watch(() => props.modelValue, (newDialogShown, oldDialogShown) => {
  if (newDialogShown && !oldDialogShown) {
    inputNickName.value = props.nickname
    console.log('已更新昵称')
  }
})

function confirmEditNickname() {
  let nickname = inputNickName.value.replace(/^\s+|\s+$/g, '')
  if (nickname === '') {
    nickname = props.nickname
  }

  emit('update:nickname', nickname)
  emit('update:modelValue', false)
}

function cancelEditNickname() {
  emit('update:modelValue', false)
}

</script>

<style lang="scss" scoped>
// SCSS 代码
</style>
