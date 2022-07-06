/**
 * 获取后端 WebSocket URL。
 * 由于本地不方便配置 TLS 证书，故在开发时会获取到 ws 协议 localhost:8012 的 URL，在生产时会获取到不带端口号的 wss 协议 URL。
 */
function getBackendWsUrl(): string {
  if (import.meta.env.DEV) {
    return `ws://${location.hostname}:8012/ws/chat`
  }

  return `${(location.protocol === 'https:') ? 'wss' : 'ws'}://${location.host}/ws/chat`
}

/**
 * 用于等待 WebSocket 对象切换到开启状态的函数。
 * @param ws WebSocket 对象
 * @return 一个无值的 Promise，当 WebSocket 切换到开启对象后会被 resolve。
 */
function waitWebSocketOpen(ws: WebSocket): Promise<void> {
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

  onReceiveMessage: ()
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
