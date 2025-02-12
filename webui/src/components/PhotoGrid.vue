<template>
  <div class="photo-grid">
    <div v-for="photo in photos" :key="photo.id" class="photo-item" @click="$emit('open-photo', photo)">
      <img 
        :src="`data:image/jpeg;base64,${photo.image}`" 
        :alt="photo.caption || 'Photo'" 
      />
      <div v-if="photo.caption" class="photo-caption">
        {{ photo.caption }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    photos: {
      type: Array,
      required: true,
      default: () => []
    }
  },
  emits: ['open-photo']
}
</script>

<style scoped>
.photo-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
  padding: 20px;
}

.photo-item {
  position: relative;
  aspect-ratio: 1;
  cursor: pointer;
  transition: opacity 0.2s;
}

.photo-item:hover {
  opacity: 0.9;
}

.photo-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 4px;
}

.photo-caption {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(0,0,0,0.5);
  color: white;
  padding: 8px;
  border-radius: 0 0 4px 4px;
  font-size: 14px;
}
</style>