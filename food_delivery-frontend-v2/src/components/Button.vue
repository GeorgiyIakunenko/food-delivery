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
        return 'bg-primary text-white'
      case 'secondary':
        return 'bg-secondary text-white'
      default:
        return 'px-2.5 py-1.5 font-sans font-medium rounded-xl border-0 bg-neutral-1 '
    }
  })

</script>
<template>
  <button class="px-3.5 py-1.5 font-sans font-medium rounded-xl border-0 bg-neutral-30 hover:bg-neutral-40 active:transform-cpu transition duration-300 ease-in-out">
    <slot></slot>
  </button>
</template>

<style scoped>

</style>