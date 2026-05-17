<template>
  <div class="min-h-screen bg-gradient-to-br from-purple-500/10 to-blue-500/10">
    <div class="container mx-auto px-4 py-8">
      <header class="flex items-center justify-between mb-8">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center text-2xl">⚙️</div>
          <div>
            <h1 class="text-2xl font-bold text-white">系统管理</h1>
            <p class="text-slate-400 text-sm">百米乐音乐平台</p>
          </div>
        </div>
        <router-link to="/" class="px-4 py-2 glass rounded-xl hover:bg-white/20 transition-all text-slate-300">
          ← 返回首页
        </router-link>
      </header>

      <div class="flex gap-4 mb-6 overflow-x-auto">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          :class="[
            'px-6 py-3 rounded-xl font-medium transition-all whitespace-nowrap',
            activeTab === tab.id
              ? 'bg-gradient-to-r from-purple-600 to-blue-600 text-white'
              : 'glass text-slate-300 hover:bg-white/20'
          ]"
        >
          {{ tab.icon }} {{ tab.name }}
        </button>
      </div>

      <div class="glass-dark rounded-2xl p-8">
        <div v-if="activeTab === 'update'">
          <UpdateManager />
        </div>
        <div v-else-if="activeTab === 'icp'">
          <ICPManager />
        </div>
        <div v-else-if="activeTab === 'storage'">
          <StorageManager />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import UpdateManager from './UpdateManager.vue'
import ICPManager from './ICPManager.vue'
import StorageManager from './StorageManager.vue'

const activeTab = ref('update')

const tabs = [
  { id: 'update', name: '更新管理', icon: '🔄' },
  { id: 'icp', name: '公安备案', icon: '📄' },
  { id: 'storage', name: '存储设置', icon: '💾' }
]
</script>
