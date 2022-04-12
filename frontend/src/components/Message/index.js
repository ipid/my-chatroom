import Message from './Message.vue'

/*
消息类型示例：

1. 普通富文本消息
{
  "type": "msg",
  "data": {
    "id": 123,
    "content": "Hello!",
    "timestampMs": 1234567,
    "sender": "Fuck",
    "isSender": true,
    "acknowledged": false
  }
}

2. 骰子消息
{
  "type": "dice",
  "data": {
    "id": 123,
    "dices": [1, 2, 3],
    "timestampMs": 1234567,
    "sender": "Fuck",
    "isSender": true,
    "acknowledged": false
  }
}

*/
export default Message

