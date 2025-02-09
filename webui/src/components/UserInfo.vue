<template>
    <div class="user-info">
      <h2>{{ username }}</h2>
      <div class="stats">
        <span>Posts: {{ postCount }}</span>
        <span>Followers: {{ followerCount }}</span>
        <span>Following: {{ followingCount }}</span>
      </div>
      <button v-if="!isOwnProfile" @click="toggleFollow">
        {{ isFollowing ? 'Unfollow' : 'Follow' }}
      </button>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      username: String,
      postCount: Number,
      followerCount: Number,
      followingCount: Number,
      isOwnProfile: Boolean,
      isFollowing: Boolean
    },
    methods: {
      toggleFollow() {
        this.$emit('toggle-follow')
      }
    }
  }
  </script>
  
  <style scoped>
  .user-info {
    text-align: center;
    margin-bottom: 20px;
  }
  .stats {
    display: flex;
    justify-content: space-around;
    margin: 10px 0;
  }
  </style>
  
  