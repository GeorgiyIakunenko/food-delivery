<script setup>
import Button from "@/components/Button.vue";
import { BanknotesIcon } from "@heroicons/vue/24/outline";
</script>

<template>
  <main class="flex flex-col bg-card-bg">
    <div class="container">
      <div class="flex flex-col gap-5">
        <div class="flex flex-col gap-3">
          <div
            class="rounded-xl bg-neutral-0 p-4 font-sans text-neutral-800 md:w-fit"
          >
            <div class="pb-3 text-lg font-bold">Delivery to:</div>
            <div class="flex gap-4">
              <img src="@/assets/images/map.svg" alt="map" />
              <div class="font-medium">Address</div>
            </div>
          </div>
          <div
            class="rounded-xl bg-neutral-0 p-4 font-sans font-medium text-neutral-800 md:w-fit"
          >
            <div class="pb-3 text-lg font-bold">Order products:</div>
            <div class="flex gap-4"></div>
            <div
              class="flex justify-between border-b-2 border-b-card-bg px-2 py-4"
            >
              <div class="">Subtotal:</div>
              <div class="">$ 36.66</div>
            </div>
            <div
              class="flex justify-between border-b-2 border-b-card-bg px-2 py-4"
            >
              <div class="">Delivery</div>
              <div class="">$ 0.00</div>
            </div>
            <div class="flex justify-between px-2 pt-4 text-lg font-bold">
              <div class="">Total:</div>
              <div class="font-medium text-primary-500">$ 36.66</div>
            </div>
          </div>
        </div>
        <div
          class="mt-auto flex flex-col gap-3 rounded-xl bg-neutral-0 p-4 font-sans font-medium text-neutral-800 md:w-fit"
        >
          <div class="">
            <div class="">
              <div
                class="flex w-fit cursor-pointer items-center gap-2 rounded-xl border-2 bg-red-50 p-4 font-sans text-neutral-800"
              >
                <BanknotesIcon class="h-8 w-8 text-green-400"></BanknotesIcon>
                <div class="flex flex-col items-center">
                  <div class="font-medium">$ 35.60</div>
                  <div class="text-sm font-medium text-neutral-200">Cash</div>
                </div>
              </div>
            </div>
          </div>
          <Button type="primary" class="mx-auto w-full">Submit</Button>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
