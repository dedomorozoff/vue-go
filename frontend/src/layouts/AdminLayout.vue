<script setup>
import { LayoutDashboard, FolderKanban, Activity, Settings, LogOut, Home } from 'lucide-vue-next'
</script>

<template>
  <div class="app-container">
    <aside class="sidebar">
      <div class="logo">
        <Activity class="icon-brand" />
        <span>VueGo Admin</span>
      </div>
      <nav class="nav-menu">
        <RouterLink to="/admin" class="nav-item" active-class="active" end>
          <LayoutDashboard :size="20" />
          <span>Dashboard</span>
        </RouterLink>
        <RouterLink to="/admin/projects" class="nav-item" active-class="active">
          <FolderKanban :size="20" />
          <span>Projects</span>
        </RouterLink>
      </nav>
      
      <div class="sidebar-footer">
        <RouterLink to="/" class="nav-item">
          <Home :size="20" />
          <span>Back to Site</span>
        </RouterLink>
        <div class="nav-item">
          <Settings :size="20" />
          <span>Settings</span>
        </div>
        <div class="nav-item logout">
          <LogOut :size="20" />
          <span>Logout</span>
        </div>
      </div>
    </aside>

    <main class="main-content">
      <header class="top-bar">
        <h1>{{ $route.meta.title || 'Admin' }}</h1>
        <div class="user-profile">
          <div class="user-info">
            <span class="user-name">Alex DeMoroz</span>
            <span class="user-role">Administrator</span>
          </div>
          <div class="avatar">AD</div>
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
  background-color: var(--bg-main);
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

.sidebar-footer {
  margin-top: auto;
  padding-top: 20px;
  border-top: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.logout:hover {
  color: #ef4444;
  background: #fef2f2;
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

.user-profile {
  display: flex;
  align-items: center;
  gap: 16px;
}

.user-info {
  text-align: right;
}

.user-name {
  display: block;
  font-weight: 600;
  font-size: 0.875rem;
  color: var(--text-main);
}

.user-role {
  display: block;
  font-size: 0.75rem;
  color: var(--text-muted);
}

.avatar {
  width: 40px;
  height: 40px;
  background: var(--primary);
  color: white;
  border-radius: 12px;
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
