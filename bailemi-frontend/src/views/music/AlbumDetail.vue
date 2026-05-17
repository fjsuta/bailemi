<template>
  <div class="min-h-screen pb-24">
    <div class="container mx-auto px-4 py-8">
      <div class="flex flex-col md:flex-row gap-6 mb-8">
        <div class="relative w-48 h-48 md:w-64 md:h-64 flex-shrink-0">
          <img :src="album.cover" class="w-full h-full object-cover rounded-2xl" />
          <div class="absolute inset-0 bg-black/30 rounded-2xl flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity">
            <button @click="playAlbum" class="w-16 h-16 bg-white/20 backdrop-blur rounded-full flex items-center justify-center text-white text-3xl hover:bg-white/30 transition-colors">
              ▶
            </button>
          </div>
        </div>
        <div class="flex-1">
          <h1 class="text-3xl md:text-4xl font-bold text-white mb-2">{{ album.name }}</h1>
          <p class="text-purple-400 mb-4">{{ album.artist }}</p>
          <div class="flex gap-6 mb-6">
            <span>{{ album.year }}年发行</span>
            <span>{{ album.songs.length }}首歌曲</span>
            <span>{{ album.language }}</span>
          </div>
          <p class="text-slate-300 mb-6">{{ album.description }}</p>
          <div class="flex gap-4">
            <button class="px-6 py-2 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl font-medium hover:opacity-90">
              ♥ 收藏专辑
            </button>
            <button class="px-6 py-2 bg-white/10 rounded-xl font-medium hover:bg-white/20">
              ➤ 分享
            </button>
          </div>
        </div>
      </div>

      <div class="glass-dark rounded-2xl p-6">
        <div class="flex items-center justify-between mb-6">
          <h2 class="text-xl font-bold">曲目列表</h2>
          <button @click="playAlbum" class="text-purple-400 hover:text-purple-300">播放全部</button>
        </div>

        <div class="space-y-1">
          <div
            v-for="(song, index) in album.songs"
            :key="song.id"
            @click="playSong(song)"
            :class="[
              'flex items-center gap-4 p-4 rounded-xl cursor-pointer transition-all',
              currentSong?.id === song.id ? 'bg-purple-500/20' : 'hover:bg-white/10'
            ]"
          >
            <div class="w-8 text-center">
              <span v-if="currentSong?.id !== song.id" class="text-slate-400">{{ index + 1 }}</span>
              <span v-else class="text-purple-400 animate-pulse">♪</span>
            </div>
            <img :src="song.cover" class="w-10 h-10 rounded-lg object-cover" />
            <div class="flex-1 min-w-0">
              <h4 class="font-medium truncate">{{ song.title }}</h4>
              <p class="text-sm text-slate-400">{{ song.artist }}</p>
            </div>
            <div class="flex items-center gap-4">
              <span class="text-sm text-slate-500">{{ song.duration }}</span>
              <span v-if="song.isFavorite" class="text-red-500">❤️</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useMusicStore } from '@/stores/music'

const musicStore = useMusicStore()

const currentSong = ref(null)

const album = ref({
  id: 1,
  name: '最伟大的作品',
  artist: '周杰伦',
  cover: 'https://picsum.photos/400/400?random=album1',
  year: '2022',
  language: '国语',
  description: '周杰伦第15张个人专辑，融合古典与现代音乐风格，展现音乐艺术的无限可能。',
  songs: [
    { id: 1, title: '最伟大的作品', artist: '周杰伦', duration: '4:32', cover: 'https://picsum.photos/100/100?random=as1', isFavorite: true },
    { id: 2, title: '还在流浪', artist: '周杰伦', duration: '4:51', cover: 'https://picsum.photos/100/100?random=as2', isFavorite: false },
    { id: 3, title: '说好不哭', artist: '周杰伦', duration: '4:03', cover: 'https://picsum.photos/100/100?random=as3', isFavorite: true },
    { id: 4, title: '红颜如霜', artist: '周杰伦', duration: '4:30', cover: 'https://picsum.photos/100/100?random=as4', isFavorite: false },
    { id: 5, title: '不爱我就拉倒', artist: '周杰伦', duration: '4:12', cover: 'https://picsum.photos/100/100?random=as5', isFavorite: false },
    { id: 6, title: 'Mojito', artist: '周杰伦', duration: '3:54', cover: 'https://picsum.photos/100/100?random=as6', isFavorite: true },
    { id: 7, title: '粉色海洋', artist: '周杰伦', duration: '3:49', cover: 'https://picsum.photos/100/100?random=as7', isFavorite: false },
    { id: 8, title: '倒影', artist: '周杰伦', duration: '4:06', cover: 'https://picsum.photos/100/100?random=as8', isFavorite: false },
    { id: 9, title: '错过的烟火', artist: '周杰伦', duration: '4:15', cover: 'https://picsum.photos/100/100?random=as9', isFavorite: false },
    { id: 10, title: '等你下课', artist: '周杰伦', duration: '4:32', cover: 'https://picsum.photos/100/100?random=as10', isFavorite: true },
    { id: 11, title: '告白气球', artist: '周杰伦', duration: '3:35', cover: 'https://picsum.photos/100/100?random=as11', isFavorite: true },
    { id: 12, title: '一路向北', artist: '周杰伦', duration: '4:55', cover: 'https://picsum.photos/100/100?random=as12', isFavorite: false }
  ]
})

const playSong = (song) => {
  currentSong.value = song
  musicStore.playSong(song, album.value.songs)
}

const playAlbum = () => {
  if (album.value.songs.length > 0) {
    playSong(album.value.songs[0])
  }
}
</script>
