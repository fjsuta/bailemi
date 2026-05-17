<template>
  <div class="min-h-screen flex items-center justify-center relative overflow-hidden">
    <div class="absolute inset-0 overflow-hidden">
      <div class="absolute -top-40 -left-40 w-80 h-80 bg-gradient-to-br from-purple-500/30 to-purple-500/10 rounded-full blur-3xl animate-float"></div>
      <div class="absolute -bottom-40 -right-40 w-80 h-80 bg-gradient-to-br from-blue-500/30 to-blue-500/10 rounded-full blur-3xl animate-float" style="animation-delay: -3s"></div>
    </div>

    <div class="glass-dark rounded-3xl p-8 w-full max-w-md mx-4 relative z-10">
      <div class="text-center mb-8">
        <div class="flex items-center justify-center gap-3 mb-4">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center">
            <svg class="w-7 h-7 text-white" viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"/>
            </svg>
          </div>
          <h1 class="text-3xl font-bold gradient-text">百米乐</h1>
        </div>
        <p class="text-slate-400">欢迎回来！继续你的音乐之旅</p>
      </div>

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
            <component :is="getProviderIcon(provider)" class="w-6 h-6" />
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
import { ref, onMounted, h } from 'vue'
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

const GoogleIcon = {
  render() {
    return h('svg', { viewBox: '0 0 24 24', fill: 'none' }, [
      h('path', { d: 'M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92a5.06 5.06 0 0 1-2.2 3.32v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.1z', fill: '#4285F4' }),
      h('path', { d: 'M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z', fill: '#34A853' }),
      h('path', { d: 'M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z', fill: '#FBBC05' }),
      h('path', { d: 'M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z', fill: '#EA4335' })
    ])
  }
}

const MicrosoftIcon = {
  render() {
    return h('svg', { viewBox: '0 0 24 24', fill: 'none' }, [
      h('rect', { x: '1', y: '1', width: '10', height: '10', fill: '#F25022' }),
      h('rect', { x: '13', y: '1', width: '10', height: '10', fill: '#7FBA00' }),
      h('rect', { x: '1', y: '13', width: '10', height: '10', fill: '#00A4EF' }),
      h('rect', { x: '13', y: '13', width: '10', height: '10', fill: '#FFB900' })
    ])
  }
}

const AppleIcon = {
  render() {
    return h('svg', { viewBox: '0 0 24 24', fill: 'currentColor' }, [
      h('path', { d: 'M17.05 20.28c-.98.95-2.05.88-3.08.4-1.09-.5-2.08-.48-3.24 0-1.44.62-2.2.44-3.06-.4C2.79 15.25 3.51 7.59 9.05 7.31c1.35.07 2.29.74 3.08.8 1.18-.24 2.31-.93 3.57-.84 1.51.12 2.65.72 3.4 1.8-3.12 1.87-2.38 5.98.48 7.13-.57 1.5-1.31 2.99-2.54 4.09zM12.03 7.25c-.15-2.23 1.66-4.07 3.74-4.25.29 2.58-2.34 4.5-3.74 4.25z' })
    ])
  }
}

const WechatIcon = {
  render() {
    return h('svg', { viewBox: '0 0 24 24', fill: 'none' }, [
      h('path', { d: 'M8.691 2.188C3.891 2.188 0 5.476 0 9.53c0 2.212 1.17 4.203 3.002 5.55a.59.59 0 0 1 .213.665l-.39 1.48c-.019.07-.048.141-.048.213 0 .163.13.295.29.295a.326.326 0 0 0 .167-.054l1.903-1.114a.864.864 0 0 1 .717-.098 10.16 10.16 0 0 0 2.837.403c.276 0 .543-.027.811-.05-.857-2.578.157-4.972 1.932-6.446 1.703-1.415 3.882-1.98 5.853-1.838-.576-3.583-4.196-6.348-8.596-6.348zM5.785 5.991c.642 0 1.162.529 1.162 1.18a1.17 1.17 0 0 1-1.162 1.178A1.17 1.17 0 0 1 4.623 7.17c0-.651.52-1.18 1.162-1.18zm5.813 0c.642 0 1.162.529 1.162 1.18a1.17 1.17 0 0 1-1.162 1.178 1.17 1.17 0 0 1-1.162-1.178c0-.651.52-1.18 1.162-1.18zm3.97 3.258c-2.031 0-3.898.753-5.309 1.915-1.66 1.368-2.562 3.275-2.562 5.298 0 1.372.465 2.666 1.252 3.78a.527.527 0 0 1 .085.396l-.229 1.098c-.032.12-.06.248-.06.377 0 .164.13.296.291.296a.326.326 0 0 0 .166-.053l1.32-.774a.774.774 0 0 1 .642-.087 8.44 8.44 0 0 0 2.404.35c4.07 0 7.395-2.91 7.395-6.382 0-3.472-3.325-6.382-7.395-6.382v.068zm-2.196 3.34c.542 0 .981.446.981.996a.989.989 0 0 1-.981.996.989.989 0 0 1-.982-.996c0-.55.44-.996.982-.996zm4.392 0c.542 0 .982.446.982.996a.989.989 0 0 1-.982.996.989.989 0 0 1-.982-.996c0-.55.44-.996.982-.996z', fill: '#07C160' })
    ])
  }
}

const QQIcon = {
  render() {
    return h('svg', { viewBox: '0 0 24 24', fill: 'none' }, [
      h('path', { d: 'M12 2C7.589 2 4 5.589 4 9.996c0 1.928.691 3.691 1.835 5.07-.18.636-.41 1.288-.688 1.928-.36.835-.76 1.645-1.17 2.355-.14.24-.04.55.21.66.73.32 1.71.56 2.63.56.74 0 1.43-.14 1.98-.39A7.94 7.94 0 0 0 12 20c1.12 0 2.18-.23 3.15-.64.57.27 1.28.42 2.04.42.92 0 1.9-.24 2.63-.56a.488.488 0 0 0 .21-.66c-.41-.71-.81-1.52-1.17-2.355-.28-.65-.51-1.31-.69-1.955A7.963 7.963 0 0 0 20 9.996C20 5.589 16.411 2 12 2zm-2.5 8a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3zm5 0a1.5 1.5 0 1 1 0-3 1.5 1.5 0 0 1 0 3z', fill: '#12B7F5' })
    ])
  }
}

const DefaultIcon = {
  render() {
    return h('svg', { viewBox: '0 0 24 24', fill: 'none', stroke: 'currentColor', 'stroke-width': '2' }, [
      h('circle', { cx: '12', cy: '12', r: '10' }),
      h('path', { d: 'M12 8v4M12 16h.01' })
    ])
  }
}

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

const getProviderIcon = (provider) => providerIcons[provider] || DefaultIcon
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
