<template>
  <div class="fixed bottom-0 left-0 right-0 z-50 glass-dark border-t border-slate-700/50">
    <div class="container mx-auto px-4">
      <div class="progress-bar h-1 bg-slate-700 relative cursor-pointer" @click="seek">
        <div class="progress-fill h-full bg-gradient-to-r from-purple-500 to-blue-500" :style="{ width: progressPercent }"></div>
        <div class="progress-fill h-full bg-white/50 transition-opacity" :style="{ width: bufferedPercent, opacity: 0.5 }"></div>
      </div>
      <div class="flex items-center justify-between py-4">
        <div class="flex items-center gap-4 flex-1 min-w-0">
          <img v-if="musicStore.currentSong" :src="musicStore.currentSong.cover_url || 'https://picsum.photos/56/56?random=song'" class="w-14 h-14 rounded-lg object-cover" />
          <div v-else class="w-14 h-14 rounded-lg bg-slate-700"></div>
          <div class="min-w-0">
            <h4 class="font-medium truncate">{{ musicStore.currentSong?.title || '暂未播放' }}</h4>
            <p class="text-slate-400 text-sm truncate">{{ getArtistName() }}</p>
          </div>
        </div>

        <div class="flex items-center gap-4">
          <button @click="musicStore.playPrevious" class="w-10 h-10 flex items-center justify-center text-slate-400 hover:text-white">
            <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polygon points="19 20 9 12 19 4 19 20"/>
              <polygon points="5 20 15 12 5 4 5 20"/>
            </svg>
          </button>
          <button
            @click="togglePlay"
            class="w-12 h-12 bg-gradient-to-r from-purple-600 to-blue-600 rounded-full flex items-center justify-center hover:from-purple-500 hover:to-blue-500 transition-all"
          >
            <svg v-if="musicStore.isPlaying" class="w-6 h-6 text-white" viewBox="0 0 24 24" fill="currentColor">
              <rect x="6" y="4" width="4" height="16"/>
              <rect x="14" y="4" width="4" height="16"/>
            </svg>
            <svg v-else class="w-6 h-6 text-white ml-1" viewBox="0 0 24 24" fill="currentColor">
              <polygon points="5 3 19 12 5 21 5 3"/>
            </svg>
          </button>
          <button @click="musicStore.playNext" class="w-10 h-10 flex items-center justify-center text-slate-400 hover:text-white">
            <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polygon points="5 4 15 12 5 20 5 4"/>
              <polygon points="19 4 9 12 19 20 19 4"/>
            </svg>
          </button>
          <button @click="toggleRepeat" :class="['w-10 h-10 flex items-center justify-center transition-colors', musicStore.repeatMode !== 'off' ? 'text-purple-400' : 'text-slate-400 hover:text-white']">
            <svg v-if="musicStore.repeatMode === 'one'" class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"/>
              <path d="M3 16a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"/>
              <circle cx="12" cy="12" r="3"/>
            </svg>
            <svg v-else-if="musicStore.repeatMode === 'all'" class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"/>
              <path d="M3 16a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"/>
            </svg>
            <svg v-else class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 12a9 9 0 0 0-9-9 9.75 9.75 0 0 0-6.74 2.74L3 8"/>
              <path d="M3 16a9 9 0 0 0 9 9 9.75 9.75 0 0 0 6.74-2.74L21 16"/>
            </svg>
          </button>
          <button @click="toggleShuffle" :class="['w-10 h-10 flex items-center justify-center transition-colors', musicStore.isShuffled ? 'text-purple-400' : 'text-slate-400 hover:text-white']">
            <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M16 3h5v5"/>
              <path d="M4 20 21 3"/>
              <path d="M21 16v5h-5"/>
              <path d="M15 15l6 6"/>
              <path d="M4 4l5 5"/>
            </svg>
          </button>
        </div>

        <div class="flex items-center gap-4 flex-1 justify-end text-sm text-slate-400">
          <span>{{ formatTime(currentTime) }}</span>
          <span>/</span>
          <span>{{ formatTime(duration) }}</span>
          <div class="flex items-center gap-2 ml-4">
            <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"/>
              <path d="M15.54 8.46a5 5 0 0 1 0 7.07"/>
              <path d="M19.07 4.93a10 10 0 0 1 0 14.14"/>
            </svg>
            <input type="range" min="0" max="100" :value="volume" @input="changeVolume" class="w-20 accent-purple-500" />
          </div>
        </div>
      </div>
    </div>
    <audio ref="audioPlayer" @timeupdate="updateTime" @loadedmetadata="updateDuration" @ended="handleEnded" @progress="updateBuffered" preload="auto" />
    <audio ref="nextAudio" preload="auto" />
  </div>
