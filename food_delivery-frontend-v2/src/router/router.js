import {createRouter, createWebHistory} from "vue-router";


const routes = [
    {
        path: '/',
        name: 'Home',
        component: () => import('@/views/Home.vue'),
    },
    {
        path: '/suppliers',
        name: 'Suppliers',
        component: () => import('@/views/Suppliers.vue'),
    },
    {
        path : '/categories',
        name : 'Categories',
        component : () => import('@/views/Categories.vue'),
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});
export default router