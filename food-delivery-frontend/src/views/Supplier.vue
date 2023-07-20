<script setup>
  import { RouterLink, RouterView, useRoute } from 'vue-router'
  import {useSupplierStore} from "../stores/supplier";
  import Supplier from "./Supplier.vue";
  const router = useRoute()

  const supplierStore = useSupplierStore()
  console.log(router.params)

</script>

<template>
  <div class="supplier">
    <h1>  {{$route.params.id}}</h1>
 <h2> {{ $route.params.name}}</h2>

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