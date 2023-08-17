<script setup>
import Input from "@/components/Input.vue";
import { computed, onMounted, ref } from "vue";
import {
  PaperAirplaneIcon,
  CheckCircleIcon,
  AdjustmentsVerticalIcon,
  XCircleIcon,
} from "@heroicons/vue/24/outline";
import Button from "@/components/Button.vue";
import Categories from "@/components/Filter/Categories.vue";
import SortBy from "@/components/Filter/SortBy.vue";
import DoublePriceSlider from "@/components/Filter/DoublePriceSlider.vue";
import { getAllCategories } from "@/api/category";
import { useFilterStore } from "@/store/filter";
import { getImageUrl } from "@/utils/url";

const searchInput = ref("");
const filterStore = useFilterStore();

const activeFilterType = ref("category");

const setActiveFilterType = (type) => {
  activeFilterType.value = type;
};

const isFilterOpen = ref(false);

const emits = defineEmits(["filterChange"]);

const submitFilter = () => {
  console.log("filter has been submitted");
  isFilterOpen.value = false;
  emits("filterChange");
};

const categoriesItems = ref([]);

const filterCategories = computed({
  get: () => filterStore.filter.categories,
  set: (value) => {
    filterStore.filter.categories = value;
  },
});

const isCategoryActive = (id) => {
  return filterCategories.value.includes(id);
};

onMounted(async () => {
  const res = await getAllCategories();
  if (res.success === true) {
    categoriesItems.value = res.data;
  }
});
</script>

<template>
  <div class="relative bg-neutral-0 px-4 font-sans">
    <Input
      label=""
      v-model="filterStore.filter.search"
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
          <div class="flex items-center gap-2">
            <Button
              v-if="filterStore.isFiltered()"
              @click="filterStore.resetFilter()"
              class="items-center gap-2 bg-neutral-50"
              >Reset all filters
              <XCircleIcon class="h-5 w-5 text-red-500"></XCircleIcon>
            </Button>
            <div class="hidden flex-wrap items-center gap-2 lg:flex">
              <Button
                @click="filterStore.resetCategoryFilter()"
                v-if="filterStore.isCategoryFiltered()"
                class="gap-2 bg-neutral-50"
                >Categories:
                <div class="flex gap-1">
                  <img
                    v-for="category in categoriesItems"
                    class="h-5 w-5"
                    :class="{ hidden: !isCategoryActive(category.id) }"
                    :alt="category.name"
                    :src="getImageUrl(category.image)"
                  />
                </div>
                <XCircleIcon class="h-5 w-5 text-red-500"></XCircleIcon>
              </Button>
              <Button
                v-else
                class="items-center gap-2 bg-card-bg"
                :disabled="true"
                >All categories
                <CheckCircleIcon
                  class="h-5 w-5 text-green-500"
                ></CheckCircleIcon>
              </Button>
              <Button
                @click="filterStore.resetPriceFilter()"
                v-if="filterStore.isPriceFiltered()"
                class="gap-2 bg-neutral-50"
                >Price range: {{ filterStore.filter.minPrice }} -
                {{ filterStore.filter.maxPrice }}
                <XCircleIcon class="h-5 w-5 text-red-500"></XCircleIcon>
              </Button>
              <Button
                v-else
                class="items-center gap-2 bg-card-bg"
                :disabled="true"
                >All prices
                <CheckCircleIcon
                  class="h-5 w-5 text-green-500"
                ></CheckCircleIcon>
              </Button>
              <Button
                @click="filterStore.resetSortFilter()"
                v-if="filterStore.isSortFiltered()"
                class="gap-2 bg-neutral-50"
                >Sort by price :
                {{
                  filterStore.filter.sortDirection === "asc"
                    ? "low to high"
                    : "high to low"
                }}
                <XCircleIcon class="h-5 w-5 text-red-500"></XCircleIcon>
              </Button>
              <Button
                v-else
                class="items-center gap-2 bg-card-bg"
                :disabled="true"
              >
                Default sort
                <CheckCircleIcon
                  class="h-5 w-5 text-green-500"
                ></CheckCircleIcon>
              </Button>
              <Button
                @click="filterStore.resetOpenNowFilter()"
                v-if="filterStore.isOpenNowFiltered()"
                class="gap-2 bg-neutral-50"
              >
                Available now
                <XCircleIcon class="h-5 w-5 text-red-500"></XCircleIcon>
              </Button>
              <Button
                v-else
                class="items-center gap-2 bg-card-bg"
                :disabled="true"
              >
                Working hours : All
                <CheckCircleIcon
                  class="h-5 w-5 text-green-500"
                ></CheckCircleIcon>
              </Button>
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
            <Categories
              v-if="activeFilterType === 'category'"
              :categories="categoriesItems"
            ></Categories>
            <SortBy v-if="activeFilterType === 'sort-by'"></SortBy>
            <DoublePriceSlider
              v-if="activeFilterType === 'price'"
            ></DoublePriceSlider>
          </div>
          <div class="mt-auto flex justify-center">
            <Button @click="submitFilter" class="w-4/5" type="primary"
              >Complete</Button
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<style scoped></style>
