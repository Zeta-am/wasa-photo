<script>
export default {
  props: {
    username: String,
    postCount: Number,
    followerCount: Number,
    followingCount: Number,
    isOwnProfile: Boolean,
    profileImage: String,
    isFollowed: Boolean,
    userId: Number,
    isBanned: Boolean
  },
  data() {
    return {
      showOptionsModal: false,
      showSettings: false,
      loading: false
    }
  },
  methods: {
    async toggleFollow() {
      if (this.loading) return
      
      this.loading = true
      try {
        const currentUserId = this.$utils.getCurrentId()
        
        if (this.isFollowed) {
          await this.$axios.delete(`/users/${currentUserId}/followings/${this.userId}`)
        } else {
          await this.$axios.put(`/users/${currentUserId}/followings/${this.userId}`)
        }
        
        await this.$emit('follow-toggled')
      } catch (e) {
        console.error('Error toggling follow:', e)
      } finally {
        this.loading = false
      }
    }
  }
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
          <!-- Own profile controls -->
          <div v-if="isOwnProfile" class="action-buttons">
            <button @click="$emit('edit-username')" class="btn btn-primary">
              Edit username
            </button>
            <div class="settings-dropdown">
              <button @click="showSettings = !showSettings" class="settings-btn">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <circle cx="12" cy="12" r="3"></circle>
                  <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
                </svg>
              </button>
              <div v-if="showSettings" class="settings-menu">
                <button @click="$emit('show-banned-list')" class="settings-item">
                  Banned Users
                </button>
              </div>
            </div>
          </div>
          <!-- Other profile controls -->
          <div v-else class="action-buttons">
            <button
              @click="toggleFollow"
              class="follow-button"
              :class="{ 
                'following': isFollowed,
                'loading': loading 
              }"
              :disabled="loading"
            >
              {{ isFollowed ? 'Following' : 'Follow' }}
            </button>
            <div class="settings-dropdown">
              <button @click="showSettings = !showSettings" class="more-options-btn">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
                  <circle cx="12" cy="6" r="2"/>
                  <circle cx="12" cy="12" r="2"/>
                </svg>
              </button>
              <div v-if="showSettings" class="settings-menu">
                <button @click="handleBan" class="settings-item text-danger">
                  Ban User
                </button>
              </div>
            </div>
          </div>
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

.follow-button {
  padding: 8px 24px;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
  background: #0095f6;
  color: white;
  transition: all 0.3s ease; /* Increased transition time */
}

.follow-button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.follow-button.following {
  background: #efefef;
  color: #262626;
}

.follow-button.loading {
  opacity: 0.7;
  cursor: wait;
}

/* Change hover state for following button */
.follow-button.following:hover {
  background: #dbdbdb;
}

/* Normal hover state */
.follow-button:not(.following):hover {
  background: #0081d6;
}

.settings-dropdown {
  position: relative;
  display: inline-block;
}

.settings-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: white;
  border: 1px solid #dbdbdb;
  border-radius: 4px;
  box-shadow: 0 0 5px rgba(0,0,0,0.1);
  z-index: 1000;
}

.settings-item {
  display: block;
  width: 100%;
  padding: 8px 16px;
  text-align: left;
  border: none;
  background: none;
  cursor: pointer;
}

.settings-item:hover {
  background: #fafafa;
}

.settings-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
  margin-left: 10px;
}

.settings-btn svg {
  width: 24px;
  height: 24px;
}

.action-buttons {
  display: flex;
  align-items: center;
  gap: 12px;
}

.more-options-btn {
  background: none;
  border: none;
  padding: 8px;
  cursor: pointer;
  display: flex;
  align-items: center;
}

.more-options-btn svg {
  fill: currentColor;
}

.text-danger {
  color: #ed4956;
}

.settings-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background: white;
  border: 1px solid #dbdbdb;
  border-radius: 4px;
  box-shadow: 0 0 5px rgba(0,0,0,0.1);
  z-index: 1000;
  min-width: 150px;
}

.settings-item {
  display: block;
  width: 100%;
  padding: 12px 16px;
  text-align: left;
  border: none;
  background: none;
  cursor: pointer;
  font-size: 14px;
}

.settings-item:hover {
  background: #fafafa;
}
</style>