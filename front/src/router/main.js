import { createRouter, createWebHistory } from 'vue-router';
import store from '../store/index.js';
import { getUserByCookie } from '../services/users.service.js';

const routes = [
  // Route for shop
  {
    path: '/:pathMatch(.*)*',
    name: '404-page',
    component: () => import('../pages/404Page'),
    meta: { roles: ['customer', 'admin'] }
  },
  {
    path: '/',
    name: 'home-page',
    component: () => import('../pages/HomePage'),
    meta: { roles: ['customer'] }
  },
  {
    path: '/products',
    name: 'products-page',
    component: () => import('../pages/ProductsCategory'),
    meta: { roles: ['customer'] }
  },
  {
    path: '/products/:slug',
    name: 'products-detail-page',
    component: () => import('../pages/ProductDetail'),
    meta: { roles: ['customer'] }
  },
  {
    path: '/shopping-cart',
    name: 'shopping-cart-page',
    component: () => import('../pages/ShoppingCart'),
    meta: { requiresAuth: true, roles: ['customer'] }
  },
  {
    path: '/categories/:slug',
    name: 'products-category',
    component: () => import('../pages/ProductsCategory'),
    meta: { roles: ['customer'] }
  },
  {
    path: '/login',
    name: 'login',
    component: () => import('../pages/LoginPage'),
    meta: { roles: ['customer', 'admin'] }
  },
  {
    path: '/register',
    name: 'register',
    component: () => import('../pages/RegisterPage'),
    meta: { roles: ['customer'] }
  },
  {
    path: '/products/search',
    name: 'search',
    component: () => import('../pages/SearchPage'),
    meta: { roles: ['customer'] }
  },
  {
    path: '/bill',
    name: 'bill',
    component: () => import('../pages/BillPage'),
    meta: { requiresAuth: true, roles: ['customer'] }
  },
  {
    path: '/bill-detail/:id',
    name: 'bill-detail',
    component: () => import('../pages/BillDetailPage'),
    meta: { requiresAuth: true, roles: ['customer'] }
  },
  // Route for admin
  {
    path: '/dashboard',
    name: 'dashboard',
    component: () => import('../admin/pages/Dashboard'),
    meta: { requiresAuth: true, roles: ['admin'] }
  },
  {
    path: '/admin/product',
    name: 'admin-product',
    component: () => import('../admin/pages/PageProduct'),
    meta: { requiresAuth: true, roles: ['admin'] }
  },
  {
    path: '/admin/order',
    name: 'admin-order',
    component: () => import('../admin/pages/PageOrder'),
    meta: { requiresAuth: true, roles: ['admin'] }
  }
];

const router = createRouter({
  history: createWebHistory(),
  scrollBehavior() {
    return { top: 0 };
  },
  routes
});

router.beforeEach(async (to, from, next) => {
  try {
    if (!store.getters.isAuthenticated) {
      const res = await getUserByCookie();
      store.dispatch('setUser', res.data);
      if (res.data.role === 'customer') {
        store.dispatch('getCartProduct');
        store.commit('updateLayout', 'LayoutShop');
      }
      if (res.data.role === 'admin') {
        store.commit('updateLayout', 'LayoutAdmin');
      }
    }
  } catch (error) {
    if (to.meta.requiresAuth) {
      next('/login');
      return;
    } else {
      next();
      return;
    }
  }

  if (to.meta.requiresAuth) {
    let role = store.state.user.role;
    if (to.meta.roles.includes(role)) {
      next();
      return;
    } else {
      next('/login');
      return;
    }
  } else {
    let role = store.state.user ? store.state.user.role : false;
    if (!role) {
      next();
      return;
    } else {
      if (to.meta.roles.includes(role)) {
        next();
        return;
      } else {
        next('/login');
        return;
      }
    }
  }
});

export default router;
