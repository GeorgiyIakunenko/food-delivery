<script setup>
import {useProductStore} from "@/stores/product";
import {useCategoryStore} from "@/stores/category";
import {useRoute, useRouter} from "vue-router";
import {onMounted} from "vue";
import {getProductsBySupplierAndCategoryIDs} from "@/api/api";
import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
import {useSupplierStore} from "@/stores/supplier";
import Product from "@/components/Product.vue";

const productStore = useProductStore()
const categoryStore = useCategoryStore()
const supplierStore = useSupplierStore()

const route = useRoute()

onMounted(() => {
  const categoryId = route.params.categoryId
  const supplierId = route.params.supplierId
  getProductsBySupplierAndCategoryIDs(categoryId, supplierId)
})

</script>

<template>
  <Header />
  <main>
    <div class="container">

        <h1>Products in {{categoryStore.CurrentCategory.name}} of {{supplierStore.CurrentSupplier.name}}</h1>
      <div class="products">
        <Product :product="product" v-for="product in productStore.products" :key="product.id"></Product>
      </div>
    </div>
  </main>
  <Footer />
</template>

<style scoped>

    .products {
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