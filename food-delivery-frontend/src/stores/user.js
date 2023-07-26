import {computed, onMounted, reactive, ref} from "vue";
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
        user.first_name = userData.firstName;
        user.last_name = userData.lastName;
        user.username = userData.username;
        user.address = userData.address;
        user.phone = userData.phone;
        user.age = userData.age;
    }

    const setUserAfterUpdate = () => {
        user.id = updatedUser.id !== 0 ? updatedUser.id : user.id;
        user.email = updatedUser.email !== '' ? updatedUser.email : user.email;
        user.first_name = updatedUser.first_name !== '' ? updatedUser.first_name : user.first_name;
        user.last_name = updatedUser.last_name !== '' ? updatedUser.last_name : user.last_name;
        user.username = updatedUser.username !== '' ? updatedUser.username : user.username;
        user.address = updatedUser.address !== '' ? updatedUser.address : user.address;
        user.phone = updatedUser.phone !== '' ? updatedUser.phone : user.phone;
        user.age = updatedUser.age !== 0 ? updatedUser.age : user.age;
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

    const updatedUser = reactive({
        id: 0,
        email: '',
        first_name: '',
        last_name: '',
        username: '',
        address: '',
        phone: '',
        age: 0,
    })

    const userUpdate = (userData) => {
        updatedUser.id = userData.id;
        updatedUser.email = userData.email;
        updatedUser.first_name = userData.first_name;
        updatedUser.last_name = userData.last_name;
        updatedUser.username = userData.username;
        updatedUser.address = userData.address;
        updatedUser.phone = userData.phone;
        updatedUser.age = userData.age;
    }

    const resetForm =  reactive({
        email: 'g10072004@gmail.com',
        reset_code: '',
        new_password: '',
    })

    const setResetForm = (resetData) => {
        resetForm.email = resetData.email;
        resetForm.reset_code = resetData.reset_code;
        resetForm.new_password = resetData.new_password;
    }

    const orders = ref([])

    const setOrders = (ordersData) => {
        orders.value = (ordersData)
    }

    return { user,
        access_token,
        refresh_token,
        setTokens,
        setUser,
        logout,
        userUpdate,
        updatedUser,
        setUserAfterUpdate,
        resetForm,
        setResetForm,
        orders,
        setOrders

    }

})