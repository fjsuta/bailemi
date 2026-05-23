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
          <p class="text-slate-400">创建你的专属音乐账号</p>
        </div>

        <!-- 注册表单 -->
        <form @submit.prevent="handleRegister" class="space-y-4">
          <!-- 用户名 -->
          <div class="space-y-2">
            <label class="text-sm font-medium text-slate-300 flex items-center gap-2">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
              </svg>
              用户名
            </label>
            <div class="relative">
              <input
                v-model="form.username"
                type="text"
                placeholder="3-20个字符，中英文数字"
                class="w-full px-4 py-3 bg-slate-800/50 border rounded-xl transition-all duration-200"
                :class="{
                  'border-slate-700 focus:border-purple-500 focus:ring-2 focus:ring-purple-500/20': !errors.username,
                  'border-red-500 focus:border-red-500 focus:ring-2 focus:ring-red-500/20': errors.username,
                  'bg-slate-800/30 cursor-not-allowed opacity-50': loading
                }"
                :disabled="loading"
                @blur="validateField('username')"
              />
              <div v-if="errors.username" class="absolute right-3 top-1/2 -translate-y-1/2">
                <svg class="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
                </svg>
              </div>
            </div>
            <p v-if="errors.username" class="text-sm text-red-400 flex items-center gap-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
              </svg>
              {{ errors.username }}
            </p>
          </div>

          <!-- 邮箱 -->
          <div class="space-y-2">
            <label class="text-sm font-medium text-slate-300 flex items-center gap-2">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"/>
              </svg>
              邮箱
            </label>
            <div class="relative">
              <input
                v-model="form.email"
                type="email"
                placeholder="your@email.com"
                class="w-full px-4 py-3 bg-slate-800/50 border rounded-xl transition-all duration-200"
                :class="{
                  'border-slate-700 focus:border-purple-500 focus:ring-2 focus:ring-purple-500/20': !errors.email,
                  'border-red-500 focus:border-red-500 focus:ring-2 focus:ring-red-500/20': errors.email,
                  'bg-slate-800/30 cursor-not-allowed opacity-50': loading
                }"
                :disabled="loading"
                @blur="validateField('email')"
              />
              <div v-if="errors.email" class="absolute right-3 top-1/2 -translate-y-1/2">
                <svg class="w-5 h-5 text-red-500" fill="currentColor" viewBox="0 0 20 20">
                  <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
                </svg>
              </div>
            </div>
            <p v-if="errors.email" class="text-sm text-red-400 flex items-center gap-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
              </svg>
              {{ errors.email }}
            </p>
          </div>

          <!-- 密码 -->
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
                placeholder="至少8位，包含大小写字母和数字"
                class="w-full px-4 py-3 bg-slate-800/50 border rounded-xl transition-all duration-200"
                :class="{
                  'border-slate-700 focus:border-purple-500 focus:ring-2 focus:ring-purple-500/20': !errors.password,
                  'border-red-500 focus:border-red-500 focus:ring-2 focus:ring-red-500/20': errors.password,
                  'bg-slate-800/30 cursor-not-allowed opacity-50': loading
                }"
                :disabled="loading"
                @blur="validateField('password')"
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
            <!-- 密码强度指示器 -->
            <div class="space-y-2" v-if="form.password">
              <div class="flex gap-1">
                <div
                  v-for="i in 4"
                  :key="i"
                  class="h-1 flex-1 rounded-full transition-all duration-300"
                  :class="getPasswordStrengthClass(i)"
                ></div>
              </div>
              <p class="text-xs text-slate-500">{{ passwordStrengthText }}</p>
            </div>
            <p v-if="errors.password" class="text-sm text-red-400 flex items-center gap-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
              </svg>
              {{ errors.password }}
            </p>
          </div>

          <!-- 确认密码 -->
          <div class="space-y-2">
            <label class="text-sm font-medium text-slate-300 flex items-center gap-2">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"/>
              </svg>
              确认密码
            </label>
            <div class="relative">
              <input
                v-model="form.confirmPassword"
                :type="showPassword ? 'text' : 'password'"
                placeholder="再次输入密码"
                class="w-full px-4 py-3 bg-slate-800/50 border rounded-xl transition-all duration-200"
                :class="{
                  'border-slate-700 focus:border-purple-500 focus:ring-2 focus:ring-purple-500/20': !errors.confirmPassword,
                  'border-red-500 focus:border-red-500 focus:ring-2 focus:ring-red-500/20': errors.confirmPassword,
                  'bg-slate-800/30 cursor-not-allowed opacity-50': loading
                }"
                :disabled="loading"
                @blur="validateField('confirmPassword')"
                @keyup.enter="handleRegister"
              />
            </div>
            <p v-if="errors.confirmPassword" class="text-sm text-red-400 flex items-center gap-1">
              <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"/>
              </svg>
              {{ errors.confirmPassword }}
            </p>
          </div>

          <!-- 服务条款 -->
          <div class="flex items-start gap-3">
            <input
              v-model="form.agreeTerms"
              type="checkbox"
              id="terms"
              class="w-4 h-4 rounded border-slate-600 bg-slate-800/50 text-purple-500 focus:ring-2 focus:ring-purple-500/20 transition-all cursor-pointer mt-0.5"
              :disabled="loading"
            />
            <label for="terms" class="text-sm text-slate-400 cursor-pointer">
              我已阅读并同意
              <a href="#" class="text-purple-400 hover:text-purple-300" @click.prevent>《服务条款》</a>
              和
              <a href="#" class="text-purple-400 hover:text-purple-300" @click.prevent>《隐私政策》</a>
            </label>
          </div>

          <!-- 注册按钮 -->
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
            {{ loading ? '注册中...' : '立即注册' }}
          </button>
        </form>

        <!-- 登录链接 -->
        <div class="text-center mt-6 text-slate-400">
          已有账号？
          <router-link
            to="/login"
            class="text-purple-400 hover:text-purple-300 font-medium transition-colors hover:underline"
          >
            立即登录
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
import { ref, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'

const router = useRouter()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()

const form = ref({
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
  agreeTerms: false
})

const errors = ref({
  username: '',
  email: '',
  password: '',
  confirmPassword: ''
})

const loading = ref(false)
const showPassword = ref(false)

// 密码强度计算
const passwordStrength = computed(() => {
  const password = form.value.password
  if (!password) return 0
  
  let strength = 0
  if (password.length >= 8) strength++
  if (password.length >= 12) strength++
  if (/[a-z]/.test(password)) strength++
  if (/[A-Z]/.test(password)) strength++
  if (/[0-9]/.test(password)) strength++
  if (/[^a-zA-Z0-9]/.test(password)) strength++
  
  return Math.min(strength, 4)
})

const passwordStrengthText = computed(() => {
  const strength = passwordStrength.value
  if (strength <= 1) return '弱 - 建议使用混合字符'
  if (strength === 2) return '中等 - 可以更安全'
  if (strength === 3) return '良好 - 继续加强'
  return '强 - 密码安全'
})

const getPasswordStrengthClass = (index) => {
  const strength = passwordStrength.value
  if (index > strength) {
    return 'bg-slate-700/50'
  }
  if (strength <= 1) {
    return 'bg-red-500'
  }
  if (strength === 2) {
    return 'bg-yellow-500'
  }
  if (strength === 3) {
    return 'bg-blue-500'
  }
  return 'bg-green-500'
}

// 表单验证
const validateField = (field) => {
  if (field === 'username') {
    if (!form.value.username.trim()) {
      errors.value.username = '请输入用户名'
    } else if (form.value.username.length < 3) {
      errors.value.username = '用户名至少3个字符'
    } else if (form.value.username.length > 20) {
      errors.value.username = '用户名最多20个字符'
    } else if (!/^[a-zA-Z0-9\u4e00-\u9fa5]+$/.test(form.value.username)) {
      errors.value.username = '用户名只能包含中英文和数字'
    } else {
      errors.value.username = ''
    }
  }
  
  if (field === 'email') {
    if (!form.value.email.trim()) {
      errors.value.email = '请输入邮箱'
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.value.email)) {
      errors.value.email = '请输入有效的邮箱地址'
    } else {
      errors.value.email = ''
    }
  }
  
  if (field === 'password') {
    if (!form.value.password) {
      errors.value.password = '请输入密码'
    } else if (form.value.password.length < 8) {
      errors.value.password = '密码至少8个字符'
    } else if (!/(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])/.test(form.value.password)) {
      errors.value.password = '密码需包含大小写字母和数字'
    } else {
      errors.value.password = ''
    }
  }
  
  if (field === 'confirmPassword') {
    if (!form.value.confirmPassword) {
      errors.value.confirmPassword = '请确认密码'
    } else if (form.value.confirmPassword !== form.value.password) {
      errors.value.confirmPassword = '两次输入的密码不一致'
    } else {
      errors.value.confirmPassword = ''
    }
  }
}

