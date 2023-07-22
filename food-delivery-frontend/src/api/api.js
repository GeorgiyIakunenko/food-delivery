import {useUserStore} from "@/stores/user";
import {getLocalStorageItem} from "@/healpers/localstorage";


const root = 'http://localhost:8080';
//const userStore = useUserStore();  getting an error if define store locally

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
        } else {
            const errorData = await response.json();
            console.log('Login Failed:', errorData);
        }
    } catch (error) {
        console.error('Error during login:', error);
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
        } else {
            const errorData = await response.json();
            console.log('Registration Failed:', errorData);
        }
    } catch (error) {
        console.error('Error during registration:', error);
    }
}


async function getUserData() {
    const url = '/user/profile';

    const accessToken = useUserStore().access_token;
    console.log(accessToken)
    if (!accessToken) {
        console.error('Access token not found in localStorage.');
        return;
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
        } else {
            const errorData = await response.json();
            console.log('User Data Failed:', errorData);
        }
    } catch (error) {
        console.error('Error during User Data:', error);
    }
}

async function logout() {
    const url = '/auth/logout';

    const accessToken = getLocalStorageItem('access_token');
    console.log(accessToken)
    if (!accessToken) {
        console.error('Access token not found in localStorage.');
        return;
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
        } else {
            const errorData = await response.json();
            console.log('Logout Failed:', errorData);
        }
    } catch (error) {
        console.error('Error during logout:', error);
    }
}


export { login, register, getUserData, logout }


