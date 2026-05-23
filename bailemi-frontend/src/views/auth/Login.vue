<template>
  <div class="min-h-screen flex items-center justify-center relative overflow-hidden bg-[#0F0F23]">
    <!-- 背景动画 -->
    <div class="absolute inset-0 overflow-hidden">
      <div class="absolute -top-40 -left-40 w-80 h-80 bg-gradient-to-br from-purple-500/30 to-purple-500/10 rounded-full blur-3xl animate-float"></div>
      <div class="absolute -bottom-40 -right-40 w-80 h-80 bg-gradient-to-br from-blue-500/30 to-blue-500/10 rounded-full blur-3xl animate-float" style="animation-delay: -3s"></div>
      <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[600px] h-[600px] bg-gradient-to-br from-indigo-500/10 to-purple-500/5 rounded-full blur-3xl"></div>
    </div>

    <div class="relative z-10 w-full max-w-md mx-4">
      <div class="glass-dark rounded-3xl p-8 backdrop-blur-xl border border-slate-800/50">
        <!-- Logo 和标题 -->
        <div class="text-center mb-8">
          <div class="flex items-center justify-center gap-3 mb-4">
            <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center">
              <svg class="w-7 h-7 text-white" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"/>
              </svg>
            </div>
            <h1 class="text-3xl font-bold gradient-text">百米乐</h1>
          </div>
          <p class="text-slate-400">欢迎回来！继续你的音乐之旅</p>
        </div>

        <!-- 登录表单 -->
        <form @submit.prevent="handleLogin" class="space-y-5">
          <!-- 账号输入 -->
          <div class="space-y-2">
            <label class="text-sm font-medium text-slate-300 flex items-center gap-2">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
              </svg>
              账号
            </label>
            <div class="relative">
              <input
                v-model="form.account"
                type="text"
                placeholder="用户名 / 邮箱 / 手机号"
                class="w-full px-4 py-3 bg-slate-800/50 border rounded-xl transition-all duration-200"
                :class="{
                  'border-slate-700 focus:border-purple-500 focus:ring-2 focus:ring-purple-500/20': !errors.account,
                  'border-red-500 focus:border-red-500 focus:ring-2 focus:ring-red-500/20': errors.account,
                  'bg-slate-800/30 cursor-not-allowed opacity-50': loading
                }"
                :disabled="loading"
                @blur="validateField('account')"
              />
              <div v-if="errors.account" class="absolute right-3 top-1/2 -translate-y-1/2">
                <svg class="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
                </svg>
              </div>
            </div>
            <p v-if="errors.account" class="text-sm text-red-400 flex items-center gap-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
              </svg>
              {{ errors.account }}
            </p>
          </div>

          <!-- 密码输入 -->
          <div class="space-y-2">
            <label class="text-sm font-medium text-slate-300 flex items-center gap-2">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"/>
              </svg>
              密码
            </label>
            <div class="relative">
              <input
                v-model="form.password"
                :type="showPassword ? 'text' : 'password'"
                placeholder="请输入密码"
                class="w-full px-4 py-3 bg-slate-800/50 border rounded-xl transition-all duration-200"
                :class="{
                  'border-slate-700 focus:border-purple-500 focus:ring-2 focus:ring-purple-500/20': !errors.password,
                  'border-red-500 focus:border-red-500 focus:ring-2 focus:ring-red-500/20': errors.password,
                  'bg-slate-800/30 cursor-not-allowed opacity-50': loading
                }"
                :disabled="loading"
                @blur="validateField('password')"
                @keyup.enter="handleLogin"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 hover:text-slate-300 transition-colors"
                :disabled="loading"
              >
                <svg v-if="showPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21"/>
                </svg>
                <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"/>
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"/>
                </svg>
              </button>
            </div>
            <p v-if="errors.password" class="text-sm text-red-400 flex items-center gap-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
              </svg>
              {{ errors.password }}
            </p>
          </div>

          <!-- 记住我和忘记密码 -->
          <div class="flex items-center justify-between text-sm">
            <label class="flex items-center gap-2 cursor-pointer group">
              <input
                v-model="form.remember"
                type="checkbox"
                class="w-4 h-4 rounded border-slate-600 bg-slate-800/50 text-purple-500 focus:ring-2 focus:ring-purple-500/20 transition-all cursor-pointer"
                :disabled="loading"
              />
              <span class="text-slate-400 group-hover:text-slate-300 transition-colors">记住我</span>
            </label>
            <button
              type="button"
              class="text-slate-400 hover:text-purple-400 transition-colors"
              :disabled="loading"
            >
              忘记密码？
            </button>
          </div>

          <!-- 登录按钮 -->
          <button
            type="submit"
            :disabled="loading || !isFormValid"
            class="w-full py-4 bg-gradient-to-r from-purple-600 to-blue-600 hover:from-purple-500 hover:to-blue-500 rounded-xl font-semibold transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2 group"
            :class="{ 'shadow-lg shadow-purple-500/30': !loading && isFormValid }"
          >
            <svg v-if="loading" class="w-5 h-5 animate-spin" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            {{ loading ? '登录中...' : '登录' }}
          </button>
        </form>

        <!-- 第三方登录 -->
        <div v-if="enabledProviders.length > 0" class="mt-6">
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-slate-700/50"></div>
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="px-3 bg-slate-900/50 text-slate-500">其他登录方式</span>
            </div>
          </div>
          
          <div class="grid grid-cols-4 gap-3 mt-4">
            <button
              v-for="provider in enabledProviders"
              :key="provider"
              @click="handleOAuthLogin(provider)"
              :disabled="loading"
              class="p-3 bg-slate-800/50 hover:bg-slate-700/50 border border-slate-700/50 hover:border-slate-600/50 rounded-xl transition-all flex items-center justify-center group"
              :title="getProviderName(provider)"
            >
              <component :is="getProviderIcon(provider)" class="w-5 h-5 text-slate-400 group-hover:text-white transition-colors" />
            </button>
          </div>
        </div>

        <!-- 注册链接 -->
        <div class="text-center mt-6 text-slate-400">
          还没有账号？
          <router-link
            to="/register"
            class="text-purple-400 hover:text-purple-300 font-medium transition-colors hover:underline"
          >
            立即注册
          </router-link>
        </div>
        
        <div class="text-center mt-3">
          <router-link
            to="/"
            class="text-slate-500 hover:text-slate-400 text-sm transition-colors"
          >
            返回首页
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'
import GoogleIcon from '@/components/icons/GoogleIcon.vue'
import MicrosoftIcon from '@/components/icons/MicrosoftIcon.vue'
import AppleIcon from '@/components/icons/AppleIcon.vue'
import WechatIcon from '@/components/icons/WechatIcon.vue'
import QQIcon from '@/components/icons/QQIcon.vue'

