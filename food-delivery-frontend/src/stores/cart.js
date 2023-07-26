
import { defineStore } from 'pinia';
import {ref, computed, onMounted, reactive} from 'vue';

export const useCartStore = defineStore('cart', () => {
    const items = ref([]);

    const cartTotal = computed(() => {
        return items.value.reduce((total, item) => total + (item.quantity * item.price), 0.00);
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

    const orderRequest = reactive({
        total_price: 0,
        payment_method: 'Cash',
        address: '',
        product: [],
    })

    const setOrderRequest = (total_price, payment_method, address, product) => {
        orderRequest.total_price = parseFloat(total_price + ".00");
        orderRequest.payment_method = payment_method;
        orderRequest.address = address;
        orderRequest.product = product;
    }

    return {
        items,
        addToCart,
        removeFromCart,
        cartTotal,
        clearCart,
        initializeCartFromLocalStorage,
        decreaseQuantity,
        orderRequest,
        setOrderRequest,
    };
});
