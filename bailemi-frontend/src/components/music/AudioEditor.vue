<template>
  <div class="space-y-6">
    <div class="flex flex-col lg:flex-row gap-6">
      <div class="flex-1">
        <label class="block text-sm font-medium text-slate-300 mb-3">
          🎵 音频波形编辑器
        </label>
        
        <div class="glass-dark rounded-xl p-4">
          <div
            ref="waveformContainer"
            class="bg-slate-900/50 rounded-lg h-48 relative"
          ></div>
          
          <div class="flex items-center gap-3 mt-4">
            <button
              @click="togglePlay"
              class="px-4 py-2 bg-gradient-to-r from-purple-600 to-blue-600 rounded-lg font-medium flex items-center gap-2 hover:opacity-90 transition-opacity"
            >
              <span>{{ isPlaying ? '⏸️' : '▶️' }}</span>
              <span>{{ isPlaying ? '暂停' : '播放' }}</span>
            </button>
            
            <span class="text-sm text-slate-400">
              {{ formatTime(currentTime) }} / {{ formatTime(duration) }}
            </span>
          </div>
        </div>

        <div class="glass-dark rounded-xl p-4 mt-4">
          <h4 class="font-semibold text-white mb-3">✂️ 区域选择 (可拖拽)</h4>
          <div class="grid grid-cols-2 gap-4 mb-4">
            <div>
              <label class="block text-xs text-slate-400 mb-1">起始时间 (秒)</label>
              <input
                v-model.number="selection.start"
                type="number"
                step="0.01"
                min="0"
                :max="duration"
                class="w-full px-3 py-2 bg-white/10 border border-white/20 rounded-lg focus:border-purple-500 focus:outline-none text-white"
                @input="updateRegion"
              />
            </div>
            <div>
              <label class="block text-xs text-slate-400 mb-1">结束时间 (秒)</label>
              <input
                v-model.number="selection.end"
                type="number"
                step="0.01"
                :min="selection.start"
                :max="duration"
                class="w-full px-3 py-2 bg-white/10 border border-white/20 rounded-lg focus:border-purple-500 focus:outline-none text-white"
                @input="updateRegion"
              />
            </div>
          </div>
          <div class="flex items-center justify-between">
            <span class="text-xs text-slate-500">
              选择时长: {{ formatTime(selection.end - selection.start) }}
            </span>
            <button
              @click="playSelection"
              class="px-4 py-2 bg-green-600 hover:bg-green-500 rounded-lg text-sm font-medium transition-colors"
            >
              播放选区
            </button>
          </div>
        </div>

        <div class="glass-dark rounded-xl p-4 mt-4">
          <h4 class="font-semibold text-white mb-3">🔧 音频处理</h4>
          <div class="space-y-4">
            <div>
              <label class="block text-sm text-slate-400 mb-2 flex justify-between">
                <span>音量</span>
                <span>{{ volume }}%</span>
              </label>
              <input
                v-model.number="volume"
                type="range"
                min="0"
                max="200"
                class="w-full accent-purple-500"
                @input="applyEffects"
              />
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label class="block text-sm text-slate-400 mb-2 flex justify-between">
                  <span>淡入 (秒)</span>
                  <span>{{ fadeIn }}</span>
                </label>
                <input
                  v-model.number="fadeIn"
                  type="range"
                  min="0"
                  max="5"
                  step="0.5"
                  class="w-full accent-blue-500"
                />
              </div>
              <div>
                <label class="block text-sm text-slate-400 mb-2 flex justify-between">
                  <span>淡出 (秒)</span>
                  <span>{{ fadeOut }}</span>
                </label>
                <input
                  v-model.number="fadeOut"
                  type="range"
                  min="0"
                  max="5"
                  step="0.5"
                  class="w-full accent-pink-500"
                />
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="w-full lg:w-72">
        <label class="block text-sm font-medium text-slate-300 mb-3">
          📤 音频文件
        </label>
        
        <div
          @dragenter.prevent
          @dragover.prevent
          @dragleave.prevent
          @drop.prevent="handleDrop"
          @click="$refs.fileInput.click()"
          class="glass-dark rounded-xl p-8 border-2 border-dashed border-white/20 hover:border-purple-400 cursor-pointer text-center transition-colors"
        >
          <input
            ref="fileInput"
            type="file"
            accept="audio/*"
            class="hidden"
            @change="handleFileSelect"
          />
          
          <div v-if="!selectedFile" class="py-8">
            <div class="text-4xl mb-3">🎶</div>
            <p class="text-slate-300 font-medium mb-1">拖拽或点击上传</p>
            <p class="text-xs text-slate-500">支持 MP3, WAV, FLAC, OGG</p>
          </div>
          
          <div v-else class="py-4">
            <div class="w-16 h-16 mx-auto mb-3 bg-gradient-to-br from-purple-500 to-blue-500 rounded-xl flex items-center justify-center text-2xl">
              🎵
            </div>
            <p class="text-slate-300 font-medium truncate">{{ selectedFile.name }}</p>
            <p class="text-xs text-slate-500 mt-1">{{ formatFileSize(selectedFile.size) }}</p>
            <button
              @click.stop="removeFile"
              class="mt-3 px-4 py-1.5 bg-red-500/20 text-red-400 rounded-lg text-sm hover:bg-red-500/30 transition-colors"
            >
              移除
            </button>
          </div>
        </div>

        <div class="mt-4 space-y-3">
          <button
            @click="previewClip"
            :disabled="!selectedFile || !region"
            class="w-full px-6 py-3 bg-gradient-to-r from-green-600 to-emerald-600 rounded-xl font-medium hover:opacity-90 transition-opacity disabled:opacity-50"
          >
            🔊 预览裁剪
          </button>
          <button
            @click="exportClip"
            :disabled="!selectedFile || !region"
            class="w-full px-6 py-3 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl font-medium hover:opacity-90 transition-opacity disabled:opacity-50"
          >
            💾 导出音频
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, nextTick, watch } from 'vue'
import WaveSurfer from 'wavesurfer.js'
import RegionsPlugin from 'wavesurfer.js/dist/plugins/regions.esm.js'

