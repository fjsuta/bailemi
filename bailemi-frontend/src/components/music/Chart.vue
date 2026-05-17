<template>
  <div class="glass-dark rounded-2xl p-6">
    <div class="flex items-center justify-between mb-6">
      <div class="flex items-center gap-3">
        <div class="w-10 h-10 bg-gradient-to-br from-orange-500 to-red-500 rounded-xl flex items-center justify-center">🏆</div>
        <h2 class="text-xl font-bold">排行榜</h2>
      </div>
      <button class="text-purple-400 hover:text-purple-300 transition-colors">查看全部 →</button>
    </div>

    <div class="flex gap-2 mb-6 overflow-x-auto pb-2">
      <button
        v-for="tab in tabs"
        :key="tab.id"
        @click="activeTab = tab.id"
        :class="[
          'px-4 py-2 rounded-xl font-medium whitespace-nowrap transition-all',
          activeTab === tab.id
            ? 'bg-gradient-to-r from-purple-600 to-blue-600 text-white'
            : 'bg-white/10 text-slate-300 hover:bg-white/20'
        ]"
      >
        {{ tab.icon }} {{ tab.name }}
      </button>
    </div>

    <div class="space-y-3">
      <div
        v-for="(song, index) in currentChart"
        :key="song.id"
        @click="playSong(song)"
        class="flex items-center gap-4 p-3 rounded-xl hover:bg-white/10 cursor-pointer transition-all group"
      >
        <div :class="[
          'w-8 h-8 rounded-lg flex items-center justify-center font-bold text-sm',
          index === 0 ? 'bg-gradient-to-br from-yellow-400 to-orange-500 text-white' :
          index === 1 ? 'bg-gradient-to-br from-gray-300 to-gray-400 text-white' :
          index === 2 ? 'bg-gradient-to-br from-orange-400 to-orange-600 text-white' :
          'bg-white/10 text-slate-400'
        ]">
          {{ index + 1 }}
        </div>
        <img :src="song.cover" class="w-12 h-12 rounded-lg object-cover" />
        <div class="flex-1 min-w-0">
          <h4 class="font-medium truncate group-hover:text-purple-400 transition-colors">{{ song.title }}</h4>
          <p class="text-sm text-slate-400 truncate">{{ song.artist }}</p>
        </div>
        <div class="flex items-center gap-2">
          <span :class="[
            'text-xs px-2 py-1 rounded-full',
            song.trend > 0 ? 'bg-green-500/20 text-green-400' :
            song.trend < 0 ? 'bg-red-500/20 text-red-400' :
            'bg-slate-500/20 text-slate-400'
          ]">
            {{ song.trend > 0 ? '+' : '' }}{{ song.trend }}
          </span>
          <span class="text-slate-500 text-sm">{{ song.playCount }}</span>
        </div>
        <div class="opacity-0 group-hover:opacity-100 transition-opacity text-white">▶</div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useMusicStore } from '@/stores/music'

const musicStore = useMusicStore()

const activeTab = ref('hot')

const tabs = [
  { id: 'hot', name: '热歌榜', icon: '🔥' },
  { id: 'new', name: '新歌榜', icon: '✨' },
  { id: 'soar', name: '飙升榜', icon: '📈' },
  { id: 'original', name: '原创榜', icon: '🎨' }
]

const charts = ref({
  hot: [
    { id: 1, title: '早安隆回', artist: '袁树雄', cover: 'https://picsum.photos/100/100?random=c1', trend: 0, playCount: '1.2亿' },
    { id: 2, title: '孤勇者', artist: '陈奕迅', cover: 'https://picsum.photos/100/100?random=c2', trend: -1, playCount: '9800万' },
    { id: 3, title: '本草纲目', artist: '周杰伦', cover: 'https://picsum.photos/100/100?random=c3', trend: 2, playCount: '8600万' },
    { id: 4, title: '起风了', artist: '买辣椒也用券', cover: 'https://picsum.photos/100/100?random=c4', trend: -2, playCount: '7500万' },
    { id: 5, title: '世界那么大还是遇见你', artist: '程响', cover: 'https://picsum.photos/100/100?random=c5', trend: 1, playCount: '6800万' },
    { id: 6, title: '可可托海的牧羊人', artist: '王琪', cover: 'https://picsum.photos/100/100?random=c6', trend: 3, playCount: '6200万' },
    { id: 7, title: '听闻远方有你', artist: '刘艺雯', cover: 'https://picsum.photos/100/100?random=c7', trend: -1, playCount: '5800万' },
    { id: 8, title: '踏山河', artist: '七叔（叶泽浩）', cover: 'https://picsum.photos/100/100?random=c8', trend: 0, playCount: '5400万' }
  ],
  new: [
    { id: 11, title: '最伟大的作品', artist: '周杰伦', cover: 'https://picsum.photos/100/100?random=n1', trend: 1, playCount: '5200万' },
    { id: 12, title: '还在流浪', artist: '周杰伦', cover: 'https://picsum.photos/100/100?random=n2', trend: 2, playCount: '4800万' },
    { id: 13, title: '红颜如霜', artist: '周杰伦', cover: 'https://picsum.photos/100/100?random=n3', trend: 3, playCount: '4500万' }
  ],
  soar: [
    { id: 21, title: '太想念', artist: '豆包', cover: 'https://picsum.photos/100/100?random=s1', trend: 12, playCount: '3200万' },
    { id: 22, title: '我曾遇到一束光', artist: '叶斯淳', cover: 'https://picsum.photos/100/100?random=s2', trend: 8, playCount: '2800万' },
    { id: 23, title: '一路生花', artist: '温奕心', cover: 'https://picsum.photos/100/100?random=s3', trend: 6, playCount: '2500万' }
  ],
  original: [
    { id: 31, title: '入海', artist: '毛不易', cover: 'https://picsum.photos/100/100?random=o1', trend: 2, playCount: '1800万' },
    { id: 32, title: '消愁', artist: '毛不易', cover: 'https://picsum.photos/100/100?random=o2', trend: 1, playCount: '1500万' }
  ]
})

const currentChart = computed(() => charts.value[activeTab.value] || charts.value.hot)

const playSong = (song) => {
  musicStore.playSong(song, currentChart.value)
}
</script>
