<script setup>
import Button from "@/components/Button.vue";
import Modal from "@/components/Modal.vue";
import { ref } from "vue";
import { isCancel } from "axios";
import { cancelOrder } from "@/api/order";

const props = defineProps({
  order: {
    type: Object,
    required: true,
  },
});

let modalTitle = "Cancel order?";
let modalMessage = "After cancelling you can't change your order products";
let modalType = "warning";
let modalOpen = ref(false);

const openModalWarning = () => {
  modalTitle = "Cancel order?";
  modalMessage = "You won't be able to change mind after cancelling";
  modalType = "warning";
  modalOpen.value = true;
};

const emits = defineEmits(["orderCancelled"]);

const submitCancelOrder = async (isCancel) => {
  if (isCancel === false || modalType === "success") {
    modalOpen.value = false;
    return;
  }
  modalOpen.value = false;
  const res = await cancelOrder(props.order.id);
  if (res.success === true) {
    modalTitle = "Order cancelled";
    modalMessage = "";
    modalType = "success";
    modalOpen.value = true;
    emits("orderCancelled");
  } else {
    modalTitle = "Something went wrong";
    modalMessage = "Please try again later";
    modalType = "error";
    modalOpen.value = true;
  }
};

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
        <div
          :class="{
            'text-red-500': order.order_status.toLowerCase() === 'cancelled',
            'text-green-500': order.order_status.toLowerCase() === 'pending',
          }"
        >
          {{ order.order_status.toUpperCase() }}
        </div>
        <Button
          v-if="order.order_status.toLowerCase() !== 'cancelled'"
          @click="openModalWarning"
          class="h-12 w-1/2"
          type="primary"
        >
          Cancel
        </Button>
      </div>
    </div>
  </div>
  <Modal
    :type="modalType"
    :title="modalTitle"
    @modalClose="submitCancelOrder"
    v-bind:open="modalOpen"
    >{{ modalMessage }}
    <router-link v-if="modalType === 'success'" to="/profile/orders">
    </router-link>
  </Modal>
</template>

<style scoped></style>
