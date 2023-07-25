import { useUserStore } from "@/stores/user";
import { getLocalStorageItem, setLocalStorageItem } from "@/healpers/localstorage";
import { useProductStore } from "@/stores/product";
import { useCategoryStore } from "@/stores/category";
import { useSupplierStore } from "@/stores/supplier";
import axios from "axios";

const api = axios.create({
    baseURL: 'http://localhost:8080',
    headers: {
        'Content-Type': 'application/json',
    },
});

api.interceptors.response.use(
    (response) => response,
    async (error) => {
        const originalRequest = error.config;

        if (error.response.message === 'token is expired' ) {
            originalRequest._retry = true;
            const refreshToken = getLocalStorageItem('refresh_token');
            try {
                const { data } = await api.post('/auth/refresh', null, {
                    headers: { 'Authorization': `Bearer ${refreshToken}` },
                });

                if (!data || !data.access_token || !data.refresh_token) {
                    console.log('Refresh token expired. Logging out.')
                    useUserStore().logout();
                    return;
                }

                useUserStore().setTokens(data.access_token, data.refresh_token);
                setLocalStorageItem('access_token', data.access_token);
                setLocalStorageItem('refresh_token', data.refresh_token);
                originalRequest.headers['Authorization'] = `Bearer ${data.access_token}`;
                return api(originalRequest);
            } catch (err) {
                console.log('Refresh token expired. Logging out.')
                useUserStore().logout();
                return Promise.reject(err);
            }
        }
        return Promise.reject(error);
    }
);

async function login(email, password) {
    try {
        const response = await api.post('/auth/login', { email, password }, {
            headers: { 'Content-Type': 'application/json' },

        });
        useUserStore().setTokens(response.data.access_token, response.data.refresh_token);
        await getUserData();
        return true;
    } catch (error) {
        console.error('Login Failed:', error.response.data);
        return false;
    }
}

async function register(userData) {
    try {
        const response = await api.post('/auth/register', userData);
        return response.status === 200;
    } catch (error) {
        console.error('Registration Failed:', error.response.data);
        return false;
    }
}

async function updateProfile(userData) {
    const accessToken = getLocalStorageItem('access_token');
    if (!accessToken) {
        console.error('Access token not found in localStorage.');
        return false;
    }
    try {
        const response = await api.put('/user/profile', userData, {
            headers: { 'Authorization': `Bearer ${accessToken}` },
        });
        useUserStore().setUserAfterUpdate();
        return true;
    } catch (error) {
        console.error('User Update Failed:', error.response.data);
        return false;
    }
}

async function getUserData() {
    const accessToken = useUserStore().access_token;
    if (!accessToken) {
        console.error('Access token not found in localStorage.');
        return false;
    }

    try {
        const response = await api.get('/user/profile', {
            headers: { 'Authorization': `Bearer ${accessToken}` },
        });
        useUserStore().setUser(response.data);
        return true;
    } catch (error) {
        console.error('Get User Data Failed:', error.response.data);
        return false;
    }
}

async function logout() {
    const accessToken = getLocalStorageItem('access_token');
    if (!accessToken) {
        console.error('Access token not found in localStorage.');
        return false;
    }

    try {
        await api.post('/auth/logout', {}, {
            headers: { 'Authorization': `Bearer ${accessToken}` },
        });
        useUserStore().logout();
        return true;
    } catch (error) {
        console.error('Logout Failed:', error.response.data);
        return false;
    }
}

async function getAllProducts() {
    try {
        const response = await api.get('/products');
        useProductStore().setProducts(response.data);
        return true;
    } catch (error) {
        console.error('Get Products Failed:', error.response.data);
        return false;
    }
}

async function getAllCategories() {
    try {
        const response = await api.get('/categories');
        useCategoryStore().setCategories(response.data);
        return true;
    } catch (error) {
        console.error('Get Categories Failed:', error.response.data);
        return false;
    }
}

async function getAllSuppliers() {
    try {
        const response = await api.get('/suppliers');
        useSupplierStore().setSuppliers(response.data);
        return true;
    } catch (error) {
        console.error('Get Suppliers Failed:', error.response.data);
        return false;
    }
}

async function getSupplierCategoriesById(id) {
    try {
        const response = await api.get(`/supplier/${id}/categories`);
        console.log(response.data)
        useCategoryStore().setCategories(response.data);
        return response.data;
    } catch (error) {
        console.error('Get Supplier Categories Failed:', error.response.data);
        return false;
    }
}

async function getSuppliersByCategoryId(id) {
    try {
        const response = await api.get(`/suppliers/category/${id}`);
        useSupplierStore().setSuppliers(response.data);
        return response.data;
    } catch (error) {
        console.error('Get Supplier Categories Failed:', error.response.data);
        return false;
    }
}

async function getSupplierById(id) {
    try {
        const response = await api.get(`/supplier/${id}`);
        useSupplierStore().setCurrentSupplier(response.data);
        return response.data;
    } catch (error) {
        console.error('Get Supplier Failed:', error.response.data);
        return false;
    }
}

async function getProductsBySupplierAndCategoryIDs(CategoryId, SupplierId) {
    try {
        const response = await api.get(`/categories/${CategoryId}/suppliers/${SupplierId}/products`);
        useProductStore().setProducts(response.data);
        return true;
    } catch (error) {
        console.error('Get Supplier Categories Failed:', error.response.data);
        return false;
    }
}

async function getCategoryById(categoryId) {
    try {
        const response = await api.get(`/category/${categoryId}`);
        useCategoryStore().setCurrentCategory(response.data);
        return true;
    } catch (error) {
        console.error('Get Category Failed:', error.response.data);
        return false;
    }
}


export {
    login,
    register,
    getUserData,
    logout,
    getAllProducts,
    getAllCategories,
    getAllSuppliers,
    updateProfile,
    getSupplierCategoriesById,
    getSupplierById,
    getProductsBySupplierAndCategoryIDs,
    getSuppliersByCategoryId,
    getCategoryById
};
