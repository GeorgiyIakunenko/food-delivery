import {defineStore} from "pinia";
import {ref, reactive} from "vue";
import axios from "axios";

export const useSupplierStore = defineStore('supplier', () => {

    const suppliers = ref(null)


    const CurrentSupplier = reactive(null)

    function setCurrentSupplier(supplier) {
        CurrentSupplier.value = supplier
    }

    async function getSupplier(id) {
        try {
            const response = await axios.get('http://localhost:8080/supplier/' + id);
            CurrentSupplier.value = response.data
        } catch (error) {
            console.error(error);
        }
    }

    async function getSuppliers() {
        try {
            const response = await axios.get('http://localhost:8080/suppliers');
            suppliers.value = response.data
        } catch (error) {
            console.error(error);
        }

    }




    return {suppliers, getSuppliers}
})