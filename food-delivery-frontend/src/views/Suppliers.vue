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
        <img src="@/assets/images/suppliers/supplier.png" alt="supplier">
        <h1>Products</h1>
      </div>
        <div class="suppliers">
          <router-link v-for="supplier in supplierStore.suppliers" :to="`/suppliers/${supplier.id}/categories`" :key="supplier.id">
            <Supplier to="categories" :supplier="supplier"></Supplier>
          </router-link>
        </div>
    </div>
  </main>
  <Footer />
</template>

<style scoped>
    .box {
      display: flex;
      align-items: center;
      justify-content: center;
      grid-gap: 50px;
      flex-wrap: wrap;
      margin-bottom: 2rem;
    }

    img {
      max-width: 400px;
    }

    .suppliers {
      display: flex;
      grid-gap: 20px;
      flex-wrap: wrap;
      justify-content: center;
    }

    h1 {
      font-family: 'DM Sans', sans-serif;
      font-size: 4rem;
      font-weight: 600;
      text-align: center;
      margin: 2rem 0;
    }

    main {
      background: #FFF1E5;
      padding-bottom: 50px;
    }


    @media (max-width: 550px) {

      .box {
        grid-gap: 10px;
      }

      img {
       max-width: 300px;
      }
    }
</style>