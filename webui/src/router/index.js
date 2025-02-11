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
  const token = localStorage.getItem('token')
  const username = localStorage.getItem('username')
  
  // Require both token and username for auth
  const isAuthenticated = token && username
  
  if (to.meta.requiresAuth && !isAuthenticated) {
    // Clear potentially corrupt storage
    localStorage.removeItem('token')
    localStorage.removeItem('username')
    next({ name: 'Login' })
  } else if (to.name === 'Login' && isAuthenticated) {
    // Redirect to home if already logged in
    next({ name: 'Home' })
  } else {
    next()
  }
})

export default router
