<script setup>

import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
import {onMounted, ref} from "vue";
import {getAllProducts} from "@/api/api";
import Product from "@/components/Product.vue";
import {useProductStore} from "@/stores/product";


const productStore = useProductStore(); // Initialize the product store


onMounted(() => {
  fetchProducts();
});

async function fetchProducts() {
  const success = await getAllProducts();
  if (success) {
    console.log(productStore.products); // Access the products state in the store
  } else {
    console.log("Failed to fetch products");
  }
}



</script>

<template>
  <Header />
  <main>
      <div class="container">
        <h1>Products</h1>
        <div class="products">
          <Product :product="product" v-for="product in productStore.products"  :key="product.id"></Product>
        </div>
      </div>
  </main>
  <Footer />
</template>

<style scoped>
    .products {
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
    }

    main {
      background: #FFF1E5;
    }
</style>