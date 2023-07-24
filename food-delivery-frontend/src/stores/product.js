import {defineStore} from "pinia";
import {ref, reactive} from "vue";


export const useProductStore = defineStore('product', () => {

    const products = ref(null)
    const CurrentProduct = reactive(null)

    function setCurrentProduct(product) {
        CurrentProduct.value = product
    }

    function setProducts(newProducts) {
        console.log(newProducts)
        products.value = newProducts
    }

    return {products, setProducts}
})


