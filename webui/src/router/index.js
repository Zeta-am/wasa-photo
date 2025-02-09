import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      component: LoginView,
      name: 'Login',
      meta: { requiresAuth: false }
    },
    {
      path: '/home',
      component: HomeView,
      name: 'Home',
      meta: { requiresAuth: true }
    },
    {
      path: '/users/:userId',
      component: ProfileView,
      name: 'Profile',
      meta: { requiresAuth: true }
    },
    {
      path: '/',
      redirect: '/login'
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth && !localStorage.token) {
    next({ name: 'Login' })
  } else {
    next()
  }
})

export default router
