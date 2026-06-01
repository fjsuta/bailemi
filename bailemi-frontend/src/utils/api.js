import axios from 'axios'
import { useAuthStore } from '@/stores/auth'

// 后端服务配置
const AUTH_BASE_URL = import.meta.env.VITE_AUTH_API_BASE_URL || 'http://localhost:8080'
const CORE_BASE_URL = import.meta.env.VITE_CORE_API_BASE_URL || 'http://localhost:8081'
const DEFAULT_BACKEND = import.meta.env.VITE_DEFAULT_BACKEND || 'spring'

// Spring Boot 认证服务实例 (用户认证、管理)
const authApi = axios.create({
  baseURL: `${AUTH_BASE_URL}/api/v1`,
  timeout: 10000
})

// Go Gateway 核心服务实例 (音乐、播放、社交)
const coreApi = axios.create({
  baseURL: `${CORE_BASE_URL}/v1`,
  timeout: 10000
})

// 通用请求拦截器
const addAuthHeader = (config) => {
  const authStore = useAuthStore()
  if (authStore.token) {
    config.headers.Authorization = `Bearer ${authStore.token}`
  }
  return config
}

// 通用响应拦截器
const handleAuthError = (error) => {
  if (error.response?.status === 401) {
    const authStore = useAuthStore()
    authStore.logout()
    const currentPath = window.location.pathname
    if (currentPath !== '/login' && currentPath !== '/register') {
      window.location.href = '/login'
    }
  }
  return Promise.reject(error)
}

// 应用拦截器
authApi.interceptors.request.use(addAuthHeader, (error) => Promise.reject(error))
authApi.interceptors.response.use((response) => response, handleAuthError)

coreApi.interceptors.request.use(addAuthHeader, (error) => Promise.reject(error))
coreApi.interceptors.response.use((response) => response, handleAuthError)

// 统一 API 对象
const api = {
  // 认证相关 (Spring Boot)
  auth: authApi,

  // 核心业务 (Go Gateway)
  core: coreApi,

  // 便捷方法：发送认证请求
  async authRequest(method, url, data = null, config = {}) {
    return authApi.request({ method, url, data, ...config })
  },

  // 便捷方法：发送核心业务请求
  async coreRequest(method, url, data = null, config = {}) {
    return coreApi.request({ method, url, data, ...config })
  },

  // 兼容旧代码的默认导出行为
  get(url, config) {
    if (url.startsWith('/auth') || url.startsWith('/admin')) {
      return authApi.get(url, config)
    }
    return coreApi.get(url, config)
  },

  post(url, data, config) {
    if (url.startsWith('/auth') || url.startsWith('/admin')) {
      return authApi.post(url, data, config)
    }
    return coreApi.post(url, data, config)
  },

  put(url, data, config) {
    if (url.startsWith('/auth') || url.startsWith('/admin')) {
      return authApi.put(url, data, config)
    }
    return coreApi.put(url, data, config)
  },

  delete(url, config) {
    if (url.startsWith('/auth') || url.startsWith('/admin')) {
      return authApi.delete(url, config)
    }
    return coreApi.delete(url, config)
  }
}

export default api
export { authApi, coreApi, AUTH_BASE_URL, CORE_BASE_URL, DEFAULT_BACKEND }
