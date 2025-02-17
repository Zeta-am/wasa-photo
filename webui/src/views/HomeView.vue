<template>
  <main class="container h-100 d-flex align-items-center justify-content-center">
    <div class="text-center">
      <ErrorMsg v-if="errormsg" :msg="errormsg"/>
      
      <div class="mt-4">
        <p class="text-muted">
          Non ci sono post da visualizzare. 
          <br>
          Questo può essere dovuto al fatto che non segui ancora nessuno o che le persone che segui non hanno ancora pubblicato contenuti.
        </p>
        <button @click="openSearch" class="btn btn-primary mt-3">
          Cerca persone da seguire
        </button>
      </div>
    </div>
    <SearchModal ref="searchModal" />
    <div class="home-container">
      <Dashboard 
        :username="username"
        :usrLink="usrLink"
        @logout="logout"
        @open-upload="$parent.openUploadModal"
        @open-search="$parent.openSearchModal"
      />
      <!-- resto del contenuto della home -->
    </div>
  </main>
</template>

<script>
import SearchModal from '@/components/SearchModal.vue'
import Dashboard from '@/components/Dashboard.vue'

export default {
  components: {
    SearchModal,
    Dashboard
  },
  data() {
    return {
      errormsg: null
    }
  },
  created() {
    // Reset any stored user ID when entering home view
    localStorage.removeItem('lastVisitedProfile')
  },
  computed: {
    username() {
      return localStorage.getItem('username')
    },
    usrLink() {
      return `/users/${localStorage.getItem('token')}`
    }
  },
  methods: {
    openSearch() {
      this.$refs.searchModal.open()
    },
    logout() {
      localStorage.removeItem('token')
      localStorage.removeItem('username')
      this.$setAuth()
      this.$router.push({ name: 'Login' })
    }
  }
}
</script>

<style scoped>
main {
  margin: 0 !important;
  padding: 2rem !important;
}

.home-container, .profile-container {
  padding-left: 60px; /* larghezza della dashboard */
}
</style>
