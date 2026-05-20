<template>
  <div class="min-h-screen bg-gradient-to-br from-slate-900 via-purple-900 to-slate-900 flex items-center justify-center p-4">
    <div class="w-full max-w-2xl bg-white/10 backdrop-blur-xl rounded-3xl border border-white/20 p-8 shadow-2xl">
      <!-- 标题 -->
      <div class="text-center mb-8">
        <div class="w-16 h-16 mx-auto mb-4 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center">
          <svg class="w-8 h-8 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M5 12h14M12 5l7 7-7 7"/>
          </svg>
        </div>
        <h1 class="text-3xl font-bold text-white mb-2">后端配置</h1>
        <p class="text-slate-400">选择或配置后端服务</p>
      </div>

      <!-- 后端列表 -->
      <div class="space-y-4 mb-6">
        <div
          v-for="backend in backendStore.backends"
          :key="backend.id"
          @click="selectBackend(backend.id)"
          :class="[
            'p-4 rounded-2xl border-2 cursor-pointer transition-all',
            backendStore.currentBackendId === backend.id
              ? 'border-purple-500 bg-purple-500/20'
              : 'border-white/10 bg-white/5 hover:border-white/20'
          ]"
        >
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-4">
              <!-- 状态指示器 -->
              <div
                :class="[
                  'w-3 h-3 rounded-full',
                  backendStore.backendStatus[backend.id]?.online ? 'bg-green-500' : 'bg-red-500'
                ]"
              ></div>
              
              <div>
                <h3 class="font-semibold text-white">{{ backend.name }}</h3>
                <p class="text-sm text-slate-400">{{ backend.url }}</p>
                <p class="text-xs text-slate-500">{{ backend.description }}</p>
              </div>
            </div>

            <div class="text-right">
              <span
                v-if="backendStore.backendStatus[backend.id]?.online"
                class="text-xs text-green-400"
              >
                {{ backendStore.backendStatus[backend.id]?.latency }}ms
              </span>
              <span v-else class="text-xs text-red-400">离线</span>
              
              <!-- 选中标记 -->
              <div
                v-if="backendStore.currentBackendId === backend.id"
                class="mt-2 text-purple-400"
              >
                <svg class="w-5 h-5" viewBox="0 0 24 24" fill="currentColor">
                  <path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41z"/>
                </svg>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 添加新后端 -->
      <div v-if="showAddForm" class="mb-6 p-4 bg-white/5 rounded-2xl border border-white/10">
        <h3 class="font-semibold text-white mb-4">添加新后端</h3>
        <div class="space-y-4">
          <div>
            <label class="block text-sm text-slate-400 mb-1">名称</label>
            <input
              v-model="newBackend.name"
              type="text"
              placeholder="例如：生产环境"
              class="w-full px-4 py-2 bg-white/10 border border-white/20 rounded-xl text-white placeholder-slate-500 focus:outline-none focus:border-purple-500"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">URL</label>
            <input
              v-model="newBackend.url"
              type="text"
              placeholder="http://localhost:8080/api/v1"
              class="w-full px-4 py-2 bg-white/10 border border-white/20 rounded-xl text-white placeholder-slate-500 focus:outline-none focus:border-purple-500"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">描述</label>
            <input
              v-model="newBackend.description"
              type="text"
              placeholder="后端描述"
              class="w-full px-4 py-2 bg-white/10 border border-white/20 rounded-xl text-white placeholder-slate-500 focus:outline-none focus:border-purple-500"
            />
          </div>
          <div class="flex gap-2">
            <button
              @click="addBackend"
              class="flex-1 py-2 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl font-medium hover:from-purple-500 hover:to-blue-500 transition-all"
            >
              添加
            </button>
            <button
              @click="showAddForm = false"
              class="flex-1 py-2 bg-white/10 border border-white/20 rounded-xl font-medium hover:bg-white/20 transition-all"
            >
              取消
            </button>
          </div>
        </div>
      </div>

      <!-- 操作按钮 -->
      <div class="flex flex-wrap gap-3">
        <button
          @click="showAddForm = true"
          v-if="!showAddForm"
          class="flex-1 py-3 bg-white/10 border border-white/20 rounded-xl font-medium hover:bg-white/20 transition-all flex items-center justify-center gap-2"
        >
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M12 5v14M5 12h14"/>
          </svg>
          添加后端
        </button>
        
        <button
          @click="checkStatus"
          :disabled="backendStore.isCheckingStatus"
          class="flex-1 py-3 bg-white/10 border border-white/20 rounded-xl font-medium hover:bg-white/20 transition-all flex items-center justify-center gap-2 disabled:opacity-50"
        >
          <svg
            :class="['w-5 h-5', backendStore.isCheckingStatus ? 'animate-spin' : '']"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
          </svg>
          {{ backendStore.isCheckingStatus ? '检测中...' : '检测状态' }}
        </button>
        
        <button
          @click="backendStore.resetToDefault"
          class="py-3 px-4 bg-white/10 border border-white/20 rounded-xl font-medium hover:bg-white/20 transition-all"
          title="重置为默认"
        >
          <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
          </svg>
        </button>
      </div>

      <!-- 进入应用按钮 -->
      <button
        @click="enterApp"
        :disabled="!backendStore.backendStatus[backendStore.currentBackendId]?.online"
        class="w-full mt-6 py-4 bg-gradient-to-r from-purple-600 to-blue-600 rounded-2xl font-semibold text-lg hover:from-purple-500 hover:to-blue-500 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {{ backendStore.backendStatus[backendStore.currentBackendId]?.online ? '进入应用' : '当前后端离线' }}
      </button>

      <!-- 提示信息 -->
      <p class="mt-4 text-center text-sm text-slate-500">
        选择一个可用的后端服务以继续使用
      </p>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useBackendStore } from '@/stores/backend'

const router = useRouter()
const backendStore = useBackendStore()

const showAddForm = ref(false)
const newBackend = ref({
  name: '',
  url: '',
  description: ''
})

const selectBackend = (id) => {
  backendStore.setCurrentBackend(id)
}

const addBackend = () => {
  if (!newBackend.value.name || !newBackend.value.url) {
    return
  }
  
  backendStore.addBackend({
    name: newBackend.value.name,
    url: newBackend.value.url,
    description: newBackend.value.description || '自定义后端'
  })
  
  newBackend.value = { name: '', url: '', description: '' }
  showAddForm.value = false
  
  // 检测新添加的后端状态
  backendStore.checkAllBackendsStatus()
}

const checkStatus = () => {
  backendStore.checkAllBackendsStatus()
}

const enterApp = () => {
  // 标记已完成引导
  localStorage.setItem('backendConfigured', 'true')
  router.push('/')
}

onMounted(() => {
  // 初始化时检测后端状态
  backendStore.checkAllBackendsStatus()
})
</script>
