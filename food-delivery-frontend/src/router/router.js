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
      component: () => import('../views/Auth.vue')
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import('../views/Profile.vue')
    },
    {
      path: '/products',
      name: 'products',
      component: () => import('../views/Products.vue')
    },
    {
        path: '/cart',
        name: 'cart',
        component: () => import('../views/Cart.vue')
    },
    {
        path: '/categories',
        name: 'categories',
        component: () => import('../views/Categories.vue')
    },
      {
          path: '/suppliers/:id/categories',
          name: 'supplier',
          component: () => import('../views/SuppliersCategories.vue')
      },
      {
          path: '/suppliers/:supplierId/category/:categoryId/products',
          name: 'supplier-product',
          component: () => import('../views/SupplierAndCategoryProducts.vue')
      },
      {
            path: '/suppliers/category/:categoryId',
            name: 'category-=suppliers',
            component: () => import('../views/CategorySuppliers.vue')
      }
  ]
})

export default router
