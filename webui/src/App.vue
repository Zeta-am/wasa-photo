<script>
import UploadPhotoModal from '@/components/UploadPhotoModal.vue'
import SearchModal from './components/SearchModal.vue';
import Dashboard from '@/components/Dashboard.vue';

export default {
  components: {
    UploadPhotoModal,
    SearchModal,
    Dashboard
  },
  data() {
    return {
      logged: null,
      currentUsername: null,
      usrLink: null,
    }
  },
  watch: {
    $route(to, from) {
      this.checkLocalStorage();
    }
  },
  methods: {
    logout() {
      window.localStorage.removeItem("token");
      window.localStorage.removeItem("username");
      this.$setAuth();
      this.$router.push({ name: 'Login' });
    },
    checkLocalStorage() {
      const storedUsr = window.localStorage.getItem('username');
      const storedTkn = window.localStorage.getItem('token');

      if (storedUsr && storedTkn) {
        try {
          this.currentUsername = storedUsr;
          this.logged = true;
          this.usrLink = `/users/${storedTkn}`;
        } catch (error) {
          this.handleLogout();
        }
      } else {
        this.logged = false;
        this.$router.push({ name: 'Login' });
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
      this.$refs.uploadModal.open();
    },
    openSearchModal() {
      this.$refs.searchModal.open();
    },
  },
  mounted() {
    this.$setAuth();
    this.checkLocalStorage();

    this.$axios.interceptors.response.use(response => {
      return response;
    }, error => {
      if (error.response.status === 401) {
        this.$router.push({ name: 'Login' });
        return;
      } else {
        return Promise.reject(error);
      }
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
      @open-search="openSearchModal" 
    />
    <main :class="['px-md-4', 'w-100']" :style="{ paddingLeft: logged ? '60px' : '0' }">
      <RouterView />
    </main>
    <UploadPhotoModal ref="uploadModal" />
    <SearchModal ref="searchModal" />
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
