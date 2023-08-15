<script setup>
import { onMounted, ref } from "vue";
import { getAllCategories } from "@/api/category";
import { getImageUrl } from "@/utils/url";
import { useFilterStore } from "@/store/filter";

const filterCategories = useFilterStore().filter.categories;

const isCategoryActive = (id) => {
  return filterCategories.includes(id);
};

const toggleCategory = (id) => {
  if (isCategoryActive(id)) {
    filterCategories.splice(filterCategories.indexOf(id), 1);
  } else {
    filterCategories.push(id);
  }
};

const categories = ref([]);

onMounted(async () => {
  const res = await getAllCategories();
  if (res.success === true) {
    categories.value = res.data;
  }
});
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
