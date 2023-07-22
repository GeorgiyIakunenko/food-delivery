<script setup>
  import { RouterLink, RouterView } from 'vue-router'
  import {useSupplierStore} from "../stores/supplier";
  import Supplier from "./Supplier.vue";
  import Header from "@/components/Header.vue";
  import Footer from "@/components/Footer.vue";

  const supplierStore = useSupplierStore()

  supplierStore.getSuppliers()
</script>
<template>
  <Header />
  <main>
    <div v-for="supplier in supplierStore.suppliers">
      <h1>{{ supplier.name }}</h1>
      <p>{{ supplier.description }}</p>
      <p>{{supplier.type}}</p>
      <button @click="$router.push({
        name: 'supplier',
        params: { id: supplier.id, name : 'aldi' }
      })" >Open</button>
    </div>
  </main>
  <Footer />
</template>

<style scoped>

</style>