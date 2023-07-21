import {useUserStore} from "@/stores/user";


const root = 'http://localhost:8080';

async function apiFetch(url, options) {
    return await fetch(root + url, options);
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
            const userStore = useUserStore();
            console.log(data.access_token, data.refresh_token, "json tokens")
            userStore.setTokens(data.access_token, data.refresh_token);
            console.log('Login Successful:', data);
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
            console.log('Registration Successful:', data);
        } else {
            const errorData = await response.json();
            console.log('Registration Failed:', errorData);
        }
    } catch (error) {
        console.error('Error during registration:', error);
    }
}


async function getUserData() {
    const url = '/auth/profile';
    const options = {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    };

    try {
        const response = await apiFetch(url, options);
        if (response.ok) {
            const data = await response.json();
            console.log('User Data:', data);
        } else {
            const errorData = await response.json();
            console.log('User Data Failed:', errorData);
        }
    } catch (error) {
        console.error('Error during User Data:', error);
    }
}


export { login, register };


