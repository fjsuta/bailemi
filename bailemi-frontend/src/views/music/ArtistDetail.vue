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
          <span class="opacity-0 group-hover:opacity-100 text-white">▶</span>
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
              <span class="text-4xl">▶</span>
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
              <span class="text-xl">▶</span>
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
  name: '周杰伦',
  avatar: 'https://picsum.photos/150/150?random=artist1',
  cover: 'https://picsum.photos/1200/400?random=cover1',
  description: '华语乐坛天王，创作型歌手',
  fans: '3200万',
  songs: '156',
  albums: '15',
  songsList: [
    { id: 1, title: '晴天', album: '叶惠美', duration: '4:29', cover: 'https://picsum.photos/100/100?random=s1' },
    { id: 2, title: '夜曲', album: '十一月的萧邦', duration: '4:23', cover: 'https://picsum.photos/100/100?random=s2' },
    { id: 3, title: '七里香', album: '七里香', duration: '4:59', cover: 'https://picsum.photos/100/100?random=s3' },
    { id: 4, title: '稻香', album: '魔杰座', duration: '3:43', cover: 'https://picsum.photos/100/100?random=s4' },
    { id: 5, title: '双截棍', album: '范特西', duration: '3:21', cover: 'https://picsum.photos/100/100?random=s5' }
  ],
  albumsList: [
    { id: 1, name: '最伟大的作品', cover: 'https://picsum.photos/200/200?random=a1', year: '2022', songs: 12 },
    { id: 2, name: '周杰伦的床边故事', cover: 'https://picsum.photos/200/200?random=a2', year: '2016', songs: 10 },
    { id: 3, name: '哎哟，不错哦', cover: 'https://picsum.photos/200/200?random=a3', year: '2014', songs: 12 },
    { id: 4, name: '十二新作', cover: 'https://picsum.photos/200/200?random=a4', year: '2012', songs: 12 }
  ],
  mvs: [
    { id: 1, title: '最伟大的作品', cover: 'https://picsum.photos/200/100?random=m1', views: '5200万', year: '2022' },
    { id: 2, title: '说好不哭', cover: 'https://picsum.photos/200/100?random=m2', views: '8900万', year: '2019' },
    { id: 3, title: '告白气球', cover: 'https://picsum.photos/200/100?random=m3', views: '1.2亿', year: '2016' }
  ]
})

const playSong = (song) => {
  musicStore.playSong(song, artist.value.songsList)
}

const goToAlbum = (albumId) => {
  console.log('Go to album:', albumId)
}
</script>
