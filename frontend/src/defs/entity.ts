// 聊天室内的用户
export interface User {
  // 用户 id，在同一聊天室中唯一
  userId: number;

  // 用户的显示名称
  name: string;
}

// 聊天室消息的基本字段
interface MessageBase {
  // 消息 id
  msgId: number;

  // 发送该消息的用户 id
  senderUserId: number;

  // 发送消息时的时间戳（单位：毫秒）
  timestampMs: number;
}

// 文本消息
export interface TextMessage extends MessageBase {
  text: {
    // 表示文本消息的内容
    content: string;
  }
}

// 骰子消息中的一个骰子
export interface Dice {
  // 表示该骰子的取值范围是 [1, max]
  max: number;

  // 表示该骰子投出来的值
  value: number;
}

// 骰子消息
export interface DiceMessage extends MessageBase {
  // 骰子
  dices: Dice[]
}

// 表示一条消息，该消息的类型不确定
export type Message = TextMessage | DiceMessage;
