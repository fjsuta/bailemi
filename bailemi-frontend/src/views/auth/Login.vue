<template>
  <div class="min-h-screen flex items-center justify-center relative overflow-hidden">
    <div class="absolute inset-0 overflow-hidden">
      <div class="absolute -top-40 -left-40 w-80 h-80 bg-gradient-to-br from-purple-500/30 to-purple-500/10 rounded-full blur-3xl animate-float"></div>
      <div class="absolute -bottom-40 -right-40 w-80 h-80 bg-gradient-to-br from-blue-500/30 to-blue-500/10 rounded-full blur-3xl animate-float" style="animation-delay: -3s"></div>
    </div>

    <div class="glass-dark rounded-3xl p-8 w-full max-w-md mx-4 relative z-10">
      <div class="text-center mb-8">
        <div class="flex items-center justify-center gap-3 mb-4">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center text-2xl">🎵</div>
          <h1 class="text-3xl font-bold gradient-text">百米乐</h1>
        </div>
        <p class="text-slate-400">欢迎回来！继续你的音乐之旅</p>
      </div>

      <!-- 第三方登录 -->
      <div v-if="enabledProviders.length > 0" class="mb-6">
        <div class="relative">
          <div class="absolute inset-0 flex items-center">
            <div class="w-full border-t border-slate-700"></div>
          </div>
          <div class="relative flex justify-center text-sm">
            <span class="px-2 bg-transparent text-slate-400">其他登录方式</span>
          </div>
        </div>
        
        <div class="grid grid-cols-5 gap-3 mt-4">
          <button
            v-for="provider in enabledProviders"
            :key="provider"
            @click="handleOAuthLogin(provider)"
            class="p-3 bg-white/10 hover:bg-white/20 rounded-xl transition-all flex items-center justify-center"
            :title="getProviderName(provider)"
          >
            <span class="text-2xl">{{ getProviderIcon(provider) }}</span>
          </button>
        </div>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-6">
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-300">账号</label>
          <input
            v-model="form.account"
            type="text"
            placeholder="用户名/邮箱/手机号"
            class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none transition-all"
          />
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-300">密码</label>
          <input
            v-model="form.password"
            type="password"
            placeholder="请输入密码"
            class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none transition-all"
          />
        </div>
        <button
          :disabled="loading"
          type="submit"
          class="w-full py-4 bg-gradient-to-r from-purple-600 to-blue-600 hover:from-purple-500 hover:to-blue-500 rounded-xl font-semibold transition-all disabled:opacity-50"
        >
          {{ loading ? '登录中...' : '登录' }}
        </button>
      </form>

      <div class="text-center mt-6 text-slate-400">
        还没有账号？
        <router-link to="/register" class="text-purple-400 hover:text-purple-300 font-medium">立即注册</router-link>
      </div>
      
      <div class="text-center mt-4">
        <router-link to="/" class="text-slate-400 hover:text-white text-sm">返回首页</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'

const router = useRouter()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()

const form = ref({
  account: '',
  password: ''
})
const loading = ref(false)
const enabledProviders = ref([])

const providerIcons = {
  google: '🔵',
  microsoft: '🟢',
  apple: '🍎',
  wechat: '💚',
  qq: '🐧'
}

const providerNames = {
  google: 'Google',
  microsoft: 'Microsoft',
  apple: 'Apple',
  wechat: '微信',
  qq: 'QQ'
}

const getProviderIcon = (provider) => providerIcons[provider] || '❓'
const getProviderName = (provider) => providerNames[provider] || provider

const handleLogin = async () => {
  if (!form.value.account || !form.value.password) {
    notificationStore.warning('请填写完整信息', '用户名和密码不能为空')
    return
  }

  loading.value = true
  try {
    const response = await api.post('/auth/login', {
      login_type: 'username',
      account: form.value.account,
      password: form.value.password
    })
    if (response.data.code === 0) {
      const { access_token, refresh_token, ...userData } = response.data.data
      authStore.setTokens(access_token, refresh_token)
      authStore.setUser(userData)
      notificationStore.success('登录成功', `欢迎回来，${userData.username || userData.nickname}！`)
      router.push('/profile')
    } else {
      notificationStore.error('登录失败', response.data.message || '用户名或密码错误')
    }
  } catch (error) {
    console.error(error)
    notificationStore.error('登录失败', '请检查网络连接或账号密码是否正确')
  } finally {
    loading.value = false
  }
}

const handleOAuthLogin = (provider) => {
  // 跳转到后端OAuth授权页面
  const callbackUrl = encodeURIComponent(window.location.origin + '/#/oauth/' + provider)
  window.location.href = `http://localhost:8080/api/v1/oauth/${provider}/authorize?callback=${callbackUrl}`
}

const loadOAuthConfig = async () => {
  try {
    const response = await api.get('/auth/oauth/config')
    if (response.data.code === 0) {
      enabledProviders.value = response.data.data?.enabled_providers || []
    }
  } catch (error) {
    console.error(error)
    enabledProviders.value = ['google', 'microsoft', 'apple', 'wechat', 'qq']
  }
}

onMounted(() => {
  loadOAuthConfig()
})
</script>
