<template>
  <div class="min-h-screen pb-24">
    <div class="container mx-auto px-4 pt-6">
      <header class="flex items-center justify-between mb-8">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center">
            <svg class="w-7 h-7 text-white" viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"/>
            </svg>
          </div>
          <h1 class="text-2xl font-bold gradient-text">百米乐</h1>
        </div>
        <nav class="flex items-center gap-6">
          <router-link to="/" class="text-slate-400 hover:text-white">首页</router-link>
          <span class="text-white font-medium">个人中心</span>
          <button @click="handleLogout" class="text-slate-400 hover:text-white">退出</button>
        </nav>
      </header>

      <div class="glass-dark rounded-3xl p-8 mb-8">
        <div class="flex flex-col md:flex-row items-center gap-6">
          <div class="relative">
            <img
              :src="userProfile?.avatar_url || `https://picsum.photos/200/200?random=${authStore.user?.username || 'user'}`"
              alt="头像"
              class="w-32 h-32 rounded-full object-cover border-4 border-purple-500/50"
            />
            <button v-if="isEditing" @click="triggerAvatarUpload" class="absolute bottom-0 right-0 w-10 h-10 bg-purple-600 rounded-full flex items-center justify-center hover:bg-purple-500 transition-all">
              <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
                <circle cx="12" cy="13" r="4"/>
              </svg>
            </button>
            <input ref="avatarInput" type="file" accept="image/*" hidden @change="handleAvatarUpload" />
          </div>
          <div class="flex-1 text-center md:text-left">
            <div v-if="!isEditing" class="flex items-center justify-center md:justify-start gap-4 mb-2">
              <h2 class="text-3xl font-bold">{{ userProfile?.username || authStore.user?.username }}</h2>
              <button @click="isEditing = true" class="px-4 py-2 glass rounded-xl hover:bg-white/20 transition-all flex items-center gap-2">
                <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                  <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                </svg>
                编辑资料
              </button>
            </div>
            <div v-else class="flex items-center justify-center md:justify-start gap-4 mb-2">
              <input v-model="editForm.nickname" placeholder="昵称" class="px-4 py-2 bg-slate-800 rounded-xl border border-slate-700" />
              <button @click="saveProfile" class="px-4 py-2 bg-purple-600 rounded-xl hover:bg-purple-500 transition-all">
                保存
              </button>
              <button @click="cancelEdit" class="px-4 py-2 bg-slate-700 rounded-xl hover:bg-slate-600 transition-all">
                取消
              </button>
            </div>
            <div class="flex items-center justify-center md:justify-start gap-4 mb-2">
              <span class="px-3 py-1 bg-gradient-to-r from-yellow-500 to-orange-500 rounded-full text-sm font-semibold">
                VIP {{ vipLevel }}
              </span>
              <span class="text-slate-400">粉丝: {{ userProfile?.stats?.fan_count || 0 }}</span>
              <span class="text-slate-400">关注: {{ userProfile?.stats?.follow_count || 0 }}</span>
            </div>
            <div v-if="!isEditing" class="text-slate-400">
              {{ userProfile?.profile?.bio || '这个人很懒，什么都没写~' }}
            </div>
            <div v-else>
              <textarea v-model="editForm.bio" placeholder="写点什么..." class="w-full max-w-lg px-4 py-2 bg-slate-800 rounded-xl border border-slate-700" rows="3" />
            </div>
          </div>
        </div>
      </div>

      <div class="mb-4">
        <div class="flex gap-4 border-b border-slate-700 overflow-x-auto">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="[
              'px-6 py-3 font-semibold transition-all whitespace-nowrap',
              activeTab === tab.id
                ? 'text-purple-400 border-b-2 border-purple-400'
                : 'text-slate-400 hover:text-white'
            ]"
          >
            {{ tab.name }}
          </button>
        </div>
      </div>

      <div class="glass-dark rounded-3xl p-6">
        <div v-if="activeTab === 'favorites'">
          <h3 class="text-xl font-bold mb-4">我的收藏</h3>
          <div v-if="favorites.length === 0" class="text-center py-12 text-slate-400">
            还没有收藏歌曲呢，快去发现好音乐吧~
          </div>
          <div v-else class="space-y-2">
            <div
              v-for="(song, index) in favorites"
              :key="song.id"
              @click="playSong(song)"
              class="flex items-center gap-4 p-3 rounded-xl hover:bg-white/10 cursor-pointer transition-all"
            >
              <span class="w-8 text-slate-400">{{ index + 1 }}</span>
              <img :src="song.cover_url || 'https://picsum.photos/56/56?random=f' + song.id" class="w-14 h-14 rounded-lg object-cover" />
              <div class="flex-1 min-w-0">
                <h4 class="font-medium truncate">{{ song.title }}</h4>
                <p class="text-slate-400 text-sm truncate">{{ typeof song.artist === 'string' ? song.artist : song.artist?.name }}</p>
              </div>
              <button class="text-slate-400 hover:text-purple-400">
                <svg class="w-5 h-5" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
                </svg>
              </button>
            </div>
          </div>
        </div>

        <div v-else-if="activeTab === 'recent'">
          <h3 class="text-xl font-bold mb-4">最近播放</h3>
          <div v-if="recentPlays.length === 0" class="text-center py-12 text-slate-400">
            还没有播放记录呢~
          </div>
          <div v-else class="space-y-2">
            <div
              v-for="(song, index) in recentPlays"
              :key="song.id"
              @click="playSong(song)"
              class="flex items-center gap-4 p-3 rounded-xl hover:bg-white/10 cursor-pointer transition-all"
            >
              <span class="w-8 text-slate-400">{{ index + 1 }}</span>
              <img :src="song.cover_url || 'https://picsum.photos/56/56?random=r' + song.id" class="w-14 h-14 rounded-lg object-cover" />
              <div class="flex-1 min-w-0">
                <h4 class="font-medium truncate">{{ song.title }}</h4>
                <p class="text-slate-400 text-sm truncate">{{ typeof song.artist === 'string' ? song.artist : song.artist?.name }}</p>
              </div>
              <span class="text-slate-500 text-sm">{{ song.played_at }}</span>
            </div>
          </div>
        </div>

        <div v-else-if="activeTab === 'oauth'">
          <h3 class="text-xl font-bold mb-4">第三方账号绑定</h3>
          <p class="text-slate-400 mb-6">绑定第三方账号可以使用以下方式登录百米乐</p>
          
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            <div 
              v-for="provider in oauthProviders" 
              :key="provider.id"
              class="glass-dark rounded-xl p-4 flex items-center justify-between"
            >
              <div class="flex items-center gap-3">
                <component :is="providerIconMap[provider.id]" class="w-8 h-8" />
                <div>
                  <h4 class="font-medium">{{ provider.name }}</h4>
                  <p class="text-sm text-slate-400">
                    {{ isProviderBound(provider.id) ? '已绑定' : '未绑定' }}
                  </p>
                </div>
              </div>
              <div>
                <button 
                  v-if="isProviderBound(provider.id)"
                  @click="handleUnbind(provider.id)"
                  class="px-3 py-1 bg-red-600 hover:bg-red-500 rounded-lg text-sm transition-all"
                >
                  解除
                </button>
                <button 
                  v-else-if="enabledProviders.includes(provider.id)"
                  @click="handleBind(provider.id)"
                  class="px-3 py-1 bg-purple-600 hover:bg-purple-500 rounded-lg text-sm transition-all"
                >
                  绑定
                </button>
                <span 
                  v-else
                  class="text-sm text-slate-500"
                >
                  未启用
                </span>
              </div>
            </div>
          </div>
        </div>

        <div v-else-if="activeTab === 'security'">
          <h3 class="text-xl font-bold mb-6">账号安全</h3>
          
          <div class="space-y-6 max-w-lg">
            <div class="glass-dark rounded-xl p-4">
              <h4 class="font-medium mb-4">修改密码</h4>
              <div class="space-y-3">
                <input 
                  v-model="passwordForm.oldPassword"
                  type="password"
                  placeholder="请输入原密码"
                  class="w-full px-4 py-2 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none"
                />
                <input 
                  v-model="passwordForm.newPassword"
                  type="password"
                  placeholder="请输入新密码（至少8位）"
                  class="w-full px-4 py-2 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none"
                />
                <input 
                  v-model="passwordForm.confirmPassword"
                  type="password"
                  placeholder="请确认新密码"
                  class="w-full px-4 py-2 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none"
                />
                <button 
                  @click="handleChangePassword"
                  :disabled="passwordLoading"
                  class="w-full py-2 bg-purple-600 hover:bg-purple-500 rounded-xl font-medium transition-all disabled:opacity-50"
                >
                  {{ passwordLoading ? '修改中...' : '确认修改' }}
                </button>
              </div>
            </div>

            <div class="glass-dark rounded-xl p-4 border border-red-500/30">
              <h4 class="font-medium mb-2 text-red-400">危险区域</h4>
              <p class="text-sm text-slate-400 mb-4">注销账号后，所有数据将被永久删除，此操作不可撤销</p>
              <button 
                @click="showDeleteModal = true"
                class="px-4 py-2 bg-red-600 hover:bg-red-500 rounded-xl font-medium transition-all"
              >
                注销账号
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="showDeleteModal" class="fixed inset-0 bg-black/70 flex items-center justify-center z-50 p-4">
      <div class="glass-dark rounded-2xl p-6 max-w-md w-full">
        <h3 class="text-xl font-bold mb-4 text-red-400">确认注销账号</h3>
        <p class="text-slate-400 mb-4">请输入您的密码以确认注销账号。此操作将永久删除您的所有数据。</p>
        <input 
          v-model="deletePassword"
          type="password"
          placeholder="请输入密码"
          class="w-full px-4 py-2 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-red-500 focus:outline-none mb-4"
        />
        <div class="flex gap-3">
          <button 
            @click="showDeleteModal = false"
            class="flex-1 py-2 bg-slate-700 hover:bg-slate-600 rounded-xl font-medium transition-all"
          >
            取消
          </button>
          <button 
            @click="handleDeleteAccount"
            :disabled="deleteLoading"
            class="flex-1 py-2 bg-red-600 hover:bg-red-500 rounded-xl font-medium transition-all disabled:opacity-50"
          >
            {{ deleteLoading ? '注销中...' : '确认注销' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, h } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useMusicStore } from '@/stores/music'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'

