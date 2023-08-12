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
    class="cursor-pointer rounded-xl bg-neutral-0 p-3 shadow-md transition-all duration-300 hover:scale-105"
  >
    <img
      class="mb-2 w-36 md:w-52"
      :src="getImageUrl(category.image)"
      :alt="category.name"
    />
    <div class="px-1">
      <div class="flex items-center gap-3">
        <div class="font-sans text-lg font-medium">
          {{ category.name }}
        </div>
      </div>
      <div class="flex items-center gap-3">
        {{ category.description }}
      </div>
    </div>
  </div>
</template>

<style scoped></style>
