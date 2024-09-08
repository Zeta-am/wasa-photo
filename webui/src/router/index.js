import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/login', component: LoginView, name: 'Login'},
		{path: '/home', component: HomeView, name: 'Home'},
		{path: '/users/:userId', component: ProfileView, name: 'Profile'},
		{path: '/', redirect: '/login'}
	]
})

router.beforeEach((to, from, next) => {
	if (to.path !== '/login' && !localStorage.token) next({ path: '/login' })
	else next()
})

export default router
