import { Message } from "./entity.js";

export interface WebSocketParam {
  // 上次登录用的用户 id
  id: number;

  // 上次登录时服务器发放的 token
  token: string;
}

// 客户端请求，表示客户端希望发送一条消息
export interface SendRequest {
  send: {
    // 要发送的消息内容，可以是普通文本，也可以是 1d9 等骰子文本
    content: string;
  }
}

// 服务器端初始回应，用于将 token、当前用户、当前历史消息等告诉用户
export interface InitialResponse {
  // 客户端当前的用户 id
  userId: number;

  // 第一次登录时，这里会存放登录用的 token；用 token 登录时此字段为空字符串
  token: string;

  // 服务器端的历史消息
  history: Message[]
}

// 服务器端消息回应，表示收到了一或多条消息
export interface MessageResponse {
  // 收到的消息
  msgs: Message[]
}

// 服务器端确认回应，表示客户端发出的请求得到了服务器端的确认
export interface AckResponse {
  // 一条消息，其内容与同一聊天室其他客户端收到的消息一致
  ack: Message;
}

// 表示一条客户端请求，其类型不确定
export type Request = SendRequest;

// 表示一条服务器端的消息，其类型不确定
export type Response = InitialResponse | MessageResponse | AckResponse;