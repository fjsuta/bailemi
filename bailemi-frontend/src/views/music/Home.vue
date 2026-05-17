<template>
  <div class="min-h-screen pb-24">
    <div class="container mx-auto px-4 pt-6">
      <header class="flex items-center justify-between mb-8">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center text-2xl">
            <svg class="w-7 h-7 text-white" viewBox="0 0 24 24" fill="currentColor">
              <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"/>
            </svg>
          </div>
          <h1 class="text-2xl font-bold gradient-text">百米乐</h1>
        </div>
        <div class="flex items-center gap-4">
          <div v-if="authStore.isAuthenticated" class="flex items-center gap-3">
            <router-link to="/upload" class="px-4 py-2 bg-gradient-to-r from-green-600 to-emerald-600 rounded-lg text-sm font-medium hover:opacity-90 transition-opacity flex items-center gap-2">
              <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M12 19V5"/>
                <path d="M5 12l7-7 7 7"/>
              </svg>
              <span>上传</span>
            </router-link>
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
          
          <button @click="openSearch" class="p-2 hover:bg-white/10 rounded-lg transition-colors" title="搜索">
            <svg class="w-6 h-6 text-slate-300 hover:text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="11" cy="11" r="8"/>
              <path d="m21 21-4.35-4.35"/>
            </svg>
          </button>
        </div>
      </header>

      <div class="text-center py-8 mb-8">
        <h2 class="text-4xl md:text-5xl font-bold gradient-text mb-4">发现好音乐，感受美好</h2>
        <p class="text-slate-400 text-lg">上传原创音乐，分享美好声音</p>
      </div>

      <Teleport to="body">
        <div v-if="showSearch" class="fixed inset-0 z-50 flex items-start justify-center pt-24">
          <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="showSearch = false"></div>
          <div class="relative w-full max-w-2xl mx-4 glass-dark rounded-2xl p-4 shadow-2xl">
            <button @click="showSearch = false" class="absolute top-3 right-3 p-1 hover:bg-white/10 rounded-lg transition-colors">
              <svg class="w-5 h-5 text-slate-400 hover:text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M18 6L6 18"/>
                <path d="m6 6 12 12"/>
              </svg>
            </button>
            <div class="relative">
              <input
                v-model="searchQuery"
                @keyup.enter="handleSearch"
                type="text"
                placeholder="搜索歌曲、歌手、歌单..."
                class="w-full px-6 py-4 bg-white/10 dark:bg-slate-800/50 border border-white/20 rounded-full focus:outline-none focus:border-purple-500 text-white placeholder-slate-400"
                ref="searchInput"
              />
              <button
                @click="handleSearch"
                class="absolute right-2 top-1/2 -translate-y-1/2 p-2 bg-gradient-to-r from-purple-600 to-blue-600 rounded-full hover:from-purple-500 hover:to-blue-500 transition-all"
              >
                <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="11" cy="11" r="8"/>
                  <path d="m21 21-4.35-4.35"/>
                </svg>
              </button>
            </div>
            
            <div v-if="searchResults.length > 0" class="mt-4 max-h-96 overflow-y-auto">
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
                <button class="p-2 bg-purple-600 rounded-full hover:bg-purple-500 transition-colors">
                  <svg class="w-4 h-4 text-white" viewBox="0 0 24 24" fill="currentColor">
                    <polygon points="5,3 19,12 5,21 5,3"/>
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </Teleport>

      <!-- 轮播图 -->
      <section class="mb-12">
        <Banner />
      </section>

      <!-- 每日推荐 -->
      <section class="mb-12">
        <Recommend />
      </section>

      <!-- 排行榜 -->
      <section class="mb-12">
        <Chart />
      </section>

      <!-- 歌单分类墙 -->
      <section class="mb-12">
        <CategoryWall />
      </section>

      <!-- 热门歌单 -->
      <section class="mb-12">
        <h3 class="text-2xl font-bold mb-6">热门歌单</h3>
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
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useMusicStore } from '@/stores/music'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'
import Banner from '@/components/music/Banner.vue'
import Recommend from '@/components/music/Recommend.vue'
import Chart from '@/components/music/Chart.vue'
import CategoryWall from '@/components/music/CategoryWall.vue'

const router = useRouter()
const authStore = useAuthStore()
const musicStore = useMusicStore()
const notificationStore = useNotificationStore()

const searchQuery = ref('')
const searchResults = ref([])
const songs = ref([])
const playlists = ref([])
const showSearch = ref(false)
const searchInput = ref(null)

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

const openSearch = () => {
  showSearch.value = true
  setTimeout(() => {
    searchInput.value?.focus()
  }, 100)
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
    songs.value = []
  }
}

const playSong = (song) => {
  musicStore.playSong(song, songs.value)
}

const handleKeydown = (e) => {
  if (e.target.tagName === 'INPUT' || e.target.tagName === 'TEXTAREA') return
  
  switch (e.code) {
    case 'Space':
      e.preventDefault()
      musicStore.togglePlay()
      break
    case 'ArrowLeft':
      e.preventDefault()
      musicStore.prevSong()
      break
    case 'ArrowRight':
      e.preventDefault()
      musicStore.nextSong()
      break
    case 'ArrowUp':
      e.preventDefault()
      musicStore.volumeUp()
      break
    case 'ArrowDown':
      e.preventDefault()
      musicStore.volumeDown()
      break
  }
}

onMounted(() => {
  loadSongs()
  playlists.value = []
  
  window.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleKeydown)
})
</script>
