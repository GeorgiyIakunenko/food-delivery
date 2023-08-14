import { defineStore } from "pinia";
import { reactive } from "vue";

export const useCartStore = defineStore("cartStore", {
  state: () => {
    const cart = reactive({
      total_price: 0,
      payment_method: 0,
      address: "",
      products: [],
    });

    const setCart = (data) => {
      cart.total_price = data.total_price;
      cart.payment_method = data.payment_method;
      cart.address = data.address;
      cart.products = data.products;
    };

    const addProduct = (product) => {
      const existingProduct = cart.products.find((p) => p.id === product.id);

      if (existingProduct) {
        existingProduct.quantity++;
      } else {
        cart.products.push({ ...product, quantity: 1 });
      }

      updateTotalPrice();
    };

    const decreaseProduct = (product) => {
      const existingProductIndex = cart.products.findIndex(
        (p) => p.id === product.id,
      );

      if (existingProductIndex !== -1) {
        const existingProduct = cart.products[existingProductIndex];

        if (existingProduct.quantity > 1) {
          existingProduct.quantity--;
        } else {
          cart.products.splice(existingProductIndex, 1);
        }

        updateTotalPrice();
      }
    };

    const updateTotalPrice = () => {
      cart.total_price = cart.products.reduce((total, product) => {
        return total + product.price * product.quantity;
      }, 0);
    };

    const removeProduct = (product) => {
      const existingProductIndex = cart.products.indexOf(product.id);
      return cart.products.splice(existingProductIndex, 1);
    };

    return {
      cart,
      setCart,
      addProduct,
      decreaseProduct,
      updateTotalPrice,
      removeProduct,
    };
  },

  persist: true,
});
