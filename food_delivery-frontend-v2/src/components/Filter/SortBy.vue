<script setup>
import {
  BanknotesIcon,
  CheckCircleIcon,
  StarIcon,
  ClockIcon,
} from "@heroicons/vue/24/outline";
import Button from "@/components/Button.vue";
import { useFilterStore } from "@/store/filter";

const SortDirection = useFilterStore().filter.sortDirection;

// todo 1: fix the desc and asc buttons
</script>

<template>
  <div class="flex flex-col items-start justify-center gap-4 py-6">
    <Button
      @click="SortDirection = 'asc'"
      class="mx-auto w-11/12 justify-between gap-2 py-2.5 text-sm"
      :class="{
        'bg-green-300': SortDirection === 'asc',
        'bg-neutral-0': SortDirection === 'desc',
      }"
    >
      <div class="flex items-center gap-2">
        <BanknotesIcon class="h-6 w-6 text-neutral-500"></BanknotesIcon>
        By price (Low to High)
      </div>
      <CheckCircleIcon class="h-6 w-6 text-neutral-500"></CheckCircleIcon>
    </Button>
    <Button
      @click="SortDirection = 'desc'"
      class="mx-auto w-11/12 justify-between gap-2 py-2.5 text-sm"
      :class="{
        'bg-green-300': SortDirection === 'desc',
        'bg-neutral-0': SortDirection === 'asc',
      }"
    >
      <div class="flex items-center gap-2">
        <BanknotesIcon class="h-6 w-6 text-neutral-500"></BanknotesIcon>
        By price (High to Low)
      </div>
      <CheckCircleIcon class="h-6 w-6 text-neutral-500"></CheckCircleIcon>
    </Button>
    <div class="mx-auto">Working Time :</div>
    <Button class="mx-auto w-11/12 justify-between gap-2 py-2.5 text-sm">
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
