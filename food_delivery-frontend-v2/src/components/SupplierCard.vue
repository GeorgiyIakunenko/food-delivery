<script setup>
import { getImageUrl } from "@/utils/url";

const props = defineProps({
  supplier: {
    type: Object,
    required: true,
  },
});
</script>
<template>
  card
  <img :src="getImageUrl(supplier.image)" alt="" />
</template>

<style scoped></style>
