import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useMusicStore = defineStore('music', () => {
  const currentSong = ref(null)
  const isPlaying = ref(false)
  const playlist = ref([])
  const currentIndex = ref(0)
  const currentTime = ref(0)
  const duration = ref(0)

  function playSong(song, songs = null) {
    if (songs && songs.length > 0) {
      playlist.value = songs
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
    currentIndex.value = (currentIndex.value + 1) % playlist.value.length
    if (playlist.value[currentIndex.value]) {
      currentSong.value = playlist.value[currentIndex.value]
      isPlaying.value = true
    }
  }

  function playPrevious() {
    if (playlist.value.length === 0) return
    currentIndex.value = (currentIndex.value - 1 + playlist.value.length) % playlist.value.length
    if (playlist.value[currentIndex.value]) {
      currentSong.value = playlist.value[currentIndex.value]
      isPlaying.value = true
    }
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
    playSong,
    togglePlay,
    playNext,
    playPrevious,
    updateTime,
    updateDuration
  }
})
