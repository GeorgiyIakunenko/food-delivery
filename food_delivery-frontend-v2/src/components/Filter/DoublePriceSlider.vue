<script setup>
import Input from "@/components/Input.vue";
import { ref } from "vue";

const min = 0;
const max = 9999;

const priceGap = 1000;

const minInput = ref(500);
const maxInput = ref(4000);

let barStyle = `left: ${(minInput.value / max) * 100}%; right: ${
  100 - (maxInput.value / max) * 100
}%;`;
const setBarStyles = (event) => {
  if (maxInput.value - minInput.value < priceGap) {
    if (event.target.name === "minInput")
      minInput.value = maxInput.value - priceGap;
    else {
      maxInput.value = minInput.value + priceGap;
    }
  } else {
    // setting the bar styles
    barStyle = `left: ${(minInput.value / max) * 100}%; right: ${
      100 - (maxInput.value / max) * 100
    }%;`;
  }
};
</script>

<template>
  <div class="mt-5 px-5 font-bold">Products price:</div>
  <div class="mt-5 rounded-xl bg-card-bg px-5 py-4">
    <div class="flex justify-between">
      <div class="px-5">
        HUF
        <input
          v-model="minInput"
          class="w-14 bg-card-bg"
          disabled
          name="minInput"
          label="MinPrice"
          type="number"
        />
      </div>
      -
      <div class="px-5">
        HUF
        <input
          v-model="maxInput"
          class="-mr-2 w-14 bg-card-bg"
          name="maxInput"
          disabled
          label="MaxPrice"
          type="number"
        />
      </div>
    </div>
    <div class="relative mb-3 mt-7">
      <div class="h-2 rounded bg-neutral-40">
        <div
          :style="barStyle"
          class="absolute left-7 right-3 h-2 rounded bg-primary-300"
        ></div>
      </div>
      <div class="input-box relative">
        <input
          @input="(event) => setBarStyles(event)"
          v-model="minInput"
          class="absolute -top-3 w-full"
          name="minInput"
          label="minInput"
          type="range"
          :min="min"
          :max="max"
        />
        <input
          @input="(event) => setBarStyles(event)"
          v-model="maxInput"
          class="absolute -top-3 w-full"
          name="maxInput"
          label="maxInput"
          type="range"
          :min="min"
          :max="max"
        />
      </div>
    </div>
  </div>
</template>

<style scoped>
.input-box input {
  -webkit-appearance: none;
  pointer-events: none;
  appearance: none;
  background: none;
}

.input-box input::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 1rem;
  height: 1rem;
  background: #ef9f27;
  border-radius: 50%;
  cursor: pointer;
  pointer-events: auto;
}
</style>