const router = useRouter()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()

const form = ref({
  account: '',
  password: '',
  remember: false
})

const errors = ref({
  account: '',
  password: ''
})

const loading = ref(false)
const showPassword = ref(false)
const enabledProviders = ref([])

const providerIcons = {
  google: GoogleIcon,
  microsoft: MicrosoftIcon,
  apple: AppleIcon,
  wechat: WechatIcon,
  qq: QQIcon
}

const providerNames = {
  google: 'Google',
  microsoft: 'Microsoft',
  apple: 'Apple',
  wechat: '微信',
  qq: 'QQ'
}

const getProviderIcon = (provider) => providerIcons[provider]
const getProviderName = (provider) => providerNames[provider] || provider

// 表单验证
const validateField = (field) => {
  if (field === 'account') {
    if (!form.value.account.trim()) {
      errors.value.account = '请输入账号'
    } else if (form.value.account.length < 3) {
      errors.value.account = '账号至少3个字符'
    } else {
      errors.value.account = ''
    }
  }
  
  if (field === 'password') {
    if (!form.value.password) {
      errors.value.password = '请输入密码'
    } else if (form.value.password.length < 6) {
      errors.value.password = '密码至少6个字符'
    } else {
      errors.value.password = ''
    }
  }
}

const validateForm = () => {
  validateField('account')
  validateField('password')
  return !errors.value.account && !errors.value.password
}

const isFormValid = computed(() => {
  return form.value.account.trim() && form.value.password.length >= 6
})

const handleLogin = async () => {
  if (!validateForm()) return

  loading.value = true
  try {
    const response = await api.post('/auth/login', {
      account: form.value.account,
      password: form.value.password
    })
    
    if (response.data.code === 0) {
      const { access_token, refresh_token, ...userData } = response.data.data
      authStore.setTokens(access_token, refresh_token)
      authStore.setUser(userData)
      
      if (form.value.remember) {
        localStorage.setItem('remember_account', form.value.account)
      }
      
      notificationStore.success('登录成功', `欢迎回来，${userData.username || userData.nickname}！`)
      router.push('/profile')
    } else {
      notificationStore.error('登录失败', response.data.message || '用户名或密码错误')
    }
  } catch (error) {
    console.error('Login error:', error)
    
    if (error.response?.data?.message) {
      notificationStore.error('登录失败', error.response.data.message)
    } else if (error.code === 'ECONNREFUSED') {
      notificationStore.error('连接失败', '无法连接到服务器，请稍后重试')
    } else {
      notificationStore.error('登录失败', '用户名或密码错误，请检查后重试')
    }
  } finally {
    loading.value = false
  }
}

const handleOAuthLogin = (provider) => {
  const callbackUrl = encodeURIComponent(window.location.origin + '/#/oauth/' + provider)
  window.location.href = `${import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'}/api/v1/oauth/${provider}/authorize?callback=${callbackUrl}`
}

const loadOAuthConfig = async () => {
  try {
    const response = await api.get('/auth/oauth/config')
    if (response.data.code === 0) {
      enabledProviders.value = response.data.data?.enabled_providers || []
    }
  } catch (error) {
    console.error('Failed to load OAuth config:', error)
    enabledProviders.value = []
  }
}

onMounted(() => {
  const savedAccount = localStorage.getItem('remember_account')
  if (savedAccount) {
    form.value.account = savedAccount
    form.value.remember = true
  }
  loadOAuthConfig()
})
</script>

<style scoped>
.glass-dark {
  @apply bg-slate-900/80 backdrop-blur-xl;
}

.gradient-text {
  @apply bg-gradient-to-r from-purple-400 via-pink-400 to-blue-400 bg-clip-text text-transparent;
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}
</style>