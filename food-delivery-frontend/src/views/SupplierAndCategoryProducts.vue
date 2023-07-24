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
<!--      <h1>Products in {{categoryStore.CurrentCategory.value.name}} of {{supplierStore.CurrentSupplier.value.name}}</h1>-->
      <div class="products">
        <Product :product="product" v-for="product in productStore.products" :key="product.id"></Product>
      </div>
    </div>
  </main>
  <Footer />
</template>

<style scoped>

</style>