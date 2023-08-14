<script setup>
import router from "@/router/router";
import { onMounted, ref } from "vue";
import { getCategorySuppliersById } from "@/api/category";
import SupplierCard from "@/components/SupplierCard.vue";

const categoryId = router.currentRoute.value.params.categoryId;

const suppliers = ref([]);

onMounted(async () => {
  const res = await getCategorySuppliersById(categoryId);
  console.log(res);
  if (res.success === true) {
    console.log(res.data);
    suppliers.value = res.data;
  }
});
</script>

<template>
  <main>
    <div class="">
      <div class="container">
        <div
          class="mx-auto mt-7 grid max-w-fit grid-cols-2 gap-4 md:grid-cols-3 md:gap-7 xl:grid-cols-4"
        >
          <SupplierCard
            @click="
              router.push(
                '/category/' +
                  categoryId +
                  '/supplier/' +
                  supplier.id +
                  '/products',
              )
            "
            v-for="supplier in suppliers"
            :supplier="supplier"
          ></SupplierCard>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
