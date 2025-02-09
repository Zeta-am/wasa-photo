<template>
  <form @submit.prevent="changeUsername" class="username-form">
    <div class="form-group">
      <input 
        v-model="newUsername" 
        type="text" 
        class="form-control"
        placeholder="New username (3-16 characters, letters, numbers, _, $)" 
        pattern="^[a-zA-Z0-9_$]{3,16}$"
        required
      >
    </div>
    <button type="submit" class="btn btn-primary" :disabled="isLoading">
      {{ isLoading ? 'Changing...' : 'Change Username' }}
    </button>
    <div v-if="error" class="alert alert-danger mt-2">{{ error }}</div>
  </form>
</template>

<script>
export default {
  data() {
    return {
      newUsername: '',
      error: null,
      isLoading: false
    }
  },
  methods: {
    async changeUsername() {
      this.isLoading = true;
      this.error = null;
      
      try {
        if (this.newUsername === localStorage.getItem('username')) {
          this.error = "New username must be different from current one";
          return;
        }

        const userId = this.$utils.getCurrentId();
        const response = await this.$axios.put(`/users/${userId}/edit`, {
          username: this.newUsername
        });

        if (response.status === 200) {
          localStorage.setItem('username', this.newUsername);
          this.$emit('username-changed', this.newUsername);
          this.newUsername = '';
          this.$router.go(); // Refresh the page to update all components
        }
      } catch (e) {
        if (e.response?.status === 400) {
          this.error = "Username already exists or is invalid";
        } else {
          this.error = e.response?.data || "Error changing username";
        }
      } finally {
        this.isLoading = false;
      }
    }
  }
}
</script>
