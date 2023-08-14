import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("@/views/Home.vue"),
  },
  {
    path: "/suppliers",
    name: "Suppliers",
    component: () => import("@/views/Suppliers.vue"),
  },
  {
    path: "/supplier/:supplierId/categories",
    name: "SupplierCategories",
    component: () => import("@/views/SupplierCategories.vue"),
  },
  {
    path: "/categories",
    name: "Categories",
    component: () => import("@/views/Categories.vue"),
  },
  {
    path: "/category/:categoryId/suppliers",
    name: "CategorySuppliers",
    component: () => import("@/views/CategorySuppliers.vue"),
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("@/views/Register.vue"),
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/views/Login.vue"),
  },
  {
    path: "/reset-password",
    name: "ResetPassword",
    component: () => import("@/views/ResetPassword.vue"),
  },
  {
    path: "/profile",
    name: "Profile",
    component: () => import("@/views/Profile.vue"),
  },
  {
    path: "/cart",
    name: "Cart",
    component: () => import("@/views/Cart.vue"),
  },
  {
    path: "/profile-info",
    name: "ProfileInfo",
    component: () => import("@/views/ProfileInfo.vue"),
  },
  {
    path: "/profile-password",
    name: "ProfilePassword",
    component: () => import("@/views/ProfilePassword.vue"),
  },
  {
    path: "/profile/orders",
    name: "Orders",
    component: () => import("@/views/Orders.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});
export default router;
