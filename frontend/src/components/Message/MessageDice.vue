<template>
  <div :class="contentClass">
    <div v-for="(dice, i) in msg.data.dice.dices" :key="i" class="message-dice__dice">
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
  margin-bottom: -8px;

  /* A dice is a square block with black background, and there is a number of white color inside */
  .message-dice__dice {
    display: inline-block;
    width: 65px;
    height: 65px;
    border-radius: 10px;
    background-color: #975c2a;
    color: #fff;
    font-size: 40px;
    line-height: 65px;
    text-align: center;
    vertical-align: middle;
    overflow: hidden;
    margin-bottom: 8px;
  }

  &.message-dice--sender {
    margin-left: -10px;

    .message-dice__dice {
      background-color: #34577a;
      margin-left: 10px;
    }
  }

  &.message-dice--receiver {
    margin-right: -5px;

    .message-dice__dice {
      margin-right: 5px;
    }
  }
}
</style>
