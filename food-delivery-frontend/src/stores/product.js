import {defineStore} from "pinia";
import {ref, reactive} from "vue";


export const useProductStore = defineStore('product', () => {

    const products = ref([])

    function setProducts(newProducts) {
        console.log(newProducts)
        products.value = newProducts
        filteredProducts.value = newProducts
    }

    const filteredProducts = ref([])

    function setFilteredProducts(newProducts) {
        filteredProducts.value = newProducts
    }


    return {products, setProducts,filteredProducts,setFilteredProducts}
})


