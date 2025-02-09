<!-- src/components/Dashboard.vue -->
<template>
  <div class="dashboard">
    <nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
      <div class="position-sticky pt-3 sidebar-sticky">
        <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
          <span>Menu</span>
        </h6>
        <ul class="nav flex-column">
          <li class="nav-item">
            <RouterLink to="/home" class="nav-link">
              <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
              Home
            </RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink :to="usrLink" class="nav-link">
              <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user"/></svg>
              Profile
            </RouterLink>
          </li>
          <li class="nav-item">
            <a class="nav-link" role="button" @click="$emit('logout')">
              <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
              Logout
            </a>
          </li>
        </ul>
        <h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
          <span>Logged as {{ username }}</span>
        </h6>
      </div>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  props: {
    username: String,
    usrLink: String
  }
}
</script>

<style scoped>
.sidebar {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  z-index: 100;
  padding: 48px 0 0;
  box-shadow: inset -1px 0 0 rgba(0, 0, 0, .1);
}

.sidebar-sticky {
  position: relative;
  top: 0;
  height: calc(100vh - 48px);
  padding-top: .5rem;
  overflow-x: hidden;
  overflow-y: auto;
}

.nav-link {
  font-weight: 500;
  color: #333;
  cursor: pointer;
}

.nav-link:hover {
  color: #007bff;
}

.feather {
  width: 16px;
  height: 16px;
  margin-right: 4px;
  vertical-align: text-bottom;
}
</style>