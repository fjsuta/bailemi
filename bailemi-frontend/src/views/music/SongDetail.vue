<template>
  <div class="min-h-screen pb-24">
    <div class="container mx-auto px-4 py-8 max-w-6xl">
      <div class="flex items-center gap-4 mb-8">
        <button @click="$router.back()" class="w-10 h-10 glass rounded-xl flex items-center justify-center hover:bg-white/20 transition-all">
          ←
        </button>
        <div>
          <h1 class="text-2xl font-bold text-white">{{ song?.title }}</h1>
          <p class="text-slate-400">{{ song?.artist }}</p>
        </div>
      </div>

      <div class="grid lg:grid-cols-3 gap-8">
        <div class="lg:col-span-1">
          <div class="glass-dark rounded-2xl overflow-hidden sticky top-8">
            <div class="aspect-square bg-slate-800">
              <img
                :src="song?.cover || 'https://picsum.photos/400/400?random=1'"
                :alt="song?.title"
                class="w-full h-full object-cover"
              />
            </div>
            <div class="p-6">
              <div class="flex items-center gap-3 mb-4">
                <div
                  class="w-12 h-12 rounded-full bg-gradient-to-br from-purple-600 to-blue-600 flex items-center justify-center text-2xl cursor-pointer hover:opacity-90 transition-opacity"
                  @click="togglePlay"
                >
                  <svg v-if="isPlaying" class="w-6 h-6 text-white" viewBox="0 0 24 24" fill="currentColor">
                    <rect x="6" y="4" width="4" height="16"/>
                    <rect x="14" y="4" width="4" height="16"/>
                  </svg>
                  <svg v-else class="w-6 h-6 text-white ml-1" viewBox="0 0 24 24" fill="currentColor">
                    <polygon points="5 3 19 12 5 21 5 3"/>
                  </svg>
                </div>
                <div class="flex-1">
                  <div class="flex items-center gap-3 mb-2">
                    <button class="w-10 h-10 flex items-center justify-center hover:text-pink-500 transition-colors"
                      :class="{ 'text-pink-500': isLiked }"
                    >
                      <svg v-if="isLiked" class="w-5 h-5" viewBox="0 0 24 24" fill="currentColor">
                        <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
                      </svg>
                      <svg v-else class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
                      </svg>
                    </button>
                    <button class="w-10 h-10 flex items-center justify-center hover:text-purple-500 transition-colors">
                      <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/>
                        <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>
                      </svg>
                    </button>
                    <button class="w-10 h-10 flex items-center justify-center hover:text-green-500 transition-colors">
                      <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"/>
                        <path d="M15.54 8.46a5 5 0 0 1 0 7.07"/>
                      </svg>
                    </button>
                    <button v-if="song?.allowDownload" class="w-10 h-10 flex items-center justify-center hover:text-blue-500 transition-colors">
                      <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
                        <polyline points="7 10 12 15 17 10"/>
                        <line x1="12" y1="15" x2="12" y2="3"/>
                      </svg>
                    </button>
                  </div>
                  <div class="flex items-center gap-2 text-sm text-slate-400">
                    <span class="flex items-center gap-1">
                      <svg class="w-3.5 h-3.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                        <circle cx="12" cy="12" r="3"/>
                      </svg>
                      {{ song?.views || 0 }}
                    </span>
                    <span>·</span>
                    <span class="flex items-center gap-1">
                      <svg class="w-3.5 h-3.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M14 9V5a3 3 0 0 0-3-3l-4 9v11h11.28a2 2 0 0 0 2-1.7l1.38-9a2 2 0 0 0-2-2.3zM7 22H4a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3"/>
                      </svg>
                      {{ song?.likes || 0 }}
                    </span>
                    <span>·</span>
                    <span>{{ song?.duration || '0:00' }}</span>
                  </div>
                </div>
              </div>

              <div class="border-t border-slate-700 pt-4">
                <div class="flex items-center gap-2 mb-3">
                  <span class="text-sm font-medium text-slate-300">版权许可:</span>
                  <div class="flex items-center gap-2 px-3 py-1 bg-purple-500/10 rounded-full text-sm">
                    <svg class="w-4 h-4 text-purple-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <path d="M9 12l2 2 4-4"/>
                      <circle cx="12" cy="12" r="10"/>
                    </svg>
                    <span class="text-purple-400 font-medium">{{ licenseData.name }}</span>
                  </div>
                </div>
                <p class="text-xs text-slate-500 mb-4">{{ licenseData.description }}</p>
                
                <div class="flex flex-wrap gap-2">
                  <span
                    v-for="tag in licenseTags"
                    :key="tag.id"
                    :class="[
                      'px-2 py-1 rounded-full text-xs',
                      tag.type === 'allow' ? 'bg-green-500/10 text-green-400' :
                      tag.type === 'forbid' ? 'bg-red-500/10 text-red-400' :
                      'bg-slate-500/10 text-slate-400'
                    ]"
                  >
                    {{ tag.label }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <div class="glass-dark rounded-xl p-4 mt-4">
            <div class="flex items-center gap-3">
              <div class="w-12 h-12 rounded-full bg-gradient-to-br from-blue-500 to-cyan-500 flex items-center justify-center">
                <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M12 1a3 3 0 0 0-3 3v8a3 3 0 0 0 6 0V4a3 3 0 0 0-3-3z"/>
                  <path d="M19 10v2a7 7 0 0 1-14 0v-2"/>
                  <line x1="12" y1="19" x2="12" y2="23"/>
                </svg>
              </div>
              <div class="flex-1">
                <h4 class="font-medium text-white">{{ song?.artist }}</h4>
                <p class="text-xs text-slate-400">创作者</p>
              </div>
              <button class="px-4 py-1.5 bg-white/10 rounded-lg text-sm hover:bg-white/20 transition-colors">
                关注
              </button>
            </div>
          </div>
        </div>

        <div class="lg:col-span-2 space-y-6">
          <div class="glass-dark rounded-xl p-6">
            <h3 class="font-semibold text-white mb-3 flex items-center gap-2">
              <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                <polyline points="14 2 14 8 20 8"/>
                <line x1="16" y1="13" x2="8" y2="13"/>
                <line x1="16" y1="17" x2="8" y2="17"/>
              </svg>
              简介
            </h3>
            <p class="text-slate-300">
              {{ song?.description || '暂无简介' }}
            </p>
          </div>

          <div class="glass-dark rounded-xl p-6">
            <div class="flex items-center justify-between mb-4">
              <h3 class="font-semibold text-white flex items-center gap-2">
                <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/>
                </svg>
                评论
              </h3>
              <span class="text-slate-400 text-sm">{{ comments.length }} 条</span>
            </div>
            
            <CommentSection />
          </div>

          <div v-if="relatedSongs.length > 0" class="glass-dark rounded-xl p-6">
            <h3 class="font-semibold text-white mb-4 flex items-center gap-2">
              <svg class="w-5 h-5" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"/>
              </svg>
              相关推荐
            </h3>
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-3">
              <div
                v-for="s in relatedSongs"
                :key="s.id"
                @click="goToDetail(s.id)"
                class="flex items-center gap-3 p-3 rounded-lg hover:bg-white/10 cursor-pointer transition-all"
              >
                <img :src="s.cover" class="w-12 h-12 rounded-lg object-cover" />
                <div class="flex-1 min-w-0">
                  <h4 class="font-medium text-white truncate">{{ s.title }}</h4>
                  <p class="text-xs text-slate-400 truncate">{{ s.artist }}</p>
                </div>
                <span class="text-slate-500 text-sm">{{ s.duration }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMusicStore } from '@/stores/music'
import CommentSection from '@/components/music/CommentSection.vue'

const route = useRoute()
const router = useRouter()
const musicStore = useMusicStore()

const isPlaying = ref(false)
const isLiked = ref(false)

const song = ref({
  id: 1,
  title: '',
  artist: '',
  album: '',
  genre: '',
  duration: '',
  cover: 'https://picsum.photos/400/400?random=1',
  description: '',
  license: 'cc-by',
  views: 0,
  likes: 0,
  isPublic: true,
  allowDownload: true
})

const comments = ref([])

const relatedSongs = ref([])

const licenseData = computed(() => {
  const data = {
    'cc0': {
      name: 'CC0 1.0',
      description: '任何人都可以随意使用、修改和商用，无需署名',
      tags: [
        { id: 1, label: '允许商业使用', type: 'allow' },
        { id: 2, label: '允许修改', type: 'allow' },
        { id: 3, label: '无需署名', type: 'allow' }
      ]
    },
    'cc-by': {
      name: 'CC BY 4.0',
      description: '可以自由使用，但必须注明原作者',
      tags: [
        { id: 1, label: '允许商业使用', type: 'allow' },
        { id: 2, label: '允许修改', type: 'allow' },
        { id: 3, label: '需要署名', type: 'condition' }
      ]
    },
    'cc-by-nc': {
      name: 'CC BY-NC 4.0',
      description: '可以修改使用，但不能用于商业用途',
      tags: [
        { id: 1, label: '允许修改', type: 'allow' },
        { id: 2, label: '禁止商业使用', type: 'forbid' },
        { id: 3, label: '需要署名', type: 'condition' }
      ]
    },
    'cc-by-nd': {
      name: 'CC BY-ND 4.0',
      description: '可以转发，但不能对原作品进行修改',
      tags: [
        { id: 1, label: '允许转发', type: 'allow' },
        { id: 2, label: '禁止修改', type: 'forbid' },
        { id: 3, label: '需要署名', type: 'condition' }
      ]
    },
    'cc-by-sa': {
      name: 'CC BY-SA 4.0',
      description: '修改后的作品必须采用相同协议',
      tags: [
        { id: 1, label: '允许商业使用', type: 'allow' },
        { id: 2, label: '允许修改', type: 'allow' },
        { id: 3, label: '需要相同协议', type: 'condition' },
        { id: 4, label: '需要署名', type: 'condition' }
      ]
    }
  }
  
  return data[song.value.license] || data['cc-by']
})

const licenseTags = computed(() => licenseData.value.tags)

const togglePlay = () => {
  isPlaying.value = !isPlaying.value
}

const goToDetail = (id) => {
  router.push(`/song/${id}`)
}

onMounted(() => {
  const id = route.params.id
  if (id) {
    // 加载歌曲详情
  }
})
</script>
