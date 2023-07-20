<script setup>
  import { RouterLink, RouterView } from 'vue-router'
  import {useSupplierStore} from "../stores/supplier";
  import Supplier from "./Supplier.vue";

  const supplierStore = useSupplierStore()

  supplierStore.getSuppliers()
</script>
<template>
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
</template>

<style scoped>

</style>