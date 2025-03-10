<script>
export default {
  props: {
    show: Boolean,
    photo: {
      type: Object,
      required: true
    },
    username: String,
    isOwner: Boolean
  },
  data() {
    return {
      showOptions: false,
      newComment: '',
      currentUserId: this.$utils.getCurrentId(),
      comments: [],
      likes: [],
      isLiked: false,
      error: null, // Add error state
      likeCount: 0 // Add like counter
    }
  },
  async created() {
    if (this.photo?.id) {
      await this.loadData();
    }
  },
  async mounted() {
    if (this.photo?.id) {
      await this.loadData();
    }
  },
  methods: {
    async loadData() {
      this.error = null;
      try {
        const userId = this.$utils.getCurrentId();
        
        try {
          const commentsRes = await this.$axios.get(`/users/${userId}/posts/${this.photo.id}/comments`);
          this.comments = commentsRes.data || [];
        } catch (e) {
          this.error = 'Error loading comments';
          this.comments = [];
        }

        try {
          const likesRes = await this.$axios.get(`/users/${userId}/posts/${this.photo.id}/likes`);
          this.likes = likesRes.data || [];
          this.likeCount = this.likes.length;
          // Usa l'uguaglianza non stretta per evitare problemi di tipo
          this.isLiked = this.likes.some(like => like.userId == this.currentUserId);
        } catch (e) {
          this.error = 'Error loading likes';
          this.likes = [];
          this.likeCount = 0;
          this.isLiked = false;
        }
      } catch (e) {
        this.error = 'Error loading data';
        this.comments = [];
        this.likes = [];
        this.likeCount = 0;
        this.isLiked = false;
      }
    },
    close() {
      this.showOptions = false;
      this.$emit('close');
    },
    async toggleLike() {
      try {
        const userId = this.$utils.getCurrentId();
        if (!this.isLiked) {
          await this.$axios.put(`/users/${userId}/posts/${this.photo.id}/likes`);
          this.likeCount++;
          this.isLiked = true;
        } else {
          await this.$axios.delete(`/users/${userId}/posts/${this.photo.id}/likes`);
          this.likeCount--;
          this.isLiked = false;
        }
        // Emetti l'evento per il nuovo stato dei like
        this.$emit('like-updated', { likeCount: this.likeCount, isLiked: this.isLiked });
      } catch (e) {
        this.error = 'Error updating like';
      }
    },
    async addComment() {
      if (!this.newComment.trim()) return;
      
      try {
        const userId = this.$utils.getCurrentId();
        const response = await this.$axios.post(`/users/${userId}/posts/${this.photo.id}/comments`, {
          caption: this.newComment
        });
        
        // Add new comment to list
        this.comments.push({
          id: response.data.id,
          userId: this.currentUserId,
          username: this.$utils.getCurrentUsername(),
          caption: this.newComment,
          timestamp: new Date().toISOString()
        });
        
        this.newComment = '';
        // Emetti l'evento con il nuovo conteggio dei commenti
        this.$emit('comment-updated', { commentCount: this.comments.length });
      } catch (e) {
        console.error('Error adding comment:', e);
      }
    },
    async deleteComment(commentId) {
      try {
        const userId = this.$utils.getCurrentId();
        await this.$axios.delete(`/users/${userId}/posts/${this.photo.id}/comments/${commentId}`);
        this.comments = this.comments.filter(c => c.id !== commentId);
        // Emetti l'evento con il nuovo conteggio dei commenti
        this.$emit('comment-updated', { commentCount: this.comments.length });
      } catch (e) {
        console.error('Error deleting comment:', e);
      }
    },
    async deletePhoto() {
      try {
        const userId = this.$utils.getCurrentId();
        await this.$axios.delete(`/users/${userId}/posts/${this.photo.id}`);
        this.close();
        this.$emit('photo-deleted');
      } catch (e) {
        console.error('Error deleting photo:', e);
      }
    }
  },
  watch: {
    // Reload data when photo changes
    'photo.id': {
      handler: 'loadData',
      immediate: true
    }
  }
}
</script>

