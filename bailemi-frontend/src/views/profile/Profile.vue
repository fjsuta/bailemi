<template>
  <div class="min-h-screen pb-24">
    <div class="container mx-auto px-4 pt-6">
      <header class="flex items-center justify-between mb-8">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center text-2xl">🎵</div>
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
            <button v-if="isEditing" @click="triggerAvatarUpload" class="absolute bottom-0 right-0 w-10 h-10 bg-purple-600 rounded-full flex items-center justify-center text-xl hover:bg-purple-500 transition-all">
              📷
            </button>
            <input ref="avatarInput" type="file" accept="image/*" hidden @change="handleAvatarUpload" />
          </div>
          <div class="flex-1 text-center md:text-left">
            <div v-if="!isEditing" class="flex items-center justify-center md:justify-start gap-4 mb-2">
              <h2 class="text-3xl font-bold">{{ userProfile?.username || authStore.user?.username }}</h2>
              <button @click="isEditing = true" class="px-4 py-2 glass rounded-xl hover:bg-white/20 transition-all">
                ✏️ 编辑资料
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
              <button class="text-slate-400 hover:text-purple-400">❤️</button>
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
                <span class="text-3xl">{{ provider.icon }}</span>
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
import { ref, computed, onMounted } from 'vue'
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

const oauthProviders = [
  { id: 'google', name: 'Google', icon: '🔵' },
  { id: 'microsoft', name: 'Microsoft', icon: '🟢' },
  { id: 'apple', name: 'Apple', icon: '🍎' },
  { id: 'wechat', name: '微信', icon: '💚' },
  { id: 'qq', name: 'QQ', icon: '🐧' }
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
