<template>
  <div class="space-y-4">
    <label class="block text-sm font-medium text-slate-300 mb-2 flex items-center gap-2">
      <svg class="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
        <polyline points="14 2 14 8 20 8"/>
      </svg>
      版权协议选择
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
          <svg class="w-6 h-6 text-purple-400" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M9 12l2 2 4-4"/>
            <circle cx="12" cy="12" r="10"/>
          </svg>
          <div class="flex-1">
            <h4 class="font-semibold text-white">{{ license.name }}</h4>
            <p class="text-xs text-slate-400">{{ license.shortDesc }}</p>
          </div>
          <div
            v-if="modelValue === license.id"
            class="w-6 h-6 rounded-full bg-purple-500 flex items-center justify-center"
          >
            <svg class="w-4 h-4 text-white" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="3">
              <polyline points="20 6 9 17 4 12"/>
            </svg>
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
    description: '作者放弃所有版权，任何人都可以随意使用、修改和商用，无需署名'
  },
  {
    id: 'cc-by',
    name: 'CC BY 4.0',
    shortDesc: '署名许可',
    description: '他人可以自由使用、修改和商用，但必须注明原作者'
  },
  {
    id: 'cc-by-nc',
    name: 'CC BY-NC 4.0',
    shortDesc: '署名-非商业性使用',
    description: '他人可以修改和使用，但不能用于商业用途，且需署名'
  },
  {
    id: 'cc-by-nd',
    name: 'CC BY-ND 4.0',
    shortDesc: '署名-禁止演绎',
    description: '他人可以转发分享，但不能对原作品进行修改，且需署名'
  },
  {
    id: 'cc-by-sa',
    name: 'CC BY-SA 4.0',
    shortDesc: '署名-相同方式共享',
    description: '他人修改后的作品，必须采用相同的 CC 协议发布，且需署名'
  }
])

const selectLicense = (id) => {
  emit('update:modelValue', id)
}
</script>
