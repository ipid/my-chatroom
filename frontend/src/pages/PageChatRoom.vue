<template>
  <div class="pcr__container">
    <div class="pcr__toolbar">
      <el-button :icon="EditIcon" round @click="editNickname">改名</el-button>
    </div>
    <div ref="msgContainer" class="pcr__messages">
      <Message
        v-for="(msg, index) in messages" :key="index"
        :content="msg.content" :is-sender="msg.isSender" :name="msg.sender" :timestampMs="msg.timestampMs"
      />
    </div>
    <div class="pcr__send-form">
      <MessageSendForm @sendMessage="onSendMessage"/>
    </div>
  </div>
  <div class="pcr__height-filler" />
  <el-dialog v-model="nicknameDialogShown" title="更改昵称" :width="dialogWidth()">
    <el-form v-model="editNicknameForm" label-width="40px">
      <el-form-item label="昵称">
        <el-input @keydown.enter.native="confirmEditNickname" v-model="editNicknameForm.nickname"/>
      </el-form-item>
      <el-form-item label="">
        <el-button :icon="CheckIcon" type="primary" round size="large" @click="confirmEditNickname">确定</el-button>
        <el-button round size="large" @click="cancelEditNickname">取消</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import Message from '../components/Message.vue'
import MessageSendForm from '../components/MessageSendForm.vue'
import {getBackendWsUrl, waitWebSocketOpen} from '../util/web-socket-util'
import {nextTick, onMounted, reactive, ref} from 'vue'
import {ElMessage} from 'element-plus'
import {Check as CheckIcon, Edit as EditIcon} from '@element-plus/icons-vue'
import 'element-plus/es/components/message/style/css'

/* 昵称相关部分 */
function dialogWidth() {
  return (window.innerWidth <= 600) ? '80%' : '480px'
}

const DEFAULT_NICKNAME = '匿名'
let nickname = DEFAULT_NICKNAME
const editNicknameForm = reactive({
  nickname,
})
const nicknameDialogShown = ref(false)

function editNickname() {
  editNicknameForm.nickname = nickname
  nicknameDialogShown.value = true
}

function confirmEditNickname() {
  nickname = editNicknameForm.nickname.replace(/^\s+|\s+$/g, '')
  if (nickname === '') {
    nickname = DEFAULT_NICKNAME
  }
  nicknameDialogShown.value = false
}

function cancelEditNickname() {
  nicknameDialogShown.value = false
}

/* 消息收发部分 */
const messages = reactive([])
const msgContainer = ref(null)

let wsConn = null

async function onSendMessage({content}) {
  if (wsConn === null) {
    ElMessage({
      type: 'error',
      message: '尚未连接到服务器，无法发送消息',
      duration: 1500,
    })
    return
  }

  wsConn.send(JSON.stringify({
    sender: nickname,
    content,
  }))

  messages.push({
    sender: nickname,
    timestampMs: Date.now(),
    content,
    isSender: true,
  })

  await nextTick()
  msgContainer.value?.scrollTo(0, 1e9)
}

onMounted(async () => {
  const connectingToast = ElMessage({
    message: '正在连接服务器……',
    duration: 0,
  })

  const ws = new WebSocket(getBackendWsUrl())
  ws.addEventListener('close', () => {
    wsConn = null

    ElMessage({
      message: '与服务器失去连接',
      type: 'error',
      duration: 1500,
    })
  })
  ws.addEventListener('message', (ev) => {
    if (typeof ev.data !== 'string') {
      return
    }

    try {
      const serverMsg = JSON.parse(ev.data)
      messages.push({
        sender: serverMsg.sender,
        timestampMs: serverMsg.timestampMs,
        content: serverMsg.content,
        isSender: false,
      })
    } catch {
      ws.close()
    }
  })

  await waitWebSocketOpen(ws)
  wsConn = ws

  connectingToast.close()
  ElMessage({
    message: '成功连接到服务器',
    type: 'success',
    duration: 1500,
  })
})
</script>

<style lang="scss" scoped>
.pcr__container {
  $border-color: #eee5f0;

  margin: 0 auto;
  max-width: 800px;
  height: calc(var(--vh, vh) * 100);
  border-left: $border-color 1px solid;
  border-right: $border-color 1px solid;
  background: white;

  display: flex;
  flex-direction: column;

  .pcr__toolbar {
    padding: 10px 10px;
    text-align: right;
    border-bottom: #e8e6e6 1px dashed;
  }

  .pcr__messages {
    position: relative;

    flex: 1 0 0;
    border-bottom: #e8e6e6 1px solid;
    padding: 20px 40px;

    overflow-x: hidden;
    overflow-y: auto;
  }

  .pcr__send-form {
    flex: 0 0 auto;
  }
}

.pcr__height-filler {
  height: calc(100vh - var(--vh, vh) * 100);
}
</style>