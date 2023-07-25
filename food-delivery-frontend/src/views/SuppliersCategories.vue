<script setup>
import {useSupplierStore} from "@/stores/supplier";
import {useCategoryStore} from "@/stores/category";
import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
import {onMounted} from "vue";
import {getSupplierById, getSupplierCategoriesById} from "@/api/api";
import Category from "@/components/Category.vue";
import {useRoute} from "vue-router";

const route = useRoute()



onMounted(() => {
  const id = route.params.id
   getSupplierById(id)
   getSupplierCategoriesById(id)
})

const supplierStore = useSupplierStore()
const categoryStore = useCategoryStore()


</script>

<template>
  <Header />
  <main>
    <div class="container">
    <h1>Categories of {{supplierStore.CurrentSupplier.name}}</h1>
      <div class="categories">
        <Category :to="'products'" :category="category" v-for="category in categoryStore.categories" :key="category.id">
        </Category>
      </div>
    </div>
  </main>
  <Footer />
</template>

<style scoped>
  main {
    background: #FFF1E5;
  }

  h1 {
    font-family: 'DM Sans', sans-serif;
    font-size: 4rem;
    text-align: center;
    font-weight: 600;
    margin-bottom: 2rem ;
    margin-left: 1rem;
  }

  .categories {
    display: flex;
    grid-gap: 20px;
    flex-wrap: wrap;
    justify-content: center;
  }


</style>