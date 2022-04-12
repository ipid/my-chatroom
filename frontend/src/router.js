import {createRouter, createWebHashHistory} from 'vue-router'
import PageChatRoom from './pages/PageChatRoom'

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    {path: '/', nickname: 'chat-room', component: PageChatRoom},
  ],
})
