<template>
  <div class="relative overflow-hidden rounded-2xl h-[400px] md:h-[500px]">
    <div
      class="flex transition-transform duration-700 ease-in-out h-full"
      :style="{ transform: `translateX(-${currentIndex * 100}%)` }"
    >
      <div
        v-for="(banner, index) in banners"
        :key="banner.id"
        class="min-w-full h-full relative"
      >
        <img
          :src="banner.image"
          :alt="banner.title"
          class="w-full h-full object-cover"
        />
        <div class="absolute inset-0 bg-gradient-to-t from-black/70 via-black/30 to-transparent">
          <div class="absolute bottom-0 left-0 right-0 p-6 md:p-10">
            <span class="px-3 py-1 bg-purple-600/80 rounded-full text-sm font-medium mb-4 inline-block">
              {{ banner.tag }}
            </span>
            <h2 class="text-2xl md:text-4xl font-bold text-white mb-2">{{ banner.title }}</h2>
            <p class="text-slate-300 mb-6">{{ banner.description }}</p>
            <button class="px-6 py-3 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl font-medium hover:opacity-90 transition-opacity">
              {{ banner.buttonText }}
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="absolute bottom-4 left-1/2 -translate-x-1/2 flex gap-2">
      <button
        v-for="(_, index) in banners"
        :key="index"
        @click="goToSlide(index)"
        :class="[
          'w-2 h-2 rounded-full transition-all',
          currentIndex === index ? 'bg-white w-6' : 'bg-white/50 hover:bg-white/70'
        ]"
      ></button>
    </div>

    <button
      @click="prevSlide"
      class="absolute left-4 top-1/2 -translate-y-1/2 w-10 h-10 glass rounded-full flex items-center justify-center hover:bg-white/20 transition-all"
    >
      <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M15 18l-6-6 6-6"/>
      </svg>
    </button>
    <button
      @click="nextSlide"
      class="absolute right-4 top-1/2 -translate-y-1/2 w-10 h-10 glass rounded-full flex items-center justify-center hover:bg-white/20 transition-all"
    >
      <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M9 18l6-6-6-6"/>
      </svg>
    </button>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const currentIndex = ref(0)
let timer = null

const banners = ref([
  {
    id: 1,
    title: '发现好音乐',
    description: '分享您的原创作品，发现更多精彩内容',
    tag: '推荐',
    image: 'https://picsum.photos/1200/500?random=1',
    buttonText: '立即探索'
  },
  {
    id: 2,
    title: '热门分类',
    description: '按流派和场景快速找到喜欢的内容',
    tag: '分类',
    image: 'https://picsum.photos/1200/500?random=2',
    buttonText: '查看分类'
  },
  {
    id: 3,
    title: '上传您的作品',
    description: '分享您的创作，支持多种音频格式',
    tag: '创作',
    image: 'https://picsum.photos/1200/500?random=3',
    buttonText: '开始上传'
  }
])

const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % banners.value.length
}

const prevSlide = () => {
  currentIndex.value = (currentIndex.value - 1 + banners.value.length) % banners.value.length
}

const goToSlide = (index) => {
  currentIndex.value = index
}

onMounted(() => {
  timer = setInterval(nextSlide, 5000)
})

onUnmounted(() => {
  if (timer) clearInterval(timer)
})
</script>
