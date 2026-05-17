<template>
  <div class="space-y-4">
    <label class="block text-sm font-medium text-slate-300 mb-2">
      📜 版权协议选择
    </label>
    <div class="grid gap-3">
      <div
        v-for="license in licenses"
        :key="license.id"
        @click="selectLicense(license.id)"
        :class="[
          'p-4 rounded-xl border cursor-pointer transition-all duration-300',
          modelValue === license.id
            ? 'border-purple-500 bg-purple-500/10'
            : 'border-white/20 hover:border-purple-400/50 hover:bg-white/5'
        ]"
      >
        <div class="flex items-center gap-3 mb-2">
          <span class="text-2xl">{{ license.icon }}</span>
          <div class="flex-1">
            <h4 class="font-semibold text-white">{{ license.name }}</h4>
            <p class="text-xs text-slate-400">{{ license.shortDesc }}</p>
          </div>
          <div
            v-if="modelValue === license.id"
            class="w-6 h-6 rounded-full bg-purple-500 flex items-center justify-center"
          >
            <span class="text-white text-xs">✓</span>
          </div>
        </div>
        <p class="text-xs text-slate-500 pl-9">{{ license.description }}</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: 'cc-by'
  }
})

const emit = defineEmits(['update:modelValue'])

const licenses = ref([
  {
    id: 'cc0',
    name: 'CC0 1.0 通用',
    shortDesc: '公共领域贡献',
    icon: '🌍',
    description: '作者放弃所有版权，任何人都可以随意使用、修改和商用，无需署名'
  },
  {
    id: 'cc-by',
    name: 'CC BY 4.0',
    shortDesc: '署名许可',
    icon: '📝',
    description: '他人可以自由使用、修改和商用，但必须注明原作者'
  },
  {
    id: 'cc-by-nc',
    name: 'CC BY-NC 4.0',
    shortDesc: '署名-非商业性使用',
    icon: '💰',
    description: '他人可以修改和使用，但不能用于商业用途，且需署名'
  },
  {
    id: 'cc-by-nd',
    name: 'CC BY-ND 4.0',
    shortDesc: '署名-禁止演绎',
    icon: '🚫',
    description: '他人可以转发分享，但不能对原作品进行修改，且需署名'
  },
  {
    id: 'cc-by-sa',
    name: 'CC BY-SA 4.0',
    shortDesc: '署名-相同方式共享',
    icon: '🔄',
    description: '他人修改后的作品，必须采用相同的 CC 协议发布，且需署名'
  }
])

const selectLicense = (id) => {
  emit('update:modelValue', id)
}
</script>
