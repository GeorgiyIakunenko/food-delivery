<script setup>
import { getImageUrl } from "@/utils/url";

const props = defineProps({
  supplier: {
    type: Object,
    required: true,
  },
});

const isOpen = () => {
  const currentTime = new Date();
  const currentTimeFormated = `${currentTime.getHours()}:${currentTime.getMinutes()}`;
  return (
    currentTimeFormated >= props.supplier.open_time &&
    currentTimeFormated <= props.supplier.close_time
  );
};
</script>
<template>
  <div
    class="cursor-pointer rounded-xl bg-neutral-0 p-3 shadow-md transition-all duration-300 hover:scale-105"
  >
    <img
      class="mb-2 w-36 md:w-52"
      :src="getImageUrl(supplier.image)"
      :alt="supplier.name"
    />
    <div class="px-1">
      <div class="flex items-center gap-3">
        <div class="font-sans text-lg font-medium">
          {{ supplier.name }}
        </div>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="14"
          height="14"
          viewBox="0 0 14 14"
          fill="none"
        >
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M7.373 0.0669637L12.758 2.05396C13.1825 2.20939 13.4651 2.61293 13.466 3.06496V6.04696C13.5712 9.86911 10.9191 13.2159 7.174 13.987C7.06144 14.0056 6.94656 14.0056 6.834 13.987C3.08895 13.2159 0.436783 9.86911 0.542 6.04696V3.06496C0.541686 2.61497 0.820676 2.21203 1.242 2.05396L6.627 0.0669637C6.86771 -0.0219957 7.13229 -0.0219957 7.373 0.0669637ZM6.811 8.84396L9.824 5.56896L9.832 5.57296C9.99724 5.36646 10.0398 5.0875 9.94354 4.84115C9.84732 4.5948 9.62699 4.4185 9.36554 4.37865C9.10408 4.3388 8.84124 4.44146 8.676 4.64796L6.176 7.27696L5.276 6.37696C4.98688 6.08784 4.51812 6.08784 4.229 6.37696C3.93988 6.66608 3.93988 7.13484 4.229 7.42396L5.71 8.89996C5.84818 9.03987 6.03636 9.11903 6.233 9.11996H6.274C6.48423 9.109 6.67971 9.00853 6.811 8.84396Z"
            fill="#00875A"
          />
        </svg>
      </div>
      <div class="flex items-center gap-3">
        <div v-if="isOpen" class="font-sans font-medium text-green-400">
          Open
        </div>
        <div v-else class="font-sans font-medium text-red-400">Closed</div>
        <div class="text-neutral-200">{{ supplier.type }}</div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
