<template>
  <div :class="contentClass">
    <div v-for="(dice, i) in msg.data.dices" :key="i" class="message-dice__dice">
      {{ dice }}
    </div>
  </div>
</template>

<script setup>
import {computed} from 'vue'

const props = defineProps({
  msg: {
    type: Object,
    required: true,
  },
})

const contentClass = computed(() => {
  if (props.msg.data.isSender) {
    return ['message-dice__container', 'message-dice--sender']
  }
  return ['message-dice__container', 'message-dice--receiver']
})
</script>

<style lang="scss" scoped>
.message-dice__container {
  /* A dice is a square block with black background, and there is a number of white color inside */
  .message-dice__dice {
    display: inline-block;
    width: 100px;
    height: 100px;
    border-radius: 10px;
    background-color: #975c2a;
    color: #fff;
    font-size: 50px;
    line-height: 100px;
    text-align: center;
    vertical-align: middle;
    overflow: hidden;
  }

  &.message-dice--sender {
    .message-dice__dice {
      margin-left: 10px;
    }
  }

  &.message-dice--receiver {
    .message-dice__dice {
      margin-right: 10px;
    }
  }
}
</style>
