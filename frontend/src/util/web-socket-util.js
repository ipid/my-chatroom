export function getBackendWsUrl() {
  if (import.meta.env.DEV) {
    return `ws://${location.hostname}:8012/ws/chat`
  }

  return `${(location.protocol === 'https:') ? 'wss' : 'ws'}://${location.host}/ws/chat`
}

export function waitWebSocketOpen(ws) {
  return new Promise((resolve, reject) => {
    if (ws.readyState === WebSocket.OPEN) {
      resolve()
    }

    function openHandler() {
      resolve()
      ws.removeEventListener('open', openHandler)
    }

    ws.addEventListener('open', openHandler)
  })
}