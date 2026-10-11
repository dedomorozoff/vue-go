import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'landing',
            component: () => import('../views/LandingView.vue'),
            meta: { title: 'Vue-Go Starter' },
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('../views/LoginView.vue'),
            meta: { title: 'Login' },
        },
        {
            path: '/admin',
            component: () => import('../layouts/AdminLayout.vue'),
            meta: { requiresAuth: true },
            children: [
                {
                    path: '',
                    name: 'home',
                    component: () => import('../views/HomeView.vue'),
                    meta: { title: 'Dashboard' }
                },
                {
                    path: 'projects',
                    name: 'projects',
                    component: () => import('../views/ProjectsView.vue'),
                    meta: { title: 'Project Management' }
                }
            ]
        },
        {
            path: '/:pathMatch(.*)*',
            name: 'not-found',
            component: () => import('../views/NotFoundView.vue'),
            meta: { title: 'Page Not Found' },
        },
    ]
})

router.beforeEach(async (to, from, next) => {
    const auth = useAuthStore()
    // Hydrate user on reload if token exists but user is missing
    if (auth.token && !auth.user) {
        await auth.fetchMe()
    }
    if (to.meta.requiresAuth && !auth.isAuthenticated) {
        next('/login')
    } else if (to.name === 'login' && auth.isAuthenticated) {
        next('/admin')
    } else {
        next()
    }
})

router.afterEach((to) => {
    document.title = to.meta.title ? `${to.meta.title} — Vue-Go` : 'Vue-Go'
})

export default router
