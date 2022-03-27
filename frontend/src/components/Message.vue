<template>
  <div :style="outerStyle" class="message__container">
    <div class="message__sender">
      <span class="message__name">{{ props.name }}</span>
      &nbsp;
      <span>{{ displayTime }}</span>
    </div>
    <br/>
    <div>
      <div :class="contentClasses">{{ props.content }}</div>
    </div>
  </div>
</template>

<script setup>
import {computed} from 'vue'
import dayjs from 'dayjs'

const props = defineProps({
  name: String,
  timestampMs: Number,
  content: String,
  isSender: {
    type: Boolean,
    default: false,
  },
})

const displayTime = computed(() => dayjs(props.timestampMs).format('HH:mm:ss'))

const contentClasses = computed(() => {
  if (props.isSender) {
    return ['message__content', 'is-sender']
  }
  return ['message__content']
})
const outerStyle = computed(() => {
  if (props.isSender) {
    return {
      'text-align': 'right',
    }
  }
  return {}
})
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

  .message__content {
    display: inline-block;
    padding: 10px 15px 10px 20px;
    background: #e8f1f9;
    border-radius: 0 16px 16px 16px;
    color: #34577a;
  }

  .message__content.is-sender {
    padding: 10px 20px 10px 20px;
    border-radius: 16px 0 16px 16px;
    background: #3e84af;
    color: white;
  }
}
</style>