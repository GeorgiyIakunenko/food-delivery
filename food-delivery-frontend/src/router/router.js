import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/suppliers',
      name: 'suppliers',
        component: () => import('../views/Suppliers.vue')
    },{
        path: '/suppliers/:id',
        name: 'supplier',
        props: true,
        component: () => import('../views/Supplier.vue')
    },
    {
      path: '/user',
      name: 'user',
      component: () => import('../views/User.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue')
    },{
        path: '/test',
        name: 'test',
        component: () => import('../components/Auth/LoginForm.vue')
    }

  ]
})

export default router
