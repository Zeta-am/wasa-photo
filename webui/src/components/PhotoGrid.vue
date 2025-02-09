<template>
  <div class="photo-grid">
    <div v-for="photo in photos" :key="photo.id" class="photo-item">
      <img 
        :src="`data:image/jpeg;base64,${photo.image}`" 
        :alt="photo.caption"
        @click="$emit('open-photo', photo)"
      />
      <div class="photo-caption">{{ photo.caption }}</div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    photos: Array
  }
}
</script>

<style scoped>
.photo-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 20px;
  padding: 20px;
}

.photo-item {
  position: relative;
  aspect-ratio: 1;
  cursor: pointer;
}

.photo-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  border-radius: 8px;
}

.photo-caption {
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  background: rgba(0,0,0,0.5);
  color: white;
  padding: 8px;
  border-radius: 0 0 8px 8px;
}
</style>