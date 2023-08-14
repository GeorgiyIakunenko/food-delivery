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
      class="mb-2 h-36 w-36 md:h-52 md:w-52"
      :src="getImageUrl(category.image)"
      :alt="category.name"
    />
    <div class="px-1">
      <div class="flex items-center justify-center gap-3">
        <div class="mb-1 font-sans text-lg font-bold">
          {{ category.name }}
        </div>
      </div>
      <div
        class="flex w-40 items-center text-center text-sm md:w-48 md:text-base"
      >
        {{ category.description.slice(0, 60) }}...
      </div>
    </div>
  </div>
</template>

<style scoped></style>
