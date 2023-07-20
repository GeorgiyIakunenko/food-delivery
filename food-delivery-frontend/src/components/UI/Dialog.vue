<script>
  export default {
    name : 'Dialog',
  }
</script>


<script setup>
  const props = defineProps({
    show : {
      type : Boolean,
      default : false
    },
  });
</script>

<template>
  <div class="dialog" v-if="show">
    <div class="dialog__content">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>
  .dialog {
    position: fixed;
    top: 0;
    left: 0;
    z-index: 100;
    width: 100%;
    height: 100%;
    background-color: rgba(0,0,0,0.5);
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .dialog__content {
    margin: auto;
    background-color: white;
    border-radius: 12px;
    min-height: 200px;
    min-width: 300px;
  }
</style>