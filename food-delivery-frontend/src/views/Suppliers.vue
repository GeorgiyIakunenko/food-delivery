<script setup>
  import { RouterLink, RouterView } from 'vue-router'
  import {useSupplierStore} from "../stores/supplier";
  import Supplier from "@/components/Supplier.vue";
  import Header from "@/components/Header.vue";
  import Footer from "@/components/Footer.vue";
  import {onMounted} from "vue";
  import {getAllSuppliers} from "@/api/api";

  const supplierStore = useSupplierStore()

  onMounted(() => {
    getAllSuppliers()
  })

</script>
<template>
  <Header />
  <main>
    <div class="container">
      <div class="box">
        <h1>Suppliers</h1>
        <div class="suppliers">
          <Supplier @click="supplierStore.se" :supplier="supplier" v-for="supplier in supplierStore.suppliers" :key="supplier.id"></Supplier>
        </div>
      </div>
    </div>
  </main>
  <Footer />
</template>

<style scoped>

    .box {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      margin-bottom: 2rem;
    }

    .suppliers {
      display: flex;
      grid-gap: 20px;
      flex-wrap: wrap;
      justify-content: center;
    }

    h1 {
      font-family: 'DM Sans', sans-serif;
      font-size: 6rem;
      color: #897ec7;
      font-weight: 600;
      margin-left: 1rem;
    }

    main {
      background: #FFF1E5;
    }
</style>