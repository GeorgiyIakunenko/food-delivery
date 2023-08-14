<script setup>
import { computed, onMounted, ref } from "vue";
import { getCategoryById } from "@/api/category";
import { getSupplierById } from "@/api/supplier";

const props = defineProps({
  categoryId: {
    type: Number,
    required: false,
  },
  supplierId: {
    type: Number,
    required: false,
  },
  isCategory: {
    type: Boolean,
    required: false,
  },
});

const homePage = computed(() => {
  return props.isCategory ? "categories" : "suppliers";
});

const categoryName = ref("");
const supplierName = ref("");

onMounted(async () => {
  console.log(props);

  if (props.categoryId) {
    const res = await getCategoryById(props.categoryId);
    console.log(res);
    if (res.success === true) {
      console.log(res.data);
      categoryName.value = res.data.name;
    }
  }
  if (props.supplierId) {
    const res = await getSupplierById(props.supplierId);
    console.log(res);
    if (res.success === true) {
      console.log(res.data);
      supplierName.value = res.data.name;
    }
  }
});
</script>

<template>
  <!-- Breadcrumb -->
  <nav
    class="mx-auto mb-10 flex w-fit justify-center rounded-xl border border-gray-200 bg-gray-50 px-3 py-3 font-sans font-medium text-gray-700 dark:border-gray-700 dark:bg-gray-800 md:mb-14 md:px-7"
    aria-label="Breadcrumb"
  >
    <ol class="inline-flex items-center space-x-1 md:space-x-3">
      <li class="inline-flex items-center">
        <router-link
          class="inline-flex items-center text-sm font-medium text-gray-700 hover:text-blue-600 dark:text-gray-400 dark:hover:text-white"
          :to="`/${homePage}`"
          exact
        >
          <svg
            class="mr-2.5 h-3 w-3"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="currentColor"
            viewBox="0 0 20 20"
          >
            <path
              d="m19.707 9.293-2-2-7-7a1 1 0 0 0-1.414 0l-7 7-2 2a1 1 0 0 0 1.414 1.414L2 10.414V18a2 2 0 0 0 2 2h3a1 1 0 0 0 1-1v-4a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v4a1 1 0 0 0 1 1h3a2 2 0 0 0 2-2v-7.586l.293.293a1 1 0 0 0 1.414-1.414Z"
            />
          </svg>
          {{ homePage.charAt(0).toUpperCase() + homePage.slice(1) }}
        </router-link>
      </li>
      <li>
        <div class="flex items-center">
          <svg
            class="mx-1 h-3 w-3 text-gray-400"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 6 10"
          >
            <path
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="m1 9 4-4-4-4"
            />
          </svg>
          <div class="inline-flex" v-if="!(supplierName && categoryName)">
            <span
              v-if="isCategory"
              class="ml-1 text-sm font-bold text-gray-700 hover:text-blue-600 dark:text-gray-400 dark:hover:text-white md:ml-2"
              >{{ categoryName }}</span
            >
            <span
              v-else
              class="ml-1 text-sm font-bold text-gray-700 hover:text-blue-600 dark:text-gray-400 dark:hover:text-white md:ml-2"
              >{{ supplierName }}</span
            >
          </div>
          <div class="inline-flex" v-else>
            <router-link
              v-if="isCategory"
              :to="`/category/${props.categoryId}/suppliers`"
              class="ml-1 text-sm text-gray-700 hover:text-blue-600 dark:text-gray-400 dark:hover:text-white md:ml-2"
              >{{ categoryName }}</router-link
            >
            <router-link
              v-else
              :to="`/supplier/${props.supplierId}/categories`"
              class="ml-1 text-sm text-gray-700 hover:text-blue-600 dark:text-gray-400 dark:hover:text-white md:ml-2"
              >{{ supplierName }}</router-link
            >
          </div>
        </div>
      </li>
      <li v-if="categoryName && supplierName" aria-current="page">
        <div class="flex items-center">
          <svg
            class="mx-1 h-3 w-3 text-gray-400"
            aria-hidden="true"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 6 10"
          >
            <path
              stroke="currentColor"
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="m1 9 4-4-4-4"
            />
          </svg>
          <span
            v-if="isCategory"
            class="ml-1 text-sm text-gray-700 hover:text-blue-600 dark:text-gray-400 dark:hover:text-white md:ml-2"
            :class="{ 'font-bold': categoryName && supplierName }"
            >{{ supplierName }}</span
          >
          <span
            v-else
            class="ml-1 text-sm text-gray-700 hover:text-blue-600 dark:text-gray-400 dark:hover:text-white md:ml-2"
            :class="{ 'font-bold': categoryName && supplierName }"
            >{{ categoryName }}</span
          >
        </div>
      </li>
    </ol>
  </nav>
</template>

<style scoped></style>
