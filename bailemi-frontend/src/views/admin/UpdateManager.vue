<template>
  <div>
    <h2 class="text-xl font-bold mb-6 text-white">🔄 自动更新管理</h2>
    
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
      <div class="glass rounded-xl p-6">
        <h3 class="font-semibold text-slate-300 mb-3">当前版本</h3>
        <p class="text-2xl font-bold text-purple-400">{{ currentVersion || '检测中...' }}</p>
      </div>
      <div class="glass rounded-xl p-6">
        <h3 class="font-semibold text-slate-300 mb-3">更新状态</h3>
        <div class="flex items-center gap-2">
          <span
            :class="[
              'w-3 h-3 rounded-full animate-pulse',
              hasUpdate ? 'bg-green-500' : 'bg-slate-500'
            ]"
          ></span>
          <span class="text-slate-300">{{ updateMessage }}</span>
        </div>
      </div>
    </div>

    <div class="flex gap-4 mb-8">
      <button
        @click="checkUpdate"
        :disabled="checking"
        class="px-6 py-3 bg-gradient-to-r from-purple-600 to-blue-600 hover:from-purple-500 hover:to-blue-500 rounded-xl font-medium transition-all disabled:opacity-50"
      >
        {{ checking ? '检查中...' : '立即检查更新' }}
      </button>
      <button
        v-if="hasUpdate"
        @click="doUpdate"
        :disabled="updating"
        class="px-6 py-3 bg-gradient-to-r from-green-600 to-emerald-600 hover:from-green-500 hover:to-emerald-500 rounded-xl font-medium transition-all disabled:opacity-50"
      >
        {{ updating ? '更新中...' : '执行更新' }}
      </button>
    </div>

    <div class="glass rounded-xl p-6">
      <h3 class="font-semibold text-slate-300 mb-4">更新日志</h3>
      <div v-if="logs.length === 0" class="text-center py-8 text-slate-500">
        暂无更新记录
      </div>
      <div v-else class="space-y-3">
        <div
          v-for="log in logs"
          :key="log.id"
          class="flex items-center justify-between p-4 bg-white/5 rounded-xl"
        >
          <div>
            <div class="flex items-center gap-2">
              <span class="font-medium text-white">{{ log.version }}</span>
              <span
                :class="[
                  'px-2 py-1 text-xs rounded-full',
                  log.status === 'completed'
                    ? 'bg-green-500/20 text-green-400'
                    : log.status === 'failed'
                    ? 'bg-red-500/20 text-red-400'
                    : 'bg-yellow-500/20 text-yellow-400'
                ]"
              >
                {{ log.status }}
              </span>
            </div>
            <p class="text-sm text-slate-400 mt-1">{{ log.description }}</p>
            <p class="text-xs text-slate-500 mt-1">{{ formatDate(log.created_at) }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'

const notification = useNotificationStore()

const currentVersion = ref('')
const hasUpdate = ref(false)
const updateMessage = ref('')
const checking = ref(false)
const updating = ref(false)
const logs = ref([])

const checkUpdate = async () => {
  checking.value = true
  try {
    const res = await api.get('/admin/update/check')
    if (res.data.code === 0) {
      currentVersion.value = res.data.data.version
      hasUpdate.value = res.data.data.has_update
      updateMessage.value = res.data.data.message
    }
    await loadLogs()
  } catch (error) {
    notification.error('检查更新失败', error.message)
  } finally {
    checking.value = false
  }
}

const doUpdate = async () => {
  if (!confirm('确定要执行更新吗？更新后需要重启服务。')) {
    return
  }
  updating.value = true
  try {
    const res = await api.post('/admin/update/do')
    if (res.data.code === 0) {
      notification.success('更新成功', '请重启服务以应用更新')
      hasUpdate.value = false
    }
    await loadLogs()
  } catch (error) {
    notification.error('更新失败', error.message)
  } finally {
    updating.value = false
  }
}

const loadLogs = async () => {
  try {
    const res = await api.get('/admin/update/logs')
    if (res.data.code === 0) {
      logs.value = res.data.data.items || []
    }
  } catch (error) {
    console.error(error)
  }
}

const formatDate = (date) => {
  return new Date(date).toLocaleString('zh-CN')
}

onMounted(() => {
  checkUpdate()
})
</script>
