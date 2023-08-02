<script setup>
import { defineProps, defineEmits } from "vue";
</script>

<template>
  <div class="flex flex-col items-start py-6">
    <div class="text-center">
      <div
        class="flex h-24 w-24 cursor-pointer items-center justify-center rounded-full bg-primary-50 transition-all duration-300 hover:bg-primary-75"
      >
        <img src="@/images/Pizza.svg" alt="Pizza" />
      </div>
      <div class="mt-3 text-sm font-medium">Pizza</div>
    </div>
  </div>
</template>

<style scoped></style>
