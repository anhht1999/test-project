import { createRouter, createWebHashHistory } from "vue-router";
import store from "../store";

const routes = [
  {
    path: "/",
    name: "HomePage",
    component: () => import("../pages/home/HomePage.vue"),
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

  {
    path: "/payment",
    name: "paymentPage",
    beforeEnter: (to, from, next)  => {
      if(store.state.users.isLoginSuccess) {
        next()
      } else {
        next("/login")
      }
    },
    component: () => import("../pages/payment/payemntPage.vue"),
  },

  {
    path: "/confirm",
    name: "ConfirmCheckout",
    component: () => import("../pages/confirm/confirmPage.vue"),
  },

  {
    path: "/user",
    name: "UserPage",
    component: () => import("../pages/user/userPage.vue"),
    children: [],
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  scrollBehavior() {
    return { top: 0 };
  },
  routes,
});

export default router;
