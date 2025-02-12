<script>
import SearchModal from './SearchModal.vue';
import UploadPhotoModal from './UploadPhotoModal.vue'

export default {
  name: 'Dashboard',
  props: {
    username: String,
    usrLink: String,
    logo: String
  },
  components: {
    UploadPhotoModal,
    SearchModal
  },
  methods: {
    openUploadModal() {
      this.$refs.uploadModal.open()
    },
    openSearchModal() {
      this.$refs.searchModal.open()
    }
  },
  emits: ['logout', 'open-upload', 'open-search']
}
</script>

<template>
  <div class="dashboard">

    <UploadPhotoModal ref="uploadModal" />
    <SearchModal ref="searchModal" />
    
    <nav id="sidebarMenu" class="sidebar">
      <div class="sidebar-content">

        <!-- Main navigation -->
        <ul class="nav flex-column align-items-center main-nav">
          <li class="nav-item">
            <RouterLink to="/home" class="nav-link" title="Home">
              <svg aria-label="Home" class="_ab6-" color="currentColor" fill="currentColor" height="24" role="img" viewBox="0 0 24 24" width="24">
                <path d="M9.005 16.545a2.997 2.997 0 0 1 2.997-2.997A2.997 2.997 0 0 1 15 16.545V22h7V11.543L12 2 2 11.543V22h7.005Z" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/>
              </svg>
            </RouterLink>
          </li>
          <li class="nav-item">
            <a class="nav-link" role="button" @click="$emit('open-search')" title="Search">
              <svg aria-label="Search" class="_ab6-" color="currentColor" fill="currentColor" height="24" role="img" viewBox="0 0 24 24" width="24">
                <path d="M19 10.5A8.5 8.5 0 1 1 10.5 2a8.5 8.5 0 0 1 8.5 8.5Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>
                <line fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" x1="16.511" x2="22" y1="16.511" y2="22"/>
              </svg>
            </a>
          </li>
          <li class="nav-item">
            <a class="nav-link" role="button" @click="$emit('open-upload')" title="Create Post">
              <svg aria-label="New post" class="_ab6-" color="currentColor" fill="currentColor" height="24" role="img" viewBox="0 0 24 24" width="24">
                <path d="M2 12v3.45c0 2.849.698 4.005 1.606 4.944.94.909 2.098 1.608 4.946 1.608h6.896c2.848 0 4.006-.7 4.946-1.608C21.302 19.455 22 18.3 22 15.45V8.552c0-2.849-.698-4.006-1.606-4.945C19.454 2.7 18.296 2 15.448 2H8.552c-2.848 0-4.006.699-4.946 1.607C2.698 4.547 2 5.703 2 8.552Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>
                <line fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" x1="6.545" x2="17.455" y1="12.001" y2="12.001"/>
                <line fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" x1="12.003" x2="12.003" y1="6.545" y2="17.455"/>
              </svg>
            </a>
          </li>
          <li class="nav-item">
            <RouterLink :to="usrLink" class="nav-link" title="Profile">
              <svg aria-label="Profile" class="_ab6-" color="currentColor" fill="currentColor" height="24" role="img" viewBox="0 0 24 24" width="24">
                <circle cx="12" cy="8" r="6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>
                <path d="M20 21c0-4.418-3.582-8-8-8s-8 3.582-8 8" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>
              </svg>
            </RouterLink>
          </li>
        </ul>

        <!-- Bottom section with logout -->
        <div class="bottom-section">
          <a class="nav-link logout-link" role="button" @click="$emit('logout')" title="Logout">
            <svg aria-label="Logout" class="_ab6-" color="currentColor" fill="currentColor" height="24" role="img" viewBox="0 0 24 24" width="24">
              <path d="M16 13v-2H7V8l-5 4 5 4v-3Z" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>
              <path d="M7 21h11a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>
            </svg>
          </a>
        </div>
      </div>
    </nav>
  </div>
</template>

<style scoped>
.dashboard {
  position: fixed;
  top: 0;
  left: 0;
  width: 60px;
  height: 100vh;
  background: rgba(255, 255, 255, 0.98);
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  backdrop-filter: blur(10px);
}

.sidebar {
  height: 100%;
  display: flex;
  flex-direction: column;
  padding: 20px 0;
}

.sidebar-content {
  height: 100%;
  display: flex;
  flex-direction: column;
}

.brand-logo {
  width: 24px;
  height: auto;
}

.main-nav {
  flex-grow: 1;
  margin-top: 20px;
}

.nav-item {
  width: 100%;
  margin-bottom: 16px;
}

.nav-link {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 12px;
  color: #262626;
  border-radius: 8px;
  transition: all 0.2s;
}

.nav-link:hover {
  color: #0095f6;
  background: #fafafa;
}

.nav-link.router-link-active {
  color: #0095f6;
}

.feather {
  width: 24px;
  height: 24px;
}

.bottom-section {
  padding: 16px 0;
  border-top: 1px solid #dbdbdb;
  margin-top: auto;
  text-align: center;
}

.username {
  font-size: 12px;
  color: #8e8e8e;
}

.logout-link {
  color: #ed4956;
}

.logout-link:hover {
  color: #c13584;
}
</style>