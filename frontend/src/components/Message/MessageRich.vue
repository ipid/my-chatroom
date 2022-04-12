<template>
  <div v-loading="!msg.data.acknowledged" :class="contentClasses">
    {{ msg.data.content }}
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

const contentClasses = computed(() => {
  if (props.msg.data.isSender) {
    return ['message__content', 'is-sender']
  }
  return ['message__content']
})
</script>

<style lang="scss" scoped>
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
</style>