const props = defineProps({
  modelValue: {
    type: File,
    default: null
  }
})

const emit = defineEmits(['update:modelValue', 'export'])

const waveformContainer = ref(null)
const fileInput = ref(null)

const wavesurfer = ref(null)
const region = ref(null)
const selectedFile = ref(null)
const isPlaying = ref(false)
const currentTime = ref(0)
const duration = ref(0)
const volume = ref(100)
const fadeIn = ref(0)
const fadeOut = ref(0)

const selection = ref({
  start: 0,
  end: 0
})

let audioBuffer = null

const initWavesurfer = () => {
  if (!waveformContainer.value) return
  
  wavesurfer.value = WaveSurfer.create({
    container: waveformContainer.value,
    height: 192,
    waveColor: '#475569',
    progressColor: '#8b5cf6',
    cursorColor: '#ffffff',
    barWidth: 2,
    barRadius: 3,
    barGap: 3,
    normalize: true
  })
  
  wavesurfer.value.registerPlugin(RegionsPlugin.create())
  
  wavesurfer.value.on('ready', () => {
    duration.value = wavesurfer.value.getDuration()
    selection.value = {
      start: 0,
      end: duration.value
    }
    updateRegion()
  })
  
  wavesurfer.value.on('audioprocess', () => {
    currentTime.value = wavesurfer.value.getCurrentTime()
  })
  
  wavesurfer.value.on('play', () => {
    isPlaying.value = true
  })
  
  wavesurfer.value.on('pause', () => {
    isPlaying.value = false
  })
  
  wavesurfer.value.on('finish', () => {
    isPlaying.value = false
    currentTime.value = 0
  })
}

const handleFileSelect = (event) => {
  const file = event.target.files[0]
  if (file) {
    loadFile(file)
  }
}

const handleDrop = (event) => {
  const file = event.dataTransfer.files[0]
  if (file && file.type.startsWith('audio/')) {
    loadFile(file)
  }
}

const loadFile = (file) => {
  selectedFile.value = file
  emit('update:modelValue', file)
  
  const url = URL.createObjectURL(file)
  wavesurfer.value.load(url)
  
  const audioContext = new (window.AudioContext || window.webkitAudioContext)()
  const reader = new FileReader()
  reader.onload = (e) => {
    audioContext.decodeAudioData(e.target.result, (buffer) => {
      audioBuffer = buffer
    })
  }
  reader.readAsArrayBuffer(file)
}

const removeFile = () => {
  selectedFile.value = null
  emit('update:modelValue', null)
  if (wavesurfer.value) {
    wavesurfer.value.empty()
  }
  if (region.value) {
    region.value.remove()
    region.value = null
  }
  audioBuffer = null
  currentTime.value = 0
  duration.value = 0
}

const togglePlay = () => {
  wavesurfer.value.playPause()
}

const playSelection = () => {
  if (!region.value) return
  wavesurfer.value.play(region.value.start, region.value.end)
}

const updateRegion = () => {
  if (!wavesurfer.value) return
  
  if (region.value) {
    region.value.remove()
  }
  
  region.value = wavesurfer.value.addRegion({
    start: selection.value.start,
    end: selection.value.end,
    color: 'rgba(139, 92, 246, 0.3)',
    drag: true,
    resize: true
  })
  
  region.value.on('update', (reg) => {
    selection.value.start = reg.start
    selection.value.end = reg.end
  })
}

