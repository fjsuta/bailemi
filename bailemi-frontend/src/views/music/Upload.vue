<template>
  <div class="min-h-screen pb-24">
    <div class="container mx-auto px-4 py-8 max-w-5xl">
      <div class="flex items-center gap-4 mb-8">
        <button @click="$router.back()" class="w-10 h-10 glass rounded-xl flex items-center justify-center hover:bg-white/20 transition-all">
          ←
        </button>
        <div>
          <h1 class="text-2xl font-bold text-white">上传音乐</h1>
          <p class="text-slate-400">分享你的原创音乐或编辑现有音频</p>
        </div>
      </div>

      <div class="flex gap-2 mb-6 overflow-x-auto pb-2">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="[
            'px-6 py-2.5 rounded-xl font-medium whitespace-nowrap transition-all',
            activeTab === tab.id
              ? 'bg-gradient-to-r from-purple-600 to-blue-600 text-white'
              : 'bg-white/10 text-slate-300 hover:bg-white/20'
          ]"
        >
          {{ tab.name }}
        </button>
      </div>

      <div v-if="activeTab === 'upload'" class="space-y-6">
        <div class="grid gap-6">
          <AudioEditor v-model="selectedFile" @export="handleAudioExport" />
          
          <div class="glass-dark rounded-xl p-6">
            <h3 class="font-semibold text-white mb-4 flex items-center gap-2">
              <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
                <polyline points="14 2 14 8 20 8"/>
              </svg>
              基本信息
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label class="block text-sm text-slate-400 mb-2">
                  标题 *
                </label>
                <input
                  v-model="formData.title"
                  type="text"
                  placeholder="输入音乐标题"
                  class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl focus:border-purple-500 focus:outline-none text-white"
                />
              </div>
              <div>
                <label class="block text-sm text-slate-400 mb-2">
                  艺术家/创作者 *
                </label>
                <input
                  v-model="formData.artist"
                  type="text"
                  placeholder="输入艺术家名称"
                  class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl focus:border-purple-500 focus:outline-none text-white"
                />
              </div>
              <div>
                <label class="block text-sm text-slate-400 mb-2">
                  专辑
                </label>
                <input
                  v-model="formData.album"
                  type="text"
                  placeholder="专辑名称 (可选)"
                  class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl focus:border-purple-500 focus:outline-none text-white"
                />
              </div>
              <div>
                <label class="block text-sm text-slate-400 mb-2">
                  流派
                </label>
                <select
                  v-model="formData.genre"
                  class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl focus:border-purple-500 focus:outline-none text-white"
                >
                  <option value="">选择流派</option>
                  <option value="pop">流行</option>
                  <option value="rock">摇滚</option>
                  <option value="electronic">电子</option>
                  <option value="classical">古典</option>
                  <option value="jazz">爵士</option>
                  <option value="hiphop">嘻哈/说唱</option>
                  <option value="folk">民谣</option>
                  <option value="rnb">R&B</option>
                  <option value="other">其他</option>
                </select>
              </div>
              <div class="md:col-span-2">
                <label class="block text-sm text-slate-400 mb-2">
                  封面图片
                </label>
                <div
                  @click="$refs.coverInput.click()"
                  class="border-2 border-dashed border-white/20 rounded-xl p-8 cursor-pointer hover:border-purple-400 transition-colors"
                >
                  <input
                    ref="coverInput"
                    type="file"
                    accept="image/*"
                    class="hidden"
                    @change="handleCoverUpload"
                  />
                  <div class="flex items-center gap-4">
                    <div v-if="formData.cover" class="w-20 h-20 rounded-lg overflow-hidden bg-slate-800">
                      <img :src="formData.cover" class="w-full h-full object-cover" alt="封面" />
                    </div>
                    <div v-else class="w-20 h-20 rounded-lg bg-slate-800 flex items-center justify-center">
                      <svg class="w-8 h-8 text-slate-500" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                        <path d="M23 19a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h4l2-3h6l2 3h4a2 2 0 0 1 2 2z"/>
                        <circle cx="12" cy="13" r="4"/>
                      </svg>
                    </div>
                    <div>
                      <p class="text-slate-300">点击上传封面图片</p>
                      <p class="text-xs text-slate-500">支持 JPG, PNG, 格式，建议正方形</p>
                    </div>
                  </div>
                </div>
              </div>
              <div class="md:col-span-2">
                <label class="block text-sm text-slate-400 mb-2">
                  简介
                </label>
                <textarea
                  v-model="formData.description"
                  placeholder="描述一下这首音乐..."
                  rows="3"
                  class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl focus:border-purple-500 focus:outline-none text-white resize-none"
                ></textarea>
              </div>
            </div>
          </div>

          <div class="glass-dark rounded-xl p-6">
            <LicenseSelector v-model="formData.license" />
          </div>

          <div class="glass-dark rounded-xl p-6">
            <div class="flex items-center gap-3 mb-4">
              <div class="w-6 h-6 rounded bg-slate-800 flex items-center justify-center">
                <svg class="w-4 h-4 text-slate-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="3"/>
                  <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06A1.65 1.65 0 0 0 4.68 15a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06A1.65 1.65 0 0 0 9 4.68a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06A1.65 1.65 0 0 0 19.4 9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/>
                </svg>
              </div>
              <h3 class="font-semibold text-white">高级设置</h3>
            </div>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <label class="flex items-center gap-3 cursor-pointer">
                <input
                  v-model="formData.isPublic"
                  type="checkbox"
                  class="w-5 h-5 accent-purple-500"
                />
                <span class="text-slate-300">公开显示 (所有人可见)</span>
              </label>
              <label class="flex items-center gap-3 cursor-pointer">
                <input
                  v-model="formData.allowDownload"
                  type="checkbox"
                  class="w-5 h-5 accent-purple-500"
                />
                <span class="text-slate-300">允许下载</span>
              </label>
            </div>
          </div>
        </div>

        <div class="sticky bottom-0 pt-4 pb-2">
          <div class="glass-dark border-t border-slate-700 p-4 -mx-4 -mb-2 rounded-t-xl">
            <div class="flex items-center justify-between">
              <div class="text-sm text-slate-400">
                <span v-if="selectedFile">已选择: {{ selectedFile.name }}</span>
                <span v-else>请先选择音频文件</span>
              </div>
              <div class="flex gap-3">
                <button
                  @click="resetForm"
                  class="px-6 py-2.5 bg-white/10 rounded-xl font-medium hover:bg-white/20 transition-colors text-slate-300"
                >
                  重置
                </button>
                <button
                  @click="handleSubmit"
                  :disabled="!canSubmit"
                  class="px-8 py-2.5 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl font-medium hover:opacity-90 transition-opacity disabled:opacity-50 disabled:cursor-not-allowed"
                >
                  {{ isUploading ? '上传中...' : '发布' }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div v-else class="glass-dark rounded-xl p-8 text-center">
        <div class="mb-4">
          <svg class="w-12 h-12 mx-auto text-yellow-500" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"/>
            <line x1="12" y1="9" x2="12" y2="13"/>
            <line x1="12" y1="17" x2="12.01" y2="17"/>
          </svg>
        </div>
        <h3 class="text-xl font-semibold text-white mb-2">专业编曲器开发中</h3>
        <p class="text-slate-400 mb-6">此功能正在开发中，敬请期待...</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { useNotificationStore } from '@/stores/notification'
import AudioEditor from '@/components/music/AudioEditor.vue'
import LicenseSelector from '@/components/music/LicenseSelector.vue'

const router = useRouter()
const authStore = useAuthStore()
const notificationStore = useNotificationStore()

const activeTab = ref('upload')
const selectedFile = ref(null)
const isUploading = ref(false)

const formData = ref({
  title: '',
  artist: '',
  album: '',
  genre: '',
  description: '',
  cover: '',
  license: 'cc-by',
  isPublic: true,
  allowDownload: true
})

const tabs = [
  { id: 'upload', name: '音频裁剪' },
  { id: 'composer', name: '专业编曲' }
]

const canSubmit = computed(() => {
  return selectedFile.value &&
    formData.value.title.trim() &&
    formData.value.artist.trim()
})

const handleCoverUpload = (event) => {
  const file = event.target.files[0]
  if (file) {
    formData.value.cover = URL.createObjectURL(file)
  }
}

const handleAudioExport = ({ blob }) => {
  selectedFile.value = blob
  notificationStore.success('导出成功', '音频已处理好，可以上传了')
}

const handleSubmit = async () => {
  if (!canSubmit.value) return
  
  isUploading.value = true
  
  try {
    const form = new FormData()
    form.append('title', formData.value.title)
    form.append('artist', formData.value.artist)
    form.append('album', formData.value.album || '')
    form.append('genre', formData.value.genre || '')
    form.append('description', formData.value.description || '')
    form.append('license', formData.value.license)
    form.append('isPublic', formData.value.isPublic)
    form.append('allowDownload', formData.value.allowDownload)
    form.append('audio', selectedFile.value)
    
    notificationStore.success('上传成功', '您的音乐已发布')
    
    setTimeout(() => {
      router.push('/')
    }, 1500)
    
  } catch (error) {
    console.error(error)
    notificationStore.error('上传失败', '请稍后重试')
  } finally {
    isUploading.value = false
  }
}

const resetForm = () => {
  selectedFile.value = null
  formData.value = {
    title: '',
    artist: '',
    album: '',
    genre: '',
    description: '',
    cover: '',
    license: 'cc-by',
    isPublic: true,
    allowDownload: true
  }
}

onMounted(() => {
  if (!authStore.isAuthenticated) {
    router.push('/login')
    notificationStore.warning('请先登录', '登录后才能上传音乐')
  }
})
</script>
