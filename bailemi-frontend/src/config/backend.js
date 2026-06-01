// 后端服务配置
// 用于管理多个后端服务的连接

const config = {
  // 开发环境
  development: {
    // Spring Boot 认证服务
    auth: {
      baseURL: import.meta.env.VITE_AUTH_API_BASE_URL || 'http://localhost:8080',
      prefix: '/api/v1',
      name: 'Spring Boot Auth Service'
    },
    // Go Gateway 核心服务
    core: {
      baseURL: import.meta.env.VITE_CORE_API_BASE_URL || 'http://localhost:8081',
      prefix: '/v1',
      name: 'Go Gateway Core Service'
    }
  },

  // 生产环境
  production: {
    auth: {
      baseURL: import.meta.env.VITE_AUTH_API_BASE_URL || 'https://api-auth.bailemi.com',
      prefix: '/api/v1',
      name: 'Spring Boot Auth Service'
    },
    core: {
      baseURL: import.meta.env.VITE_CORE_API_BASE_URL || 'https://api.bailemi.com',
      prefix: '/v1',
      name: 'Go Gateway Core Service'
    }
  }
}

const env = import.meta.env.MODE || 'development'
const currentConfig = config[env] || config.development

// API 端点路由映射
const endpointRoutes = {
  // 认证相关 -> Spring Boot
  '/auth': 'auth',
  '/admin': 'auth',
  '/oauth': 'auth',

  // 核心业务 -> Go Gateway
  '/user': 'core',
  '/song': 'core',
  '/album': 'core',
  '/artist': 'core',
  '/playlist': 'core',
  '/search': 'core',
  '/play': 'core',
  '/rank': 'core',
  '/comment': 'core',
  '/genres': 'core'
}

// 根据端点路径获取对应的服务
export function getServiceByEndpoint(endpoint) {
  for (const [prefix, service] of Object.entries(endpointRoutes)) {
    if (endpoint.startsWith(prefix)) {
      return service
    }
  }
  // 默认使用核心服务
  return 'core'
}

// 获取完整 API URL
export function getApiUrl(endpoint, service = null) {
  const targetService = service || getServiceByEndpoint(endpoint)
  const serviceConfig = currentConfig[targetService]
  return `${serviceConfig.baseURL}${serviceConfig.prefix}${endpoint}`
}

export { currentConfig as backendConfig }
export default currentConfig
