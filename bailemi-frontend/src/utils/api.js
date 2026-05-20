import axios from 'axios'
import { useAuthStore } from '@/stores/auth'
import { useBackendStore } from '@/stores/backend'

// 创建 axios 实例的函数
const createApiInstance = (baseURL) => {
  const instance = axios.create({
    baseURL,
    timeout: 10000
  })

  instance.interceptors.request.use(
    (config) => {
      const authStore = useAuthStore()
      if (authStore.token) {
        config.headers.Authorization = `Bearer ${authStore.token}`
      }
      return config
    },
    (error) => Promise.reject(error)
  )

  instance.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response?.status === 401) {
        const authStore = useAuthStore()
        authStore.logout()
        if (window.location.pathname !== '/login' && window.location.pathname !== '/register' && window.location.pathname !== '/backend-config') {
          window.location.href = '/login'
        }
      }
      return Promise.reject(error)
    }
  )

  return instance
}

// 默认 API 实例
let api = createApiInstance('http://localhost:8080/api/v1')

// 初始化 API（在应用启动时调用）
export const initApi = () => {
  const backendStore = useBackendStore()
  const baseURL = backendStore.currentBaseURL
  api = createApiInstance(baseURL)
  return api
}

// 重新配置 API（切换后端时调用）
export const reconfigureApi = (baseURL) => {
  api = createApiInstance(baseURL)
  return api
}

// 获取当前 API 实例
export const getApi = () => api

// 导出默认实例
export default api
