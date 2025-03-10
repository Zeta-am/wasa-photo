<script>
export default {
  name: 'SearchResult',
  props: {
    user: {
      type: Object,
      required: true
    }
  },
  emits: ['select']
}
</script>

<template>
  <div class="result-item" @click="$emit('select', user)">
    <div class="user-info">
      <div class="username">{{ user.username }}</div>
    </div>
  </div>
</template>


<style scoped>
.result-item {
  display: flex;
  align-items: center;
  padding: 12px;
  cursor: pointer;
  border-bottom: 1px solid #efefef;
}

.result-item:hover {
  background: #fafafa;
}

.avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  margin-right: 12px;
}

.user-info {
  flex: 1;
}

.username {
  font-weight: 500;
  color: #262626;
}
</style>
