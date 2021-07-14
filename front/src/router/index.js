import { createRouter, createWebHashHistory } from "vue-router";
import HomePage from "../pages/home/HomePage.vue";

const routes = [
  {
    path: "/",
    name: "HomePage",
    component: HomePage,
  },
  {
    path: "/product",
    name: "ProductsPage",
    component: () => import("../pages/products/ProductsPage.vue"),
  },
  {
    path: "/product/:id",
    name: "ProductDetail",
    component: () => import("../pages/product-detail/ProductDetail.vue"),
  },
  {
    path: "/cart",
    name: "Cart",
    component: () => import("../pages/carts/Cart.vue"),
  },

  {
    path: "/about",
    name: "AboutPage",
    component: () => import("../pages/about/AboutPage.vue"),
  },

  {
    path: "/login",
    name: "LoginPage",
    component: () => import("../pages/login/LoginPage.vue"),
  },
  {
    path: "/register",
    name: "RegisterPage",
    component: () => import("../pages/register/RegisterPage.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  scrollBehavior(){
      return {top: 0};
  },
  routes,
});

export default router;
