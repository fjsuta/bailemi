<template>
  <div>
    <h2 class="text-xl font-bold mb-6 text-white flex items-center gap-2">
      <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
        <polyline points="7 10 12 15 17 10"/>
        <line x1="12" y1="15" x2="12" y2="3"/>
      </svg>
      存储配置管理
    </h2>

    <div class="glass rounded-xl p-6 mb-6">
      <div class="flex items-center justify-between mb-6">
        <div>
          <h3 class="font-semibold text-slate-300">存储类型选择</h3>
          <p class="text-sm text-slate-500">选择文件存储方式</p>
        </div>
        <select
          v-model="config.type"
          class="px-4 py-2 bg-slate-800/50 border border-slate-700 rounded-xl text-white"
        >
          <option value="local">本地存储</option>
          <option value="aliyun_oss">阿里云 OSS</option>
          <option value="tencent_cos">腾讯云 COS</option>
          <option value="qiniu">七牛云</option>
          <option value="upyun">又拍云</option>
          <option value="huawei_obs">华为云 OBS</option>
          <option value="baidu_bos">百度云 BOS</option>
        </select>
      </div>

      <div v-if="config.type === 'local'" class="space-y-4">
        <div>
          <label class="block text-sm text-slate-400 mb-1">存储路径</label>
          <input
            v-model="config.local.path"
            type="text"
            placeholder="./uploads"
            class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
          />
        </div>
      </div>

      <div v-else-if="config.type === 'aliyun_oss'" class="space-y-4">
        <div class="flex items-center gap-4 mb-4">
          <label class="flex items-center gap-2">
            <input type="checkbox" v-model="config.aliyun_oss.enabled" class="rounded" />
            <span class="text-slate-300">启用阿里云 OSS</span>
          </label>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-slate-400 mb-1">AccessKey ID</label>
            <input
              v-model="config.aliyun_oss.access_key_id"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">AccessKey Secret</label>
            <input
              v-model="config.aliyun_oss.access_key_secret"
              type="password"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">Bucket</label>
            <input
              v-model="config.aliyun_oss.bucket"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">Region</label>
            <input
              v-model="config.aliyun_oss.region"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm text-slate-400 mb-1">Endpoint</label>
            <input
              v-model="config.aliyun_oss.endpoint"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
        </div>
      </div>

      <div v-else-if="config.type === 'tencent_cos'" class="space-y-4">
        <div class="flex items-center gap-4 mb-4">
          <label class="flex items-center gap-2">
            <input type="checkbox" v-model="config.tencent_cos.enabled" class="rounded" />
            <span class="text-slate-300">启用腾讯云 COS</span>
          </label>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-slate-400 mb-1">SecretId</label>
            <input
              v-model="config.tencent_cos.secret_id"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">SecretKey</label>
            <input
              v-model="config.tencent_cos.secret_key"
              type="password"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">Bucket</label>
            <input
              v-model="config.tencent_cos.bucket"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">Region</label>
            <input
              v-model="config.tencent_cos.region"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
        </div>
      </div>

      <div v-else-if="config.type === 'qiniu'" class="space-y-4">
        <div class="flex items-center gap-4 mb-4">
          <label class="flex items-center gap-2">
            <input type="checkbox" v-model="config.qiniu.enabled" class="rounded" />
            <span class="text-slate-300">启用七牛云</span>
          </label>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-slate-400 mb-1">AccessKey</label>
            <input
              v-model="config.qiniu.access_key"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">SecretKey</label>
            <input
              v-model="config.qiniu.secret_key"
              type="password"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">Bucket</label>
            <input
              v-model="config.qiniu.bucket"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">CDN 域名</label>
            <input
              v-model="config.qiniu.domain"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
        </div>
      </div>

      <div v-else-if="config.type === 'upyun'" class="space-y-4">
        <div class="flex items-center gap-4 mb-4">
          <label class="flex items-center gap-2">
            <input type="checkbox" v-model="config.upyun.enabled" class="rounded" />
            <span class="text-slate-300">启用又拍云</span>
          </label>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-slate-400 mb-1">服务名称</label>
            <input
              v-model="config.upyun.bucket"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">操作员</label>
            <input
              v-model="config.upyun.operator"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">密码</label>
            <input
              v-model="config.upyun.password"
              type="password"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">CDN 域名</label>
            <input
              v-model="config.upyun.domain"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
        </div>
      </div>

      <div v-else-if="config.type === 'huawei_obs'" class="space-y-4">
        <div class="flex items-center gap-4 mb-4">
          <label class="flex items-center gap-2">
            <input type="checkbox" v-model="config.huawei_obs.enabled" class="rounded" />
            <span class="text-slate-300">启用华为云 OBS</span>
          </label>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-slate-400 mb-1">AccessKey</label>
            <input
              v-model="config.huawei_obs.access_key"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">SecretKey</label>
            <input
              v-model="config.huawei_obs.secret_key"
              type="password"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">Bucket</label>
            <input
              v-model="config.huawei_obs.bucket"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">Region</label>
            <input
              v-model="config.huawei_obs.region"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm text-slate-400 mb-1">Endpoint</label>
            <input
              v-model="config.huawei_obs.endpoint"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
        </div>
      </div>

      <div v-else-if="config.type === 'baidu_bos'" class="space-y-4">
        <div class="flex items-center gap-4 mb-4">
          <label class="flex items-center gap-2">
            <input type="checkbox" v-model="config.baidu_bos.enabled" class="rounded" />
            <span class="text-slate-300">启用百度云 BOS</span>
          </label>
        </div>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-slate-400 mb-1">AccessKey</label>
            <input
              v-model="config.baidu_bos.access_key"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">SecretKey</label>
            <input
              v-model="config.baidu_bos.secret_key"
              type="password"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">Bucket</label>
            <input
              v-model="config.baidu_bos.bucket"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">Region</label>
            <input
              v-model="config.baidu_bos.region"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div class="md:col-span-2">
            <label class="block text-sm text-slate-400 mb-1">Endpoint</label>
            <input
              v-model="config.baidu_bos.endpoint"
              type="text"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
        </div>
      </div>
    </div>

    <div class="flex justify-end">
      <button
        @click="saveConfig"
        :disabled="saving"
        class="px-8 py-3 bg-gradient-to-r from-purple-600 to-blue-600 hover:from-purple-500 hover:to-blue-500 rounded-xl font-medium transition-all disabled:opacity-50"
      >
        {{ saving ? '保存中...' : '保存配置' }}
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'