const router = useRouter()
const authStore = useAuthStore()
const musicStore = useMusicStore()
const notificationStore = useNotificationStore()

const activeTab = ref('favorites')
const isEditing = ref(false)
const userProfile = ref(null)
const avatarInput = ref(null)
const vipLevel = ref(1)
const enabledProviders = ref([])
const boundProviders = ref([])

const showDeleteModal = ref(false)
const deletePassword = ref('')
const deleteLoading = ref(false)
const passwordLoading = ref(false)

const passwordForm = ref({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const tabs = [
  { id: 'favorites', name: '我的收藏' },
  { id: 'recent', name: '最近播放' },
  { id: 'oauth', name: '账号绑定' },
  { id: 'security', name: '账号安全' }
]

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

const providerIconMap = {
  google: GoogleIcon,
  microsoft: MicrosoftIcon,
  apple: AppleIcon,
  wechat: WechatIcon,
  qq: QQIcon
}

const oauthProviders = [
  { id: 'google', name: 'Google' },
  { id: 'microsoft', name: 'Microsoft' },
  { id: 'apple', name: 'Apple' },
  { id: 'wechat', name: '微信' },
  { id: 'qq', name: 'QQ' }
]

const editForm = ref({
  nickname: '',
  bio: ''
})

const favorites = ref([
  { id: 1, title: '双截棍', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=1' },
  { id: 2, title: '晴天', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=2' },
  { id: 3, title: '七里香', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=3' },
  { id: 4, title: '稻香', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=4' }
])

const recentPlays = ref([
  { id: 1, title: '简单爱', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=5', played_at: '2分钟前' },
  { id: 2, title: '夜曲', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=6', played_at: '30分钟前' },
  { id: 3, title: '青花瓷', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=7', played_at: '1小时前' }
])

const isProviderBound = (providerId) => {
  return boundProviders.value.some(b => b.provider === providerId)
}

const loadUserProfile = async () => {
  try {
    const response = await api.get('/user/me')
    if (response.data.code === 0) {
      userProfile.value = response.data.data
      editForm.value.nickname = response.data.data.profile?.nickname || ''
      editForm.value.bio = response.data.data.profile?.bio || ''
    }
  } catch (error) {
    console.error(error)
  }
}

const loadOAuthConfig = async () => {
  try {
    const response = await api.get('/auth/oauth/config')
    if (response.data.code === 0) {
      enabledProviders.value = response.data.data?.enabled_providers || []
    }
  } catch (error) {
    console.error(error)
    enabledProviders.value = oauthProviders.map(p => p.id)
  }
}

const loadBoundProviders = async () => {
  try {
    const response = await api.get('/oauth/bindings')
    if (response.data.code === 0) {
      boundProviders.value = response.data.data || []
    }
  } catch (error) {
    console.error(error)
  }
}

const handleBind = (provider) => {
  const callbackUrl = encodeURIComponent(window.location.origin + '/#/oauth/' + provider)
  window.location.href = `http://localhost:8080/api/v1/oauth/${provider}/authorize?callback=${callbackUrl}`
}

const handleUnbind = async (provider) => {
  if (!confirm(`确定要解除${provider}账号的绑定吗？`)) return
  
  try {
    await api.delete(`/oauth/bind/${provider}`)
    notificationStore.success('解除绑定成功', '')
    loadBoundProviders()
  } catch (error) {
    console.error(error)
    notificationStore.error('解除绑定失败', '请稍后重试')
  }
}

const saveProfile = async () => {
  try {
    await api.put('/user/profile', {
      nickname: editForm.value.nickname,
      bio: editForm.value.bio
    })
    notificationStore.success('资料更新成功', '')
    isEditing.value = false
    loadUserProfile()
  } catch (error) {
    console.error(error)
    notificationStore.error('更新失败', '请重试')
  }
}

const cancelEdit = () => {
  isEditing.value = false
  editForm.value.nickname = userProfile.value?.profile?.nickname || ''
  editForm.value.bio = userProfile.value?.profile?.bio || ''
}

const triggerAvatarUpload = () => {
  avatarInput.value?.click()
}

const handleAvatarUpload = async (e) => {
  const file = e.target.files[0]
  if (!file) return

  const formData = new FormData()
  formData.append('avatar', file)

  try {
    const response = await api.post('/user/avatar', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    if (response.data.code === 0) {
      notificationStore.success('头像更新成功', '')
      loadUserProfile()
    }
  } catch (error) {
    console.error(error)
    notificationStore.error('头像上传失败', '')
  }
}

const handleChangePassword = async () => {
  if (passwordForm.value.newPassword !== passwordForm.value.confirmPassword) {
    notificationStore.warning('密码不一致', '两次输入的新密码不一致')
    return
  }
  
  if (passwordForm.value.newPassword.length < 8) {
    notificationStore.warning('密码太短', '新密码长度至少8位')
    return
  }

  passwordLoading.value = true
  try {
    await api.post('/user/password/change', {
      old_password: passwordForm.value.oldPassword,
      new_password: passwordForm.value.newPassword
    })
    notificationStore.success('密码修改成功', '下次登录请使用新密码')
    passwordForm.value = { oldPassword: '', newPassword: '', confirmPassword: '' }
  } catch (error) {
    console.error(error)
    notificationStore.error('密码修改失败', error.response?.data?.message || '请检查原密码是否正确')
  } finally {
    passwordLoading.value = false
  }
}

const handleDeleteAccount = async () => {
  if (!deletePassword.value) {
    notificationStore.warning('请输入密码', '需要密码确认才能注销账号')
    return
  }

  deleteLoading.value = true
  try {
    await api.post('/user/delete', {
      password: deletePassword.value
    })
    notificationStore.success('账号已注销', '感谢您使用百米乐，再见！')
    authStore.logout()
    router.push('/')
  } catch (error) {
    console.error(error)
    notificationStore.error('账号注销失败', error.response?.data?.message || '请检查密码是否正确')
  } finally {
    deleteLoading.value = false
  }
}

const handleLogout = () => {
  authStore.logout()
  notificationStore.success('已退出登录', '欢迎再次使用百米乐')
  router.push('/')
}

const playSong = (song) => {
  const allSongs = [...favorites.value, ...recentPlays.value]
  musicStore.playSong(song, allSongs)
}

onMounted(() => {
  loadUserProfile()
  loadOAuthConfig()
  loadBoundProviders()
})
</script>
