<script setup>
import Button from "@/components/Button.vue";
import { BanknotesIcon } from "@heroicons/vue/24/outline";
import { useCartStore } from "@/store/cart";
import CartProductCard from "@/components/CartProductCard.vue";
import { useUserStore } from "@/store/user";
import { computed, onMounted, ref } from "vue";
import { getProfile } from "@/api/user";
import { createOrder } from "@/api/order";
import Modal from "@/components/Modal.vue";

const cartStore = useCartStore();
const userStore = useUserStore();

const deliveryCost = 250;

const totalCost = computed(() => {
  return cartStore.cart.total_price + deliveryCost;
});

let modalTitle = "Submit order? ";
let modalMessage = "After submitting you can't change your order products";
let modalType = "warning";
let modalOpen = ref(false);

const openModalWarning = () => {
  modalTitle = "Submit order?";
  modalMessage = "After submitting you can't change your order products";
  modalType = "warning";
  modalOpen.value = true;
};

const submitOrder = async (isSubmit) => {
  if (isSubmit === false) {
    modalOpen.value = false;
    return;
  }

  if (modalType === "success") {
    modalOpen.value = false;
    return;
  }

  const order = {
    total_price: totalCost.value,
    payment_method: cartStore.cart.payment_method,
    address: userAddress,
    product: cartStore.cart.products.map((product) => {
      return {
        product_id: product.id,
        quantity: product.quantity,
      };
    }),
  };

  const res = await createOrder(order);
  console.log(res);
  if (res.success === true) {
    modalTitle = "Order submitted";
    modalMessage = "Your can check your order in your orders page";
    modalType = "success";
    modalOpen.value = true;
    cartStore.clearCart();
  }
};

onMounted(async () => {
  if (userStore.accessToken) {
    await getProfile();
  }
});

const userAddress = userStore.user.Address;
</script>

<template>
  <main class="flex flex-col bg-card-bg">
    <div class="container">
      <div class="flex flex-col gap-5">
        <div class="flex flex-col gap-3">
          <div
            v-if="cartStore.cart.products.length !== 0 && userStore.accessToken"
            class="rounded-xl bg-neutral-0 p-4 font-sans text-neutral-800"
          >
            <div class="pb-3 text-lg font-bold">Delivery to:</div>
            <div class="flex gap-4">
              <img src="@/assets/images/map.svg" alt="map" />
              <input v-model="userAddress" class="h-fit p-2 font-medium" />
            </div>
          </div>
          <div
            class="rounded-xl bg-neutral-0 p-4 font-sans font-medium text-neutral-800"
          >
            <div
              v-if="cartStore.cart.products.length !== 0"
              class="pb-3 text-lg font-bold"
            >
              Cart products:
            </div>
            <div class="my-7 flex flex-wrap gap-5">
              <div
                v-if="cartStore.cart.products.length === 0"
                class="mx-auto flex flex-col items-center justify-center gap-5"
              >
                <div class="mx-auto text-center text-xl font-bold">
                  Your cart is empty
                </div>
                <router-link to="/"
                  ><Button type="primary">Go To Products</Button></router-link
                >
              </div>
              <CartProductCard
                v-else
                v-for="product in cartStore.cart.products"
                :product="product"
              />
            </div>

            <div class="flex gap-4"></div>
            <div v-if="cartStore.cart.products.length !== 0" class="">
              <div
                class="flex justify-between border-b-2 border-b-card-bg px-2 py-4"
              >
                <div class="">Subtotal:</div>
                <div class="">HUF {{ cartStore.cart.total_price }}</div>
              </div>
              <div
                class="flex justify-between border-b-2 border-b-card-bg px-2 py-4"
              >
                <div class="">Delivery</div>
                <div class="">HUF {{ deliveryCost }}</div>
              </div>
              <div class="flex justify-between px-2 pt-4 text-lg font-bold">
                <div class="">Total:</div>
                <div class="font-medium text-primary-500">
                  HUF {{ totalCost }}
                </div>
              </div>
            </div>
          </div>
        </div>
        <div
          v-if="cartStore.cart.products.length !== 0"
          class="mt-auto flex flex-col gap-3 rounded-xl bg-neutral-0 p-4 font-sans font-medium text-neutral-800"
        >
          <div class="">
            <div class="">
              <div
                class="flex w-fit cursor-pointer items-center gap-2 rounded-xl border-2 bg-red-50 p-4 font-sans text-neutral-800"
              >
                <BanknotesIcon class="h-8 w-8 text-green-400"></BanknotesIcon>
                <div class="flex flex-col items-center">
                  <div class="font-medium">HUF {{ totalCost }}</div>
                  <div class="text-sm font-medium text-neutral-200">Cash</div>
                </div>
              </div>
            </div>
            <div class="ml-2 mt-2 text-xs text-red-500">
              Only cash is available right now
            </div>
          </div>

          <Button
            @click="openModalWarning"
            v-if="userStore.accessToken"
            type="primary"
            class="mx-auto w-full"
            >Submit</Button
          >
          <router-link v-else to="/login" exact
            ><Button type="primary" class="mx-auto w-full"
              >Login to continue</Button
            ></router-link
          >
        </div>
      </div>
    </div>
    <Modal
      :type="modalType"
      :title="modalTitle"
      @modalClose="submitOrder"
      v-bind:open="modalOpen"
      >{{ modalMessage }}
      <router-link v-if="modalType === 'success'" to="/profile/orders">
        <Button class="mb-1 mt-3" type="primary">Go to orders</Button>
      </router-link>
    </Modal>
  </main>
</template>

<style scoped></style>
