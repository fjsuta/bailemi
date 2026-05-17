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
                  {{ isPlaying ? '⏸️' : '▶️' }}
                </div>
                <div class="flex-1">
                  <div class="flex items-center gap-3 mb-2">
                    <button
                      class="w-10 h-10 flex items-center justify-center hover:text-pink-500 transition-colors"
                      :class="{ 'text-pink-500': isLiked }"
                    >
                      {{ isLiked ? '❤️' : '🤍' }}
                    </button>
                    <button class="w-10 h-10 flex items-center justify-center hover:text-purple-500 transition-colors">
                      🔗
                    </button>
                    <button class="w-10 h-10 flex items-center justify-center hover:text-green-500 transition-colors">
                      🔊
                    </button>
                    <button v-if="song?.allowDownload" class="w-10 h-10 flex items-center justify-center hover:text-blue-500 transition-colors">
                      ⬇️
                    </button>
                  </div>
                  <div class="flex items-center gap-2 text-sm text-slate-400">
                    <span>👁 {{ song?.views || 0 }}</span>
                    <span>·</span>
                    <span>👍 {{ song?.likes || 0 }}</span>
                    <span>·</span>
                    <span>{{ song?.duration || '0:00' }}</span>
                  </div>
                </div>
              </div>

              <div class="border-t border-slate-700 pt-4">
                <div class="flex items-center gap-2 mb-3">
                  <span class="text-sm font-medium text-slate-300">版权许可:</span>
                  <div class="flex items-center gap-2 px-3 py-1 bg-purple-500/10 rounded-full text-sm">
                    <span>{{ licenseData.icon }}</span>
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
              <div class="w-12 h-12 rounded-full bg-gradient-to-br from-blue-500 to-cyan-500 flex items-center justify-center text-xl">
                🎤
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
            <h3 class="font-semibold text-white mb-3">📝 简介</h3>
            <p class="text-slate-300">
              {{ song?.description || '暂无简介' }}
            </p>
          </div>

          <div class="glass-dark rounded-xl p-6">
            <div class="flex items-center justify-between mb-4">
              <h3 class="font-semibold text-white">💬 评论</h3>
              <span class="text-slate-400 text-sm">{{ comments.length }} 条</span>
            </div>
            
            <CommentSection />
          </div>

          <div v-if="relatedSongs.length > 0" class="glass-dark rounded-xl p-6">
            <h3 class="font-semibold text-white mb-4">🎵 相关推荐</h3>
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
      icon: '🌍',
      name: 'CC0 1.0',
      description: '任何人都可以随意使用、修改和商用，无需署名',
      tags: [
        { id: 1, label: '允许商业使用', type: 'allow' },
        { id: 2, label: '允许修改', type: 'allow' },
        { id: 3, label: '无需署名', type: 'allow' }
      ]
    },
    'cc-by': {
      icon: '📝',
      name: 'CC BY 4.0',
      description: '可以自由使用，但必须注明原作者',
      tags: [
        { id: 1, label: '允许商业使用', type: 'allow' },
        { id: 2, label: '允许修改', type: 'allow' },
        { id: 3, label: '需要署名', type: 'condition' }
      ]
    },
    'cc-by-nc': {
      icon: '💰',
      name: 'CC BY-NC 4.0',
      description: '可以修改使用，但不能用于商业用途',
      tags: [
        { id: 1, label: '允许修改', type: 'allow' },
        { id: 2, label: '禁止商业使用', type: 'forbid' },
        { id: 3, label: '需要署名', type: 'condition' }
      ]
    },
    'cc-by-nd': {
      icon: '🚫',
      name: 'CC BY-ND 4.0',
      description: '可以转发，但不能对原作品进行修改',
      tags: [
        { id: 1, label: '允许转发', type: 'allow' },
        { id: 2, label: '禁止修改', type: 'forbid' },
        { id: 3, label: '需要署名', type: 'condition' }
      ]
    },
    'cc-by-sa': {
      icon: '🔄',
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
