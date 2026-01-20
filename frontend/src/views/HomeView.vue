<script setup>
import { ref, onMounted } from 'vue'
import { CheckCircle2, Clock, Layers, Activity } from 'lucide-vue-next'

const stats = ref([
  { label: 'Total Projects', value: '12', icon: Layers, color: '#6366f1' },
  { label: 'Active Tasks', value: '24', icon: Activity, color: '#f59e0b' },
  { label: 'Completed', value: '158', icon: CheckCircle2, color: '#10b981' },
])

const backendStatus = ref('Checking...')

onMounted(async () => {
  try {
    const res = await fetch('/api/health')
    const data = await res.json()
    backendStatus.value = data.message
  } catch (e) {
    backendStatus.value = 'Offline'
  }
})
</script>

<template>
  <div class="dashboard">
    <div class="stats-grid">
      <div v-for="stat in stats" :key="stat.label" class="stat-card glass-card">
        <div class="stat-icon" :style="{ backgroundColor: stat.color + '20', color: stat.color }">
          <component :is="stat.icon" :size="24" />
        </div>
        <div class="stat-info">
          <p class="stat-label">{{ stat.label }}</p>
          <p class="stat-value">{{ stat.value }}</p>
        </div>
      </div>
    </div>

    <div class="status-card glass-card">
      <h3>System Status</h3>
      <div class="status-indicator">
        <div class="pulse" :class="{ 'pulse-green': backendStatus.includes('alive') }"></div>
        <span>Backend API: {{ backendStatus }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 24px;
  margin-bottom: 40px;
}

.stat-card {
  padding: 24px;
  display: flex;
  align-items: center;
  gap: 20px;
}

.stat-icon {
  width: 56px;
  height: 56px;
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.stat-label {
  color: var(--text-muted);
  font-size: 0.875rem;
  margin: 0 0 4px 0;
}

.stat-value {
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
}

.status-card {
  padding: 24px;
}

.status-indicator {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-top: 16px;
}

.pulse {
  width: 12px;
  height: 12px;
  background: #ef4444;
  border-radius: 50%;
  box-shadow: 0 0 0 rgba(239, 68, 68, 0.4);
  animation: pulse 2s infinite;
}

.pulse-green {
  background: #10b981;
  box-shadow: 0 0 0 rgba(16, 185, 129, 0.4);
}

@keyframes pulse {
  0% { transform: scale(0.95); box-shadow: 0 0 0 0 rgba(16, 185, 129, 0.7); }
  70% { transform: scale(1); box-shadow: 0 0 0 10px rgba(16, 185, 129, 0); }
  100% { transform: scale(0.95); box-shadow: 0 0 0 0 rgba(16, 185, 129, 0); }
}
</style>
