<script setup>

import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
import {useCartStore} from "@/stores/cart";

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
            <h1>Cart</h1>
            <div v-for="item in cartStore.items">
              <h1>{{item.name}}</h1>
              <h2>{{item.price}}</h2>
              <p>{{item.quantity}}</p>
              <button @click="removeFromCart(item)">Remove</button>
            </div>
        </div>
    </main>
    <Footer />
</template>

<style scoped>
</style>