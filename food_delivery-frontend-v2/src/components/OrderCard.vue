<script setup>
import Button from "@/components/Button.vue";

const props = defineProps({
  order: {
    type: Object,
    required: true,
  },
});

const date = new Date(props.order.created_at);

const formattedDateTime = date.toLocaleString("en-US", {
  year: "numeric",
  month: "long",
  day: "numeric",
  hour: "2-digit",
  minute: "2-digit",
});
</script>

<template>
  <div
    class="flex gap-2 rounded-xl bg-card-bg px-5 py-3 font-sans shadow-md transition-all duration-300 hover:scale-105 hover:transform"
  >
    <div class="flex flex-col gap-1 font-medium">
      <div class="flex flex-col gap-3">
        <div class="text-lg">Order #{{ order.id }}</div>
        <div class="">
          <span class="font-bold">Order time: </span> {{ formattedDateTime }}
        </div>
        <div class="">
          <span class="font-bold">Order address: </span> {{ order.address }}
        </div>
        <div class="">
          <span class="font-bold">Payment method: </span>
          {{ order.payment_method }}
        </div>
        <div class="">
          <span class="font-bold">Total price:</span> HUF
          {{ order.total_price }}
        </div>
      </div>

      <div class="mt-4 flex h-10 items-center justify-between">
        <div class="">
          {{ order.order_status }}
        </div>
        <Button class="h-12 w-1/2" type="primary"> Cancel </Button>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
