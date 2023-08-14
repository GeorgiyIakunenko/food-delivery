<script setup>
import { getImageUrl } from "@/utils/url";

const props = defineProps({
  category: {
    type: Object,
    required: true,
  },
});

console.log(props);
</script>
<template>
  <div
    class="flex cursor-pointer flex-col items-center justify-center rounded-xl bg-amber-50 p-3 shadow-md transition-all duration-300 hover:scale-105"
  >
    <img
      class="mb-2 w-36 md:w-52"
      :src="getImageUrl(category.image)"
      :alt="category.name"
    />
    <div class="px-1">
      <div class="flex items-center justify-center gap-3">
        <div class="mb-1 font-sans text-lg font-medium">
          {{ category.name }}
        </div>
      </div>
      <div class="flex w-48 items-center text-center text-sm md:text-base">
        {{ category.description.slice(0, 50) }}...
      </div>
    </div>
  </div>
</template>

<style scoped></style>
