<template>
  <div class="fixed bottom-4 left-1/2 -translate-x-1/2 z-50">
    <div class="bg-white/10 backdrop-blur-md border border-white/20 rounded-3xl px-6 py-4 shadow-2xl">
      <!-- 进度条 -->
      <div class="h-1.5 bg-white/20 rounded-full relative cursor-pointer mb-4" @click="seek">
        <div class="absolute inset-y-0 left-0 bg-gradient-to-r from-purple-500 to-blue-500 rounded-full transition-all" :style="{ width: progressPercent }"></div>
        <div class="absolute inset-y-0 left-0 bg-white/30 rounded-full" :style="{ width: bufferedPercent, opacity: 0.3 }"></div>
      </div>

      <div class="flex items-center gap-6">
        <!-- 歌曲信息 -->
        <div class="flex items-center gap-3 min-w-0">
          <div class="w-12 h-12 rounded-xl overflow-hidden bg-white/10 flex-shrink-0">
            <img v-if="musicStore.currentSong" :src="musicStore.currentSong.cover_url || 'https://picsum.photos/48/48?random=song'" class="w-full h-full object-cover" />
            <div v-else class="w-full h-full flex items-center justify-center">
              <svg class="w-6 h-6 text-white/50" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M9 18V5l12-2v13"></path>
                <circle cx="6" cy="18" r="3"></circle>
                <circle cx="18" cy="16" r="3"></circle>
              </svg>
            </div>
          </div>
          <div class="min-w-0">
            <h4 class="font-medium text-white truncate">{{ musicStore.currentSong?.title || '暂未播放' }}</h4>
            <p class="text-sm text-white/60 truncate">{{ getArtistName() }}</p>
          </div>
        </div>

        <!-- 播放控制 -->
        <div class="flex items-center gap-2">
          <button 
            @click="musicStore.playPrevious" 
            class="w-10 h-10 flex items-center justify-center text-white/70 hover:text-white hover:bg-white/10 rounded-2xl transition-all"
          >
            <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polygon points="19 20 9 12 19 4 19 20"></polygon>
              <line x1="5" y1="19" x2="5" y2="5"></line>
            </svg>
          </button>
          <button
            @click="togglePlay"
            class="w-14 h-14 bg-gradient-to-r from-purple-600 to-blue-600 rounded-full flex items-center justify-center hover:from-purple-500 hover:to-blue-500 transition-all shadow-lg hover:shadow-xl"
          >
            <svg v-if="musicStore.isPlaying" class="w-7 h-7 text-white" viewBox="0 0 24 24" fill="currentColor">
              <rect x="6" y="4" width="4" height="16" rx="1"></rect>
              <rect x="14" y="4" width="4" height="16" rx="1"></rect>
            </svg>
            <svg v-else class="w-7 h-7 text-white ml-1" viewBox="0 0 24 24" fill="currentColor">
              <polygon points="5 3 19 12 5 21 5 3"></polygon>
            </svg>
          </button>
          <button 
            @click="musicStore.playNext" 
            class="w-10 h-10 flex items-center justify-center text-white/70 hover:text-white hover:bg-white/10 rounded-2xl transition-all"
          >
            <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polygon points="5 4 15 12 5 20 5 4"></polygon>
              <line x1="19" y1="5" x2="19" y2="19"></line>
            </svg>
          </button>
        </div>

        <!-- 模式控制 -->
        <div class="flex items-center gap-1">
          <button 
            @click="toggleShuffle" 
            :class="['w-9 h-9 flex items-center justify-center rounded-2xl transition-all', musicStore.isShuffled ? 'text-purple-400 bg-white/10' : 'text-white/50 hover:text-white/80 hover:bg-white/5']"
          >
            <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="16 3 21 3 21 8"></polyline>
              <line x1="4" y1="20" x2="21" y2="3"></line>
              <polyline points="21 16 21 21 16 21"></polyline>
              <line x1="15" y1="15" x2="21" y2="21"></line>
              <line x1="4" y1="4" x2="9" y2="9"></line>
            </svg>
          </button>
          <button 
            @click="toggleRepeat" 
            :class="['w-9 h-9 flex items-center justify-center rounded-2xl transition-all', musicStore.repeatMode !== 'off' ? 'text-purple-400 bg-white/10' : 'text-white/50 hover:text-white/80 hover:bg-white/5']"
          >
            <svg v-if="musicStore.repeatMode === 'one'" class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M17 1l4 4-4 4"></path>
              <path d="M3 11V9a4 4 0 0 1 4-4h14"></path>
              <path d="M7 23l-4-4 4-4"></path>
              <path d="M21 13v2a4 4 0 0 1-4 4H3"></path>
              <circle cx="12" cy="12" r="1"></circle>
            </svg>
            <svg v-else class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M17 1l4 4-4 4"></path>
              <path d="M3 11V9a4 4 0 0 1 4-4h14"></path>
              <path d="M7 23l-4-4 4-4"></path>
              <path d="M21 13v2a4 4 0 0 1-4 4H3"></path>
            </svg>
          </button>
        </div>

        <!-- 时间和音量 -->
        <div class="flex items-center gap-3">
          <span class="text-sm text-white/60 min-w-[45px] text-right">{{ formatTime(currentTime) }}</span>
          <span class="text-sm text-white/40">/</span>
          <span class="text-sm text-white/60 min-w-[45px]">{{ formatTime(duration) }}</span>
          
          <div class="flex items-center gap-2 ml-4 pl-4 border-l border-white/10">
            <button 
              @click="toggleMute"
              class="w-9 h-9 flex items-center justify-center text-white/70 hover:text-white hover:bg-white/10 rounded-2xl transition-all"
            >
              <svg v-if="volume === 0" class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon>
                <line x1="23" y1="9" x2="17" y2="15"></line>
                <line x1="17" y1="9" x2="23" y2="15"></line>
              </svg>
              <svg v-else-if="volume < 50" class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon>
                <path d="M19.07 4.93a10 10 0 0 1 0 14.14M15.54 8.46a5 5 0 0 1 0 7.07"></path>
              </svg>
              <svg v-else class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"></polygon>
                <path d="M19.07 4.93a10 10 0 0 1 0 14.14M15.54 8.46a5 5 0 0 1 0 7.07"></path>
              </svg>
            </button>
            <input 
              type="range" 
              min="0" 
              max="100" 
              :value="volume" 
              @input="changeVolume" 
              class="w-24 h-1.5 bg-white/20 rounded-full appearance-none cursor-pointer accent-purple-500"
            />
          </div>
        </div>
      </div>
    </div>

    <audio ref="audioPlayer" @timeupdate="updateTime" @loadedmetadata="updateDuration" @ended="handleEnded" @progress="updateBuffered" preload="auto" />
    <audio ref="nextAudio" preload="auto" />
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import { useMusicStore } from '@/stores/music'

