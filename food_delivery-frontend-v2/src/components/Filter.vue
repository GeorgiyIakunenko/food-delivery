<script setup>
import Input from "@/components/Input.vue";
import { ref } from "vue";
import {
  PaperAirplaneIcon,
  ChevronDownIcon,
  AdjustmentsVerticalIcon,
} from "@heroicons/vue/24/outline";
import Button from "@/components/Button.vue";

const searchInput = ref("");

const isFilterOpen = ref(false);
</script>

<template>
  <div class="relative bg-neutral-0 px-4 font-sans">
    <Input
      label=""
      v-model="searchInput"
      name="search"
      placeholder="Search"
      class="mb-4"
    ></Input>

    <div class="">
      <div class="flex justify-between">
        <div class="flex items-center gap-3 text-xs">
          <PaperAirplaneIcon
            class="h-6 w-6 -rotate-12 text-neutral-500"
          ></PaperAirplaneIcon>
          <div class="div flex flex-col gap-1">
            <div class="font-medium text-primary-400">Delivery to</div>
            <div class="flex items-center gap-2">
              <div class="">Address</div>
              <ChevronDownIcon
                class="h-4 w-4 text-primary-400"
              ></ChevronDownIcon>
            </div>
          </div>
        </div>
        <div class="">
          <Button
            @click="isFilterOpen = !isFilterOpen"
            class="gap-2 py-2.5 text-sm"
            :class="{ 'font-bold': isFilterOpen }"
          >
            <AdjustmentsVerticalIcon
              class="h-6 w-6 text-neutral-500"
            ></AdjustmentsVerticalIcon>
            Filter
          </Button>
        </div>
      </div>

      <!-- tabs -->
      <div class="overflow-hidden">
        <div
          class="transition-all duration-300"
          :class="{ '-translate-y-52': !isFilterOpen }"
        >
          <div class="mt-4 h-0.5 w-full bg-card-bg"></div>
          <div
            class="flex justify-between border-b px-2 pb-4 text-sm font-medium"
          >
            <div class="border-b-2 border-b-primary-400 px-2 py-4">
              Category
            </div>
            <div class="px-0.5 py-4">Sort By</div>
            <div class="px-0.5 py-4">Price</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped></style>
