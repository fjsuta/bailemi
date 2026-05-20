<template>
  <div class="space-y-8">
    <!-- Tab 切换 -->
    <div class="flex items-center gap-2 bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl p-2">
      <button
        v-for="tab in tabs"
        :key="tab.value"
        @click="currentTab = tab.value"
        :class="[
          'px-6 py-3 rounded-xl font-medium transition-all',
          currentTab === tab.value 
            ? 'bg-gradient-to-r from-purple-600 to-blue-600 text-white' 
            : 'text-slate-400 hover:text-white hover:bg-white/10'
        ]"
      >
        {{ tab.label }}
      </button>
    </div>

    <!-- 排行榜列表 -->
    <div class="space-y-2">
      <div
        v-for="(song, index) in songs"
        :key="song.id"
        @click="playSong(song, index)"
        class="flex items-center gap-4 p-4 bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl hover:bg-white/15 hover:scale-[1.02] transition-all cursor-pointer group"
      >
        <!-- 排名 -->
        <div class="w-12 h-12 flex items-center justify-center">
          <span
            v-if="index < 3"
            :class="[
              'text-2xl font-bold',
              index === 0 ? 'text-yellow-400' : '',
              index === 1 ? 'text-gray-300' : '',
              index === 2 ? 'text-orange-400' : ''
            ]"
          >
            {{ index + 1 }}
          </span>
          <span v-else class="text-xl text-slate-400">{{ index + 1 }}</span>
        </div>

        <!-- 封面 -->
        <img
          :src="song.cover_url || `https://picsum.photos/56/56?random=${song.id}`"
          class="w-14 h-14 rounded-xl object-cover"
          :alt="song.title"
        />

        <!-- 歌曲信息 -->
        <div class="flex-1 min-w-0">
          <h4 class="font-semibold text-white truncate">{{ song.title }}</h4>
          <p class="text-sm text-slate-400 truncate">{{ song.artist?.name || '未知歌手' }}</p>
        </div>

        <!-- 播放量 -->
        <div class="text-right hidden md:block">
          <p class="text-sm text-slate-400">{{ formatPlayCount(song.play_count) }}</p>
          <p class="text-xs text-slate-500">播放</p>
        </div>

        <!-- 操作按钮 -->
        <button
          @click.stop="addToQueue(song)"
          class="p-3 bg-purple-600/80 hover:bg-purple-600 rounded-xl transition-all opacity-0 group-hover:opacity-100"
          title="加入播放队列"
        >
          <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="currentColor">
            <polygon points="5,3 19,12 5,21 5,3"/>
          </svg>
        </button>
      </div>
    </div>

    <!-- 加载更多 -->
    <div v-if="hasMore" class="flex justify-center">
      <button
        @click="loadMore"
        :disabled="loading"
        class="px-8 py-3 bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl hover:bg-white/15 transition-all disabled:opacity-50"
      >
        {{ loading ? '加载中...' : '加载更多' }}
      </button>
    </div>

    <!-- 加载状态 -->
    <div v-if="loading && songs.length === 0" class="flex items-center justify-center py-20">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-purple-500"></div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useMusicStore } from '@/stores/music'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'

const musicStore = useMusicStore()
const notificationStore = useNotificationStore()

const tabs = [
  { label: '热歌榜', value: 'hot' },
  { label: '新歌榜', value: 'new' },
  { label: '飙升榜', value: 'rising' }
]

const currentTab = ref('hot')
const songs = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = 20
const hasMore = ref(true)

const fetchCharts = async (reset = false) => {
  if (loading.value) return
  
  loading.value = true
  if (reset) {
    page.value = 1
    songs.value = []
    hasMore.value = true
  }

  try {
    const response = await api.get(`/charts?type=${currentTab.value}&limit=${pageSize}&page=${page.value}`)
    if (response.data.code === 0) {
      const newSongs = response.data.data?.songs || []
      songs.value = reset ? newSongs : [...songs.value, ...newSongs]
      hasMore.value = newSongs.length >= pageSize
      if (newSongs.length > 0) {
        page.value++
      }
    }
  } catch (error) {
    console.error('获取排行榜失败:', error)
    notificationStore.error('获取排行榜失败', '请稍后重试')
  } finally {
    loading.value = false
  }
}

const playSong = (song, index) => {
  musicStore.playSong(song, songs.value)
}

const addToQueue = (song) => {
  musicStore.addToQueue(song)
  notificationStore.success('已添加到播放队列', song.title)
}

const loadMore = () => {
  fetchCharts(false)
}

const formatPlayCount = (count) => {
  if (!count) return '0'
  if (count >= 100000000) {
    return (count / 100000000).toFixed(1) + '亿'
  }
  if (count >= 10000) {
    return (count / 10000).toFixed(1) + '万'
  }
  return count.toString()
}

watch(currentTab, () => {
  fetchCharts(true)
})

onMounted(() => {
  fetchCharts(true)
})
</script>

<style scoped>
</style>
