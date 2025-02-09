<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Dashboard from '@/components/Dashboard.vue'
</script>
<script>
export default {
  components: {
    Dashboard
  },
	data: function() {
		return {
			logged: null,
			currentUsername: null,
			usrLink: null,
		}
	},
	watch: {
		$route (to, from) {
			this.checkLocalStorage();
		}
	},
	methods: {
		logout() {
			localStorage.removeItem("token");
			localStorage.removeItem("username");
			this.$setAuth();
			this.$router.push({name: 'Login'})
		},
		checkLocalStorage() {
			const storedUsr = localStorage.getItem('username');
			const storedTkn = localStorage.getItem('token');

			if (storedUsr && storedTkn) {
				try {
					this.currentUsername = storedUsr;
					this.logged = true;  // Make sure this is set
					this.usrLink = `/users/${storedTkn}`; // Use token instead of username
				} catch (error) {
					this.handleLogout();
				}
			} else {
				this.logged = false;
			}
		}
	},
	mounted() {
		this.$setAuth();
		this.checkLocalStorage();

		this.$axios.interceptors.response.use(response => {
			return response;
		}, error => {
				// If the user is Unauthorized, redirect to login
			if (error.response.status === 401) {
				this.$router.push({ name: 'Login' })
				return;
			}
			else 
				return Promise.reject(error) // Leave other error handlers
		});
	}
}
</script>

<template>
  <div class="container-fluid">
    <div class="row g-0">
      <Dashboard 
        v-if="logged"
        :username="currentUsername" 
        :usrLink="`/users/${localStorage.getItem('token')}`"
        @logout="logout"
        class="col-md-3 col-lg-2 d-md-block"
      />
      <main :class="[
        logged ? 'col-md-9 ms-sm-auto col-lg-10' : 'col-12',
        'px-md-4'
      ]">
        <RouterView />
      </main>
    </div>
  </div>
</template>

<style>
#app {
  min-height: 100vh;
}

.container-fluid {
  padding: 0;
}

.min-vh-100 {
  min-height: 100vh;
}
</style>
