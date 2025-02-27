<script>
import UserInfo from '@/components/UserInfo.vue'
import PhotoGrid from '@/components/PhotoGrid.vue'
import ChangeUsernameModal from '@/components/ChangeUsernameModal.vue'
import BannedListModal from '@/components/BannedListModal.vue'
import Dashboard from '@/components/Dashboard.vue'

export default {
  name: 'ProfileView',
  components: {
    UserInfo,
    PhotoGrid,
    ChangeUsernameModal,
    BannedListModal,
    Dashboard
  },
  data() {
    return {
      profile: null,
      photos: [],
      loading: false,
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
  watch: {
    '$route.params.userId': {
      handler() {
        this.loadProfileData()
      },
      immediate: true
    }
  },
  computed: {
    username() {
      return localStorage.getItem('username')
    },
    usrLink() {
      return `/users/${localStorage.getItem('token')}`
    }
  },
  methods: {
    async loadProfileData() {
      this.loading = true
      this.error = null
      
      try {
        const userId = this.$route.params.userId
        if (!userId || this.$route.name === 'home') {
          return
        }

        const [profileResponse, photosResponse] = await Promise.all([
          this.$axios.get(`/users/${userId}`),
          this.$axios.get(`/users/${userId}/posts`)
        ])

        this.profile = {
          ...profileResponse.data,
          isFollowed: profileResponse.data.followed
        }
        this.photos = Array.isArray(photosResponse.data) ? photosResponse.data : []
        this.isOwnProfile = userId === this.$utils.getCurrentId().toString()

        localStorage.setItem('lastVisitedProfile', userId)
      } catch (e) {
        console.error('Profile load error:', e)
        this.error = e.response ? e.response.data : 'Error loading profile'
        this.photos = []
      } finally {
        this.loading = false
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
    },

    handlePhotoDeleted() {
      this.selectedPhoto = null;
      this.loadProfileData(); // Refresh photos after deletion
    },

    logout() {
      localStorage.removeItem('token')
      localStorage.removeItem('username')
      this.$setAuth()
      this.$router.push({ name: 'Login' })
    },

    beforeRouteLeave(to, from, next) {
      if (to.name === 'home') {
        localStorage.removeItem('lastVisitedProfile')
      }
      next()
    },

    // updateFollowState(newFollowState, followerChange) {
    //   this.profile.isFollowed = newFollowState;
    //   this.profile.followerCount += followerChange;
    //   this.loadProfileData();
    // },

    async followUser() {
      try {
        const userId = this.$route.params.userId
        await this.$axios.put(`/users/${this.$utils.getCurrentId()}/followings/${userId}`)
        this.profile.followerNo++
        this.profile.isFollowed = true
      } catch (e) {
        this.error = e.response?.data || 'Error following user'
        console.error('Follow error:', e);
      }
    },

    async unfollowUser() {
      try {
        const userId = this.$route.params.userId
        await this.$axios.delete(`/users/${this.$utils.getCurrentId()}/followings/${userId}`)
        this.profile.followerNo--
        this.profile.isFollowed = false
      } catch (e) {
        this.error = e.response?.data || 'Error unfollowing user'
        console.error('Unfollow error:', e);
      }
    },
  },
}
</script>

<template>
  <div class="profile-container">
    <div class="profile-view">
      <ErrorMsg v-if="error" :msg="error"/>
      <LoadingSpinner :loading="loading">
        <div v-if="profile" class="profile-content">
          <!-- Profile Header -->
          <UserInfo
              :username="profile.username"
              :postCount="profile.postNo"
              :followerCount="profile.followerNo"
              :followingCount="profile.followingNo"
              :isOwnProfile="isOwnProfile"
              :profileImage="profile.profileImage"
              :isFollowed="profile.isFollowed"
              :userId="profile.id"
              @edit-username="openChangeUsername"
              @show-followers="showFollowers"
              @show-following="showFollowing"
              @profile-updated="loadProfileData"
              @follow-user="followUser"
              @unfollow-user="unfollowUser"
              @show-banned-list="$refs.bannedListModal.open()"
              @toggle-ban="toggleBan"
            />

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

          <!-- Photos Grid -->
          <PhotoGrid 
            :photos="photos"
            @open-photo="openPhoto"
            v-if="photos && photos.length > 0"
            class="mt-4"
          />
          <div v-else class="text-center mt-4">
            <p class="text-muted">No photos yet</p>
          </div>
        </div>
      </LoadingSpinner>

      <!-- Photo Modal -->
      <PhotoModal 
        v-if="selectedPhoto"
        :show="!!selectedPhoto"
        :photo="selectedPhoto"
        :username="profile.username"
        :isOwner="isOwnProfile"
        @close="closePhoto"
        @photo-updated="loadProfileData"
        @photo-deleted="handlePhotoDeleted"
      />

      <ChangeUsernameModal 
        ref="changeUsernameModal"
        @username-changed="handleUsernameChanged"
      />

      <BannedListModal ref="bannedListModal" />
    </div>
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

.home-container, .profile-container {
  padding-left: 60px; /* larghezza della dashboard */
}
</style>