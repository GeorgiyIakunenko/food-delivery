<script setup>
import { getImageUrl } from "@/utils/url";
import { useCartStore } from "@/store/cart";

const props = defineProps({
  product: {
    type: Object,
    required: true,
  },
});
</script>

<template>
  <div
    class="flex w-full items-center gap-6 rounded-xl border px-3 py-2 md:w-80"
  >
    <img
      class="h-20 w-20 rounded-lg object-contain"
      :src="getImageUrl(product.image)"
      :alt="product.name"
    />
    <div class="flex flex-col items-start gap-2">
      <div class="">{{ product.name }}</div>
      <div class="flex items-center gap-5">
        <div class="flex items-center gap-3 rounded-xl bg-card-bg px-2 py-1">
          <svg
            class="cursor-pointer"
            @click="useCartStore().decreaseProduct(product)"
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 14 14"
            fill="none"
          >
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M7 0C3.13401 0 0 3.13401 0 7C0 10.866 3.13401 14 7 14C10.866 14 14 10.866 14 7C14 5.14348 13.2625 3.36301 11.9497 2.05025C10.637 0.737498 8.85652 0 7 0ZM9.757 8.023H4.242C3.70612 8.00599 3.28051 7.56665 3.28051 7.0305C3.28051 6.49435 3.70612 6.05501 4.242 6.038H9.757C10.1191 6.0265 10.4587 6.21312 10.6432 6.52493C10.8276 6.83674 10.8276 7.22426 10.6432 7.53607C10.4587 7.84788 10.1191 8.0345 9.757 8.023Z"
              fill="#C1C7D0"
            />
          </svg>
          <div class="w-5 text-center">
            <div class="">{{ product.quantity }}</div>
          </div>
          <svg
            @click="useCartStore().addProduct(product)"
            class="cursor-pointer"
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 14 14"
            fill="none"
          >
            <path
              fill-rule="evenodd"
              clip-rule="evenodd"
              d="M7 0C3.13401 0 0 3.13401 0 7C0 10.866 3.13401 14 7 14C10.866 14 14 10.866 14 7C14 5.14348 13.2625 3.36301 11.9497 2.05025C10.637 0.737498 8.85652 0 7 0ZM9.758 8.023H7.992V9.788C7.992 10.3359 7.54787 10.78 7 10.78C6.45213 10.78 6.008 10.3359 6.008 9.788V8.023H4.242C3.70612 8.00599 3.28051 7.56665 3.28051 7.0305C3.28051 6.49435 3.70612 6.05501 4.242 6.038H6.008V4.273C6.008 3.72513 6.45213 3.281 7 3.281C7.54787 3.281 7.992 3.72513 7.992 4.273V6.038H9.758C10.2939 6.05501 10.7195 6.49435 10.7195 7.0305C10.7195 7.56665 10.2939 8.00599 9.758 8.023Z"
              fill="#EF9F27"
            />
          </svg>
        </div>
        <div class="w-20 text-sm text-primary-500">
          HUF {{ product.price * product.quantity }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
