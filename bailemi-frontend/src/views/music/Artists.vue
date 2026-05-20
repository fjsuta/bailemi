<template>
  <div class="min-h-screen pb-40">
    <!-- 侧边栏 -->
    <Sidebar />

    <div
      :style="sidebarStore.layout === 'left'
        ? { paddingLeft: sidebarStore.width + 'px', paddingTop: '0', transition: 'padding-left 0.15s ease-out, padding-top 0.15s ease-out' }
        : { paddingLeft: '0', paddingTop: sidebarStore.height + 'px', transition: 'padding-left 0.15s ease-out, padding-top 0.15s ease-out' }"
    >
      <div class="container mx-auto px-4 py-6">
        <!-- 页面标题 -->
        <div class="mb-8">
          <h1 class="text-3xl font-bold gradient-text mb-2">歌手</h1>
          <p class="text-slate-400">发现你喜欢的音乐人</p>
        </div>

        <!-- 筛选栏 -->
        <div class="mb-6 space-y-4">
          <!-- 排序选项 -->
          <div class="flex items-center gap-2 bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl p-2">
            <button
              v-for="sort in sortOptions"
              :key="sort.value"
              @click="currentSort = sort.value"
              :class="[
                'px-4 py-2 rounded-xl font-medium transition-all',
                currentSort === sort.value
                  ? 'bg-gradient-to-r from-purple-600 to-blue-600 text-white'
                  : 'text-slate-400 hover:text-white hover:bg-white/10'
              ]"
            >
              {{ sort.label }}
            </button>
          </div>

          <!-- 字母索引 -->
          <div class="flex items-center gap-1 flex-wrap bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl p-2">
            <button
              v-for="letter in letters"
              :key="letter"
              @click="currentLetter = letter"
              :class="[
                'w-8 h-8 rounded-lg font-medium transition-all text-sm',
                currentLetter === letter
                  ? 'bg-gradient-to-r from-purple-600 to-blue-600 text-white'
                  : 'text-slate-400 hover:text-white hover:bg-white/10'
              ]"
            >
              {{ letter }}
            </button>
          </div>
        </div>

        <!-- 歌手网格 -->
        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
          <div
            v-for="artist in artists"
            :key="artist.id"
            @click="goToArtist(artist.id)"
            class="group cursor-pointer"
          >
            <div class="bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl p-4 hover:bg-white/15 hover:scale-105 transition-all">
              <!-- 头像 -->
              <div class="relative mb-4">
                <img
                  :src="artist.avatar_url || `https://picsum.photos/200/200?random=${artist.id}`"
                  class="w-full aspect-square object-cover rounded-xl"
                  :alt="artist.name"
                />
                <!-- 认证标识 -->
                <div
                  v-if="artist.is_verified"
                  class="absolute bottom-2 right-2 w-6 h-6 bg-blue-500 rounded-full flex items-center justify-center"
                >
                  <svg class="w-4 h-4 text-white" viewBox="0 0 24 24" fill="currentColor">
                    <path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/>
                  </svg>
                </div>
              </div>

              <!-- 信息 -->
              <h3 class="font-semibold text-white truncate mb-1">{{ artist.name }}</h3>
              <p class="text-sm text-slate-400 truncate">
                {{ artist.fan_count ? formatCount(artist.fan_count) + ' 粉丝' : '音乐人' }}
              </p>
            </div>
          </div>
        </div>

        <!-- 加载更多 -->
        <div v-if="hasMore" class="flex justify-center mt-8">
          <button
            @click="loadMore"
            :disabled="loading"
            class="px-8 py-3 bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl hover:bg-white/15 transition-all disabled:opacity-50"
          >
            {{ loading ? '加载中...' : '加载更多' }}
          </button>
        </div>

        <!-- 加载状态 -->
        <div v-if="loading && artists.length === 0" class="flex items-center justify-center py-20">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-purple-500"></div>
        </div>

        <!-- 空状态 -->
        <div v-if="!loading && artists.length === 0" class="text-center py-20">
          <svg class="w-16 h-16 mx-auto text-slate-500 mb-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/>
            <circle cx="9" cy="7" r="4"/>
            <path d="M22 21v-2a4 4 0 0 0-3-3.87"/>
            <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
          </svg>
          <p class="text-slate-400">暂无歌手</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useSidebarStore } from '@/stores/sidebar'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'
import Sidebar from '@/components/Sidebar.vue'

const router = useRouter()
const sidebarStore = useSidebarStore()
const notificationStore = useNotificationStore()

const sortOptions = [
  { label: '热门', value: 'hot' },
  { label: '最新', value: 'new' },
  { label: '字母', value: 'name' }
]

const letters = ['全部', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']

const currentSort = ref('hot')
const currentLetter = ref('全部')
const artists = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = 20
const hasMore = ref(true)

const fetchArtists = async (reset = false) => {
  if (loading.value) return

  loading.value = true
  if (reset) {
    page.value = 1
    artists.value = []
    hasMore.value = true
  }

  try {
    const letter = currentLetter.value === '全部' ? '' : currentLetter.value
    const response = await api.get(`/artist?page=${page.value}&pageSize=${pageSize}&sort=${currentSort.value}&letter=${letter}`)
    if (response.data.code === 0) {
      const newArtists = response.data.data?.list || []
      artists.value = reset ? newArtists : [...artists.value, ...newArtists]
      hasMore.value = newArtists.length >= pageSize
      if (newArtists.length > 0) {
        page.value++
      }
    }
  } catch (error) {
    console.error('获取歌手列表失败:', error)
    notificationStore.error('获取歌手列表失败', '请稍后重试')
  } finally {
    loading.value = false
  }
}

const goToArtist = (id) => {
  router.push(`/artist/${id}`)
}

const loadMore = () => {
  fetchArtists(false)
}

const formatCount = (count) => {
  if (!count) return '0'
  if (count >= 100000000) {
    return (count / 100000000).toFixed(1) + '亿'
  }
  if (count >= 10000) {
    return (count / 10000).toFixed(1) + '万'
  }
  return count.toString()
}

watch([currentSort, currentLetter], () => {
  fetchArtists(true)
})

onMounted(() => {
  fetchArtists(true)
})
</script>

<style scoped>
.gradient-text {
  background: linear-gradient(135deg, #8B5CF6 0%, #3B82F6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}
</style>
