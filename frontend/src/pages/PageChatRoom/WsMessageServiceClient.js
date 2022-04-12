import {REQ_TYPE_MESSAGE} from '../../constants/def'

function getBackendWsUrl() {
  if (import.meta.env.DEV) {
    return `ws://${location.hostname}:8012/ws/chat`
  }

  return `${(location.protocol === 'https:') ? 'wss' : 'ws'}://${location.host}/ws/chat`
}

function waitWebSocketOpen(ws) {
  return new Promise((resolve, reject) => {
    if (ws.readyState === WebSocket.OPEN) {
      resolve()
    }

    function openHandler() {
      resolve()
      ws.removeEventListener('open', openHandler)
    }

    function closeHandler() {
      reject()
      ws.removeEventListener('close', closeHandler)
    }

    ws.addEventListener('open', openHandler)
    ws.addEventListener('close', closeHandler)
  })
}

export default class WsMessageServiceClient {
  #wsConn = null

  onReceiveMessage = (msg) => {
  }
  onConnected = () => {
  }
  onDisconnected = () => {
  }

  async connect() {
    if (this.isConnected()) {
      return
    }

    const ws = new WebSocket(getBackendWsUrl())
    ws.addEventListener('close', () => {
      this.#wsConn = null
      this.onDisconnected()
    })
    ws.addEventListener('message', (ev) => {
      if (typeof ev.data !== 'string') {
        return
      }

      try {
        this.onReceiveMessage(JSON.parse(ev.data))
      } catch (e) {
        console.error(e)
        ws.close()
      }
    })

    await waitWebSocketOpen(ws)

    this.#wsConn = ws
    this.onConnected()
  }

  isConnected() {
    return this.#wsConn !== null
  }

  sendMessage(messageObject) {
    if (this.#wsConn === null) {
      throw new Error('服务器未连接，无法发送消息')
    }

    this.#wsConn.send(JSON.stringify(messageObject))
  }

  close() {
    if (this.#wsConn !== null) {
      this.#wsConn.close()
    }
  }
}