const notification = useNotificationStore()

const config = ref({
  type: 'local',
  local: {
    path: './uploads'
  },
  aliyun_oss: {
    enabled: false,
    access_key_id: '',
    access_key_secret: '',
    bucket: '',
    region: '',
    endpoint: ''
  },
  tencent_cos: {
    enabled: false,
    secret_id: '',
    secret_key: '',
    bucket: '',
    region: ''
  },
  qiniu: {
    enabled: false,
    access_key: '',
    secret_key: '',
    bucket: '',
    domain: ''
  },
  upyun: {
    enabled: false,
    bucket: '',
    operator: '',
    password: '',
    domain: ''
  },
  huawei_obs: {
    enabled: false,
    access_key: '',
    secret_key: '',
    bucket: '',
    region: '',
    endpoint: ''
  },
  baidu_bos: {
    enabled: false,
    access_key: '',
    secret_key: '',
    bucket: '',
    region: '',
    endpoint: ''
  }
})

const saving = ref(false)

const saveConfig = async () => {
  saving.value = true
  try {
    const res = await api.post('/admin/storage', config.value)
    if (res.data.code === 0) {
      notification.success('保存成功', '存储配置已更新')
    }
  } catch (error) {
    notification.error('保存失败', error.message)
  } finally {
    saving.value = false
  }
}

const loadConfig = async () => {
  try {
    const res = await api.get('/admin/storage')
    if (res.data.code === 0) {
      config.value = { ...config.value, ...res.data.data }
    }
  } catch (error) {
    console.error(error)
  }
}

onMounted(() => {
  loadConfig()
})
</script>
