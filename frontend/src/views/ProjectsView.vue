<script setup>
import { ref, onMounted } from 'vue'
import { Plus, MoreVertical, Search } from 'lucide-vue-next'

const projects = ref([])
const loading = ref(true)

onMounted(async () => {
  try {
    const res = await fetch('/api/projects')
    projects.value = await res.json()
  } catch (e) {
    console.error('Failed to load projects')
  } finally {
    loading.value = false
  }
})

const getStatusColor = (status) => {
  switch (status) {
    case 'Completed': return '#10b981'
    case 'Active': return '#6366f1'
    default: return '#64748b'
  }
}
</script>

<template>
  <div class="projects">
    <div class="action-bar">
      <div class="search-box">
        <Search :size="20" />
        <input type="text" placeholder="Search projects..." />
      </div>
      <button class="btn-primary">
        <Plus :size="20" />
        <span>New Project</span>
      </button>
    </div>

    <div v-if="loading" class="loading">Loading projects...</div>
    
    <div v-else class="project-list glass-card">
      <table class="data-table">
        <thead>
          <tr>
            <th>Title</th>
            <th>Description</th>
            <th>Status</th>
            <th>Created At</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="project in projects" :key="project.id">
            <td><span class="project-title">{{ project.title }}</span></td>
            <td><span class="project-desc">{{ project.description }}</span></td>
            <td>
              <span class="status-badge" :style="{ backgroundColor: getStatusColor(project.status) + '15', color: getStatusColor(project.status) }">
                {{ project.status }}
              </span>
            </td>
            <td>{{ new Date(project.created_at).toLocaleDateString() }}</td>
            <td class="actions"><MoreVertical :size="18" /></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.action-bar {
  display: flex;
  justify-content: space-between;
  margin-bottom: 24px;
}

.search-box {
  background: white;
  border: 1px solid var(--border);
  padding: 8px 16px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  gap: 12px;
  width: 300px;
}

.search-box input {
  border: none;
  outline: none;
  width: 100%;
}

.btn-primary {
  display: flex;
  align-items: center;
  gap: 8px;
}

.project-list {
  padding: 8px;
  overflow: hidden;
}

.data-table {
  width: 100%;
  border-collapse: collapse;
}

.data-table th {
  text-align: left;
  padding: 16px;
  color: var(--text-muted);
  font-weight: 600;
  font-size: 0.875rem;
  border-bottom: 1px solid var(--border);
}

.data-table td {
  padding: 16px;
  border-bottom: 1px solid var(--border);
}

.project-title {
  font-weight: 600;
  color: var(--text-main);
}

.project-desc {
  color: var(--text-muted);
}

.status-badge {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.75rem;
  font-weight: 600;
}

.actions {
  color: var(--text-muted);
  text-align: right;
  cursor: pointer;
}
</style>
