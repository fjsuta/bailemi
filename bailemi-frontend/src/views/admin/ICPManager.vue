<template>
  <div>
    <h2 class="text-xl font-bold mb-6 text-white flex items-center gap-2">
      <svg class="w-5 h-5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
        <polyline points="14 2 14 8 20 8"/>
      </svg>
      公安备案管理
    </h2>

    <div class="max-w-3xl">
      <div class="glass rounded-xl p-6 mb-6">
        <h3 class="font-semibold text-slate-300 mb-4">ICP备案信息</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-slate-400 mb-1">ICP备案号</label>
            <input
              v-model="icpInfo.icp_number"
              type="text"
              placeholder="如：京ICP备xxxxxxxx号"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">公安备案号</label>
            <input
              v-model="icpInfo.psb_record"
              type="text"
              placeholder="如：京公网安备xxxxxxxxxx号"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">网站域名</label>
            <input
              v-model="icpInfo.domain"
              type="text"
              placeholder="yourdomain.com"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">公司名称</label>
            <input
              v-model="icpInfo.company_name"
              type="text"
              placeholder="公司全称"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">法人姓名</label>
            <input
              v-model="icpInfo.legal_person"
              type="text"
              placeholder="法人姓名"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">联系人</label>
            <input
              v-model="icpInfo.contact"
              type="text"
              placeholder="联系人姓名"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">联系电话</label>
            <input
              v-model="icpInfo.contact_phone"
              type="text"
              placeholder="联系电话"
              class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
            />
          </div>
        </div>
        <div class="mt-4">
          <label class="block text-sm text-slate-400 mb-1">联系地址</label>
          <textarea
            v-model="icpInfo.address"
            placeholder="详细地址"
            rows="2"
            class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
          ></textarea>
        </div>
      </div>

      <div class="glass rounded-xl p-6 mb-6">
        <h3 class="font-semibold text-slate-300 mb-4">备案证书图片</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label class="block text-sm text-slate-400 mb-1">ICP备案证书</label>
            <div class="border-2 border-dashed border-slate-700 rounded-xl p-6 text-center hover:border-purple-500 transition-all">
              <input type="file" accept="image/*" class="hidden" ref="icpImageInput" @change="uploadICPImage" />
              <button @click="$refs.icpImageInput.click()" class="text-slate-400 hover:text-white">
                {{ icpInfo.icp_image_url ? '重新上传' : '点击上传图片' }}
              </button>
              <img v-if="icpInfo.icp_image_url" :src="icpInfo.icp_image_url" class="mt-4 max-h-32 rounded mx-auto" />
            </div>
          </div>
          <div>
            <label class="block text-sm text-slate-400 mb-1">公安备案证书</label>
            <div class="border-2 border-dashed border-slate-700 rounded-xl p-6 text-center hover:border-purple-500 transition-all">
              <input type="file" accept="image/*" class="hidden" ref="psbImageInput" @change="uploadPSBImage" />
              <button @click="$refs.psbImageInput.click()" class="text-slate-400 hover:text-white">
                {{ icpInfo.psb_image_url ? '重新上传' : '点击上传图片' }}
              </button>
              <img v-if="icpInfo.psb_image_url" :src="icpInfo.psb_image_url" class="mt-4 max-h-32 rounded mx-auto" />
            </div>
          </div>
        </div>
      </div>

      <div class="glass rounded-xl p-6 mb-6">
        <h3 class="font-semibold text-slate-300 mb-4">其他设置</h3>
        <div class="flex items-center justify-between mb-4">
          <span class="text-slate-300">备案状态</span>
          <select
            v-model="icpInfo.status"
            class="px-4 py-2 bg-slate-800/50 border border-slate-700 rounded-xl text-white"
          >
            <option value="pending">待备案</option>
            <option value="processing">备案中</option>
            <option value="completed">已备案</option>
          </select>
        </div>
        <div>
          <label class="block text-sm text-slate-400 mb-1">备注信息</label>
          <textarea
            v-model="icpInfo.remark"
            placeholder="其他备注信息"
            rows="2"
            class="w-full px-4 py-3 bg-slate-800/50 border border-slate-700 rounded-xl focus:border-purple-500 focus:outline-none text-white"
          ></textarea>
        </div>
      </div>

      <div class="flex justify-end">
        <button
          @click="saveInfo"
          :disabled="saving"
          class="px-8 py-3 bg-gradient-to-r from-purple-600 to-blue-600 hover:from-purple-500 hover:to-blue-500 rounded-xl font-medium transition-all disabled:opacity-50"
        >
          {{ saving ? '保存中...' : '保存配置' }}
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useNotificationStore } from '@/stores/notification'
import api from '@/utils/api'

const notification = useNotificationStore()

const icpInfo = ref({
  icp_number: '',
  psb_record: '',
  domain: '',
  company_name: '',
  legal_person: '',
  contact: '',
  contact_phone: '',
  address: '',
  icp_image_url: '',
  psb_image_url: '',
  status: 'pending',
  remark: ''
})

const saving = ref(false)

const saveInfo = async () => {
  saving.value = true
  try {
    const res = await api.post('/admin/icp', icpInfo.value)
    if (res.data.code === 0) {
      notification.success('保存成功', '备案信息已保存')
    }
  } catch (error) {
    notification.error('保存失败', error.message)
  } finally {
    saving.value = false
  }
}

const uploadICPImage = (e) => {
  const file = e.target.files[0]
  if (file) {
    // 模拟上传
    icpInfo.value.icp_image_url = URL.createObjectURL(file)
    notification.success('上传成功', 'ICP证书图片已更新')
  }
}

const uploadPSBImage = (e) => {
  const file = e.target.files[0]
  if (file) {
    // 模拟上传
    icpInfo.value.psb_image_url = URL.createObjectURL(file)
    notification.success('上传成功', '公安备案证书图片已更新')
  }
}

const loadInfo = async () => {
  try {
    const res = await api.get('/admin/icp')
    if (res.data.code === 0) {
      icpInfo.value = { ...icpInfo.value, ...res.data.data }
    }
  } catch (error) {
    console.error(error)
  }
}

onMounted(() => {
  loadInfo()
})
</script>
