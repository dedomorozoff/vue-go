<script setup>
import { LayoutDashboard, FolderKanban, Activity, Settings } from 'lucide-vue-next'
</script>

<template>
  <div class="app-container">
    <aside class="sidebar">
      <div class="logo">
        <Activity class="icon-brand" />
        <span>VueGo Admin</span>
      </div>
      <nav class="nav-menu">
        <RouterLink to="/" class="nav-item" active-class="active">
          <LayoutDashboard :size="20" />
          <span>Dashboard</span>
        </RouterLink>
        <RouterLink to="/projects" class="nav-item" active-class="active">
          <FolderKanban :size="20" />
          <span>Projects</span>
        </RouterLink>
      </nav>
      <div class="sidebar-footer">
        <div class="nav-item">
          <Settings :size="20" />
          <span>Settings</span>
        </div>
      </div>
    </aside>

    <main class="main-content">
      <header class="top-bar">
        <h1>{{ $route.name === 'home' ? 'Dashboard' : 'Project Management' }}</h1>
        <div class="user-profile">
          <div class="avatar">AL</div>
        </div>
      </header>
      <div class="page-container">
        <RouterView v-slot="{ Component }">
          <transition name="fade" mode="out-in">
            <component :is="Component" />
          </transition>
        </RouterView>
      </div>
    </main>
  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  width: 260px;
  background: white;
  border-right: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  padding: 24px;
  position: sticky;
  top: 0;
  height: 100vh;
}

.logo {
  display: flex;
  align-items: center;
  gap: 12px;
  font-weight: 700;
  font-size: 1.25rem;
  margin-bottom: 40px;
  color: var(--primary);
}

.icon-brand {
  color: var(--primary);
}

.nav-menu {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  text-decoration: none;
  color: var(--text-muted);
  border-radius: 12px;
  transition: all 0.2s;
  cursor: pointer;
}

.nav-item:hover {
  background: var(--bg-main);
  color: var(--text-main);
}

.nav-item.active {
  background: #f0f1ff;
  color: var(--primary);
  font-weight: 500;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
}

.top-bar {
  height: 80px;
  padding: 0 40px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: white;
  border-bottom: 1px solid var(--border);
}

.top-bar h1 {
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
}

.avatar {
  width: 40px;
  height: 40px;
  background: var(--primary);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  font-size: 0.875rem;
}

.page-container {
  padding: 40px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
