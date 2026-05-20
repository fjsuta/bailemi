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
        <div class="mb-8 flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold gradient-text mb-2">歌单</h1>
            <p class="text-slate-400">发现精彩歌单，收藏你的音乐</p>
          </div>
          <button
            @click="showCreateModal = true"
            class="px-6 py-3 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl font-medium hover:from-purple-500 hover:to-blue-500 transition-all flex items-center gap-2"
          >
            <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 5v14M5 12h14"/>
            </svg>
            创建歌单
          </button>
        </div>

        <!-- 标签筛选 -->
        <div class="mb-6">
          <div class="flex items-center gap-2 flex-wrap bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl p-2">
            <button
              v-for="tag in tags"
              :key="tag.value"
              @click="currentTag = tag.value"
              :class="[
                'px-4 py-2 rounded-xl font-medium transition-all',
                currentTag === tag.value
                  ? 'bg-gradient-to-r from-purple-600 to-blue-600 text-white'
                  : 'text-slate-400 hover:text-white hover:bg-white/10'
              ]"
            >
              {{ tag.label }}
            </button>
          </div>
        </div>

        <!-- 排序选项 -->
        <div class="mb-6 flex items-center gap-4">
          <span class="text-slate-400">排序：</span>
          <div class="flex items-center gap-2">
            <button
              v-for="sort in sortOptions"
              :key="sort.value"
              @click="currentSort = sort.value"
              :class="[
                'px-4 py-2 rounded-xl font-medium transition-all',
                currentSort === sort.value
                  ? 'bg-white/20 text-white'
                  : 'text-slate-400 hover:text-white hover:bg-white/10'
              ]"
            >
              {{ sort.label }}
            </button>
          </div>
        </div>

        <!-- 歌单网格 -->
        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
          <div
            v-for="playlist in playlists"
            :key="playlist.id"
            @click="goToPlaylist(playlist.id)"
            class="group cursor-pointer"
          >
            <div class="bg-white/10 backdrop-blur-md border border-white/20 rounded-2xl overflow-hidden hover:bg-white/15 hover:scale-105 transition-all">
              <!-- 封面 -->
              <div class="relative aspect-square">
                <img
                  :src="playlist.cover_url || `https://picsum.photos/300/300?random=${playlist.id}`"
                  class="w-full h-full object-cover"
                  :alt="playlist.title"
                />
                <!-- 播放量 -->
                <div class="absolute bottom-2 right-2 px-2 py-1 bg-black/60 backdrop-blur-sm rounded-lg text-xs text-white flex items-center gap-1">
                  <svg class="w-3 h-3" viewBox="0 0 24 24" fill="currentColor">
                    <path d="M8 5v14l11-7z"/>
                  </svg>
                  {{ formatCount(playlist.play_count) }}
                </div>
                <!-- 播放按钮 -->
                <div class="absolute inset-0 bg-black/40 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-all">
                  <button
                    @click.stop="playPlaylist(playlist.id)"
                    class="w-12 h-12 bg-purple-600 rounded-full flex items-center justify-center hover:bg-purple-500 transition-all"
                  >
                    <svg class="w-6 h-6 text-white" viewBox="0 0 24 24" fill="currentColor">
                      <polygon points="5,3 19,12 5,21 5,3"/>
                    </svg>
                  </button>
                </div>
              </div>

              <!-- 信息 -->
              <div class="p-4">
                <h3 class="font-semibold text-white truncate mb-2">{{ playlist.title }}</h3>
                <div class="flex items-center justify-between text-sm">
                  <span class="text-slate-400">{{ playlist.song_count }} 首</span>
                  <span class="text-slate-400">{{ playlist.creator?.username || '官方' }}</span>
                </div>
                <!-- 标签 -->
                <div v-if="playlist.tags && playlist.tags.length > 0" class="flex items-center gap-1 mt-2 flex-wrap">
                  <span
                    v-for="tag in playlist.tags.slice(0, 3)"
                    :key="tag"
                    class="px-2 py-0.5 bg-white/10 rounded text-xs text-slate-400"
                  >
                    {{ tag }}
                  </span>
                </div>
              </div>
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
        <div v-if="loading && playlists.length === 0" class="flex items-center justify-center py-20">
          <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-purple-500"></div>
        </div>

        <!-- 空状态 -->
        <div v-if="!loading && playlists.length === 0" class="text-center py-20">
          <svg class="w-16 h-16 mx-auto text-slate-500 mb-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M9 19V6l12-3v13M9 19c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zm12-3c0 1.105-1.343 2-3 2s-3-.895-3-2 1.343-2 3-2 3 .895 3 2zM9 10l12-3"/>
          </svg>
          <p class="text-slate-400">暂无歌单</p>
        </div>
      </div>
    </div>

    <!-- 创建歌单弹窗 -->
    <Teleport to="body">
      <div v-if="showCreateModal" class="fixed inset-0 z-50 flex items-center justify-center">
        <div class="absolute inset-0 bg-black/60 backdrop-blur-sm" @click="showCreateModal = false"></div>
        <div class="relative w-full max-w-md mx-4 bg-white/10 backdrop-blur-xl rounded-2xl border border-white/20 p-6 shadow-2xl">
          <div class="flex items-center justify-between mb-6">
            <h2 class="text-xl font-bold">创建歌单</h2>
            <button @click="showCreateModal = false" class="p-2 hover:bg-white/10 rounded-xl transition-all">
              <svg class="w-5 h-5 text-slate-400 hover:text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M18 6L6 18"/>
                <path d="m6 6 12 12"/>
              </svg>
            </button>
          </div>

          <form @submit.prevent="createPlaylist" class="space-y-4">
            <div>
              <label class="block text-sm font-medium mb-2">歌单名称</label>
              <input
                v-model="newPlaylist.title"
                type="text"
                placeholder="给你的歌单起个名字"
                class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl focus:outline-none focus:border-purple-500 text-white placeholder-slate-400"
                required
              />
            </div>

            <div>
              <label class="block text-sm font-medium mb-2">描述</label>
              <textarea
                v-model="newPlaylist.description"
                placeholder="描述一下你的歌单..."
                rows="3"
                class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl focus:outline-none focus:border-purple-500 text-white placeholder-slate-400 resize-none"
              ></textarea>
            </div>

            <div>
              <label class="block text-sm font-medium mb-2">标签（用逗号分隔）</label>
              <input
                v-model="newPlaylist.tagsInput"
                type="text"
                placeholder="例如：学习, 运动, 放松"
                class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl focus:outline-none focus:border-purple-500 text-white placeholder-slate-400"
              />
            </div>

            <div class="flex items-center gap-2">
              <input
                v-model="newPlaylist.isPublic"
                type="checkbox"
                id="isPublic"
                class="w-4 h-4 rounded border-white/20 bg-white/10 text-purple-600 focus:ring-purple-500"
              />
              <label for="isPublic" class="text-sm">公开歌单</label>
            </div>

            <button
              type="submit"
              :disabled="creating"
              class="w-full py-3 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl font-medium hover:from-purple-500 hover:to-blue-500 transition-all disabled:opacity-50"
            >
              {{ creating ? '创建中...' : '创建' }}
            </button>
          </form>
        </div>
      </div>
    </Teleport>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useSidebarStore } from '@/stores/sidebar'
