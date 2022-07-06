import {createRouter, createWebHashHistory} from 'vue-router'
import PageChatRoom from './pages/PageChatRoom'

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    {path: '/', name: 'chat-room', component: PageChatRoom},
  ],
})
