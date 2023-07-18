<script setup>
  import { RouterLink, RouterView, useRoute } from 'vue-router'
  import {useSupplierStore} from "../stores/supplier";
  import Supplier from "./Supplier.vue";

  const route = useRoute()

  const SupplierID = route.params.id

  const supplierStore = useSupplierStore()
  console.log(supplierStore.suppliers)
  console.log(SupplierID)
  //const supplier = supplierStore.suppliers($route.params.id - 1)

</script>

<template>
  <div class="supplier">
    <h1>  {{$route.params.id}}</h1>
<!--    <h1>{{supplier.name}}</h1>
    <p>{{supplier.description}}</p>
    <div>{{supplier.type}}</div>-->
  </div>
</template>

<style scoped>
    .supplier {
      border: 1px solid black;
      padding: 10px;
      margin: 10px;
      background-color: beige;
      border-radius: 10px;
    }

    .supplier h1 {
      font-size: 1.5rem;
    }


</style>