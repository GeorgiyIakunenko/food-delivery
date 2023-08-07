<script setup>
import Input from "@/components/Input.vue";
import { onMounted, ref } from "vue";
import {
  PaperAirplaneIcon,
  ChevronDownIcon,
  AdjustmentsVerticalIcon,
} from "@heroicons/vue/24/outline";
import Button from "@/components/Button.vue";
import Categories from "@/components/Filter/Categories.vue";
import SortBy from "@/components/Filter/SortBy.vue";
import DoublePriceSlider from "@/components/Filter/DoublePriceSlider.vue";
import { getAllCategories } from "@/api/category";

const searchInput = ref("");

const activeFilterType = ref("category");

const setActiveFilterType = (type) => {
  activeFilterType.value = type;
};

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
          <!--          <div class="div flex flex-col gap-1">
            <div class="font-medium text-primary-400">Delivery to</div>
            <div class="flex items-center gap-2">
              <div class="">Address</div>
              <ChevronDownIcon
                class="h-4 w-4 text-primary-400"
              ></ChevronDownIcon>
            </div>
          </div>-->
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
      <div class="mt-4 h-0.5 w-full bg-card-bg"></div>
      <!-- tabs -->
      <div class="overflow-hidden">
        <div
          class="flex flex-col transition-all duration-300"
          :class="{
            'h-0 -translate-y-104': !isFilterOpen,
            'h-100': isFilterOpen,
          }"
        >
          <div
            class="flex justify-between border-b px-2 pb-4 text-sm font-medium"
          >
            <div
              @click="setActiveFilterType('category')"
              class="cursor-pointer px-2 py-4"
              :class="{
                'border-b-2 border-b-primary-400 font-bold':
                  activeFilterType === 'category',
              }"
            >
              Category
            </div>
            <div
              @click="setActiveFilterType('sort-by')"
              :class="{
                'border-b-2 border-b-primary-400 font-bold':
                  activeFilterType === 'sort-by',
              }"
              class="cursor-pointer px-0.5 py-4"
            >
              Sort By
            </div>
            <div
              @click="setActiveFilterType('price')"
              class="cursor-pointer px-0.5 py-4"
              :class="{
                'border-b-2 border-b-primary-400 font-bold':
                  activeFilterType === 'price',
              }"
            >
              Price
            </div>
          </div>
          <div class="">
            <Categories v-if="activeFilterType === 'category'"></Categories>
            <SortBy v-if="activeFilterType === 'sort-by'"></SortBy>
            <DoublePriceSlider
              v-if="activeFilterType === 'price'"
            ></DoublePriceSlider>
          </div>
          <div class="mt-auto flex justify-center">
            <Button
              @click="isFilterOpen = !isFilterOpen"
              class="w-4/5"
              type="primary"
              >Complete</Button
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped></style>
