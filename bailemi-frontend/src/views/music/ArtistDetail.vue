<template>
  <div class="min-h-screen pb-24">
    <div class="relative h-80 md:h-96">
      <img :src="artist.cover" class="w-full h-full object-cover" />
      <div class="absolute inset-0 bg-gradient-to-t from-black via-black/50 to-transparent">
        <div class="absolute bottom-0 left-0 right-0 p-6 md:p-10">
          <div class="flex items-end gap-6">
            <img :src="artist.avatar" class="w-24 h-24 md:w-32 md:h-32 rounded-full object-cover border-4 border-white/20" />
            <div>
              <h1 class="text-3xl md:text-5xl font-bold text-white mb-2">{{ artist.name }}</h1>
              <p class="text-slate-300">{{ artist.description }}</p>
              <div class="flex gap-6 mt-3">
                <span>{{ artist.fans }} 粉丝</span>
                <span>{{ artist.songs }} 歌曲</span>
                <span>{{ artist.albums }} 专辑</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="container mx-auto px-4 py-6">
      <div class="flex gap-4 mb-6 overflow-x-auto">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="[
            'px-6 py-2 rounded-xl font-medium whitespace-nowrap transition-all',
            activeTab === tab.id
              ? 'bg-gradient-to-r from-purple-600 to-blue-600 text-white'
              : 'bg-white/10 text-slate-300 hover:bg-white/20'
          ]"
        >
          {{ tab.name }}
        </button>
      </div>

      <div v-if="activeTab === 'songs'" class="space-y-3">
        <div
          v-for="(song, index) in artist.songsList"
          :key="song.id"
          @click="playSong(song)"
          class="flex items-center gap-4 p-4 glass rounded-xl hover:bg-white/20 cursor-pointer transition-all group"
        >
          <span class="text-slate-400 w-8">{{ index + 1 }}</span>
          <img :src="song.cover" class="w-12 h-12 rounded-lg object-cover" />
          <div class="flex-1 min-w-0">
            <h4 class="font-medium truncate group-hover:text-purple-400">{{ song.title }}</h4>
            <p class="text-sm text-slate-400">{{ song.album }}</p>
          </div>
          <span class="text-slate-500 text-sm">{{ song.duration }}</span>
          <span class="opacity-0 group-hover:opacity-100 text-white">
            <svg class="w-4 h-4" viewBox="0 0 24 24" fill="currentColor">
              <polygon points="5 3 19 12 5 21 5 3"/>
            </svg>
          </span>
        </div>
      </div>

      <div v-else-if="activeTab === 'albums'" class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        <div
          v-for="album in artist.albumsList"
          :key="album.id"
          @click="goToAlbum(album.id)"
          class="cursor-pointer group"
        >
          <div class="relative aspect-square rounded-xl overflow-hidden mb-3">
            <img :src="album.cover" class="w-full h-full object-cover group-hover:scale-105 transition-transform" />
            <div class="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 flex items-center justify-center">
              <svg class="w-8 h-8 text-white" viewBox="0 0 24 24" fill="currentColor">
                <polygon points="5 3 19 12 5 21 5 3"/>
              </svg>
            </div>
          </div>
          <h4 class="font-medium truncate">{{ album.name }}</h4>
          <p class="text-sm text-slate-400">{{ album.year }} · {{ album.songs }}首</p>
        </div>
      </div>

      <div v-else-if="activeTab === 'mv'" class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div
          v-for="mv in artist.mvs"
          :key="mv.id"
          class="flex gap-4 p-4 glass rounded-xl cursor-pointer hover:bg-white/20 transition-all"
        >
          <div class="relative w-32 h-18 flex-shrink-0">
            <img :src="mv.cover" class="w-full h-full object-cover rounded-lg" />
            <div class="absolute inset-0 flex items-center justify-center">
              <svg class="w-4 h-4 text-white" viewBox="0 0 24 24" fill="currentColor">
                <polygon points="5 3 19 12 5 21 5 3"/>
              </svg>
            </div>
          </div>
          <div class="flex-1">
            <h4 class="font-medium">{{ mv.title }}</h4>
            <p class="text-sm text-slate-400">{{ mv.views }}播放 · {{ mv.year }}</p>
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

const activeTab = ref('songs')

const tabs = [
  { id: 'songs', name: '热门歌曲' },
  { id: 'albums', name: '专辑' },
  { id: 'mv', name: 'MV' }
]

const artist = ref({
  id: 1,
  name: '',
  avatar: 'https://picsum.photos/150/150?random=artist1',
  cover: 'https://picsum.photos/1200/400?random=cover1',
  description: '',
  fans: '',
  songs: '',
  albums: '',
  songsList: [],
  albumsList: [],
  mvs: []
})

const playSong = (song) => {
  musicStore.playSong(song, artist.value.songsList)
}

const goToAlbum = (albumId) => {
  console.log('Go to album:', albumId)
}
</script>
