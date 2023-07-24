import {useUserStore} from "@/stores/user";
import {getLocalStorageItem} from "@/healpers/localstorage";
import {useProductStore} from "@/stores/product";
import {useCategoryStore} from "@/stores/category";



const root = 'http://localhost:8080'

//console.log(root)

async function apiFetch(url, options) {
    return fetch(root + url, options);
}

async function login(email, password) {
    const url = '/auth/login';
    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
    };
    console.log(options.body)

    try {
        const response = await apiFetch(url, options);
        if (response.ok) {
            const data = await response.json();
            useUserStore().setTokens(data.access_token, data.refresh_token);
            await getUserData();
            useUserStore().startLogoutTimer();
            return true;
        } else {
            const errorData = await response.json();
            console.log('Login Failed:', errorData);
            return false;
        }
    } catch (error) {
        console.error('Error during login:', error);
        return false;
    }
}

async function register(userData) {
    const url = '/auth/register';
    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(userData),
    };

    try {
        const response = await apiFetch(url, options);
        if (response.ok) {
            const data = await response.json();
            return true;
        } else {
            const errorData = await response.json();
            console.log('Registration Failed:', errorData);
            return false;
        }
    } catch (error) {
        console.error('Error during registration:', error);
        return false;
    }
}

async function getUserData() {
    const url = '/user/profile';

    const accessToken = useUserStore().access_token;
    console.log(accessToken)
    if (!accessToken) {
        console.error('Access token not found in localStorage.');
        return false;
    }

    const options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`, // Add the access token to the Authorization header
        },
    };

    try {
        const response = await apiFetch(url, options);
        if (response.ok) {
            const data = await response.json();
            useUserStore().setUser(data);
            return true;
        } else {
            const errorData = await response.json();
            console.log('User Data Failed:', errorData);
            return false;
        }
    } catch (error) {
        console.error('Error during User Data:', error);
        return false;
    }
}

async function logout() {
    const url = '/auth/logout';

    const accessToken = getLocalStorageItem('access_token');
    console.log(accessToken)
    if (!accessToken) {
        console.error('Access token not found in localStorage.');
        return false
    }

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${accessToken}`, // Add the access token to the Authorization header
        },
    };

    try {
        const response = await apiFetch(url, options);
        if (response.ok) {
            useUserStore().logout();
            return true;
        } else {
            const errorData = await response.json();
            console.log('Logout Failed:', errorData);
            return false;
        }
    } catch (error) {
        console.error('Error during logout:', error);
        return false;
    }
}

async function getAllProducts() {
    const url = '/products';
    const options = {
        method: 'GET',
    }

    try {
        const response = await apiFetch(url, options);
        if (response.ok) {
            const products = await response.json();
            console.log(products)
            useProductStore().setProducts(products)
            return true;
        } else {
            const errorData = await response.json();
            console.log('Get Products Failed:', errorData);
            return false;
        }
    } catch (error) {
        console.error('Error during getting products:', error);
        return false;
    }
}

async function getAllCategories() {
    const url = '/categories';
    const options = {
        method: 'GET',
    }

    try {
        const response = await apiFetch(url, options);
        if (response.ok) {
            const categories = await response.json();
            console.log(categories)
            useCategoryStore().setCategories(categories)
            return true;
        } else {
            const errorData = await response.json();
            console.log('Get Categories Failed:', errorData);
            return false;
        }
    } catch (error) {
        console.error('Error during getting categories:', error);
        return false;
    }
}


export { login, register, getUserData, logout, getAllProducts, getAllCategories }


