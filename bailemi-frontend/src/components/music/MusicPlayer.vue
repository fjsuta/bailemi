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
          <button @click="musicStore.playPrevious" class="w-10 h-10 flex items-center justify-center text-slate-400 hover:text-white text-2xl">
            ⏮️
          </button>
          <button
            @click="togglePlay"
            class="w-12 h-12 bg-gradient-to-r from-purple-600 to-blue-600 rounded-full flex items-center justify-center text-2xl hover:from-purple-500 hover:to-blue-500 transition-all"
          >
            {{ musicStore.isPlaying ? '⏸️' : '▶️' }}
          </button>
          <button @click="musicStore.playNext" class="w-10 h-10 flex items-center justify-center text-slate-400 hover:text-white text-2xl">
            ⏭️
          </button>
          <button @click="toggleRepeat" :class="['w-10 h-10 flex items-center justify-center text-xl transition-colors', musicStore.repeatMode !== 'off' ? 'text-purple-400' : 'text-slate-400 hover:text-white']">
            {{ repeatIcon }}
          </button>
          <button @click="toggleShuffle" :class="['w-10 h-10 flex items-center justify-center text-xl transition-colors', musicStore.isShuffled ? 'text-purple-400' : 'text-slate-400 hover:text-white']">
            🔀
          </button>
        </div>

        <div class="flex items-center gap-4 flex-1 justify-end text-sm text-slate-400">
          <span>{{ formatTime(currentTime) }}</span>
          <span>/</span>
          <span>{{ formatTime(duration) }}</span>
          <div class="flex items-center gap-2 ml-4">
            <span>🔊</span>
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
const repeatIcon = computed(() => {
 switch (musicStore.repeatMode) {
 case 'one': return '🔂';
 case 'all': return '🔁';
 default: return '🔄';
 }
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
