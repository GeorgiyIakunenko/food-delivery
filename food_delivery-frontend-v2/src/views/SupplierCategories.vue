<script setup>
import CategoryCard from "@/components/CategoryCard.vue";
import router from "@/router/router";
import { onMounted, ref } from "vue";
import { getSupplierCategoriesById } from "@/api/supplier";
import Breadcrumb from "@/components/Breadcrumb.vue";

const supplierId = router.currentRoute.value.params.supplierId;

const categories = ref([]);

onMounted(async () => {
  const res = await getSupplierCategoriesById(supplierId);
  console.log(res);
  if (res.success === true) {
    console.log(res.data);
    categories.value = res.data;
  }
});
</script>

<template>
  <main>
    <div class="">
      <Breadcrumb
        :supplier-id="+supplierId"
        :is-category="false"
        class=""
      ></Breadcrumb>
      <div class="container">
        <div
          class="mx-auto mt-7 grid max-w-fit grid-cols-2 gap-4 md:grid-cols-3 md:gap-7 xl:grid-cols-4"
        >
          <CategoryCard
            @click="
              router.push(
                '/supplier/' +
                  supplierId +
                  '/category/' +
                  category.id +
                  '/products',
              )
            "
            v-for="category in categories"
            :category="category"
          ></CategoryCard>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
