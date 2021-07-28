import { createRouter, createWebHistory } from "vue-router";
import store from "../store";

const routes = [
  {
    path: "/",
    name: "HomePage",
    component: () => import("../pages/home/HomePage.vue"),
    meta: { roles: ['2'] }
  },
  {
    path: "/product",
    name: "ProductsPage",
    component: () => import("../pages/products/ProductsPage.vue"),
    meta: { roles: ['2'] }
  },
  {
    path: "/product/:id",
    name: "ProductDetail",
    component: () => import("../pages/product-detail/ProductDetail.vue"),
    meta: { roles: ['2'] }
  },
  {
    path: "/cart",
    name: "Cart",
    component: () => import("../pages/carts/Cart.vue"),
    meta: { roles: ['2'] }
  },

  {
    path: "/about",
    name: "AboutPage",
    component: () => import("../pages/about/AboutPage.vue"),
    meta: { roles: ['2'] }
  },

  {
    path: "/login",
    name: "LoginPage",
    component: () => import("../pages/login/LoginPage.vue"),
    meta: { roles: ['2', '1'] }
  },
  {
    path: "/register",
    name: "RegisterPage",
    component: () => import("../pages/register/RegisterPage.vue"),
    meta: { roles: ['2'] }
  },

  {
    path: "/payment",
    name: "paymentPage",
    component: () => import("../pages/payment/payemntPage.vue"),
    meta: { requiresAuth: true, roles: ['2'] }
  },

  {
    path: "/confirm",
    name: "ConfirmCheckout",
    component: () => import("../pages/confirm/confirmPage.vue"),
    meta: { requiresAuth: true, roles: ['2'] }
  },

  {
    path: "/user",
    name: "UserPage",
    component: () => import("../pages/user/userPage.vue"),
    meta: { requiresAuth: true, roles: [2] }
  },
  {
    path: "/admin",
    name: "MenuAdmin",
    component: () => import("../pages/admin/MenuAdmin.vue"),
    meta: { requiresAuth: true, roles: [1] }

  },
];

const router = createRouter({
  history: createWebHistory(),
  scrollBehavior() {
    return { top: 0 };
  },
  routes,
});

router.beforeEach(async (to, from, next) => {
    if (to.meta.requiresAuth) { 
      if (store.state.users.user.role === 2 ) {
        next()
      }
      if (store.state.users.user.role === 1 ) {
        next()
      }
    } else {
      next()
    }

  
    // if (to.meta.requiresAuth) {
    //   next('/login');
    //   return;
    // } else {
    //   next();
    //   return;
    // }

  // if (to.meta.requiresAuth) {
  //   let role = store.state.users.user.role;
  //   if (to.meta.roles.includes(role)) {
  //     next();
  //     return;
  //   } else {
  //     next('/login');
  //     return;
  //   }
  // } else {
  //   let role = store.state.users ? store.state.users.user.role : false;
  //   if (!role) {
  //     next();
  //     return;
  //   } else {
  //     if (to.meta.roles.includes(role)) {
  //       next();
  //       return;
  //     } else {
  //       next('/login');
  //       return;
  //     }
  //   }
  // }
});

export default router;
