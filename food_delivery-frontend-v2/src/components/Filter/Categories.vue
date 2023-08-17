<script setup>
import { computed } from "vue";
import { getImageUrl } from "@/utils/url";
import { useFilterStore } from "@/store/filter";

const props = defineProps({
  categories: {
    type: Array,
    required: true,
  },
});

const filterStore = useFilterStore();

const filterCategories = computed({
  get: () => filterStore.filter.categories,
  set: (value) => {
    filterStore.filter.categories = value;
  },
});

const isCategoryActive = (id) => {
  return filterCategories.value.includes(id);
};

const toggleCategory = (id) => {
  if (isCategoryActive(id)) {
    filterCategories.value.splice(filterCategories.value.indexOf(id), 1);
  } else {
    filterCategories.value.push(id);
  }
};
</script>

<template>
  <div class="flex flex-col items-start overflow-x-auto py-14 md:items-center">
    <div class="flex gap-4">
      <div
        v-for="category in categories"
        @click="toggleCategory(category.id)"
        class="flex h-24 w-24 cursor-pointer items-center justify-center rounded-full bg-card-bg transition-all duration-200"
        :class="{
          'bg-primary-100 hover:bg-primary-100': isCategoryActive(category.id),
          'bg-primary-50 hover:bg-primary-75': !isCategoryActive(category.id),
        }"
      >
        <img class="h-16 w-16" :src="getImageUrl(category.image)" alt="Pizza" />
      </div>
    </div>
  </div>
</template>

<style scoped>
::-webkit-scrollbar {
  height: 10px;
}

/* Track */
::-webkit-scrollbar-track {
  background: #f1f1f1;
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 5px;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>
