<script setup>
import { ref } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useRouter } from 'vue-router'
import { LogIn, User, Lock, AlertCircle } from 'lucide-vue-next'

const auth = useAuthStore()
const router = useRouter()

const username = ref('')
const password = ref('')
const error = ref('')

const handleLogin = async () => {
  if (!username.value || !password.value) {
    error.value = 'Please fill in all fields'
    return
  }

  const success = await auth.login(username.value, password.value)
  if (success) {
    router.push('/admin')
  } else {
    error.value = auth.error
  }
}
</script>

<template>
  <div class="login-page">
    <div class="login-card glass-card">
      <div class="login-header">
        <div class="logo-circle">
          <Lock :size="32" color="white" />
        </div>
        <h2>Admin Login</h2>
        <p>Enter your credentials to access the dashboard</p>
      </div>

      <form @submit.prevent="handleLogin" class="login-form">
        <div v-if="error" class="error-msg">
          <AlertCircle :size="18" />
          <span>{{ error }}</span>
        </div>

        <div class="input-group">
          <label>Username</label>
          <div class="input-wrapper">
            <User :size="18" class="input-icon" />
            <input v-model="username" type="text" placeholder="admin" autocomplete="username" />
          </div>
        </div>

        <div class="input-group">
          <label>Password</label>
          <div class="input-wrapper">
            <Lock :size="18" class="input-icon" />
            <input v-model="password" type="password" placeholder="••••••••" autocomplete="current-password" />
          </div>
        </div>

        <button type="submit" class="btn-primary btn-block" :disabled="auth.loading">
          <LogIn v-if="!auth.loading" :size="20" />
          <span v-if="auth.loading">Logging in...</span>
          <span v-else>Sign In</span>
        </button>
      </form>
      
      <div class="login-footer">
        <RouterLink to="/" class="back-link">Back to landing page</RouterLink>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: radial-gradient(circle at top right, rgba(99, 102, 241, 0.1), transparent),
              radial-gradient(circle at bottom left, rgba(245, 158, 11, 0.1), transparent);
  padding: 20px;
}

.login-card {
  width: 100%;
  max-width: 440px;
  padding: 40px;
}

.login-header {
  text-align: center;
  margin-bottom: 32px;
}

.logo-circle {
  width: 64px;
  height: 64px;
  background: var(--primary);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
  box-shadow: 0 8px 16px rgba(99, 102, 241, 0.3);
}

.login-header h2 {
  font-size: 1.75rem;
  margin-bottom: 8px;
}

.login-header p {
  color: var(--text-muted);
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.error-msg {
  background: #fef2f2;
  color: #ef4444;
  padding: 12px;
  border-radius: 8px;
  font-size: 0.875rem;
  display: flex;
  align-items: center;
  gap: 8px;
  border: 1px solid rgba(239, 68, 68, 0.1);
}

.input-group label {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  margin-bottom: 8px;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.input-icon {
  position: absolute;
  left: 12px;
  color: var(--text-muted);
}

.input-wrapper input {
  width: 100%;
  padding: 12px 12px 12px 40px;
  border: 1px solid var(--border);
  border-radius: 10px;
  font-size: 1rem;
  transition: all 0.2s;
  outline: none;
}

.input-wrapper input:focus {
  border-color: var(--primary);
  box-shadow: 0 0 0 4px rgba(99, 102, 241, 0.1);
}

.btn-block {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  height: 48px;
}

.login-footer {
  margin-top: 32px;
  text-align: center;
}

.back-link {
  color: var(--text-muted);
  text-decoration: none;
  font-size: 0.875rem;
  transition: color 0.2s;
}

.back-link:hover {
  color: var(--primary);
}
</style>