const musicStore = useMusicStore()
const audioPlayer = ref(null)
const nextAudio = ref(null)
const currentTime = ref(0)
const duration = ref(0)
const buffered = ref(0)
const volume = ref(70)
const isMuted = ref(false)
const previousVolume = ref(70)
const PRELOAD_THRESHOLD = 5

const progressPercent = computed(() => {
  if (!duration.value) return '0%'
  return `${(currentTime.value / duration.value) * 100}%`
})

const bufferedPercent = computed(() => {
  if (!duration.value) return '0%'
  return `${(buffered.value / duration.value) * 100}%`
})

const getArtistName = () => {
  const artist = musicStore.currentSong?.artist
  if (typeof artist === 'string') return artist
  return artist?.name || '未知歌手'
}

const togglePlay = () => {
  if (!musicStore.currentSong) return
  if (musicStore.isPlaying) {
    audioPlayer.value?.pause()
  } else {
    audioPlayer.value?.play()
  }
  musicStore.togglePlay()
}

const updateTime = () => {
  if (!audioPlayer.value) return
  currentTime.value = audioPlayer.value.currentTime
  preloadNextSong()
}

const updateDuration = () => {
  if (audioPlayer.value) {
    duration.value = audioPlayer.value.duration
  }
}

const updateBuffered = () => {
  if (!audioPlayer.value) return
  const buf = audioPlayer.value.buffered
  if (buf.length > 0) {
    buffered.value = buf.end(buf.length - 1)
  }
}

const seek = (e) => {
  if (!audioPlayer.value) return
  const rect = e.target.getBoundingClientRect()
  const percent = (e.clientX - rect.left) / rect.width
  audioPlayer.value.currentTime = percent * duration.value
}

const handleEnded = () => {
  if (musicStore.repeatMode === 'one') {
    audioPlayer.value.currentTime = 0
    audioPlayer.value.play()
    return
  }
  musicStore.playNext()
}

const formatTime = (seconds) => {
  if (!seconds) return '0:00'
  const mins = Math.floor(seconds / 60)
  const secs = Math.floor(seconds % 60).toString().padStart(2, '0')
  return `${mins}:${secs}`
}

const preloadNextSong = () => {
  if (!musicStore.isPlaying || !duration.value) return
  const remaining = duration.value - currentTime.value
  if (remaining <= PRELOAD_THRESHOLD && !nextAudio.value?.src) {
    const nextSong = musicStore.getNextSong()
    if (nextSong) {
      const playUrl = nextSong.play_url || 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3'
      nextAudio.value.src = playUrl
    }
  }
}

const toggleRepeat = () => {
  const modes = ['off', 'all', 'one']
  const currentIndex = modes.indexOf(musicStore.repeatMode)
  const nextIndex = (currentIndex + 1) % modes.length
  musicStore.setRepeatMode(modes[nextIndex])
}

const toggleShuffle = () => {
  musicStore.toggleShuffle()
}

const changeVolume = (e) => {
  volume.value = parseInt(e.target.value)
  if (audioPlayer.value) {
    audioPlayer.value.volume = volume.value / 100
  }
  if (volume.value > 0) {
    isMuted.value = false
    previousVolume.value = volume.value
  }
}

const toggleMute = () => {
  if (isMuted.value) {
    volume.value = previousVolume.value
    isMuted.value = false
  } else {
    previousVolume.value = volume.value
    volume.value = 0
    isMuted.value = true
  }
  if (audioPlayer.value) {
    audioPlayer.value.volume = volume.value / 100
  }
}

watch(() => musicStore.currentSong, (newSong, oldSong) => {
  if (newSong && audioPlayer.value) {
    const playUrl = newSong.play_url || 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3'
    audioPlayer.value.src = playUrl
    audioPlayer.value.volume = volume.value / 100
    if (musicStore.isPlaying) {
      audioPlayer.value.play()
    }
  }
  nextAudio.value.src = ''
})

watch(() => musicStore.isPlaying, (isPlaying) => {
  if (!audioPlayer.value || !musicStore.currentSong) return
  if (isPlaying) {
    audioPlayer.value.play()
  } else {
    audioPlayer.value.pause()
  }
})

onMounted(() => {
  if (audioPlayer.value) {
    audioPlayer.value.volume = volume.value / 100
  }
})
</script>
