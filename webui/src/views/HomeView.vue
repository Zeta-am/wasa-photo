<template>
  <main class="container h-100 d-flex align-items-center justify-content-center">
    <div class="text-center" v-if="posts.length === 0">
      <ErrorMsg v-if="errormsg" :msg="errormsg"/>
    </div>
    <SearchModal ref="searchModal" />
    <div class="home-container">
      <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''" />
      <LoadingSpinner :loading="isLoading" />

      <div class="post-list">
        <PostCard
          v-for="(post, index) in posts"
          :key="index"
          :post="post"
        />
      </div>
      <div v-if="posts.length === 0" class="row">
        <div class="mt-4 centered-content">
          <p class="text-muted">
            There are no posts to view. 
          <br>
            This may be due to the fact that you don’t follow anyone yet or that the people you follow have not posted content yet.
          </p>
        <button @click="openSearch" class="btn btn-primary mt-3">
          Looking for people to follow
        </button>
      </div>
    
      </div>
    </div>
  </main>
</template>

<script>
import SearchModal from '@/components/SearchModal.vue'
import PostCard from '@/components/PostCard.vue'

export default {
  components: {
    SearchModal,
    PostCard
  },
  data() {
    return {
      errormsg: null,
      posts: [],
      errorMsg: "",
      isLoading: false
    }
  },
  created() {
    // Reset any stored user ID when entering home view
    localStorage.removeItem('lastVisitedProfile')
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
    openSearch() {
      this.$refs.searchModal.open()
    },
    logout() {
      localStorage.removeItem('token')
      localStorage.removeItem('username')
      this.$setAuth()
      this.$router.push({ name: 'Login' })
    },
    async getMyStream() {
      this.isLoading = true;
      try {
        const userId = this.$utils.getCurrentId();
        const response = await this.$axios.get(`/users/${userId}/stream`);
        if (response.data) {
          this.posts = response.data;
        } else {
          this.posts = [];
        }
      } catch (e) {
        this.errorMsg = 'Error loading stream';
        console.error(e.toString());
      } finally {
        this.isLoading = false;
      }
    }
  },
  mounted() {
    if (!localStorage.token) {
      this.$router.replace('/login');
    }
    this.getMyStream();
  }
}
</script>

<style scoped>
main {
  margin: 0 !important;
  padding: 2rem !important;
}

.home-container, .profile-container {
  padding-left: 60px; /* larghezza della dashboard */
}

.post-list {
  display: flex;
  flex-direction: column;
  gap: 20px;
  width: 100%;
}

.d-flex {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.centered-content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
}

.photo-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.85);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}
</style>
