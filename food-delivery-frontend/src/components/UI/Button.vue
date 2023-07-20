<script>
  export default {
    name : 'Button',
  }
</script>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  leftIconUrl: String,
  leftIconAlt: String,
  rightIconUrl: String,
  rightIconAlt: String,
  intent: {
    type: String,
    validator: (value) => {
      return ['primary', 'secondary', 'danger', 'text'].includes(value);
    },
    default: 'secondary',
  },
});

const buttonClass = computed(() => {
  return {
    'primary': props.intent === 'primary',
    'secondary': props.intent === 'secondary',
    'danger': props.intent === 'danger',
    'text': props.intent === 'text',
  };
});

//const name = 'Button';

</script>
<template>
  <button :class="['button', buttonClass]">
    <img v-if="props.leftIconUrl" :src="props.leftIconUrl" :alt="leftIconAlt" class="left-icon" />
    <slot></slot>
    <img v-if="props.rightIconUrl" :src="props.rightIconUrl" :alt="rightIconAlt" class="right-icon" />
  </button>
</template>

<style scoped>
.button {
  background-color: transparent;
  border-radius: 10px;
  cursor: pointer;
  display: flex;
  border: none;
  padding: 5px 20px;
  text-align: center;
  text-decoration: none;
  font-size: 16px;
}

.primary {
  background-color: var(--primary-color);
  color : #fff;
}

.secondary {
  background-color: orange;
  color: black;
}

.danger {
  background-color: red;
  color: #fff;
}

.text {
  background-color: rgba(92,90,87,0.1);
  color: #21201F;
}

.left-icon {
  margin-right: 10px;
}

.right-icon {
  margin-left: 10px;
}
</style>
