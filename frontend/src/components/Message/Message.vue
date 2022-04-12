<template>
  <div :style="outerStyle" class="message__container">
    <div class="message__sender">
      <span class="message__name">{{ props.msg.data.sender }}</span> &nbsp;
      <span>{{ displayTime }}</span>
    </div>
    <br/>
    <div>
      <transition name="el-fade-in">
        <component :is="MessageComponent"/>
      </transition>
    </div>
  </div>
</template>

<script lang="jsx" setup>
import {computed} from 'vue'
import dayjs from 'dayjs'
import MessageDice from './MessageDice.vue'
import MessageRich from './MessageRich.vue'
import {RESP_TYPE_DICE, RESP_TYPE_NEW_RICH_MESSAGE} from '../../constants/def'

const props = defineProps({
  msg: {
    type: Object,
    required: true,
  },
})

const displayTime = computed(() => dayjs(props.msg.data.timestampMs).format('HH:mm:ss'))

const outerStyle = computed(() => {
  if (props.msg.data.isSender) {
    return {
      'text-align': 'right',
    }
  }
  return {}
})

function MessageComponent() {
  switch (props.msg?.type) {
    case RESP_TYPE_DICE:
      return <MessageDice msg={props.msg}/>
    case RESP_TYPE_NEW_RICH_MESSAGE:
      return <MessageRich msg={props.msg}/>
    default:
      throw new Error('<Message> 组件表示，我不认识这个消息类型: ' + props.msg?.type)
  }
}
</script>

<style lang="scss">
.message__container {
  margin: 10px 0;

  .message__sender {
    display: inline-block;
    margin-bottom: 5px;

    font-size: 0.875rem;
    color: #50626d;

    .message__name {
      font-size: 1.4rem;
      font-weight: bold;
      color: #3f577a;
    }
  }
}
</style>
