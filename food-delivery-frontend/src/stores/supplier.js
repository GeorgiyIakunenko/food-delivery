import {defineStore} from "pinia";
import {ref, reactive} from "vue";
import axios from "axios";

export const useSupplierStore = defineStore('supplier', () => {

    const suppliers = ref(null)


    const CurrentSupplier = reactive(null)

    function setCurrentSupplier(supplier) {
        CurrentSupplier.value = supplier
    }

    function setSuppliers(newSuppliers) {
        suppliers.value = newSuppliers
    }

    return {suppliers, setCurrentSupplier, CurrentSupplier, setSuppliers }
})