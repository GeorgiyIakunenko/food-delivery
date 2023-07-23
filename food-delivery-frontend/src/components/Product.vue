<script setup>
  import {useCartStore} from "@/stores/cart";

  const props = defineProps({
    product: {
      type: Object,
      required: true
    }
  })

  const cartStore = useCartStore()

  const addToCart = (product) => {
    cartStore.addToCart(product)
  }

  const removeFromCart = (product) => {
    cartStore.removeFromCart(product)
  }

</script>

<template>
  <div class="product-card">
    <div class="product-card__top">
      <div class="product__supplier">
        <div class="product__supplier-box">
          <img class="product__supplier-img" src="@/assets/images/suppliers/aldi.jpg" alt="{{product.supplier.name}}">
          <div class="product__supplier-name">{{product.supplier.name}}</div>
        </div>
        <div class="product__supplier-working-hours">
          <div class="product__supplier-working-hours-time">{{product.supplier.open_time}} - {{ product.supplier.close_time}}</div>
        </div>
      </div>

      <div class="product-card__category">{{ product.category.name}}</div>
    </div>

    <div class="product__image">
      <img src="@/assets/images/products/apple.png" alt="product.name">
    </div>
    <div class="product__info">
      <h3>{{ product.name }}</h3>
      <p>{{ product.description }}</p>
    </div>
    <div class="product-card__bottom">
      <p> HUF {{ product.price }}</p>
      <button @click="addToCart(props.product)" class="btn"><img src="@/assets/images/icons/plus.svg" alt=""></button>
    </div>
  </div>
</template>

<style scoped>
  .product-card {
    font-family: "DM Sans", sans-serif;
    display: flex;
    color: #0D0D0D;
    background-color: white;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    width: 250px;
    height: 400px;
    box-shadow: 0px 1px 8px 0px rgba(0, 0, 0, 0.12);
    border-radius: 10px;
    padding: 1.75rem;
    margin: 1rem;
  }

  .product-card__top {
    display: flex;
    justify-content: space-between;
    width: 100%;
    margin-bottom: 1rem;
  }

  img {
    height: 200px;
    width: 200px;
    object-fit: cover;
  }

  .product__supplier-img {
    height: 20px;
    width: 20px;
  }

  .product__supplier-box {
    display: flex;
    align-items: center;
  }

  .product__supplier-name {
    font-size: 0.75rem;
    font-weight: 600;
    color: #151111;
  }

  .product-card__category {
    font-size: 0.75rem;
    font-weight: 600;
    text-transform: Capitalize;
    color: #151111;
  }

  .product__info {
    width: 100%;
    text-transform: Capitalize;
  }

  .product__info h3 {

    font-size: 1.5rem;
    font-weight: 600;
    text-align: center;
    color: #151111;
    margin-bottom: 0.5rem;
  }



  .product-card__bottom {
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 100%;
    margin-top: 1rem;
    color : #FF9431;
    font-weight: bold;
  }

.product__info p {
    font-size: 0.75rem;
    font-weight: 400;
    color: #151111;
  text-align: left;
    margin-bottom: 0.5rem;
  }

  .btn {
    background-color: transparent;
    border: none;
    cursor: pointer;
    padding: 0;
    margin: 0;
  }

  .btn img {
    width: 30px;
    height: 30px;
  }

</style>