import {reactive} from 'vue'
import {ElMessage} from 'element-plus'
import {
  REQ_TYPE_MESSAGE,
  RESP_TYPE_MESSAGE_ACKNOWLEDGE,
  RESP_TYPE_NEW_MESSAGE,
} from '../../constants/def'
import WsMessageServiceClient from './WsMessageServiceClient'

export class MessageRenderManager {
  #messages
  #nickname
  #ackQueue
  #sendMsgCount
  #onRecvMessage
  #wsClient

  constructor(nickname) {
    this.#messages = reactive([])

    this.#nickname = nickname
    this.#ackQueue = []
    this.#sendMsgCount = 0
    this.#onRecvMessage = new Set()

    const wsClient = new WsMessageServiceClient()
    // const wsClient = {
    //   async connect() {
    //   },
    //   isConnected() {
    //     return true
    //   },
    //   sendMessage() {
    //   },
    //   close() {
    //   },
    // }
    // console.error('DEBUG: wsClient 已使用 Mock 实现，记得一定要改回来')

    wsClient.onConnected = () => {
      ElMessage({
        message: '成功连接到服务器',
        type: 'success',
        duration: 1500,
      })
    }
    wsClient.onDisconnected = () => {
      ElMessage({
        message: '与服务器失去连接',
        type: 'error',
        duration: 1500,
      })
    }
    wsClient.onReceiveMessage = (msg) => {
      this.handleMessage(msg)
    }

    this.#wsClient = wsClient
  }

  get messages() {
    return this.#messages
  }

  get onRecvMessage() {
    return this.#onRecvMessage
  }

  async connect() {
    const connectingToast = ElMessage({
      message: '正在连接服务器……',
      duration: 0,
    })

    try {
      await this.#wsClient.connect()
    } catch {
      // 什么也不做
    } finally {
      connectingToast.close()
    }
  }

  #insertMessage(msg) {
    let target

    // O(n) 寻找 id 比 msg 小的第一个元素
    for (target = this.#messages.length - 1; target >= 0; target--) {
      const id = this.#messages[target].data.id
      if (id >= 0 && id < msg.data.id) {
        break
      }
    }

    // 往后寻找第一个不为假消息的元素
    for (target += 1; target < this.#messages.length; target++) {
      const id = this.#messages[target].data.id
      if (id >= 0) {
        break
      }
    }
    // O(n) 插入
    this.#messages.splice(target, 0, msg)
  }

  handleMessage(msg) {
    if (msg.type === RESP_TYPE_NEW_MESSAGE) {
      msg.data.isSender = false
      msg.data.acknowledged = true
      this.#insertMessage(msg)
      for (const handler of this.#onRecvMessage) {
        handler(msg)
      }

    } else if (msg.type === RESP_TYPE_MESSAGE_ACKNOWLEDGE) {
      msg.data.isSender = true
      msg.data.acknowledged = true

      if (this.#ackQueue.length <= 0) {
        throw new Error('#ackQueue 怎么可能为空？你这服务器玩阴的是吧，他奶奶的，直接来吧')
      }

      const fakeMsgId = this.#ackQueue[0]
      this.#ackQueue.shift()

      let fakeMsgIndex
      for (fakeMsgIndex = this.#messages.length - 1; fakeMsgIndex >= 0; fakeMsgIndex--) {
        if (this.#messages[fakeMsgIndex].data.id === fakeMsgId) {
          break
        }
      }
      if (fakeMsgIndex < 0) {
        throw new Error(`不可能，找不到 id 为 ${fakeMsgId} 的 fakeMsg，一定是哪里出问题了`)
      }
      this.#messages.splice(fakeMsgIndex, 1)

      msg.type = RESP_TYPE_NEW_MESSAGE
      this.#insertMessage(msg)

    } else {
      throw new Error(`未知的消息类型: ${msg.type}`)
    }
  }

  sendMessage(richTextContent) {
    try {
      this.#wsClient.sendMessage({
        type: REQ_TYPE_MESSAGE,
        data: {
          sender: this.#nickname.value,
          content: richTextContent,
        },
      })
    } catch (e) {
      ElMessage({
        message: e.message,
        type: 'error',
        duration: 1500,
      })

      return
    }

    this.#sendMsgCount++
    const fakeMsg = {
      type: RESP_TYPE_NEW_MESSAGE,
      data: {
        id: -this.#sendMsgCount,
        timestampMs: Date.now(),
        sender: this.#nickname.value,
        isSender: true,
        acknowledged: false,
        richText: {
          content: richTextContent,
        },
      },
    }
    this.#messages.push(fakeMsg)
    this.#ackQueue.push(fakeMsg.data.id)
  }
}
