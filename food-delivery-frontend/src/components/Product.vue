<script setup>
  const props = defineProps({
    product: {
      type: Object,
      required: true
    }
  })
</script>

<template>
  <div class="product-card">
    <div class="product__supplier">{{product.supplier.name}}</div>
    <div class="product__image">
      <img :src="product.image" alt="product.name">
    </div>
    <div class="product__info">
      <h3>{{ product.name }}</h3>
      <p>{{ product.description }}</p>
      <p>{{ product.price }}</p>
    </div>
  </div>
</template>

<style scoped>
  .product-card {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 300px;
    height: 400px;
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 1rem;
    margin: 1rem;
  }
</style>