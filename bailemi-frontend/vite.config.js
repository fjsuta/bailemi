import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { fileURLToPath, URL } from 'node:url'

export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd(), '')

  const authApiUrl = env.VITE_AUTH_API_BASE_URL || 'http://localhost:8080'
  const coreApiUrl = env.VITE_CORE_API_BASE_URL || 'http://localhost:8081'

  return {
    plugins: [vue()],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    server: {
      port: 3000,
      proxy: {
        // Spring Boot 认证服务代理
        '/api/auth': {
          target: authApiUrl,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '/api')
        },
        '/api/admin': {
          target: authApiUrl,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '/api')
        },
        // Go Gateway 核心服务代理
        '/api/v1': {
          target: coreApiUrl,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api\/v1/, '/v1')
        }
      }
    },
    base: '/bailemi/'
  }
})
