import {defineStore} from "pinia";
import {ref, reactive} from "vue";


export const useProductStore = defineStore('product', () => {

    const products = ref([])

    function setProducts(newProducts) {
        console.log(newProducts)
        products.value = newProducts
    }

    return {products, setProducts}
})


