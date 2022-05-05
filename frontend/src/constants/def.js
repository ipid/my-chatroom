/* 客户端请求类型 */

const sendMessageExample = {
  'type': 'send-msg',
  'data': {
    'content': 'Hello!',
    'sender': 'Fuck',
  },
}

export const REQ_TYPE_MESSAGE = 'send-msg'

/* 服务端返回类型 */

const richMessageExample = {
  'type': 'msg',
  'data': {
    'id': 123,
    'content': 'Hello!',
    'timestampMs': 1234567,
    'sender': 'Fuck',
    'acknowledged': true,
  },
}

export const RESP_TYPE_NEW_MESSAGE = 'msg'

const diceMessageExample = {
  'type': 'dice',
  'data': {
    'id': 123,
    'dices': [1, 22, 333],
    'timestampMs': 1234567,
    'sender': 'Fuck',
    'isSender': true,
    'acknowledged': true,
  },
}

const msgAckExample = {
  'type': 'msg-ack',
  'data': {
    'newMsg': {
      'type': 'dice',
      'data': {
        'id': 123,
        'dices': [1, 2, 3],
        'timestampMs': 1234567,
        'sender': 'Fuck',
      },
    },
  },
}

export const RESP_TYPE_MESSAGE_ACKNOWLEDGE = 'msg-ack'
