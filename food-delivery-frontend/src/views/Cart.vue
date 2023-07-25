<script setup>

import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
import {useCartStore} from "@/stores/cart";
import CartProduct from "@/components/CartProduct.vue";
import {ref} from "vue";
import {useUserStore} from "@/stores/user";

const userStore = useUserStore()
const cartStore = useCartStore()

const addToCart = (product) => {
    cartStore.addToCart(product)
}

const removeFromCart = (item) => {
    cartStore.removeFromCart(item)
}
</script>

<template>
    <Header />
    <main>
        <div class="container">
          <div class="box">
            <h1>Cart</h1>
<!--            <img src="@/assets/images/" alt="">-->
          </div>
            <div class="cart-products">
              <CartProduct :product="item"  v-for="item in cartStore.items" :key="item.id" />
            </div>
            <div class="cart-total">
              <h2>Total: {{cartStore.cartTotal}} HUF</h2>
              <button v-if="userStore.access_token !== ''" class="btn">Checkout</button>
              <router-link v-else  to="/login"> <button  class="btn login-to-checkout">Login to Checkout</button></router-link>
            </div>
        </div>
    </main>
    <Footer />
</template>

<style scoped>

  main {
    padding-bottom: 15px;
  }

  .cart-products {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    grid-gap: 20px;
  }

  h1 {
    text-align: center;
    margin: 2rem 0;
  }

  .cart-total {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 2rem;
  }

  .btn {
    border: none;
    background: #7c6326;

    color: white;
    padding: 1rem 2rem;
    border-radius: 26px;
    font-family: "DM Sans", sans-serif;
    font-weight: bold;
    cursor: pointer;
  }

  .login-to-checkout {
    background: #483a17;
  }

  .box {
    display: flex;
    align-items: center;
    justify-content: center;
    grid-gap: 50px;
    flex-wrap: wrap;
    margin-bottom: 2rem;
  }

  @media (max-width: 550px) {
    .cart-products {
      grid-template-columns: 1fr;
    }

    .cart-total {
      flex-direction: column;
      grid-gap: 20px;
      align-items: center;
    }
  }
</style>