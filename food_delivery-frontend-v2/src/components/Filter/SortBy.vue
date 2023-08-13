<script setup>
import {
  BanknotesIcon,
  CheckCircleIcon,
  StarIcon,
  ClockIcon,
} from "@heroicons/vue/24/outline";
import Button from "@/components/Button.vue";
import { useFilterStore } from "@/store/filter";
import colors from "tailwindcss/colors";

const isSortActive = (sort) => {
  return useFilterStore().filter.sortDirection === sort;
};

const toggleSort = (sort) => {
  if (isSortActive(sort)) {
    useFilterStore().filter.sortDirection = "";
  } else {
    useFilterStore().filter.sortDirection = sort;
  }
};

// todo 1: fix the desc and asc buttons
</script>

<template>
  <div class="flex flex-col items-start justify-center gap-4 py-6">
    <Button
      @click="toggleSort('asc')"
      class="mx-auto w-11/12 justify-between gap-2 py-2.5 text-sm"
      :style="{
        backgroundColor: isSortActive('asc')
          ? colors.green[400]
          : isSortActive('desc')
          ? colors.neutral[0]
          : null,
      }"
    >
      <div class="flex items-center gap-2">
        <BanknotesIcon class="h-6 w-6 text-neutral-500"></BanknotesIcon>
        By price (Low to High)
      </div>
      <CheckCircleIcon class="h-6 w-6 text-neutral-500"></CheckCircleIcon>
    </Button>
    <Button
      @click="toggleSort('desc')"
      class="mx-auto w-11/12 justify-between gap-2 py-2.5 text-sm"
      :style="{
        backgroundColor: isSortActive('desc')
          ? colors.green[400]
          : isSortActive('asc')
          ? colors.neutral[0]
          : null,
      }"
    >
      <div class="flex items-center gap-2">
        <BanknotesIcon class="h-6 w-6 text-neutral-500"></BanknotesIcon>
        By price (High to Low)
      </div>
      <CheckCircleIcon class="h-6 w-6 text-neutral-500"></CheckCircleIcon>
    </Button>
    <div class="mx-auto">Working Time :</div>
    <Button
      @click="
        useFilterStore().filter.openNow = !useFilterStore().filter.openNow
      "
      class="mx-auto w-11/12 justify-between gap-2 py-2.5 text-sm"
      :style="{
        backgroundColor: useFilterStore().filter.openNow
          ? colors.green[400]
          : colors.neutral[0],
      }"
    >
      <div class="flex items-center gap-2">
        <ClockIcon class="h-6 w-6 text-neutral-500"></ClockIcon>
        Open Now
      </div>
      <CheckCircleIcon class="h-6 w-6 text-neutral-500"></CheckCircleIcon>
    </Button>
  </div>
</template>

<style scoped></style>

<style scoped></style>
