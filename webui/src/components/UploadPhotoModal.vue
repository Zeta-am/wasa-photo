<script>

export default {
    name: 'UploadPhotoModal',
    data() {
    return {
        show: false,
        photoCaption: '',
        selectedFile: null,
        isLoading: false,
        error: null,
    }
    },
    methods: {
    handlePhotoSelect(event) {
        this.selectedFile = event.target.files[0]
    },
    async uploadPhoto() {
        if (!this.selectedFile) {
        this.error = 'Please select a photo'
        return
        }

        this.isLoading = true
        this.error = null

        try {
        const formData = new FormData()
        formData.append('image', this.selectedFile)
        
        const userId = this.$utils.getCurrentId()
        await this.$axios.post(`/users/${userId}/posts?caption=${this.photoCaption}`, formData, {
            headers: {
            'Content-Type': 'multipart/form-data'
            }
        })

        this.close()
        this.$router.go() // Refresh current route
        } catch (e) {
        this.error = e.response?.data || 'Error uploading photo'
        } finally {
        this.isLoading = false
        }
    },
    open() {
        this.show = true
        this.error = null
        this.photoCaption = ''
        this.selectedFile = null
        if (this.$refs.photoInput) {
        this.$refs.photoInput.value = ''
        }
    },
    close() {
        this.show = false
    }
    }
}
</script>

<template>
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <h3>Upload Photo</h3>
      <form @submit.prevent="uploadPhoto">
        <div class="form-group mb-3">
          <input 
            type="file" 
            ref="photoInput"
            @change="handlePhotoSelect"
            accept="image/*"
            class="form-control"
            required
          >
        </div>
        <div class="form-group mb-3">
          <input 
            v-model="photoCaption"
            type="text"
            class="form-control"
            placeholder="Add a caption (max 200 characters)"
            maxlength="200"
          >
        </div>
        <div class="buttons">
          <button type="button" class="btn btn-secondary me-2" @click="close">Cancel</button>
          <button type="submit" class="btn btn-primary" :disabled="isLoading">
            {{ isLoading ? 'Uploading...' : 'Share' }}
          </button>
        </div>
      </form>
      <div v-if="error" class="alert alert-danger mt-3">{{ error }}</div>
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
  background: rgba(0, 0, 0, 0.85);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 2000; /* Higher than dashboard */
}

.modal-content {
  background: white;
  border-radius: 12px;
  width: 90%;
  max-width: 500px;
  padding: 0;
  overflow: hidden;
}

h3 {
  margin: 0;
  padding: 14px 16px;
  border-bottom: 1px solid #dbdbdb;
  font-size: 16px;
  font-weight: 600;
  text-align: center;
}

form {
  padding: 16px;
}

.form-group {
  margin-bottom: 16px;
}

.form-control {
  border: 1px solid #dbdbdb;
  border-radius: 4px;
  padding: 8px 12px;
}

.buttons {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 8px 16px;
  border-top: 1px solid #dbdbdb;
}

.btn {
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  font-size: 14px;
}

.btn-primary {
  background: #0095f6;
  border: none;
  color: white;
}

.btn-primary:hover {
  background: #1877f2;
}

.btn-secondary {
  background: none;
  border: none;
  color: #262626;
}

.alert {
  margin: 0;
  padding: 12px 16px;
  border-top: 1px solid #dbdbdb;
}
</style>