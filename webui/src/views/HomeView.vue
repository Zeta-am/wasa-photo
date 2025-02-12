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
  </main>
</template>

<script>
import SearchModal from '@/components/SearchModal.vue'

export default {
  components: {
    SearchModal
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
  methods: {
    openSearch() {
      this.$refs.searchModal.open()
    }
  }
}
</script>

<style scoped>
main {
  margin: 0 !important;
  padding: 2rem !important;
}
</style>
