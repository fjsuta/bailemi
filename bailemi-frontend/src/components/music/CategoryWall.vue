<template>
  <div class="glass-dark rounded-2xl p-6">
    <div class="flex items-center gap-3 mb-6">
      <div class="w-10 h-10 bg-gradient-to-br from-green-500 to-teal-500 rounded-xl flex items-center justify-center">
          <svg class="w-5 h-5 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="12" cy="12" r="10"/>
            <circle cx="12" cy="12" r="3"/>
          </svg>
        </div>
      <h2 class="text-xl font-bold">歌单分类</h2>
    </div>

    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
      <button
        v-for="category in categories"
        :key="category.id"
        @click="filterByCategory(category.id)"
        class="group relative aspect-square rounded-xl overflow-hidden"
      >
        <img :src="category.image" class="w-full h-full object-cover group-hover:scale-110 transition-transform" />
        <div class="absolute inset-0 bg-gradient-to-t from-black/80 via-black/20 to-transparent">
          <div class="absolute bottom-0 left-0 right-0 p-4">
            <h4 class="font-bold text-white">{{ category.name }}</h4>
            <p class="text-xs text-slate-300">{{ category.count }}个歌单</p>
          </div>
        </div>
      </button>
    </div>

    <div class="mt-6">
      <h3 class="font-semibold mb-4">场景歌单</h3>
      <div class="flex flex-wrap gap-3">
        <button
          v-for="scene in scenes"
          :key="scene.id"
          @click="filterByScene(scene.id)"
          :class="[
            'px-4 py-2 rounded-full font-medium transition-all',
            activeScene === scene.id
              ? 'bg-gradient-to-r from-purple-600 to-blue-600 text-white'
              : 'bg-white/10 text-slate-300 hover:bg-white/20'
          ]"
        >
          {{ scene.name }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const activeScene = ref('study')

const categories = ref([
  { id: 'chinese', name: '华语', count: 128, image: 'https://picsum.photos/300/300?random=c1' },
  { id: 'english', name: '欧美', count: 96, image: 'https://picsum.photos/300/300?random=c2' },
  { id: 'korean', name: '韩语', count: 64, image: 'https://picsum.photos/300/300?random=c3' },
  { id: 'japanese', name: '日语', count: 52, image: 'https://picsum.photos/300/300?random=c4' },
  { id: 'pop', name: '流行', count: 156, image: 'https://picsum.photos/300/300?random=c5' },
  { id: 'rock', name: '摇滚', count: 78, image: 'https://picsum.photos/300/300?random=c6' },
  { id: 'classical', name: '古典', count: 42, image: 'https://picsum.photos/300/300?random=c7' },
  { id: 'electronic', name: '电子', count: 89, image: 'https://picsum.photos/300/300?random=c8' }
])

const scenes = ref([
  { id: 'study', name: '学习' },
  { id: 'workout', name: '运动' },
  { id: 'sleep', name: '睡眠' },
  { id: 'focus', name: '专注' },
  { id: 'party', name: '聚会' },
  { id: 'driving', name: '驾车' },
  { id: 'romantic', name: '浪漫' },
  { id: 'travel', name: '旅行' }
])

const filterByCategory = (categoryId) => {
  console.log('Filter by category:', categoryId)
}

const filterByScene = (sceneId) => {
  activeScene.value = sceneId
  console.log('Filter by scene:', sceneId)
}
</script>
