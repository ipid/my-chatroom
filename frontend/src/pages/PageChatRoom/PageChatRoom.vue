<template>
  <div class="pcr__container">
    <div class="pcr__toolbar">
      <el-button :icon="EditIcon" round @click="editNickname">改名</el-button>
    </div>
    <div ref="msgContainer" class="pcr__messages">
      <Message v-for="(msg, index) in msgManager.messages" :key="index" :msg="msg"/>
    </div>
    <div class="pcr__send-toolbar">
      <el-popover
        placement="top"
        title="提示"
        :width="200"
        trigger="click"
        content="可以在聊天框中发送形如 2d9、3d12 的内容，服务端会将其转换为骰子。（最大范围：99d99）"
      >
        <template #reference>
          <div class="pcr__dice-icon iconfont">&#xe6b3;</div>
        </template>
      </el-popover>
    </div>
    <div class="pcr__send-form">
      <MessageSendForm @sendMessage="onSendMessage"/>
    </div>
  </div>

  <ChangeNickname
    v-model="nicknameDialogShown"
    :nickname="nickname"
    @update:nickname="onNicknameUpdated"
  />
</template>

<script setup lang="ts">
import {onMounted, ref, nextTick} from 'vue'
import {Edit as EditIcon} from '@element-plus/icons-vue'
import 'element-plus/es/components/message/style/css'
import Message from '../../components/Message/Message.vue'
import MessageSendForm from './MessageSendForm.vue'
import ChangeNickname from './ChangeNickname.vue'
import {MessageRenderManager} from './MessageRenderManager'

/* --- 昵称相关部分 --- */

const nickname = ref('匿名')
const nicknameDialogShown = ref(false)

function editNickname() {
  nicknameDialogShown.value = true
}

function onNicknameUpdated(newNickname) {
  nickname.value = newNickname
}

/* --- 消息接收部分 --- */

const msgManager = new MessageRenderManager(nickname)
const msgContainer = ref(null)

onMounted(async () => {
  await msgManager.connect()
})

msgManager.onRecvMessage.add(async () => {
  await nextTick()
  msgContainer.value.scrollTop = 1e9
})

/* --- 消息发送部分 --- */

async function onSendMessage(content) {
  msgManager.sendMessage(content)
  await nextTick()
  msgContainer.value.scrollTop = 1e9
}

</script>

<style lang="scss" scoped>
.pcr__container {
  $border-color: #e8e6e6;

  margin: 0 auto;
  max-width: 800px;
  height: 100%;
  border-left: $border-color 1px solid;
  border-right: $border-color 1px solid;
  background: white;

  display: flex;
  flex-direction: column;

  .pcr__toolbar {
    padding: 10px 10px;
    text-align: right;
    border-bottom: $border-color 1px dashed;
  }

  .pcr__messages {
    position: relative;

    flex: 1 0 0;
    border-bottom: $border-color 1px solid;
    padding: 20px 40px;

    overflow-x: hidden;
    overflow-y: auto;
  }

  .pcr__send-toolbar {
    flex: 0 0 auto;
    padding: 5px 30px 0;

    & * {
      user-select: none;
      -webkit-tap-highlight-color: transparent;
    }

    .pcr__dice-icon {
      color: darken($border-color, 15%);
      display: inline-block;
      padding: 5px;
      font-size: 20px;
      cursor: pointer;

      border-radius: 10px;
      transition: background-color 0.2s ease-in-out;

      &:hover {
        background: #f0f0f0;
      }
    }
  }

  .pcr__send-form {
    flex: 0 0 auto;
  }
}
</style>
