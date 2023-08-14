<script setup>
import { onMounted, ref } from "vue";
import router from "@/router/router";
import { getProductsByCategoryAndSupplierIDs } from "@/api/products";
import ProductCard from "@/components/ProductCard.vue";
import Breadcrumb from "@/components/Breadcrumb.vue";

const products = ref([]);

const supplierId = router.currentRoute.value.params.supplierId;
const categoryId = router.currentRoute.value.params.categoryId;

onMounted(async () => {
  const res = await getProductsByCategoryAndSupplierIDs(categoryId, supplierId);
  console.log(res);
  if (res.success === true) {
    console.log(res.data);
    products.value = res.data;
  }
});
</script>

<template>
  <main>
    <div class="">
      <div class="container">
        <Breadcrumb
          :category-id="+categoryId"
          :supplier-id="+supplierId"
          :is-category="false"
          class=""
        ></Breadcrumb>
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