const applyEffects = () => {
  if (wavesurfer.value) {
    wavesurfer.value.setVolume(volume.value / 100)
  }
}

const previewClip = async () => {
  if (!region.value) return
  
  const start = region.value.start
  const end = region.value.end
  
  wavesurfer.value.play(start, end)
  
  if (fadeIn.value > 0) {
    wavesurfer.value.setVolume(0)
    let vol = 0
    const fadeInterval = setInterval(() => {
      vol += 10
      wavesurfer.value.setVolume(Math.min(vol, volume.value) / 100)
      if (vol >= volume.value) clearInterval(fadeInterval)
    }, fadeIn.value * 1000 / 10)
  }
}

const exportClip = async () => {
  if (!audioBuffer || !region.value) return
  
  const start = region.value.start
  const end = region.value.end
  
  const originalDuration = audioBuffer.duration
  const clipDuration = end - start
  const sampleRate = audioBuffer.sampleRate
  const channelCount = audioBuffer.numberOfChannels
  const length = clipDuration * sampleRate
  
  const offlineContext = new OfflineAudioContext(
    channelCount,
    length,
    sampleRate
  )
  
  const bufferSource = offlineContext.createBufferSource()
  bufferSource.buffer = audioBuffer
  
  const gainNode = offlineContext.createGain()
  
  const fadeInSamples = fadeIn.value * sampleRate
  const fadeOutSamples = fadeOut.value * sampleRate
  
  gainNode.gain.setValueAtTime(0, 0)
  
  if (fadeIn.value > 0) {
    gainNode.gain.linearRampToValueAtTime(volume.value / 100, fadeIn.value)
  } else {
    gainNode.gain.setValueAtTime(volume.value / 100, 0)
  }
  
  if (fadeOut.value > 0) {
    const fadeOutStart = clipDuration - fadeOut.value
    gainNode.gain.setValueAtTime(volume.value / 100, fadeOutStart)
    gainNode.gain.linearRampToValueAtTime(0, clipDuration)
  }
  
  bufferSource.connect(gainNode)
  gainNode.connect(offlineContext.destination)
  bufferSource.start(0, start, end - start)
  
  const renderedBuffer = await offlineContext.startRendering()
  
  const audioBlob = audioBufferToWaveBlob(renderedBuffer, sampleRate)
  
  emit('export', {
    blob: audioBlob,
    duration: clipDuration
  })
}

const audioBufferToWaveBlob = (buffer, sampleRate) => {
  const numChannels = buffer.numberOfChannels
  const length = buffer.length * numChannels * 2 + 44
  const arrayBuffer = new ArrayBuffer(length)
  const view = new DataView(arrayBuffer)
  
  const writeString = (offset, string) => {
    for (let i = 0; i < string.length; i++) {
      view.setUint8(offset + i, string.charCodeAt(i))
    }
  }
  
  writeString(0, 'RIFF')
  view.setUint32(4, 36 + buffer.length * numChannels * 2, true)
  writeString(8, 'WAVE')
  writeString(12, 'fmt ')
  view.setUint32(16, 16, true)
  view.setUint16(20, 1, true)
  view.setUint16(22, numChannels, true)
  view.setUint32(24, sampleRate, true)
  view.setUint32(28, sampleRate * numChannels * 2, true)
  view.setUint16(32, numChannels * 2, true)
  view.setUint16(34, 16, true)
  writeString(36, 'data')
  view.setUint32(40, buffer.length * numChannels * 2, true)
  
  let offset = 44
  for (let channel = 0; channel < numChannels; channel++) {
    const channelData = buffer.getChannelData(channel)
    for (let i = 0; i < buffer.length; i++) {
      const sample = Math.max(-1, Math.min(1, channelData[i]))
      view.setInt16(offset, sample < 0 ? sample * 0x8000 : sample * 0x7FFF, true)
      offset += 2
    }
  }
  
  return new Blob([arrayBuffer], { type: 'audio/wav' })
}

const formatTime = (seconds) => {
  if (!seconds || isNaN(seconds)) return '0:00'
  const mins = Math.floor(seconds / 60)
  const secs = Math.floor(seconds % 60).toString().padStart(2, '0')
  return `${mins}:${secs}`
}

const formatFileSize = (bytes) => {
  if (!bytes) return '0 B'
  const units = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(1024))
  return `${(bytes / Math.pow(1024, i)).toFixed(2)} ${units[i]}`
}

watch(() => props.modelValue, (newVal) => {
  if (newVal && newVal !== selectedFile.value) {
    loadFile(newVal)
  }
})

onMounted(() => {
  nextTick(() => {
    initWavesurfer()
  })
})

onUnmounted(() => {
  if (wavesurfer.value) {
    wavesurfer.value.destroy()
  }
})
</script>
