<script setup>
import router from "@/router/router";
import { onMounted, ref } from "vue";
import { getOrders } from "@/api/order";

const orders = ref([]);

onMounted(async () => {
  const res = await getOrders();
  if (res.success === true) {
    orders.value = res.data;
  }
  console.log(orders.value);
});
</script>

<template>
  <main>
    <div class="">
      Orders
      <div class="container">
        <div
          class="mx-auto mt-7 grid max-w-fit grid-cols-2 gap-4 md:grid-cols-3 md:gap-7 xl:grid-cols-4"
        >
          <div v-for="order in orders" class="">{{ order.total_price }}</div>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
