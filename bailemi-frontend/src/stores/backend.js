import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import axios from 'axios'

export const useBackendStore = defineStore('backend', () => {
  // 默认后端配置
  const defaultBackends = [
    {
      id: 'backend1',
      name: '本地开发后端',
      url: 'http://localhost:8080/api/v1',
      description: '本地开发环境',
      isDefault: true
    },
    {
      id: 'backend2', 
      name: '测试后端',
      url: 'http://localhost:8081/api/v1',
      description: '测试环境',
      isDefault: false
    }
  ]

  // 从 localStorage 加载配置
  const loadBackends = () => {
    const saved = localStorage.getItem('backendConfigs')
    if (saved) {
      try {
        return JSON.parse(saved)
      } catch (e) {
        console.error('加载后端配置失败:', e)
      }
    }
    return defaultBackends
  }

  const backends = ref(loadBackends())
  const currentBackendId = ref(localStorage.getItem('currentBackendId') || defaultBackends[0].id)
  const backendStatus = ref({}) // 存储各后端状态
  const isCheckingStatus = ref(false)

  // 当前选中的后端
  const currentBackend = computed(() => {
    return backends.value.find(b => b.id === currentBackendId.value) || backends.value[0]
  })

  // 获取当前后端URL
  const currentBaseURL = computed(() => {
    return currentBackend.value?.url || defaultBackends[0].url
  })

  // 设置当前后端
  const setCurrentBackend = (id) => {
    const backend = backends.value.find(b => b.id === id)
    if (backend) {
      currentBackendId.value = id
      localStorage.setItem('currentBackendId', id)
      return true
    }
    return false
  }

  // 添加新后端
  const addBackend = (backend) => {
    const newBackend = {
      ...backend,
      id: 'backend_' + Date.now()
    }
    backends.value.push(newBackend)
    saveBackends()
    return newBackend.id
  }

  // 更新后端配置
  const updateBackend = (id, config) => {
    const index = backends.value.findIndex(b => b.id === id)
    if (index !== -1) {
      backends.value[index] = { ...backends.value[index], ...config }
      saveBackends()
      return true
    }
    return false
  }

  // 删除后端
  const removeBackend = (id) => {
    const index = backends.value.findIndex(b => b.id === id)
    if (index !== -1 && backends.value.length > 1) {
      backends.value.splice(index, 1)
      // 如果删除的是当前选中的，切换到第一个
      if (currentBackendId.value === id) {
        setCurrentBackend(backends.value[0].id)
      }
      saveBackends()
      return true
    }
    return false
  }

  // 保存到 localStorage
  const saveBackends = () => {
    localStorage.setItem('backendConfigs', JSON.stringify(backends.value))
  }

  // 检测后端状态
  const checkBackendStatus = async (backend) => {
    try {
      const response = await axios.get(`${backend.url}/health`, {
        timeout: 5000,
        headers: {
          'Content-Type': 'application/json'
        }
      })
      return {
        online: true,
        latency: response.data.latency || 0,
        message: '正常'
      }
    } catch (error) {
      // 尝试检查根路径
      try {
        const start = Date.now()
        await axios.get(backend.url.replace('/api/v1', ''), {
          timeout: 5000
        })
        return {
          online: true,
          latency: Date.now() - start,
          message: '正常'
        }
      } catch (e) {
        return {
          online: false,
          latency: 0,
          message: error.message || '无法连接'
        }
      }
    }
  }

  // 检测所有后端状态
  const checkAllBackendsStatus = async () => {
    isCheckingStatus.value = true
    const status = {}
    
    for (const backend of backends.value) {
      status[backend.id] = await checkBackendStatus(backend)
    }
    
    backendStatus.value = status
    isCheckingStatus.value = false
    return status
  }

  // 重置为默认配置
  const resetToDefault = () => {
    backends.value = [...defaultBackends]
    currentBackendId.value = defaultBackends[0].id
    saveBackends()
    localStorage.setItem('currentBackendId', defaultBackends[0].id)
  }

  // 初始化时检测状态
  const init = async () => {
    await checkAllBackendsStatus()
  }

  return {
    backends,
    currentBackendId,
    currentBackend,
    currentBaseURL,
    backendStatus,
    isCheckingStatus,
    setCurrentBackend,
    addBackend,
    updateBackend,
    removeBackend,
    checkBackendStatus,
    checkAllBackendsStatus,
    resetToDefault,
    init
  }
})
