<script setup>
import {
  BanknotesIcon,
  CheckCircleIcon,
  StarIcon,
  ClockIcon,
} from "@heroicons/vue/24/outline";
import Button from "@/components/Button.vue";
</script>

<template>
  <div class="flex flex-col items-start justify-center gap-4 py-6">
    <Button class="mx-auto w-11/12 justify-between gap-2 py-2.5 text-sm">
      <div class="flex items-center gap-2">
        <BanknotesIcon class="h-6 w-6 text-neutral-500"></BanknotesIcon>
        By price
      </div>
      <CheckCircleIcon class="h-6 w-6 text-neutral-500"></CheckCircleIcon>
    </Button>
    <Button class="mx-auto w-11/12 justify-between gap-2 py-2.5 text-sm">
      <div class="flex items-center gap-2">
        <StarIcon class="h-6 w-6 text-neutral-500"></StarIcon>
        By Rating (High to Low)
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
