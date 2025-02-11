<script>
import UserInfo from '@/components/UserInfo.vue'
import PhotoGrid from '@/components/PhotoGrid.vue'
import ChangeUsernameModal from '@/components/ChangeUsernameModal.vue'

export default {
  name: 'ProfileView',
  components: {
    UserInfo,
    PhotoGrid,
    ChangeUsernameModal
  },
  data() {
    return {
      profile: null,
      photos: [],
      loading: true,
      error: null,
      isOwnProfile: false,
      photoCaption: '',
      selectedPhoto: null,
      showList: false,
      listTitle: '',
      userList: [],
      selectedFile: null
    }
  },
  created() {
    this.loadProfileData()
  },
  methods: {
    async loadProfileData() {
      this.loading = true;
      this.error = null;
      
      try {
        const userId = this.$route.params.userId;
        const currentId = this.$utils.getCurrentId();
        
        this.isOwnProfile = Number(userId) === Number(currentId);

        // Set auth header before making requests
        this.$setAuth();

        const [profileRes, photosRes] = await Promise.all([
          this.$axios.get(`/users/${userId}`),
          this.$axios.get(`/users/${userId}/posts`)
        ]);

        this.profile = profileRes.data;
        this.photos = photosRes.data;
      } catch (e) {
        if (e.response?.status === 401) {
          this.$router.push('/login');
        } else {
          this.error = e.response?.data || 'Error loading profile';
        }
      } finally {
        this.loading = false;
      }
    },

    handlePhotoSelect(event) {
      this.selectedFile = event.target.files[0]
    },

    async showFollowers() {
      try {
        const userId = this.$route.params.userId
        const response = await this.$axios.get(`/users/${userId}/followers`)
        this.userList = response.data
        this.listTitle = 'Followers'
        this.showList = true
      } catch (e) {
        this.error = e.response?.data || 'Error loading followers'
      }
    },

    async showFollowing() {
      try {
        const userId = this.$route.params.userId
        const response = await this.$axios.get(`/users/${userId}/followings`)
        this.userList = response.data
        this.listTitle = 'Following'
        this.showList = true
      } catch (e) {
        this.error = e.response?.data || 'Error loading following'
      }
    },

    async showBanned() {
      try {
        const userId = this.$route.params.userId
        const response = await this.$axios.get(`/users/${userId}/banList`)
        this.userList = response.data
        this.listTitle = 'Banned Users'
        this.showList = true
      } catch (e) {
        this.error = e.response?.data || 'Error loading banned users'
      }
    },

    async toggleBan(user) {
      try {
        const userId = this.$utils.getCurrentId()
        if (user.banned) {
          await this.$axios.delete(`/users/${userId}/banList/${user.id}`)
        } else {
          await this.$axios.put(`/users/${userId}/banList/${user.id}`)
        }
        await this.loadProfileData()
      } catch (e) {
        this.error = e.response?.data || 'Error toggling ban'
      }
    },

    closeList() {
      this.showList = false
      this.userList = []
    },

    openPhoto(photo) {
      this.selectedPhoto = photo
    },

    closePhoto() {
      this.selectedPhoto = null
    },

    async handleUsernameChanged(newUsername) {
      try {
        const userId = this.$utils.getCurrentId()
        await this.$axios.put(`/users/${userId}`, {
          username: newUsername
        })
        
        this.profile.username = newUsername
        localStorage.setItem('username', newUsername)
        await this.loadProfileData()
      } catch (e) {
        this.error = e.response?.data || 'Error updating username'
      }
    },

    openChangeUsername() {
      this.$refs.changeUsernameModal.open()
    }
  }
}
</script>

<template>
  <div class="profile-view">
    <ErrorMsg v-if="error" :msg="error"/>
    <LoadingSpinner :loading="loading">
      <div v-if="profile" class="profile-content">
        <!-- Profile Header -->
        <div class="profile-header">
          <UserInfo
            :username="profile.username"
            :postCount="profile.postNo"
            :followerCount="profile.followerNo"
            :followingCount="profile.followingNo"
            :isOwnProfile="isOwnProfile"
            :profileImage="profile.profileImage"
            @edit-username="openChangeUsername"
            @show-followers="showFollowers"
            @show-following="showFollowing"
            @profile-updated="loadProfileData"
          />
        </div>

        <!-- Lists Modals -->
        <div v-if="showList" class="user-list-modal">
          <div class="modal-content">
            <h3>{{ listTitle }}</h3>
            <ul class="list-group">
              <li v-for="user in userList" :key="user.id" class="list-group-item d-flex justify-content-between align-items-center">
                {{ user.username }}
                <button 
                  v-if="isOwnProfile && !user.banned" 
                  @click="toggleBan(user)"
                  class="btn btn-sm"
                  :class="user.banned ? 'btn-success' : 'btn-danger'"
                >
                  {{ user.banned ? 'Unban' : 'Ban' }}
                </button>
              </li>
            </ul>
            <button @click="closeList" class="btn btn-secondary mt-3">Close</button>
          </div>
        </div>

        <!-- Photo Grid -->
        <PhotoGrid 
          :photos="photos" 
          @open-photo="openPhoto" 
          class="mt-4"
        />
      </div>
    </LoadingSpinner>

    <!-- Photo Modal -->
    <div v-if="selectedPhoto" class="photo-modal">
      <div class="modal-content">
        <img :src="'data:image/jpeg;base64,' + selectedPhoto.image" :alt="selectedPhoto.caption">
        <p>{{ selectedPhoto.caption }}</p>
        <button @click="closePhoto" class="btn btn-secondary">Close</button>
      </div>
    </div>

    <ChangeUsernameModal 
      ref="changeUsernameModal"
      @username-changed="handleUsernameChanged"
    />
  </div>
</template>


<style scoped>
.profile-view {
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

.profile-content {
  background: white;
  border-radius: 8px;
  padding: 24px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.user-list-modal,
.photo-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0,0,0,0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  padding: 20px;
  border-radius: 8px;
  max-width: 500px;
  width: 90%;
  max-height: 80vh;
  overflow-y: auto;
}

.photo-modal .modal-content img {
  max-width: 100%;
  height: auto;
}

.stats-buttons {
  display: flex;
  justify-content: center;
  gap: 10px;
}

.action-buttons {
  margin: 20px 0;
  padding: 20px;
  border: 1px solid #eee;
  border-radius: 8px;
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
</style>