const validateForm = () => {
  validateField('username')
  validateField('email')
  validateField('password')
  validateField('confirmPassword')
  
  if (!form.value.agreeTerms) {
    notificationStore.warning('请阅读并同意服务条款', '需要同意服务条款才能注册')
    return false
  }
  
  return !errors.value.username && 
         !errors.value.email && 
         !errors.value.password && 
         !errors.value.confirmPassword
}

const isFormValid = computed(() => {
  return form.value.username.trim() &&
         form.value.email.trim() &&
         form.value.password.length >= 8 &&
         form.value.confirmPassword.length >= 8 &&
         form.value.password === form.value.confirmPassword &&
         form.value.agreeTerms
})

const handleRegister = async () => {
  if (!validateForm()) return

  loading.value = true
  try {
    const response = await api.post('/auth/register', {
      username: form.value.username,
      email: form.value.email,
      password: form.value.password
    })
    
    if (response.data.code === 0) {
      const { access_token, refresh_token, ...userData } = response.data.data
      authStore.setTokens(access_token, refresh_token)
      authStore.setUser(userData)
      
      notificationStore.success('注册成功', `欢迎加入百米乐，${userData.username}！`)
      router.push('/profile')
    } else {
      notificationStore.error('注册失败', response.data.message || '注册失败，请稍后重试')
    }
  } catch (error) {
    console.error('Register error:', error)
    
    if (error.response?.data?.message) {
      notificationStore.error('注册失败', error.response.data.message)
    } else if (error.code === 'ECONNREFUSED') {
      notificationStore.error('连接失败', '无法连接到服务器，请稍后重试')
    } else {
      notificationStore.error('注册失败', '注册失败，请检查网络连接')
    }
  } finally {
    loading.value = false
  }
}
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