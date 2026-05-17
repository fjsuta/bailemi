<template>
  <div class="glass-dark rounded-2xl p-6">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 bg-gradient-to-br from-purple-500 to-blue-500 rounded-xl flex items-center justify-center">🎵</div>
        <div>
          <h2 class="text-xl font-bold">每日推荐</h2>
          <p class="text-sm text-slate-400">{{ today }}</p>
        </div>
      </div>
      <button class="text-purple-400 hover:text-purple-300 transition-colors">查看全部 →</button>
    </div>

    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
      <div
        v-for="song in recommendedSongs"
        :key="song.id"
        @click="playSong(song)"
        class="group cursor-pointer"
      >
        <div class="relative aspect-square rounded-xl overflow-hidden mb-3">
          <img :src="song.cover" class="w-full h-full object-cover group-hover:scale-105 transition-transform" />
          <div class="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center">
            <div class="w-12 h-12 bg-white/20 backdrop-blur rounded-full flex items-center justify-center text-white text-2xl">▶</div>
          </div>
          <span v-if="song.isHot" class="absolute top-2 right-2 px-2 py-1 bg-red-500/80 rounded text-xs text-white">HOT</span>
        </div>
        <h4 class="font-medium text-sm truncate group-hover:text-purple-400 transition-colors">{{ song.title }}</h4>
        <p class="text-xs text-slate-400 truncate">{{ song.artist }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useMusicStore } from '@/stores/music'

const musicStore = useMusicStore()

const today = computed(() => {
  const now = new Date()
  return `${now.getMonth() + 1}月${now.getDate()}日 为你推荐`
})

const recommendedSongs = ref([
  { id: 1, title: '晴天', artist: '周杰伦', cover: 'https://picsum.photos/200/200?random=r1', isHot: true },
  { id: 2, title: '稻香', artist: '周杰伦', cover: 'https://picsum.photos/200/200?random=r2', isHot: false },
  { id: 3, title: '夜曲', artist: '周杰伦', cover: 'https://picsum.photos/200/200?random=r3', isHot: true },
  { id: 4, title: '双截棍', artist: '周杰伦', cover: 'https://picsum.photos/200/200?random=r4', isHot: false },
  { id: 5, title: '七里香', artist: '周杰伦', cover: 'https://picsum.photos/200/200?random=r5', isHot: false },
  { id: 6, title: '简单爱', artist: '周杰伦', cover: 'https://picsum.photos/200/200?random=r6', isHot: true }
])

const playSong = (song) => {
  musicStore.playSong(song, recommendedSongs.value)
}
</script>
