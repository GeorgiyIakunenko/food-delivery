<script setup>
import { onMounted, ref } from "vue";
import { getAllCategories } from "@/api/category";
import CategoryCard from "@/components/CategoryCard.vue";

const categories = ref([]);

onMounted(async () => {
  const res = await getAllCategories();
  console.log(res);
  if (res.success === true) {
    categories.value = res.data;
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
          <CategoryCard
            v-for="category in categories"
            :category="category"
          ></CategoryCard>
        </div>
      </div>
    </div>
  </main>
</template>

<style scoped></style>
