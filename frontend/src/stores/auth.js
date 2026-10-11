import { defineStore } from 'pinia'
import client from '../api/client'

function safeParse(raw) {
  try {
    return raw ? JSON.parse(raw) : null
  } catch {
    return null
  }
}

export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: safeParse(localStorage.getItem('user')),
        token: localStorage.getItem('token') || null,
        loading: false,
        error: null,
    }),
    getters: {
        isAuthenticated: (state) => !!state.token,
    },
    actions: {
        async login(username, password) {
            this.loading = true
            this.error = null
            try {
                const response = await client.post('/api/auth/login', { username, password })
                this.token = response.data.token
                this.user = response.data.user
                localStorage.setItem('token', this.token)
                localStorage.setItem('user', JSON.stringify(this.user))
                return true
            } catch (err) {
                this.error = err.response?.data?.error || 'Login failed'
                return false
            } finally {
                this.loading = false
            }
        },
        async fetchMe() {
            if (!this.token) return false
            try {
                const res = await client.get('/api/auth/me')
                this.user = res.data
                localStorage.setItem('user', JSON.stringify(this.user))
                return true
            } catch {
                this.logout()
                return false
            }
        },
        logout() {
            this.token = null
            this.user = null
            localStorage.removeItem('token')
            localStorage.removeItem('user')
        },
    },
})
