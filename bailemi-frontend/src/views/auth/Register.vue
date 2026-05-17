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
        <p class="text-slate-400">创建你的专属音乐账号</p>
      </div>

      <form @submit.prevent="handleRegister" class="space-y-5">
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-300">用户名</label>
          <input
            v-model="form.username"
            type="text"
            placeholder="请输入用户名"
            class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none transition-all"
          />
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-300">邮箱</label>
          <input
            v-model="form.email"
            type="email"
            placeholder="请输入邮箱"
            class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none transition-all"
          />
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-300">密码</label>
          <input
            v-model="form.password"
            type="password"
            placeholder="至少8位密码"
            class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none transition-all"
          />
        </div>
        <div class="space-y-2">
          <label class="text-sm font-medium text-slate-300">验证码</label>
          <div class="flex gap-3">
            <input
              v-model="form.verify_code"
              type="text"
              placeholder="请输入验证码"
              class="flex-1 px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none transition-all"
            />
            <button type="button" class="px-4 py-3 bg-slate-700 hover:bg-slate-600 rounded-xl transition-all">获取验证码</button>
          </div>
        </div>
        <button
          :disabled="loading"
          type="submit"
          class="w-full py-4 bg-gradient-to-r from-purple-600 to-blue-600 hover:from-purple-500 hover:to-blue-500 rounded-xl font-semibold transition-all"
        >
          {{ loading ? '注册中...' : '立即注册' }}
        </button>
      </form>

      <div class="text-center mt-6 text-slate-400">
        已有账号？
        <router-link to="/login" class="text-purple-400 hover:text-purple-300 font-medium">立即登录</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import api from '@/utils/api'

const router = useRouter()
const authStore = useAuthStore()

const form = ref({
  username: '',
  email: '',
  phone: '',
  password: '',
  verify_code: '123456'
})
const loading = ref(false)

const handleRegister = async () => {
  loading.value = true
  try {
    const response = await api.post('/auth/register', form.value)
    if (response.data.code === 0) {
      const { access_token, refresh_token, ...userData } = response.data.data
      authStore.setTokens(access_token, refresh_token)
      authStore.setUser(userData)
      router.push('/profile')
    } else {
      alert(response.data.message)
    }
  } catch (error) {
    console.error(error)
    alert('注册失败，请检查网络')
  } finally {
    loading.value = false
  }
}
</script>
