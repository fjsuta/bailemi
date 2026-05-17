<template>
  <div class="min-h-screen pb-24">
    <div class="container mx-auto px-4 pt-6">
      <header class="flex items-center justify-between mb-8">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center text-2xl">🎵</div>
          <h1 class="text-2xl font-bold gradient-text">百米乐</h1>
        </div>
        <div class="flex items-center gap-4">
          <div v-if="authStore.isAuthenticated" class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-full bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center overflow-hidden">
              <img v-if="authStore.user?.avatar_url" :src="authStore.user.avatar_url" class="w-full h-full object-cover" />
              <span v-else>{{ authStore.user?.username?.[0]?.toUpperCase() }}</span>
            </div>
            <span class="text-slate-300">{{ authStore.user?.username || authStore.user?.nickname }}</span>
            <button @click="handleLogout" class="px-3 py-1 bg-red-600 hover:bg-red-500 rounded-lg text-sm transition-all">退出</button>
          </div>
          <div v-else class="flex items-center gap-3">
            <router-link to="/login" class="text-slate-300 hover:text-white">登录</router-link>
            <router-link to="/register" class="px-4 py-2 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl hover:from-purple-500 hover:to-blue-500 transition-all">注册</router-link>
          </div>
        </div>
      </header>

      <div class="text-center py-16 mb-8">
        <h2 class="text-5xl font-bold gradient-text mb-4">发现好音乐，感受美好</h2>
        <p class="text-slate-400 text-lg mb-8">百万首歌曲，随时随地，想听就听</p>
        
        <!-- 搜索框 -->
        <div class="max-w-2xl mx-auto">
          <div class="relative">
            <input
              v-model="searchQuery"
              @keyup.enter="handleSearch"
              type="text"
              placeholder="搜索歌曲、歌手、歌单..."
              class="w-full px-6 py-4 bg-white/10 dark:bg-slate-800/50 border border-white/20 rounded-full focus:outline-none focus:border-purple-500 text-white placeholder-slate-400"
            />
            <button
              @click="handleSearch"
              class="absolute right-2 top-1/2 -translate-y-1/2 px-6 py-2 bg-gradient-to-r from-purple-600 to-blue-600 rounded-full hover:from-purple-500 hover:to-blue-500 transition-all"
            >
              🔍 搜索
            </button>
          </div>
          
          <!-- 搜索结果 -->
          <div v-if="searchResults.length > 0" class="mt-4 glass-dark rounded-2xl p-4 max-h-96 overflow-y-auto">
            <h4 class="font-semibold mb-3">搜索结果</h4>
            <div
              v-for="song in searchResults"
              :key="song.id"
              @click="playSong(song)"
              class="flex items-center gap-3 p-3 hover:bg-white/10 rounded-lg cursor-pointer transition-all"
            >
              <img :src="song.cover_url || 'https://picsum.photos/56/56?random=' + song.id" class="w-12 h-12 rounded-lg object-cover" />
              <div class="flex-1">
                <h5 class="font-medium">{{ song.title }}</h5>
                <p class="text-sm text-slate-400">{{ typeof song.artist === 'string' ? song.artist : song.artist?.name }}</p>
              </div>
              <button class="px-3 py-1 bg-purple-600 rounded-full text-sm hover:bg-purple-500">▶</button>
            </div>
          </div>
        </div>
      </div>

      <section class="mb-12">
        <h3 class="text-2xl font-bold mb-6">热门歌曲</h3>
        <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-4">
          <div
            v-for="song in songs"
            :key="song.id"
            @click="playSong(song)"
            class="glass-dark rounded-2xl p-4 cursor-pointer hover:scale-105 transition-all"
          >
            <div class="relative mb-3">
              <img :src="song.cover_url || 'https://picsum.photos/200/200?random=' + song.id" class="w-full aspect-square object-cover rounded-xl" />
              <div class="absolute inset-0 flex items-center justify-center bg-black/50 rounded-xl opacity-0 hover:opacity-100 transition-all">
                <div class="w-12 h-12 bg-gradient-to-br from-purple-500 to-blue-500 rounded-full flex items-center justify-center text-xl">▶</div>
              </div>
            </div>
            <h4 class="font-semibold truncate">{{ song.title }}</h4>
            <p class="text-slate-400 text-sm truncate">{{ typeof song.artist === 'string' ? song.artist : song.artist?.name }}</p>
          </div>
        </div>
      </section>

      <section class="mb-12">
        <h3 class="text-2xl font-bold mb-6">推荐歌单</h3>
        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          <div v-for="playlist in playlists" :key="playlist.id" class="glass-dark rounded-2xl p-4 cursor-pointer hover:scale-105 transition-all">
            <img :src="playlist.cover_url || 'https://picsum.photos/200/200?random=p' + playlist.id" class="w-full aspect-square object-cover rounded-xl mb-3" />
            <h4 class="font-semibold truncate">{{ playlist.title }}</h4>
            <p class="text-slate-400 text-sm">{{ playlist.description }}</p>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useMusicStore } from '@/stores/music'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'

const router = useRouter()
const authStore = useAuthStore()
const musicStore = useMusicStore()
const notificationStore = useNotificationStore()

const searchQuery = ref('')
const searchResults = ref([])
const songs = ref([])
const playlists = ref([])

const handleSearch = async () => {
  if (!searchQuery.value.trim()) {
    notificationStore.warning('请输入搜索内容', '')
    return
  }
  
  try {
    const response = await api.get(`/search?keyword=${encodeURIComponent(searchQuery.value)}`)
    if (response.data.code === 0) {
      searchResults.value = response.data.data?.songs || response.data.data || []
      if (searchResults.value.length === 0) {
        notificationStore.info('搜索结果', '未找到相关歌曲')
      }
    }
  } catch (error) {
    console.error(error)
    notificationStore.error('搜索失败', '请稍后重试')
  }
}

const handleLogout = () => {
  authStore.logout()
  notificationStore.success('已退出登录', '欢迎再次使用百米乐')
  router.push('/')
}

const loadSongs = async () => {
  try {
    const response = await api.get('/song/hot')
    if (response.data.code === 0) {
      songs.value = response.data.data?.list || response.data.data || []
    }
  } catch (error) {
    console.error(error)
    songs.value = [
      { id: 1, title: '双截棍', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=1' },
      { id: 2, title: '简单爱', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=2' },
      { id: 3, title: '晴天', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=3' },
      { id: 4, title: '七里香', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=4' },
      { id: 5, title: '稻香', artist: '周杰伦', cover_url: 'https://picsum.photos/200/200?random=5' }
    ]
  }
}

const playSong = (song) => {
  musicStore.playSong(song, songs.value)
}

onMounted(() => {
  loadSongs()
  playlists.value = [
    { id: 1, title: '华语流行精选', description: '精选华语流行歌曲', cover_url: 'https://picsum.photos/200/200?random=p1' },
    { id: 2, title: '经典老歌回忆', description: '经典老歌回忆杀', cover_url: 'https://picsum.photos/200/200?random=p2' },
    { id: 3, title: '新歌速递', description: '最新最热歌曲', cover_url: 'https://picsum.photos/200/200?random=p3' },
    { id: 4, title: '深夜电台', description: '陪你度过深夜的歌', cover_url: 'https://picsum.photos/200/200?random=p4' }
  ]
})
</script>
