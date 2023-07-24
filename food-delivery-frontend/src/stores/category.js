import {defineStore} from "pinia";
import {ref} from "vue";

export const useCategoryStore = defineStore('category', () => {
    const categories = ref(null)

    const CurrentCategory = ref(null)

    function setCategories(newCategories) {
        categories.value = newCategories
    }

    function setCurrentCategory(category) {
        CurrentCategory.value = category
    }

    return {categories, setCategories, CurrentCategory, setCurrentCategory}

})