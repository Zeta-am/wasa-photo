<!-- src/components/ChangeUsernameModal.vue -->
<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <h6>Change username</h6>
      </div>
      <div class="modal-body">
        <input 
          v-model="newUsername" 
          type="text"
          class="form-control"
          placeholder="New username"
          pattern="^[a-zA-Z0-9_$]{3,16}$"
          required
        >
        <small class="text-muted">3-16 characters, letters, numbers, _ or $</small>
      </div>
      <div class="modal-footer">
        <button class="btn-cancel" @click="close">Cancel</button>
        <button class="btn-done" @click="changeUsername" :disabled="isLoading">
          {{ isLoading ? 'Changing...' : 'Done' }}
        </button>
      </div>
      <div v-if="error" class="error-message">{{ error }}</div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      show: false,
      newUsername: '',
      error: null,
      isLoading: false
    }
  },
  methods: {
    async changeUsername() {
      if (!this.newUsername) return;
      this.isLoading = true;
      this.error = null;
      
      try {
        if (this.newUsername === localStorage.getItem('username')) {
          this.error = "New username must be different from current one";
          return;
        }

        const userId = this.$utils.getCurrentId();
        const response = await this.$axios.put(`/users/${userId}`, {
          username: this.newUsername
        });

        if (response.status === 200) {
          localStorage.setItem('username', this.newUsername);
          this.$emit('username-changed', this.newUsername);
          this.close();
          this.$router.go();
        }
      } catch (e) {
        this.error = e.response?.status === 400 
          ? "Username already exists or is invalid"
          : e.response?.data || "Error changing username";
      } finally {
        this.isLoading = false;
      }
    },
    open() {
      this.show = true;
      this.error = null;
      this.newUsername = '';
    },
    close() {
      this.show = false;
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
  padding: 24px;
}

.form-control {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #dbdbdb;
  border-radius: 4px;
  margin-bottom: 8px;
}

.modal-footer {
  display: flex;
  border-top: 1px solid #dbdbdb;
}

.btn-cancel, .btn-done {
  flex: 1;
  padding: 12px;
  border: none;
  background: none;
  font-weight: 600;
  cursor: pointer;
}

.btn-done {
  color: #0095f6;
}

.btn-done:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.error-message {
  color: #ed4956;
  padding: 12px;
  text-align: center;
  border-top: 1px solid #dbdbdb;
}
</style>