<script setup>
import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
import {useCategoryStore} from "@/stores/category";
import {getAllCategories} from "@/api/api";
import Category from "@/components/Category.vue";
import {computed} from "vue";



getAllCategories()

const CategoryStore = useCategoryStore()

const Categories = computed(() => CategoryStore.categories)


</script>

<template>
  <Header />
  <main>
    <div class="container">
      <div class="box">
        <img src="@/assets/images/categories/categories.png" alt="categories">
        <h1>Categories</h1>
      </div>
      <div class="categories">
          <Category :category="category" v-for="category in Categories" :key="category.id"></Category>
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

  .categories {
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
    padding: 60px 0;
  }

  @media screen and (max-width: 768px) {
    h1 {
      font-size: 3rem;
    }
  }


</style>