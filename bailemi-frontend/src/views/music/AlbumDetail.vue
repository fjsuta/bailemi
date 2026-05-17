<template>
  <div class="min-h-screen pb-24">
    <div class="container mx-auto px-4 py-8">
      <div class="flex flex-col md:flex-row gap-6 mb-8">
        <div class="relative w-48 h-48 md:w-64 md:h-64 flex-shrink-0">
          <img :src="album.cover" class="w-full h-full object-cover rounded-2xl" />
          <div class="absolute inset-0 bg-black/30 rounded-2xl flex items-center justify-center opacity-0 hover:opacity-100 transition-opacity">
            <button @click="playAlbum" class="w-16 h-16 bg-white/20 backdrop-blur rounded-full flex items-center justify-center hover:bg-white/30 transition-colors">
              <svg class="w-7 h-7 text-white ml-1" viewBox="0 0 24 24" fill="currentColor">
                <polygon points="5 3 19 12 5 21 5 3"/>
              </svg>
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
              <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
              </svg>
              收藏专辑
            </button>
            <button class="px-6 py-2 bg-white/10 rounded-xl font-medium hover:bg-white/20">
              <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/>
                <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/>
              </svg>
              分享
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
              <svg v-else class="w-4 h-4 text-purple-400 animate-pulse" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"/>
              </svg>
            </div>
            <img :src="song.cover" class="w-10 h-10 rounded-lg object-cover" />
            <div class="flex-1 min-w-0">
              <h4 class="font-medium truncate">{{ song.title }}</h4>
              <p class="text-sm text-slate-400">{{ song.artist }}</p>
            </div>
            <div class="flex items-center gap-4">
              <span class="text-sm text-slate-500">{{ song.duration }}</span>
              <svg v-if="song.isFavorite" class="w-4 h-4 text-red-500" viewBox="0 0 24 24" fill="currentColor">
                <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"/>
              </svg>
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
  name: '',
  artist: '',
  cover: 'https://picsum.photos/400/400?random=album1',
  year: '',
  language: '',
  description: '',
  songs: []
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
