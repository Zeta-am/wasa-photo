<script>
import SearchResult from './SearchResult.vue'

export default {
  name: 'SearchModal',
  components: {
    SearchResult
  },
  data() {
    return {
      show: false,
      searchQuery: '',
      results: [], // Assicurati sia inizializzato come array
      loading: false,
      timeout: null
    }
  },
  methods: {
    open() {
      this.show = true
    },
    close() {
      this.show = false
      this.searchQuery = ''
      this.results = [] // Reset sempre a un array vuoto
    },
    handleInput() {
      clearTimeout(this.timeout)
      this.timeout = setTimeout(() => {
        this.handleSearch()
      }, 300)
    },
    async handleSearch() {
      if (!this.searchQuery) {
        this.results = []
        return
      }
      
      this.loading = true
      
      try {
        const response = await this.$axios.get('/users', {
          params: {
            username: this.searchQuery
          }
        })
        // Forza il valore in results ad essere un array
        this.results = Array.isArray(response.data.users) ? response.data.users : []
      } catch (error) {
        this.results = []
      } finally {
        this.loading = false
      }
    },
    selectUser(user) {
      // Naviga al profilo e forza il refresh dei dati
      this.$router.push(`/users/${user.id}`).then(() => {
        // Forza il refresh della pagina dopo la navigazione
        window.location.reload()
      })
      this.close()
    }
  }
}
</script>

<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <h6>Search</h6>
      </div>

      <div class="modal-body">
        <!-- Search Input -->
        <div class="search-input-wrapper">
          <input 
            type="text" 
            class="form-control" 
            placeholder="Search" 
            v-model="searchQuery"
            @input="handleInput"
          >
        </div>

        <!-- Results -->
        <div class="results-container">
          <div v-if="loading" class="text-center p-3">
            Loading...
          </div>
          <div v-else-if="results.length === 0 && searchQuery" class="text-center p-3">
            No users found
          </div>
          <div v-else>
            <SearchResult 
              v-for="user in results" 
              :key="user.id"
              :user="user"
              @select="selectUser"
            />
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button class="btn-cancel" @click="close">Close</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.65);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  border-radius: 12px;
  width: 400px;
  max-height: 80vh;
  overflow: hidden;
}

.modal-header {
  padding: 12px;
  text-align: center;
  border-bottom: 1px solid #dbdbdb;
}

.modal-header h6 {
  margin: 0;
  font-weight: 600;
}

.modal-body {
  padding: 16px;
  overflow-y: auto;
  max-height: calc(80vh - 120px);
}

.search-input-wrapper {
  margin-bottom: 16px;
}

.results-container {
  max-height: 400px;
  overflow-y: auto;
}

.result-item {
  display: flex;
  align-items: center;
  padding: 8px;
  cursor: pointer;
  border-radius: 8px;
}

.result-item:hover {
  background: #fafafa;
}

.avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  margin-right: 12px;
}

.username {
  font-weight: 600;
}

.name {
  font-size: 14px;
  color: #8e8e8e;
}

.modal-footer {
  padding: 12px;
  border-top: 1px solid #dbdbdb;
  text-align: center;
}

.btn-cancel {
  width: 100%;
  padding: 8px;
  border: none;
  background: none;
  font-weight: 600;
  cursor: pointer;
}

.loading, .no-results {
  text-align: center;
  color: #8e8e8e;
  padding: 16px;
}
</style>
