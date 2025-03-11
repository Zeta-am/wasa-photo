<template>
  <!-- Mostra la modale solo se "show" è true -->
  <div v-if="show" class="modal-overlay">
    <div class="modal-content">
      <!-- Titolo stile Instagram -->
      <h3>Create new post</h3>

      <form @submit.prevent="uploadPhoto">
        <!-- Visualizza la drag area solo se non è stata selezionata un'immagine -->
        <div v-if="!selectedFile" class="drag-area"
          :class="{ 'drag-over': isDragOver }"
          @dragenter.prevent="isDragOver = true"
          @dragover.prevent="isDragOver = true"
          @dragleave.prevent="isDragOver = false"
          @drop.prevent="onDropFile"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="60"
            height="60"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              fill="#8e8e8e"
              d="M12 7a5 5 0 100 10 5 5 0 000-10zm0-5c1.2 0 2.1.6 2.7 1.7l.6 1.3h3.2c1.6 0 2.8 1.3 2.8 2.8v9.4c0 1.6-1.3 2.8-2.8 2.8H5.3c-1.6 0-2.8-1.3-2.8-2.8V7.8c0-1.6 1.3-2.8 2.8-2.8h3.2l.6-1.3C9.9 2.6 10.8 2 12 2z"
            />
          </svg>

          <!-- Testo e pulsante di selezione file -->
          <p>Drag photos and videos here</p>
          <button
            type="button"
            class="btn-select"
            @click="triggerFileInput"
          >
            Select from computer
          </button>
        </div>

        <!-- Input file nascosto (utilizzato sia dalla drag area che dal pulsante) -->
        <input
          ref="photoInput"
          type="file"
          accept="image/*"
          style="display: none;"
          @change="handlePhotoSelect"
        />

        <!-- Visualizza anteprima e caption solo se è stata caricata una foto -->
        <div v-if="selectedFile">
          <!-- Anteprima della foto selezionata -->
          <div class="preview" style="margin-bottom: 16px; text-align: center;">
            <img :src="imagePreview" alt="Photo preview" style="max-width: 100%; max-height: 200px;" />
          </div>

          <!-- Campo per la caption -->
          <div class="form-group mb-3">
            <input
              v-model="photoCaption"
              type="text"
              class="form-control"
              placeholder="Add a caption (max 200 characters)"
              maxlength="200"
            />
          </div>
        </div>

        <!-- Pulsanti di azione -->
        <div class="buttons">
          <button
            type="button"
            class="btn btn-secondary me-2"
            @click="close"
          >
            Cancel
          </button>
          <!-- Il pulsante Share compare solo se è stata caricata un'immagine -->
          <button
            v-if="selectedFile"
            type="submit"
            class="btn btn-primary"
            :disabled="isLoading"
          >
            {{ isLoading ? 'Uploading...' : 'Share' }}
          </button>
        </div>
      </form>

      <!-- Eventuale messaggio di errore -->
      <div v-if="error" class="alert alert-danger mt-3">{{ error }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'UploadPhotoModal',
  data() {
    return {
      show: false,
      photoCaption: '',
      selectedFile: null,
      imagePreview: null,
      isLoading: false,
      error: null,
      isDragOver: false,
    }
  },
  methods: {
    onDropFile(e) {
      this.isDragOver = false;
      const file = e.dataTransfer.files[0];
      if (!file) return;
      if (!file.type.startsWith('image/')) {
        this.error = 'Please drop a valid image file';
        return;
      }
      this.selectedFile = file;
      this.imagePreview = URL.createObjectURL(file);
    },
    triggerFileInput() {
      this.$refs.photoInput.click();
    },
    handlePhotoSelect(event) {
      const file = event.target.files[0];
      if (file) {
        this.selectedFile = file;
        this.imagePreview = URL.createObjectURL(file);
      }
    },
    async uploadPhoto() {
      if (!this.selectedFile) {
        this.error = 'Please select a photo';
        return;
      }
      this.isLoading = true;
      this.error = null;
      try {
        const formData = new FormData();
        formData.append('image', this.selectedFile);
        const userId = this.$utils.getCurrentId();
        await this.$axios.post(
          `/users/${userId}/posts?caption=${this.photoCaption}`,
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data'
            }
          }
        );
        this.close();
        this.$router.go();
      } catch (e) {
        this.error = e.response?.data || 'Error uploading photo';
      } finally {
        this.isLoading = false;
      }
    },
    open() {
      this.show = true;
      this.error = null;
      this.photoCaption = '';
      this.selectedFile = null;
      this.imagePreview = null;
      if (this.$refs.photoInput) {
        this.$refs.photoInput.value = '';
      }
    },
    close() {
      this.show = false;
    }
  }
}
</script>

<style scoped>
/* Sfondo scuro semi-trasparente */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 2000;
}

/* Contenitore principale della modale */
.modal-content {
  background: white;
  border-radius: 12px;
  width: 90%;
  max-width: 500px;
  padding: 0;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

/* Titolo in alto, stile Instagram */
h3 {
  margin: 0;
  padding: 14px 16px;
  border-bottom: 1px solid #dbdbdb;
  font-size: 16px;
  font-weight: 600;
  text-align: center;
}

/* Form con un po' di padding */
form {
  padding: 16px;
}

/* Area di drag & drop */
.drag-area {
  border: 2px dashed #dbdbdb;
  border-radius: 8px;
  text-align: center;
  padding: 30px 0;
  margin-bottom: 16px;
  transition: background-color 0.2s;
}

.drag-area.drag-over {
  background-color: #f0f0f0;
}

/* Testo e icona al centro */
.drag-area svg {
  margin-bottom: 8px;
}

.drag-area p {
  margin: 0 0 16px;
  color: #8e8e8e;
  font-size: 14px;
}

/* Pulsante per selezionare file */
.btn-select {
  background-color: #0095f6;
  color: #fff;
  border: none;
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.btn-select:hover {
  background-color: #1877f2;
}

/* Gruppo per la caption */
.form-group {
  margin-bottom: 16px;
}

/* Input di testo */
.form-control {
  border: 1px solid #dbdbdb;
  border-radius: 4px;
  padding: 8px 12px;
  width: 100%;
}

/* Barra di pulsanti in basso */
.buttons {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  padding: 8px 16px;
  border-top: 1px solid #dbdbdb;
}

/* Stile base pulsanti */
.btn {
  padding: 8px 16px;
  border-radius: 8px;
  font-weight: 600;
  font-size: 14px;
  cursor: pointer;
}

/* Pulsante "Share" */
.btn-primary {
  background: #0095f6;
  border: none;
  color: white;
}

.btn-primary:hover {
  background: #1877f2;
}

/* Pulsante "Cancel" */
.btn-secondary {
  background: none;
  border: none;
  color: #262626;
}

/* Alert di errore */
.alert {
  margin: 0;
  padding: 12px 16px;
  border-top: 1px solid #dbdbdb;
  color: #d32f2f;
}
</style>
