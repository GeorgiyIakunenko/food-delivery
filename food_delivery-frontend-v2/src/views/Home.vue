<script setup>
import { useUserStore } from "@/store/user";
import { onMounted, ref, watch } from "vue";
import Filter from "@/components/Filter/Filter.vue";
import { getFilteredProducts } from "@/api/products";
import ProductCard from "@/components/ProductCard.vue";
import { useFilterStore } from "@/store/filter";
import debounce from "lodash.debounce";
import Button from "@/components/Button.vue";

const userStore = useUserStore();

const products = ref([]);

onMounted(async () => {
  const res = await getFilteredProducts(useFilterStore().filter);
  console.log(res);
  if (res.success === true) {
    products.value = res.data;
  }
});

watch(
  useFilterStore().filter,
  debounce(() => {
    filterProducts();
  }, 400),
);

const filterProducts = async () => {
  const res = await getFilteredProducts(useFilterStore().filter);
  console.log(res);
  if (res.success === true) {
    products.value = res.data;
  }
};
</script>

<template>
  <main class="flex flex-col pt-24">
    <div class="">
      <div class="container">
        <Filter @filter-change="filterProducts"></Filter>
      </div>
    </div>
    <div
      class="mx-auto mt-14 flex w-56 grow flex-col items-center justify-center gap-5 px-2 md:w-fit"
      v-if="products < 1"
    >
      <div class="text-center font-sans text-xl font-medium">
        No products found based on your filters!
      </div>
      <Button
        class="ml-2"
        @click="useFilterStore().resetFilter()"
        type="primary"
        >Reset filters</Button
      >
    </div>
    <div v-else class="">
      <div class="container">
        <div
          class="mx-auto mt-7 grid max-w-fit grid-cols-2 gap-4 md:grid-cols-3 md:gap-7 xl:grid-cols-4"
        >
          <div v-for="product in products" class="">
            <ProductCard :product="product"></ProductCard>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
