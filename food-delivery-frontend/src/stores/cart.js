
import { defineStore } from 'pinia';
import {ref, computed, onMounted} from 'vue';

export const useCartStore = defineStore('cart', () => {
    const items = ref([]);

    const cartTotal = computed(() => {
        return items.value.reduce((total, item) => total + (item.quantity * item.price), 0);
    });


    function addToCart(product) {
        const existingProduct = items.value.find((item) => item.id === product.id);

        if (existingProduct) {
            existingProduct.quantity++;
            existingProduct.itemTotal = existingProduct.quantity * existingProduct.price;
        } else {
            items.value.push({ ...product, quantity: 1, itemTotal: product.price });
        }

        updateLocalStorage();
    }

    function removeFromCart(product) {
        items.value = items.value.filter((item) => item.id !== product.id);
        updateLocalStorage();
    }

    function decreaseQuantity(product) {
        const existingProduct = items.value.find((item) => item.id === product.id);

        if (existingProduct) {
            if (existingProduct.quantity === 1) {
                removeFromCart(product);
                return;
            }
            existingProduct.quantity--;
        } else {
            console.log("Error in DecreaseQuantity")
        }

        updateLocalStorage();
    }

    function clearCart() {
        items.value = [];
        updateLocalStorage();
    }

    function updateLocalStorage() {
        localStorage.setItem('cartItems', JSON.stringify(items.value));
    }

    onMounted(() => {
        initializeCartFromLocalStorage();
    });

    function initializeCartFromLocalStorage() {
        const cartItems = localStorage.getItem('cartItems');
        if (cartItems) {
            items.value = JSON.parse(cartItems);
        }
    }


    return {
        items,
        addToCart,
        removeFromCart,
        cartTotal,
        clearCart,
        initializeCartFromLocalStorage,
        decreaseQuantity,
    };
});
