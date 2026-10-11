import { createApp } from 'vue'
import { createPinia } from 'pinia'
import client from './api/client'
import './style.css'
import App from './App.vue'
import router from './router'

// Handle unauthorized responses globally
client.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            localStorage.removeItem('token')
            localStorage.removeItem('user')
            if (router.currentRoute.value.path !== '/login') {
                router.push('/login')
            }
        }
        return Promise.reject(error)
    }
)

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.mount('#app')
