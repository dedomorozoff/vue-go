import { createRouter, createWebHistory } from 'vue-router'
import LandingView from '../views/LandingView.vue'
import AdminLayout from '../layouts/AdminLayout.vue'
import HomeView from '../views/HomeView.vue'
import ProjectsView from '../views/ProjectsView.vue'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            name: 'landing',
            component: LandingView
        },
        {
            path: '/admin',
            component: AdminLayout,
            children: [
                {
                    path: '',
                    name: 'home',
                    component: HomeView,
                    meta: { title: 'Dashboard' }
                },
                {
                    path: 'projects',
                    name: 'projects',
                    component: ProjectsView,
                    meta: { title: 'Project Management' }
                }
            ]
        }
    ]
})

export default router
