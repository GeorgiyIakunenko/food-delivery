<script setup>

import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
import {onMounted, ref, watch} from "vue";
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

const filterSearch = ref("");

const filterProducts = () => {
  console.log(filterSearch.value)
  const search = filterSearch.value.toLowerCase().trim();
  console.log(search)
  const filteredProducts = productStore.products.filter((item) => {
    return item.name.toLowerCase().includes(search);
  });

  console.log(filteredProducts)

  productStore.setFilteredProducts(filteredProducts);
}

watch(filterSearch, () => {
  filterProducts();
});


</script>

<template>
  <Header />
  <main>
      <div class="container">
        <h1>Products</h1>
        <div class="filter-box">
          <div class="search">
            <label for="search">Search:</label>
            <input v-model="filterSearch" type="text" name="search" id="search">
          </div>
        </div>
        <div class="products">
          <Product  v-for="product in productStore.filteredProducts" :product="product"  :key="product.id"></Product>
        </div>
      </div>
  </main>
  <Footer />
</template>

<style scoped>

.filter-box {
  display: flex;
  align-items: center;
  justify-content: center;
  grid-gap: 50px;
  margin-bottom: 50px;
}

.search input {
  background-color: transparent;
}

    h1 {
      text-align: center;
      margin: 2rem 0;
    }

    .products {
      display: flex;
      flex-wrap: wrap;
      justify-content: center;
    }

    main {
      background: #FFF1E5;
    }
</style>