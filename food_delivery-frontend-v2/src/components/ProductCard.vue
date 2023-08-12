<script setup>
import { getImageUrl } from "@/utils/url";
import { useCartStore } from "@/store/cart";
import { computed } from "vue";

const props = defineProps({
  product: {
    type: Object,
    required: true,
  },
});

const cartStore = useCartStore();

const isInCart = computed(() => {
  return cartStore.cart.products.some((p) => p.id === props.product.id);
});

const isOpen = () => {
  const currentTime = new Date();
  const currentTimeFormated = `${currentTime.getHours()}:${currentTime.getMinutes()}`;
  return (
    currentTimeFormated >= props.product.supplier.open_time &&
    currentTimeFormated <= props.product.supplier.close_time
  );
};
</script>

<template>
  <div class="font-sans">
    <div class="">
      <div class="rounded-xl bg-neutral-0 p-2 font-roboto shadow-md">
        <div
          class="my-2 -ml-2 flex w-fit items-center gap-2 rounded-br-lg rounded-tr-lg px-2 py-1 text-xs text-neutral-0"
          :class="{ 'bg-green-700': isOpen, 'bg-red-200 text-black': !isOpen }"
        >
          <img
            class="fit h-5 w-5 object-contain"
            :src="getImageUrl(product.supplier.image)"
            :alt="product.supplier.name"
          />
          <div v-if="isOpen" class="">Available Now</div>
          <div v-else class="">Open at {{ product.supplier.open_time }}</div>
        </div>
        <img
          class="fit mx-auto mb-5 h-40 w-40 object-contain md:h-56 md:w-56"
          :src="getImageUrl(product.image)"
          :alt="product.name"
        />
        <div class="mb-2 flex items-center px-1">
          <div class="font-bold">
            <div class="lg:text-lg">{{ product.name }}</div>
            <div class="text-sm text-primary-500">HUF {{ product.price }}</div>
          </div>
          <div
            @click="useCartStore().addProduct(product)"
            class="ml-auto cursor-pointer"
          >
            <img
              v-if="!isInCart"
              class="h-7 w-7"
              src="@/assets/images/plus.svg"
              alt="plus"
            />
            <img
              v-else
              class="h-7 w-7"
              src="@/assets/images/check.svg"
              alt="check"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
