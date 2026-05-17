import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMusicStore = defineStore('music', () => {
  const currentSong = ref(null)
  const isPlaying = ref(false)
  const playlist = ref([])
  const currentIndex = ref(0)
  const currentTime = ref(0)
  const duration = ref(0)
  const repeatMode = ref('off')
  const isShuffled = ref(false)
  const volume = ref(70)
  const originalPlaylist = ref([])

  function playSong(song, songs = null) {
    if (songs && songs.length > 0) {
      playlist.value = songs
      originalPlaylist.value = [...songs]
      currentIndex.value = songs.findIndex(s => s.id === song.id)
      if (currentIndex.value === -1) {
        currentIndex.value = 0
      }
    }
    currentSong.value = song
    isPlaying.value = true
  }

  function togglePlay() {
    isPlaying.value = !isPlaying.value
  }

  function playNext() {
    if (playlist.value.length === 0) return
    
    if (repeatMode.value === 'one') {
      return
    }
    
    if (isShuffled.value && playlist.value.length > 1) {
      let newIndex = currentIndex.value
      while (newIndex === currentIndex.value) {
        newIndex = Math.floor(Math.random() * playlist.value.length)
      }
      currentIndex.value = newIndex
    } else {
      currentIndex.value = (currentIndex.value + 1) % playlist.value.length
    }
    
    if (playlist.value[currentIndex.value]) {
      currentSong.value = playlist.value[currentIndex.value]
      isPlaying.value = true
    }
  }

  function playPrevious() {
    if (playlist.value.length === 0) return
    
    if (repeatMode.value === 'one') {
      return
    }
    
    if (isShuffled.value && playlist.value.length > 1) {
      let newIndex = currentIndex.value
      while (newIndex === currentIndex.value) {
        newIndex = Math.floor(Math.random() * playlist.value.length)
      }
      currentIndex.value = newIndex
    } else {
      currentIndex.value = (currentIndex.value - 1 + playlist.value.length) % playlist.value.length
    }
    
    if (playlist.value[currentIndex.value]) {
      currentSong.value = playlist.value[currentIndex.value]
      isPlaying.value = true
    }
  }

  function getNextSong() {
    if (playlist.value.length === 0) return null
    
    let nextIndex
    if (isShuffled.value && playlist.value.length > 1) {
      nextIndex = currentIndex.value
      while (nextIndex === currentIndex.value) {
        nextIndex = Math.floor(Math.random() * playlist.value.length)
      }
    } else {
      nextIndex = (currentIndex.value + 1) % playlist.value.length
    }
    
    return playlist.value[nextIndex] || null
  }

  function setRepeatMode(mode) {
    repeatMode.value = mode
  }

  function toggleShuffle() {
    isShuffled.value = !isShuffled.value
    if (isShuffled.value) {
      playlist.value = [...playlist.value].sort(() => Math.random() - 0.5)
      currentIndex.value = playlist.value.findIndex(s => s.id === currentSong.value?.id)
    } else {
      playlist.value = [...originalPlaylist.value]
      currentIndex.value = playlist.value.findIndex(s => s.id === currentSong.value?.id)
    }
  }

  function volumeUp() {
    volume.value = Math.min(volume.value + 10, 100)
  }

  function volumeDown() {
    volume.value = Math.max(volume.value - 10, 0)
  }

  function updateTime(time) {
    currentTime.value = time
  }

  function updateDuration(dur) {
    duration.value = dur
  }

  return {
    currentSong,
    isPlaying,
    playlist,
    currentIndex,
    currentTime,
    duration,
    repeatMode,
    isShuffled,
    volume,
    playSong,
    togglePlay,
    playNext,
    playPrevious,
    getNextSong,
    setRepeatMode,
    toggleShuffle,
    volumeUp,
    volumeDown,
    updateTime,
    updateDuration
  }
})
