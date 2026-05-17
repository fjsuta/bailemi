import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/music/Home.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/Login.vue'),
    meta: { requiresAuth: false, guestOnly: true }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/auth/Register.vue'),
    meta: { requiresAuth: false, guestOnly: true }
  },
  {
    path: '/profile',
    name: 'Profile',
    component: () => import('@/views/profile/Profile.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('@/views/admin/Index.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/artist/:id',
    name: 'ArtistDetail',
    component: () => import('@/views/music/ArtistDetail.vue')
  },
  {
    path: '/album/:id',
    name: 'AlbumDetail',
    component: () => import('@/views/music/AlbumDetail.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to, from, next) => {
  const authStore = useAuthStore()
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  const guestOnly = to.matched.some(record => record.meta.guestOnly)

  if (requiresAuth && !authStore.isAuthenticated) {
    next('/login')
  } else if (guestOnly && authStore.isAuthenticated) {
    next('/profile')
  } else {
    next()
  }
})

export default router
