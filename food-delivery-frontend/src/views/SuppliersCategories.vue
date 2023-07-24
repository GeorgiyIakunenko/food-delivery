<script setup>
  import {useSupplierStore} from "@/stores/supplier";
  import {useCategoryStore} from "@/stores/category";
  import Header from "@/components/Header.vue";
  import Footer from "@/components/Footer.vue";
  import {onMounted} from "vue";
  import {getSupplierById, getSupplierCategoriesById} from "@/api/api";
  import Category from "@/components/Category.vue";
  import {useRoute} from "vue-router";

  const supplierStore = useSupplierStore()
  const categoryStore = useCategoryStore()

  const route = useRoute()


  onMounted(() => {
    const id = route.params.id
    getSupplierCategoriesById(id)
    getSupplierById(id)
  })
</script>

<template>
  <Header />
  <main>
    <div class="container">
      <h1>Categories of {{supplierStore.CurrentSupplier.value.name}}</h1>
      <div class="categories">
          <Category :category="category" v-for="category in categoryStore.categories" :key="category.id"></Category>
      </div>
    </div>
  </main>
  <Footer />
</template>

<style scoped>

</style>