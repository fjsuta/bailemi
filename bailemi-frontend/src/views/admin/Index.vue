<template>
  <div class="min-h-screen bg-gradient-to-br from-purple-500/10 to-blue-500/10">
    <div class="container mx-auto px-4 py-8">
      <header class="flex items-center justify-between mb-8">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-2xl bg-gradient-to-br from-purple-500 to-blue-500 flex items-center justify-center">
            <svg class="w-6 h-6 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="3"/>
              <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06A1.65 1.65 0 0 0 4.68 15a1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06A1.65 1.65 0 0 0 9 4.68a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06A1.65 1.65 0 0 0 19.4 9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/>
            </svg>
          </div>
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
          {{ tab.name }}
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
  { id: 'update', name: '更新管理' },
  { id: 'icp', name: '公安备案' },
  { id: 'storage', name: '存储设置' }
]
</script>
