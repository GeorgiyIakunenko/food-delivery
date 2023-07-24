<script>
  export default {
    name : 'Button',
  }
</script>

<script setup>
import { computed } from 'vue';

const props = defineProps({
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
    <slot></slot>
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
  transition: all 0.3s ease-in-out;
}


.primary {
  background-color: #4CAF50;
  color: #fff;
}

.primary:hover {
  background-color: rgba(0, 0, 255, 0.5);
}


.secondary {
  background-color: #FF5722;
  color: #fff;
}

.secondary:hover {
  background-color: #E64A19;
}

.danger {
  background-color: red;
  color: #fff;
}

.danger:hover {
  background-color: #CC0000;
}

.text {
  background-color: rgba(92,90,87,0.1);
  color: #21201F;
}

.text:hover {
  background-color: rgba(92, 90, 87, 0.3);
}

</style>
