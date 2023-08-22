<script setup>
import { onMounted, ref } from "vue";
import { getOrders } from "@/api/order";
import OrderCard from "@/components/OrderCard.vue";
import Button from "@/components/Button.vue";

const orders = ref([]);

const getOrdersFunc = async () => {
  const res = await getOrders();
  if (res.success === true) {
    orders.value = res.data;
  }
};

onMounted(async () => {
  await getOrdersFunc();
});
</script>

<template>
  <main>
    <div class="">
      <h1 class="text-center font-sans text-xl font-bold">Your orders:</h1>
      <div class="container">
        <div
          v-if="orders.value.length === 0"
          class="mt-14 flex flex-col items-center gap-7 text-center font-sans text-xl font-bold"
        >
          You don't have orders yet!
          <router-link to="/cart" exact
            ><Button type="primary">Go to cart</Button></router-link
          >
        </div>
        <div
          v-else
          class="mx-auto mt-10 grid max-w-fit grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3 lg:gap-7"
        >
          <OrderCard
            @order-cancelled="getOrdersFunc"
            v-for="orderItem in orders"
            :order="orderItem"
          ></OrderCard>
        </div>
        <Button
          class="mx-auto mt-10 w-3/5"
          type="secondary"
          @click="$router.push('/profile')"
          >Back</Button
        >
      </div>
    </div>
  </main>
</template>

<style scoped></style>
