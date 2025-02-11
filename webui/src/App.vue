<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Dashboard from '@/components/Dashboard.vue'
import UploadPhotoModal from '@/components/UploadPhotoModal.vue'
</script>
<script>
export default {
  components: {
    Dashboard,
    UploadPhotoModal
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
			window.localStorage.removeItem("token");
			window.localStorage.removeItem("username");
			this.$setAuth();
			this.$router.push({name: 'Login'})
		},
		checkLocalStorage() {
			const storedUsr = window.localStorage.getItem('username');
			const storedTkn = window.localStorage.getItem('token');

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
		},
		handleLogout() {
			window.localStorage.removeItem('token');
			window.localStorage.removeItem('username');
			this.logged = false;
			this.currentUsername = null;
			this.usrLink = null;
			this.$setAuth();
			this.$router.push({ name: 'Login' });
		},
    openUploadModal() {
      this.$refs.uploadModal.open()
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
  <div class="container-fluid h-100">
    <Dashboard 
      v-if="logged"
      :username="currentUsername" 
      :usrLink="usrLink"
      @logout="logout"
      @open-upload="openUploadModal"
    />
    <main :class="['px-md-4', 'w-100']">
      <RouterView />
    </main>
    <UploadPhotoModal ref="uploadModal" />
  </div>
</template>

<style>
html, body, #app {
  height: 100%;
}

.container-fluid {
  padding: 0;
}

main {
  min-height: 100vh;
}
</style>
