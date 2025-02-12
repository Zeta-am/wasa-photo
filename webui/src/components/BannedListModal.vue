<!-- BannedListModal.vue -->
<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <h6 class="text-danger">Banned Users</h6>
        <!-- Removed close button -->
      </div>
      
      <div class="modal-body">
        <div v-if="loading" class="text-center p-3">
          Loading...
        </div>
        <div v-else-if="error" class="text-center p-3 text-danger">
          {{ error }}
        </div>
        <div v-else-if="users.length === 0" class="text-center p-3">
          No banned users
        </div>
        <div v-else class="banned-list">
          <div v-for="user in users" :key="user.id" class="banned-item">
            <div class="user-info">
              <img :src="user.profileImage || 'default-avatar.png'" class="avatar" alt="Profile picture">
              <span class="username">{{ user.username }}</span>
            </div>
            <button 
              @click="unbanUser(user.id)" 
              class="unban-btn"
              :disabled="unbanLoading === user.id"
            >
              {{ unbanLoading === user.id ? 'Unbanning...' : 'Unban' }}
            </button>
          </div>
        </div>
      </div>

      <div class="modal-footer">
        <button class="btn-cancel" @click="close">Close</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      show: false,
      loading: false,
      error: null,
      users: [],
      unbanLoading: null
    }
  },
  methods: {
    async open() {
      this.show = true
      this.users = [] // Reset users when opening
      await this.loadBannedUsers()
    },
    close() {
      this.show = false
      this.error = null
      this.users = [] // Reset on close
    },
    async loadBannedUsers() {
      this.loading = true
      this.error = null
      
      try {
        const userId = this.$utils.getCurrentId()
        const response = await this.$axios.get(`/users/${userId}/banList`)
        
        console.log('API Response:', response.data) // Debug log
        
        // The API likely returns the array directly, not nested under 'users'
        this.users = Array.isArray(response.data) ? response.data : []
        
      } catch (error) {
        this.error = 'Error loading banned users'
        this.users = []
        console.error('Error loading banned users:', error)
      } finally {
        this.loading = false
      }
    },
    async unbanUser(bannedId) {
      if (!bannedId) return
      
      this.unbanLoading = bannedId
      try {
        const userId = this.$utils.getCurrentId()
        await this.$axios.delete(`/users/${userId}/banList/${bannedId}`)
        await this.loadBannedUsers() // Reload the full list after unban
      } catch (error) {
        this.error = 'Error unbanning user'
        console.error('Error unbanning user:', error)
      } finally {
        this.unbanLoading = null
      }
    }
  }
}
</script>

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
  color: #dc3545;
}

.modal-body {
  padding: 0;
  overflow-y: auto;
  max-height: calc(80vh - 120px);
}

.banned-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  border-bottom: 1px solid #efefef;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
}

.username {
  font-weight: 600;
}

.unban-btn {
  padding: 8px 16px;
  border: none;
  background: #efefef;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  color: #262626;
}

.unban-btn:hover {
  background: #dbdbdb;
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
</style>
