import {defineStore} from "pinia";
import {ref} from "vue";
import axios from "axios";

export const useSupplierStore = defineStore('supplier', () => {

    const suppliers = ref(null)


    async function getSuppliers() {
        try {
            const response = await axios.get('http://localhost:8080/suppliers');
            suppliers.value = response.data
            console.log(suppliers.value)
        } catch (error) {
            console.error(error);
        }

    }


    return {suppliers, getSuppliers}
})