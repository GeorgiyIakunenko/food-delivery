<script setup>
  import {useSupplierStore} from "@/stores/supplier";
  import {useCategoryStore} from "@/stores/category";
  import Header from "@/components/Header.vue";
  import Footer from "@/components/Footer.vue";
  import Category from "@/components/Category.vue";
  import Supplier from "@/components/Supplier.vue";
  import {onMounted} from "vue";
  import {useRoute} from "vue-router";
  import {getCategoryById, getSuppliersByCategoryId} from "@/api/api";

  const supplierStore = useSupplierStore()
  const categoryStore = useCategoryStore()

  const route = useRoute()

  onMounted(() => {
    const id = route.params.categoryId
    getSuppliersByCategoryId(id)
    getCategoryById(id)
  })

</script>

<template>
  <Header />
  <main>
    <div class="container">
      <h1>Suppliers in  {{categoryStore.CurrentCategory.name}}</h1>
      <div class="categories">
        <Supplier to="products" :supplier="supplier" v-for="supplier in supplierStore.suppliers" :key="supplier.id"> </Supplier>
      </div>
    </div>
  </main>
  <Footer />
</template>

<style scoped>



    .categories {
      display: flex;
      grid-gap: 20px;
      flex-wrap: wrap;
      justify-content: center;
    }

    h1 {
      font-family: 'DM Sans', sans-serif;
      font-size: 4rem;
      text-align: center;
      font-weight: 600;
      margin-bottom: 2rem ;
      margin-left: 1rem;
    }
    main {
      background: #FFF1E5;
    }
</style>