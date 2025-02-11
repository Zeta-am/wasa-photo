<script>
export default {
  props: {
    username: String,
    postCount: Number,
    followerCount: Number,
    followingCount: Number,
    isOwnProfile: Boolean,
    profileImage: String
  },
  data() {
    return {
      showOptionsModal: false
    }
  },
  methods: {
    handleProfileClick() {
      if (!this.isOwnProfile) return

      if (this.profileImage) {
        this.showOptionsModal = true
      } else {
        this.$refs.profileInput.click()
      }
    },
    async handleProfileUpload(event) {
      const file = event.target.files[0]
      if (!file) return

      try {
        const formData = new FormData()
        formData.append('image', file)
        
        const userId = this.$utils.getCurrentId()
        await this.$axios.put(`/users/${userId}/profile-image`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        })
        
        this.$emit('profile-updated')
        this.closeOptionsModal()
      } catch (e) {
        console.error('Error uploading profile picture:', e)
      }
    },
    async removeProfilePhoto() {
      try {
        const userId = this.$utils.getCurrentId()
        await this.$axios.delete(`/users/${userId}/profile-image`)
        this.$emit('profile-updated')
        this.closeOptionsModal()
      } catch (e) {
        console.error('Error removing profile picture:', e)
      }
    },
    closeOptionsModal() {
      this.showOptionsModal = false
    }
  },
  emits: ['edit-username', 'show-followers', 'show-following', 'profile-updated']
}
</script>

<template>
  <div class="user-info">
    <div class="user-header">
      <!-- Profile Icon without interaction if not owner -->
      <div class="profile-icon" :class="{ 'profile-icon-interactive': isOwnProfile }" 
           @click="isOwnProfile && handleProfileClick">
        <div v-if="profileImage" class="profile-image-container">
          <img :src="'data:image/jpeg;base64,' + profileImage" alt="Profile" class="profile-image">
        </div>
        <svg v-else aria-label="Profile" class="_ab6-" color="currentColor" fill="currentColor" height="150" role="img" viewBox="0 0 24 24" width="150">
          <circle cx="12" cy="8" r="6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1"/>
          <path d="M20 21c0-4.418-3.582-8-8-8s-8 3.582-8 8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1"/>
        </svg>
        <input 
          v-if="isOwnProfile && !profileImage"
          type="file"
          ref="profileInput"
          @change="handleProfileUpload"
          accept="image/*"
          class="profile-input"
        >
      </div>

      <!-- User Info Section -->
      <div class="user-details">
        <!-- Username Row -->
        <div class="username-row">
          <h2 class="username">{{ username }}</h2>
          <!-- Only show edit button if owner -->
          <button v-if="isOwnProfile" 
                  @click="$emit('edit-username')" 
                  class="edit-button">
            Edit username
          </button>
        </div>

        <!-- Stats Row -->
        <div class="stats-row">
          <div class="stat-item">
            <span class="stat-value">{{ postCount }}</span>
            <span class="stat-label">posts</span>
          </div>
          <div class="stat-item clickable" @click="$emit('show-followers')">
            <span class="stat-value">{{ followerCount }}</span>
            <span class="stat-label">followers</span>
          </div>
          <div class="stat-item clickable" @click="$emit('show-following')">
            <span class="stat-value">{{ followingCount }}</span>
            <span class="stat-label">following</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Profile Picture Options Modal -->
    <div v-if="showOptionsModal" class="profile-options-modal" @click.self="closeOptionsModal">
      <div class="options-content">
        <h6>Change Profile Photo</h6>
        <div class="options-list">
          <label class="option-item upload">
            Upload Photo
            <input 
              type="file" 
              @change="handleProfileUpload" 
              accept="image/*"
              class="hidden-input"
            >
          </label>
          <button v-if="profileImage" class="option-item remove" @click="removeProfilePhoto">
            Remove Current Photo
          </button>
          <button class="option-item cancel" @click="closeOptionsModal">
            Cancel
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.user-info {
  padding: 30px 20px;
  border-bottom: 1px solid #dbdbdb;
}

.user-header {
  display: flex;
  align-items: flex-start;
  gap: 80px;
}

.profile-icon {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid #dbdbdb;
  flex-shrink: 0;
  position: relative;
  overflow: hidden;
}

.profile-icon-interactive {
  cursor: pointer;
}

.profile-icon-interactive:hover::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.3);
  border-radius: 50%;
}

.profile-image-container {
  width: 100%;
  height: 100%;
}

.profile-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 50%;
}

.profile-input {
  display: none;
}

.profile-options-modal {
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

.options-content {
  background: white;
  border-radius: 12px;
  width: 400px;
  overflow: hidden;
}

.options-content h6 {
  text-align: center;
  padding: 18px;
  margin: 0;
  border-bottom: 1px solid #dbdbdb;
  font-weight: 600;
}

.options-list {
  display: flex;
  flex-direction: column;
}

.option-item {
  padding: 14px;
  text-align: center;
  border: none;
  background: none;
  border-bottom: 1px solid #dbdbdb;
  font-size: 14px;
  cursor: pointer;
  width: 100%;
}

.option-item.upload {
  color: #0095f6;
  font-weight: 600;
}

.option-item.remove {
  color: #ed4956;
  font-weight: 600;
}

.hidden-input {
  display: none;
}

.user-details {
  flex-grow: 1;
}

.username-row {
  display: flex;
  align-items: center;
  gap: 20px;
  margin-bottom: 20px;
}

.username {
  font-size: 28px;
  font-weight: 300;
  margin: 0;
}

.edit-button {
  padding: 8px 16px;
  border: 1px solid #dbdbdb;
  background: none;
  border-radius: 4px;
  font-weight: 600;
  cursor: pointer;
}

.edit-button:hover {
  background: #fafafa;
}

.stats-row {
  display: flex;
  gap: 40px;
}

.stat-item {
  display: flex;
  gap: 5px;
}

.stat-item.clickable {
  cursor: pointer;
}

.stat-item.clickable:hover .stat-value,
.stat-item.clickable:hover .stat-label {
  color: #8e8e8e;
}

.stat-value {
  font-weight: 600;
}

.stat-label {
  color: #262626;
}
</style>