<template>
  <div v-if="show" class="photo-modal-overlay" @click.self="close">
    <!-- Add error message display -->
    <div v-if="error" class="error-message">
      {{ error }}
    </div>
    
    <div class="photo-modal-content">
      <!-- Left side - Photo -->
      <div class="photo-side">
        <img :src="'data:image/jpeg;base64,' + photo.image" :alt="photo.caption">
      </div>

      <!-- Right side - Interactions -->
      <div class="interaction-side">
        <!-- Header -->
        <div class="modal-header">
          <div class="user-info">
            <div class="username">{{ username }}</div>
          </div>
          <button v-if="isOwner" class="options-button" @click="showOptions = true">
            <svg aria-label="More options" class="_ab6-" color="currentColor" fill="currentColor" height="24" role="img" viewBox="0 0 24 24" width="24">
              <circle cx="12" cy="12" r="1.5"></circle>
              <circle cx="6" cy="12" r="1.5"></circle>
              <circle cx="18" cy="12" r="1.5"></circle>
            </svg>
          </button>
        </div>

        <!-- Comments section -->
        <div class="comments-section">
          <!-- Caption -->
          <div class="caption-container" v-if="photo.caption">
            <span class="username">{{ username }}</span>
            <span class="caption">{{ photo.caption }}</span>
          </div>
          
          <!-- Comments list -->
          <div class="comments-list" v-if="comments && comments.length > 0">
            <div v-for="comment in comments" :key="comment.id" class="comment">
              <span class="username">{{ comment.username }}</span>
              <span class="comment-text">{{ comment.caption }}</span>
              <button 
                v-if="comment.userId == currentUserId" 
                class="delete-comment"
                @click="deleteComment(comment.id)"
              >
                ×
              </button>
            </div>
          </div>
        </div>

        <!-- Actions section -->
        <div class="actions-section">
          <div class="action-buttons">
            <button class="action-button" @click="toggleLike">
              <svg v-if="!isLiked" aria-label="Like" class="_ab6-" color="currentColor" fill="currentColor" height="24" role="img" viewBox="0 0 24 24" width="24">
                <path d="M16.792 3.904A4.989 4.989 0 0 1 21.5 9.122c0 3.072-2.652 4.959-5.197 7.222-2.512 2.243-3.865 3.469-4.303 3.752-.477-.309-2.143-1.823-4.303-3.752C5.141 14.072 2.5 12.167 2.5 9.122a4.989 4.989 0 0 1 4.708-5.218 4.21 4.21 0 0 1 3.675 1.941c.84 1.175.98 1.763 1.12 1.763s.278-.588 1.11-1.766a4.17 4.17 0 0 1 3.679-1.938m0-2a6.04 6.04 0 0 0-4.797 2.127 6.052 6.052 0 0 0-4.787-2.127A6.985 6.985 0 0 0 .5 9.122c0 3.61 2.55 5.827 5.015 7.97.283.246.569.494.853.747l1.027.918a44.998 44.998 0 0 0 3.518 3.018 2 2 0 0 0 2.174 0 45.263 45.263 0 0 0 3.626-3.115l.922-.824c.293-.26.59-.519.885-.774 2.334-2.025 4.98-4.32 4.98-7.94a6.985 6.985 0 0 0-6.708-7.218Z"></path>
              </svg>
              <svg v-else aria-label="Unlike" class="_ab6-" color="rgb(255, 48, 64)" fill="rgb(255, 48, 64)" height="24" role="img" viewBox="0 0 48 48" width="24">
                <path d="M34.6 3.1c-4.5 0-7.9 1.8-10.6 5.6-2.7-3.7-6.1-5.5-10.6-5.5C6 3.1 0 9.6 0 17.6c0 7.3 5.4 12 10.6 16.5.6.5 1.3 1.1 1.9 1.7l2.3 2c4.4 3.9 6.6 5.9 7.6 6.5.5.3 1.1.5 1.6.5s1.1-.2 1.6-.5c1-.6 2.8-2.2 7.8-6.8l2-1.8c.7-.6 1.3-1.2 2-1.7C42.7 29.6 48 25 48 17.6c0-8-6-14.5-13.4-14.5z"></path>
              </svg>
            </button>
          </div>
          
          <!-- Like count -->
          <div class="likes-count" v-if="likeCount > 0">
            {{ likeCount }} {{ likeCount === 1 ? 'like' : 'likes' }}
          </div>

          <!-- Add comment section -->
          <div class="add-comment">
            <input 
              v-model="newComment"
              type="text"
              placeholder="Add a comment..."
              @keyup.enter="addComment"
            >
            <button 
              :disabled="!newComment.trim()"
              @click="addComment"
              class="post-comment"
            >
              Post
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Options Modal -->
    <div v-if="showOptions" class="options-modal" @click.self="showOptions = false">
      <div class="options-content">
        <button class="delete-option" @click="deletePhoto">
          Delete
        </button>
        <button class="cancel-option" @click="showOptions = false">
          Cancel
        </button>
      </div>
    </div>
  </div>
</template>


<style scoped>
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

.photo-modal-content {
  display: flex;
  background: white;
  max-width: 1300px;
  width: 95%;
  height: 90vh;
  border-radius: 4px;
  overflow: hidden;
}

.photo-side {
  flex: 1;
  background: black;
  display: flex;
  align-items: center;
}

.photo-side img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.interaction-side {
  width: 400px;
  display: flex;
  flex-direction: column;
  border-left: 1px solid #dbdbdb;
}

.modal-header {
  padding: 14px 16px;
  border-bottom: 1px solid #dbdbdb;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.comments-section {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
}

.caption-container {
  margin-bottom: 16px;
}

.comment {
  margin-bottom: 12px;
  position: relative;
}

.username {
  font-weight: 600;
  margin-right: 8px;
}

.delete-comment {
  position: absolute;
  right: 0;
  top: 50%;
  transform: translateY(-50%);
  background: none;
  border: none;
  color: #ed4956;
  cursor: pointer;
  font-size: 16px;
  opacity: 0;
  transition: opacity 0.3s ease;
}

.comment:hover .delete-comment {
  opacity: 1;
}

.actions-section {
  border-top: 1px solid #dbdbdb;
  padding: 16px;
}

.action-buttons {
  margin-bottom: 8px;
}

.action-button {
  background: none;
  border: none;
  padding: 8px;
  cursor: pointer;
}

.likes-count {
  font-weight: 600;
  margin-bottom: 8px;
}

.add-comment {
  display: flex;
  gap: 8px;
}

.add-comment input {
  flex: 1;
  border: none;
  padding: 8px;
  outline: none;
}

.post-comment {
  background: none;
  border: none;
  color: #0095f6;
  font-weight: 600;
  cursor: pointer;
}

.post-comment:disabled {
  opacity: 0.3;
  cursor: default;
}

.options-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.65);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1001;
}

.options-content {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  width: 400px;
}

.options-content button {
  width: 100%;
  padding: 14px;
  border: none;
  background: none;
  font-size: 14px;
  border-bottom: 1px solid #dbdbdb;
  cursor: pointer;
}

.delete-option {
  color: #ed4956;
  font-weight: 600;
}

.options-button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 8px;
}
</style>