<template>
  <div class="glass-dark rounded-2xl p-6">
    <div class="flex items-center justify-between mb-6">
      <h2 class="text-xl font-bold">评论区</h2>
      <span class="text-slate-400">{{ comments.length }} 条评论</span>
    </div>

    <div class="flex gap-4 mb-6">
      <img src="https://picsum.photos/48/48?random=user" class="w-10 h-10 rounded-full object-cover" />
      <div class="flex-1">
        <textarea
          v-model="newComment"
          placeholder="分享你的感受..."
          rows="3"
          class="w-full px-4 py-3 bg-white/10 border border-white/20 rounded-xl focus:border-purple-500 focus:outline-none resize-none"
        ></textarea>
        <div class="flex justify-end mt-2">
          <button @click="submitComment" :disabled="!newComment.trim()" class="px-6 py-2 bg-gradient-to-r from-purple-600 to-blue-600 rounded-xl font-medium disabled:opacity-50">
            发表评论
          </button>
        </div>
      </div>
    </div>

    <div class="space-y-4">
      <div
        v-for="comment in comments"
        :key="comment.id"
        class="glass rounded-xl p-4"
      >
        <div class="flex gap-4">
          <img :src="comment.avatar" class="w-10 h-10 rounded-full object-cover flex-shrink-0" />
          <div class="flex-1 min-w-0">
            <div class="flex items-center gap-3 mb-2">
              <span class="font-medium">{{ comment.username }}</span>
              <span class="text-xs text-slate-500">{{ comment.time }}</span>
            </div>
            <p class="text-slate-300 mb-3">{{ comment.content }}</p>
            <div class="flex items-center gap-6">
              <button @click="toggleLike(comment)" :class="['flex items-center gap-1 transition-colors', comment.liked ? 'text-red-500' : 'text-slate-400 hover:text-red-400']">
                <span>{{ comment.liked ? '❤️' : '🤍' }}</span>
                <span>{{ comment.likes }}</span>
              </button>
              <button @click="toggleReply(comment)" class="flex items-center gap-1 text-slate-400 hover:text-purple-400 transition-colors">
                <span>💬</span>
                <span>回复</span>
              </button>
            </div>

            <div v-if="comment.showReply" class="mt-4 pl-4 border-l-2 border-slate-700">
              <div class="flex gap-3">
                <img src="https://picsum.photos/32/32?random=me" class="w-8 h-8 rounded-full object-cover" />
                <div class="flex-1">
                  <input
                    v-model="comment.replyText"
                    placeholder="写下你的回复..."
                    class="w-full px-3 py-2 bg-white/10 rounded-lg focus:outline-none focus:border-purple-500"
                  />
                  <div class="flex justify-end mt-2">
                    <button @click="submitReply(comment)" :disabled="!comment.replyText.trim()" class="px-4 py-1 bg-purple-600/80 rounded-lg text-sm disabled:opacity-50">
                      回复
                    </button>
                  </div>
                </div>
              </div>

              <div v-if="comment.replies.length > 0" class="mt-3 space-y-2">
                <div
                  v-for="reply in comment.replies"
                  :key="reply.id"
                  class="flex gap-2"
                >
                  <img :src="reply.avatar" class="w-6 h-6 rounded-full object-cover" />
                  <div class="bg-white/5 rounded-lg px-3 py-2">
                    <span class="font-medium text-purple-400">{{ reply.username }}</span>
                    <span class="text-slate-300"> {{ reply.content }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const newComment = ref('')

const comments = ref([])

const submitComment = () => {
  if (!newComment.value.trim()) return
  
  comments.value.unshift({
    id: Date.now(),
    username: '我',
    avatar: 'https://picsum.photos/48/48?random=me',
    time: '刚刚',
    content: newComment.value,
    likes: 0,
    liked: false,
    showReply: false,
    replyText: '',
    replies: []
  })
  
  newComment.value = ''
}

const toggleLike = (comment) => {
  comment.liked = !comment.liked
  comment.likes += comment.liked ? 1 : -1
}

const toggleReply = (comment) => {
  comment.showReply = !comment.showReply
}

const submitReply = (comment) => {
  if (!comment.replyText.trim()) return
  
  comment.replies.push({
    id: Date.now(),
    username: '我',
    avatar: 'https://picsum.photos/32/32?random=me',
    content: comment.replyText
  })
  
  comment.replyText = ''
}
</script>
