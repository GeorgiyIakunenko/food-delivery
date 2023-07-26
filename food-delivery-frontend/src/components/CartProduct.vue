<script setup>
  import {useCartStore} from "@/stores/cart";
  import {computed} from "vue";
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

  const itemTotalPrice = computed(() => {
    return props.product.price * props.product.quantity
  })

  const removeFromCart = (product) => {
    cartStore.removeFromCart(product)
  }

  const decreaseProductQuantity = (product) => {
    cartStore.decreaseQuantity(product)
  }

  const prodImageUrl = new URL('/' + props.product.image, import.meta.url)

</script>

<template>
    <div class="cart-product">

        <div class="cart-product__info">
          <div class="cart-product__image">
            <img class="cart-product__image" :src="prodImageUrl" :alt="product.name">
          </div>
          <h3>{{ product.name }}</h3>
        </div>
        <div class="cart-product__bottom">
            <p> HUF {{ itemTotalPrice }}</p>
          <div class="cart-btn-box">
            <button @click="decreaseProductQuantity(props.product)" class="btn">
              <svg width="52" height="52" viewBox="0 0 52 52" fill="none" xmlns="http://www.w3.org/2000/svg">
              <g filter="url(#filter0_d_2_1271)">
                <circle cx="26" cy="26" r="16" fill="#FF9431"/>
              </g>
              <line x1="19" y1="26" x2="33" y2="26" stroke="white" stroke-width="2" stroke-linecap="round"/>
            </svg>
            </button>
            {{props.product.quantity}}
            <button @click="addToCart(props.product)" class="btn">
              <svg xmlns="http://www.w3.org/2000/svg" width="52" height="52" viewBox="0 0 52 52" fill="none">
              <g filter="url(#filter0_d_2_1270)">
                <circle cx="26" cy="26" r="16" fill="#FF9431"/>
              </g>
              <line x1="18.9988" y1="25.8571" x2="33.0013" y2="25.8571" stroke="white" stroke-width="2" stroke-linecap="round"/>
              <line x1="25.7144" y1="33" x2="25.7144" y2="19" stroke="white" stroke-width="2" stroke-linecap="round"/>
              <defs>
                <filter id="filter0_d_2_1270" x="0" y="0" width="52" height="52" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB">
                  <feFlood flood-opacity="0" result="BackgroundImageFix"/>
                  <feColorMatrix in="SourceAlpha" type="matrix" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" result="hardAlpha"/>
                  <feOffset/>
                  <feGaussianBlur stdDeviation="5"/>
                  <feComposite in2="hardAlpha" operator="out"/>
                  <feColorMatrix type="matrix" values="0 0 0 0 0.94902 0 0 0 0 0.6 0 0 0 0 0.290196 0 0 0 0.15 0"/>
                  <feBlend mode="normal" in2="BackgroundImageFix" result="effect1_dropShadow_2_1270"/>
                  <feBlend mode="normal" in="SourceGraphic" in2="effect1_dropShadow_2_1270" result="shape"/>
                </filter>
              </defs>
            </svg>
            </button>
          </div>
        </div>
    </div>
</template>

<style scoped>
  .cart-product {
    font-family: "DM Sans", sans-serif;
    display: flex;
    color: #0D0D0D;
    background-color: white;
    align-items: center;
    justify-content: space-between;
    padding: 0.5rem;
    border-radius: 1rem;
    box-shadow: 0 0 10px rgba(0,0,0,0.2);
    text-transform: capitalize;
  }

  .cart-btn-box {
    display: flex;
    align-items: center;
    border-radius: 26px;
    background: #FFF1E5;
  }

  .cart-product__image {
    min-width: 100px;
    height: 100px;
  }

  .cart-product__bottom {
    text-align: center;
    font-weight: bold;

  }

  .cart-product__info {
    text-align: center;
  }

  .btn {
    border: none;
    background: transparent;
    cursor: pointer;
  }

  @media (max-width: 560px) {
    .cart-product {
      text-align: center;
      flex-direction: column;
    }
    .cart-product__image {
      margin-bottom: 1rem;
    }
    .cart-product__bottom {
      display: flex;
      flex-direction: column;
      align-items: center;
    }

  }

</style>