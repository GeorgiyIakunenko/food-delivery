<script setup>
  import {computed} from "vue";

  const props = defineProps({
    type: {
      type: String,
      validator(value) {
        return ['standard', 'primary', 'secondary'].includes(value)
      },
      default: 'standard'
    }
  })

  const buttonClass = computed(() => {
    switch (props.type) {
      case 'primary':
        return 'bg-primary-400 border-neutral-0 border-2 hover:bg-blue-75 hover:text-black active:transform-cpu transition duration-300 ease-in-out text-white'
      case 'secondary':
        return 'bg-secondary-btn hover:bg-blue-75 hover:text-black active:transform-cpu transition duration-300 ease-in-out'
      default:
        return 'bg-neutral-30 hover:bg-neutral-40 active:transform-cpu transition duration-300 ease-in-out border-0'
    }
  })

</script>
<template>
  <button class="px-3.5 py-1.5 font-sans rounded-xl " :class="buttonClass">
    <slot></slot>
  </button>
</template>

<style scoped>
</style>