import { useNotificationStore } from '@/stores/notification'
import { useAuthStore } from '@/stores/auth'
import api from '@/utils/api'
import Sidebar from '@/components/Sidebar.vue'

const router = useRouter()
const sidebarStore = useSidebarStore()
const notificationStore = useNotificationStore()
const authStore = useAuthStore()

const tags = [
  { label: '全部', value: 'all' },
  { label: '华语', value: '华语' },
  { label: '欧美', value: '欧美' },
  { label: '摇滚', value: '摇滚' },
  { label: '民谣', value: '民谣' },
  { label: '流行', value: '流行' },
  { label: '电子', value: '电子' },
  { label: '学习', value: '学习' },
  { label: '运动', value: '运动' },
  { label: '睡眠', value: '睡眠' }
]

const sortOptions = [
  { label: '最热', value: 'hot' },
  { label: '最新', value: 'new' }
]

const currentTag = ref('all')
const currentSort = ref('hot')
const playlists = ref([])
const loading = ref(false)
const page = ref(1)
const pageSize = 20
const hasMore = ref(true)

const showCreateModal = ref(false)
const creating = ref(false)
const newPlaylist = ref({
  title: '',
  description: '',
  tagsInput: '',
  isPublic: true
})

const fetchPlaylists = async (reset = false) => {
  if (loading.value) return

  loading.value = true
  if (reset) {
    page.value = 1
    playlists.value = []
    hasMore.value = true
  }

  try {
    const tag = currentTag.value === 'all' ? '' : currentTag.value
    const response = await api.get(`/playlist?page=${page.value}&pageSize=${pageSize}&tag=${tag}&sort=${currentSort.value}`)
    if (response.data.code === 0) {
      const newPlaylists = response.data.data?.list || []
      playlists.value = reset ? newPlaylists : [...playlists.value, ...newPlaylists]
      hasMore.value = newPlaylists.length >= pageSize
      if (newPlaylists.length > 0) {
        page.value++
      }
    }
  } catch (error) {
    console.error('获取歌单列表失败:', error)
    notificationStore.error('获取歌单列表失败', '请稍后重试')
  } finally {
    loading.value = false
  }
}

const createPlaylist = async () => {
  if (!authStore.isAuthenticated) {
    notificationStore.warning('请先登录', '')
    router.push('/login')
    return
  }

  creating.value = true
  try {
    const tags = newPlaylist.value.tagsInput
      .split(/[,，]/)
      .map(tag => tag.trim())
      .filter(tag => tag)

    const response = await api.post('/playlist', {
      title: newPlaylist.value.title,
      description: newPlaylist.value.description || undefined,
      tags: tags,
      is_public: newPlaylist.value.isPublic
    })

    if (response.data.code === 0) {
      notificationStore.success('创建成功', '歌单已创建')
      showCreateModal.value = false
      newPlaylist.value = { title: '', description: '', tagsInput: '', isPublic: true }
      fetchPlaylists(true)
    }
  } catch (error) {
    console.error('创建歌单失败:', error)
    notificationStore.error('创建失败', error.response?.data?.message || '请稍后重试')
  } finally {
    creating.value = false
  }
}

const goToPlaylist = (id) => {
  router.push(`/playlist/${id}`)
}

const playPlaylist = async (id) => {
  try {
    const response = await api.get(`/playlist/${id}`)
    if (response.data.code === 0 && response.data.data?.songs?.length > 0) {
      const songs = response.data.data.songs
      // 播放第一首歌，传入整个歌单
      // musicStore.playSong(songs[0], songs)
      notificationStore.success('开始播放', `正在播放歌单：${response.data.data.title}`)
    } else {
      notificationStore.warning('歌单为空', '该歌单没有歌曲')
    }
  } catch (error) {
    console.error('播放歌单失败:', error)
    notificationStore.error('播放失败', '请稍后重试')
  }
}

const loadMore = () => {
  fetchPlaylists(false)
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

watch([currentTag, currentSort], () => {
  fetchPlaylists(true)
})

onMounted(() => {
  fetchPlaylists(true)
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
