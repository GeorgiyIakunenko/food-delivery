import {computed, ref} from "vue";
import {defineStore} from "pinia";

export const useUserStore = defineStore('user', () => {



    const id = ref(1)
    const email = ref('')
    const first_name = ref('')
    const last_name = ref('')
    const username = ref('')
    const address = ref('')
    const phone = ref('')
    const age = ref(0)

    const isUserLogged = computed(() => {
        return email.value !== ''
    })

    return {id, email, first_name, last_name, username, address, phone, age, isUserLogged}
})