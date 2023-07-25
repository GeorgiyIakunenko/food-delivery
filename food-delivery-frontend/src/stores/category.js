import {defineStore} from "pinia";
import {reactive, ref} from "vue";

export const useCategoryStore = defineStore('category', () => {
    const categories = ref([])

    const CurrentCategory = reactive({
        id: 0,
        name: '',
        image: '',
        description : '',
})

    function setCategories(newCategories) {
        categories.value = newCategories
    }

    function setCurrentCategory(category) {
        CurrentCategory.id = category.id
        CurrentCategory.name = category.name
        CurrentCategory.image = category.image
        CurrentCategory.description = category.description
    }

    return {categories, setCategories, CurrentCategory, setCurrentCategory}

})