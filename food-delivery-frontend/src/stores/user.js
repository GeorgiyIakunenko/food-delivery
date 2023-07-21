import {computed, reactive, ref} from "vue";
import {defineStore} from "pinia";

export const useUserStore = defineStore('user', () => {


    const user = reactive({
        id: 0,
        email: '',
        first_name: '',
        last_name: '',
        username: '',
        address: '',
        phone: '',
        age: 0,
    })


    const accessToken = ref("")
    const refreshToken = ref("")

    const setTokens = (accessTokenNew, refreshTokenNew ) => {

        accessToken.value = accessTokenNew;
        refreshToken.value = refreshTokenNew;
        console.log(accessToken, refreshToken)
    };

    const setUser = (userData) => {
        user.id = userData.id;
        user.email = userData.email;
        user.first_name = userData.first_name;
        user.last_name = userData.last_name;
        user.username = userData.username;
        user.address = userData.address;
        user.phone = userData.phone;
        user.age = userData.age;
    }

    return { user, accessToken, refreshToken, setTokens, setUser}

})