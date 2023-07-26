import {defineStore} from "pinia";
import {ref, reactive} from "vue";
import axios from "axios";

export const useSupplierStore = defineStore('supplier', () => {

    const suppliers = ref([])

    const CurrentSupplier = reactive(
        {
            id: 0,
            name: '',
            type: '',
            image: '',
            address: '',
            open_time: '',
            close_time: '',
        }
    )

    function setCurrentSupplier(supplier) {
        CurrentSupplier.id = supplier.id
        CurrentSupplier.name = supplier.name
        CurrentSupplier.type = supplier.type
        CurrentSupplier.image = supplier.image
        CurrentSupplier.address = supplier.address
        CurrentSupplier.open_time = supplier.open_time
        CurrentSupplier.close_time = supplier.close_time
    }

    function setSuppliers(newSuppliers) {
        suppliers.value = newSuppliers
    }

    return {suppliers, setCurrentSupplier, CurrentSupplier, setSuppliers }
})