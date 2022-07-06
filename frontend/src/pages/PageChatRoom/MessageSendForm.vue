<template>
  <div class="message-send-form__container">
    <el-input
      v-model="inputMessage"
      :prefix-icon="ChatDotRound" class="message-send-form__input"
      placeholder="输入要发送的消息..." size="large"
      @keydown.enter.native="sendMessage"
    />
    <el-button :disabled="inputMessage === ''" round size="large" type="primary" @click="sendMessage">
      <span class="iconfont">&#xe6d0;</span>&nbsp;&nbsp;发送
    </el-button>
  </div>
</template>

<script setup lang="ts">
import {ref} from 'vue'
import {ChatDotRound} from '@element-plus/icons-vue'

const emit = defineEmits(['sendMessage'])

const inputMessage = ref('')

function sendMessage() {
  if (inputMessage.value === '') {
    return
  }

  emit('sendMessage', inputMessage.value)
  inputMessage.value = ''
}
</script>

<style lang="scss" scoped>
.message-send-form__container {
  display: flex;
  align-items: center;
  padding: 20px 30px;

  .message-send-form__input {
    flex: 1 0 0;
    margin-right: 20px;
  }
}
</style>
