<script setup>
import SupplierCard from "@/components/SupplierCard.vue";
import { onMounted, ref } from "vue";
import { getAllSuppliers } from "@/api/supplier";

const suppliers = ref([]);

onMounted(async () => {
  const res = await getAllSuppliers();
  console.log(res);
  if (res.success === true) {
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
          <div v-for="supplier in suppliers" class="">
            <SupplierCard :supplier="supplier"></SupplierCard>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
