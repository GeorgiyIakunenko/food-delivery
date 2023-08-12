<script setup>
import Input from "@/components/Input.vue";
import { ref } from "vue";
import { useFilterStore } from "@/store/filter";

const min = 0;
const max = 5000;

const priceGap = 500;

const filterStore = useFilterStore();

let barStyle = `left: ${(filterStore.filter.minPrice / max) * 100}%; right: ${
  100 - (filterStore.filter.maxPrice / max) * 100
}%;`;
const setBarStyles = (event) => {
  if (filterStore.filter.maxPrice - filterStore.filter.minPrice < priceGap) {
    if (event.target.name === "minInput") {
      filterStore.filter.minPrice = +filterStore.filter.maxPrice - priceGap;
    }
    if (event.target.name === "maxInput") {
      filterStore.filter.maxPrice = +filterStore.filter.minPrice + 1000;
    }
  } else {
    // setting the bar styles
    const right = 100 - (filterStore.filter.maxPrice / max) * 100;
    const left = (filterStore.filter.minPrice / max) * 100;

    barStyle = `left: ${left}%; right: ${right}%;`;
  }
};
</script>

<template>
  <div class="mt-5 px-5 font-bold">Products price:</div>
  <div class="mt-5 rounded-xl bg-card-bg px-5 py-4">
    <div class="flex justify-between">
      <div class="w-32 px-5">HUF {{ filterStore.filter.minPrice }}</div>
      <div>-</div>
      <div class="w-32 px-5">HUF {{ filterStore.filter.maxPrice }}</div>
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
          v-model="filterStore.filter.minPrice"
          class="absolute -top-3 w-full"
          name="minInput"
          label="minInput"
          type="range"
          :min="min"
          :max="max"
        />
        <input
          @input="(event) => setBarStyles(event)"
          v-model="filterStore.filter.maxPrice"
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
