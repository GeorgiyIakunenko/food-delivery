import {computed, reactive, ref} from "vue";
import {defineStore} from "pinia";
import {getLocalStorageItem, setLocalStorageItem} from "@/healpers/localstorage";

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


    const access_token = ref(getLocalStorageItem('access_token') || '');
    const refresh_token = ref(getLocalStorageItem('refresh_token') || '');

    const setTokens = (accessTokenNew, refreshTokenNew) => {
        access_token.value = accessTokenNew;
        refresh_token.value = refreshTokenNew;
        setLocalStorageItem('access_token', accessTokenNew);
        setLocalStorageItem('refresh_token', refreshTokenNew);
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

    const logout = () => {
        access_token.value = '';
        refresh_token.value = '';
        setLocalStorageItem('access_token', '');
        setLocalStorageItem('refresh_token', '');
        setUser({
            id: 0,
            email: '',
            first_name: '',
            last_name: '',
            username: '',
            address: '',
            phone: '',
            age: 0,
        })

    }


    return { user, access_token, refresh_token, setTokens, setUser, logout}

})