import {computed, ref} from "vue";
import {defineStore} from "pinia";

export const useUserStore = defineStore('user', () => {
    const user = ref(null)

    const userCount = ref(0)

    function AgeUser() {
        user.value.age++
    }

    const isLoggedIn = computed(() => user.value !== null)

    function IncrementUserCount() {
        userCount.value++
        console.log(userCount.value)
    }

    function setUser(newUser) {
        user.value = newUser
    }

    return {user, userCount, IncrementUserCount, setUser, isLoggedIn, AgeUser}
})