</template>

<script setup>import { ref, computed, watch, onMounted, onUnmounted } from 'vue';
import { useMusicStore } from '@/stores/music';
const musicStore = useMusicStore();
const audioPlayer = ref(null);
const nextAudio = ref(null);
const currentTime = ref(0);
const duration = ref(0);
const buffered = ref(0);
const volume = ref(70);
const PRELOAD_THRESHOLD = 5;
const progressPercent = computed(() => {
 if (!duration.value)
 return '0%';
 return `${(currentTime.value / duration.value) * 100}%`;
});
const bufferedPercent = computed(() => {
 if (!duration.value)
 return '0%';
 return `${(buffered.value / duration.value) * 100}%`;
});

const getArtistName = () => {
 const artist = musicStore.currentSong?.artist;
 if (typeof artist === 'string')
 return artist;
 return artist?.name || '未知歌手';
};
const togglePlay = () => {
 if (!musicStore.currentSong)
 return;
 if (musicStore.isPlaying) {
 audioPlayer.value?.pause();
 }
 else {
 audioPlayer.value?.play();
 }
 musicStore.togglePlay();
};
const updateTime = () => {
 if (!audioPlayer.value)
 return;
 currentTime.value = audioPlayer.value.currentTime;
 preloadNextSong();
};
const updateDuration = () => {
 if (audioPlayer.value) {
 duration.value = audioPlayer.value.duration;
 }
};
const updateBuffered = () => {
 if (!audioPlayer.value)
 return;
 const buf = audioPlayer.value.buffered;
 if (buf.length > 0) {
 buffered.value = buf.end(buf.length - 1);
 }
};
const seek = (e) => {
 if (!audioPlayer.value)
 return;
 const rect = e.target.getBoundingClientRect();
 const percent = (e.clientX - rect.left) / rect.width;
 audioPlayer.value.currentTime = percent * duration.value;
};
const handleEnded = () => {
 if (musicStore.repeatMode === 'one') {
 audioPlayer.value.currentTime = 0;
 audioPlayer.value.play();
 return;
 }
 musicStore.playNext();
};
const formatTime = (seconds) => {
 if (!seconds)
 return '0:00';
 const mins = Math.floor(seconds / 60);
 const secs = Math.floor(seconds % 60).toString().padStart(2, '0');
 return `${mins}:${secs}`;
};
const preloadNextSong = () => {
 if (!musicStore.isPlaying || !duration.value)
 return;
 const remaining = duration.value - currentTime.value;
 if (remaining <= PRELOAD_THRESHOLD && !nextAudio.value?.src) {
 const nextSong = musicStore.getNextSong();
 if (nextSong) {
 const playUrl = nextSong.play_url || 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3';
 nextAudio.value.src = playUrl;
 }
 }
};
const toggleRepeat = () => {
 const modes = ['off', 'all', 'one'];
 const currentIndex = modes.indexOf(musicStore.repeatMode);
 const nextIndex = (currentIndex + 1) % modes.length;
 musicStore.setRepeatMode(modes[nextIndex]);
};
const toggleShuffle = () => {
 musicStore.toggleShuffle();
};
const changeVolume = (e) => {
 volume.value = parseInt(e.target.value);
 if (audioPlayer.value) {
 audioPlayer.value.volume = volume.value / 100;
 }
};
watch(() => musicStore.currentSong, (newSong, oldSong) => {
 if (newSong && audioPlayer.value) {
 const playUrl = newSong.play_url || 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3';
 audioPlayer.value.src = playUrl;
 audioPlayer.value.volume = volume.value / 100;
 if (musicStore.isPlaying) {
 audioPlayer.value.play();
 }
 }
 nextAudio.value.src = '';
});
watch(() => musicStore.isPlaying, (isPlaying) => {
 if (!audioPlayer.value || !musicStore.currentSong)
 return;
 if (isPlaying) {
 audioPlayer.value.play();
 }
 else {
 audioPlayer.value.pause();
 }
});
onMounted(() => {
 if (audioPlayer.value) {
 audioPlayer.value.volume = volume.value / 100;
 }
});
</script>
