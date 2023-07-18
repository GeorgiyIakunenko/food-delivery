<script setup>
  const props = defineProps({
    leftIcon : Object,
  })
</script>

<template>
  <button> <component :is="props.leftIcon"  /><slot></slot></button>
</template>

<style scoped>
  button {
    background-color: #f1f1f1;
    border: none;
    color: black;
    padding: 10px 20px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
  }
</style>