<script setup>
import { ref, onMounted, watch } from 'vue'
import { Plus, Search, Pencil, Trash2, X } from 'lucide-vue-next'
import { projectsApi } from '../api/projects'

const projects = ref([])
const loading = ref(true)
const error = ref('')
const search = ref('')
const statusFilter = ref('')
const showModal = ref(false)
const editing = ref(null)
const form = ref({ title: '', description: '', status: 'Active' })
const saving = ref(false)

let debounce
watch([search, statusFilter], () => {
  clearTimeout(debounce)
  debounce = setTimeout(loadProjects, 300)
})

async function loadProjects() {
  loading.value = true
  error.value = ''
  try {
    const res = await projectsApi.list({ search: search.value, status: statusFilter.value })
    projects.value = res.data.data || res.data
  } catch (e) {
    error.value = e.response?.data?.error || 'Failed to load projects'
  } finally {
    loading.value = false
  }
}

onMounted(loadProjects)

function openCreate() {
  editing.value = null
  form.value = { title: '', description: '', status: 'Active' }
  showModal.value = true
}

function openEdit(p) {
  editing.value = p
  form.value = { title: p.title, description: p.description, status: p.status }
  showModal.value = true
}

async function saveProject() {
  if (!form.value.title.trim()) return
  saving.value = true
  try {
    if (editing.value) {
      await projectsApi.update(editing.value.id, form.value)
    } else {
      await projectsApi.create(form.value)
    }
    showModal.value = false
    await loadProjects()
  } catch (e) {
    error.value = e.response?.data?.error || 'Failed to save project'
  } finally {
    saving.value = false
  }
}

async function deleteProject(p) {
  if (!confirm(`Delete "${p.title}"?`)) return
  try {
    await projectsApi.remove(p.id)
    await loadProjects()
  } catch (e) {
    error.value = e.response?.data?.error || 'Failed to delete project'
  }
}

const getStatusColor = (status) => {
  switch (status) {
    case 'Completed': return '#10b981'
    case 'Active': return '#6366f1'
    case 'Pending': return '#f59e0b'
    default: return '#64748b'
  }
}
</script>

<template>
  <div class="projects">
    <div class="action-bar">
      <div class="search-box">
        <Search :size="20" />
        <input v-model="search" type="text" placeholder="Search projects..." />
      </div>
      <select v-model="statusFilter" class="status-filter">
        <option value="">All statuses</option>
        <option value="Active">Active</option>
        <option value="Pending">Pending</option>
        <option value="Completed">Completed</option>
      </select>
      <button class="btn-primary" @click="openCreate">
        <Plus :size="20" />
        <span>New Project</span>
      </button>
    </div>

    <div v-if="error" class="error-banner">{{ error }}</div>
    <div v-if="loading" class="loading">Loading projects...</div>
    <div v-else-if="!projects.length" class="empty-state glass-card">
      <p>No projects found. Create your first one!</p>
      <button class="btn-primary" @click="openCreate">New Project</button>
    </div>

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
            <td>{{ project.created_at ? new Date(project.created_at).toLocaleDateString() : '—' }}</td>
            <td class="actions">
              <button class="icon-btn" @click="openEdit(project)" title="Edit"><Pencil :size="18" /></button>
              <button class="icon-btn danger" @click="deleteProject(project)" title="Delete"><Trash2 :size="18" /></button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="showModal" class="modal-overlay" @click.self="showModal = false">
      <div class="modal glass-card">
        <div class="modal-header">
          <h3>{{ editing ? 'Edit Project' : 'New Project' }}</h3>
          <button class="icon-btn" @click="showModal = false"><X :size="20" /></button>
        </div>
        <form @submit.prevent="saveProject" class="modal-form">
          <label>Title
            <input v-model="form.title" type="text" placeholder="Project title" required maxlength="200" />
          </label>
          <label>Description
            <textarea v-model="form.description" placeholder="Short description" rows="3"></textarea>
          </label>
          <label>Status
            <select v-model="form.status">
              <option>Active</option>
              <option>Pending</option>
              <option>Completed</option>
            </select>
          </label>
          <div class="modal-actions">
            <button type="button" class="btn-secondary" @click="showModal = false">Cancel</button>
            <button type="submit" class="btn-primary" :disabled="saving">{{ saving ? 'Saving...' : 'Save' }}</button>
          </div>
        </form>
      </div>
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
  white-space: nowrap;
}

.icon-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: var(--text-muted);
  padding: 6px;
  border-radius: 8px;
}
.icon-btn:hover { background: var(--bg-main); color: var(--text-main); }
.icon-btn.danger:hover { background: #fef2f2; color: #ef4444; }

.status-filter {
  border: 1px solid var(--border);
  border-radius: 10px;
  padding: 8px 12px;
  background: white;
}

.error-banner {
  background: #fef2f2;
  color: #ef4444;
  border: 1px solid rgba(239,68,68,.2);
  padding: 12px 16px;
  border-radius: 10px;
  margin-bottom: 16px;
}

.empty-state {
  padding: 48px;
  text-align: center;
  color: var(--text-muted);
  display: flex;
  flex-direction: column;
  gap: 16px;
  align-items: center;
}

.modal-overlay {
  position: fixed;
  inset: 0;
  background: rgba(15,23,42,.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
}
.modal {
  width: 100%;
  max-width: 480px;
  padding: 24px;
}
.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}
.modal-form {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.modal-form label {
  display: flex;
  flex-direction: column;
  gap: 8px;
  font-size: .875rem;
  font-weight: 500;
}
.modal-form input, .modal-form textarea, .modal-form select {
  border: 1px solid var(--border);
  border-radius: 10px;
  padding: 10px 12px;
  font-size: 1rem;
  outline: none;
}
.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}
.btn-secondary {
  background: white;
  border: 1px solid var(--border);
  padding: 10px 20px;
  border-radius: 8px;
  cursor: pointer;
}
</style>
