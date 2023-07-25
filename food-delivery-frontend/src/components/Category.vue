<script setup>

  import Button from "@/components/UI/Button.vue";
  import {useSupplierStore} from "@/stores/supplier";
  import {useCategoryStore} from "@/stores/category";

  const props = defineProps({
    category: {
      type: Object,
      required: true
    },
    to: {
      type: String,
      required: false,
      default: 'products'
    }
  })

  const supplierStore = useSupplierStore()
  const categoryStore = useCategoryStore()

  const categoryImageUrl = new URL('/' + props.category.image, import.meta.url)

</script>

<template>
  <div class="category">
    <img :src="categoryImageUrl" :alt="category.name">
    <p> {{category.description}} </p>
    <div class="category__bottom-box">
      <h2>{{ category.name }}</h2>
      <router-link v-if="to === 'suppliers'" :to="`/suppliers/category/${category.id}`"><Button intent="secondary">Open</Button></router-link>
      <router-link v-if="to === 'products'" :to="`/suppliers/${supplierStore.CurrentSupplier.id}/category/${category.id}/products`"><Button intent="secondary">Open</Button></router-link>
    </div>
  </div>
</template>

<style scoped>
  .category {
    background-color: white;
    width: 300px;
    height: 470px;
    display: flex;
    font-family: "DM Sans", sans-serif;
    text-transform: capitalize;
    flex-direction: column;
    border-radius: 10px;
    padding: 10px;
    overflow: hidden;
    box-shadow: 0 0 10px rgba(0,0,0,0.2);
    transition: all 0.3s ease-in-out;
  }

  .category:hover {
    transform: scale(1.05);
  }

  .category img {
    width: 100%;
    height: 200px;
    object-fit: contain;
  }

  .category__bottom-box {
    margin-top: auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    font-weight: bold;
    color : #fff;
    background-color: rgba(37, 114, 76, 0.37);
    border-radius: 10px;
  }

  p {
    font-size: 1rem;
    font-weight: 600;
    padding: 1rem;
  }

  .category__bottom-box h2 {
    font-size: 1.5rem;
    font-weight: 500;
  }

  .category__bottom-box p {
    font-size: 1rem;
    font-weight: 400;
    padding: 1rem;
  }

  .category__bottom-box .btn {
    padding: 0.5rem 1rem;
    border-radius: 5px;
    background-color: #FF9431;
    color: #fff;
    font-size: 1rem;
    font-weight: 500;
    text-decoration: none;
    transition: all 0.3s ease-in-out;
  }

  .category__bottom-box .btn:hover {
    background-color: #FF9431;
    color: #fff;
    transform: scale(1.05);
  }
</style>