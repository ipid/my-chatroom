import {createApp} from 'vue'
import {createRouter, createWebHashHistory} from 'vue-router'
import PageChatRoom from './pages/PageChatRoom.vue'
import App from './App.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {path: '/', name: 'chat-room', component: PageChatRoom},
  ],
})

createApp(App).use(router).mount('#app')
