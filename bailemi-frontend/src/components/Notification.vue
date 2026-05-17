<template>
  <Teleport to="body">
    <TransitionGroup name="notification" tag="div" class="fixed top-4 right-4 z-50 space-y-2 max-w-md">
      <div
        v-for="notification in notifications"
        :key="notification.id"
        :class="[
          'glass-dark rounded-lg p-4 shadow-lg border-l-4',
          notification.type === 'success' ? 'border-green-500' : '',
          notification.type === 'error' ? 'border-red-500' : '',
          notification.type === 'warning' ? 'border-yellow-500' : '',
          notification.type === 'info' ? 'border-blue-500' : ''
        ]"
      >
        <div class="flex items-start gap-3">
          <div class="flex-shrink-0">
            <span v-if="notification.type === 'success'" class="text-2xl">✅</span>
            <span v-else-if="notification.type === 'error'" class="text-2xl">❌</span>
            <span v-else-if="notification.type === 'warning'" class="text-2xl">⚠️</span>
            <span v-else class="text-2xl">ℹ️</span>
          </div>
          <div class="flex-1 min-w-0">
            <h4 class="font-semibold text-sm">{{ notification.title }}</h4>
            <p class="text-sm opacity-80 mt-1">{{ notification.message }}</p>
          </div>
          <button
            @click="removeNotification(notification.id)"
            class="flex-shrink-0 text-gray-400 hover:text-gray-200"
          >
            ✕
          </button>
        </div>
      </div>
    </TransitionGroup>
  </Teleport>
</template>

<script setup>
import { useNotificationStore } from '@/stores/notification'
import { computed } from 'vue'

const notificationStore = useNotificationStore()
const notifications = computed(() => notificationStore.notifications)

const removeNotification = (id) => {
  notificationStore.removeNotification(id)
}
</script>

<style scoped>
.notification-enter-active,
.notification-leave-active {
  transition: all 0.3s ease;
}

.notification-enter-from {
  opacity: 0;
  transform: translateX(100px);
}

.notification-leave-to {
  opacity: 0;
  transform: translateX(100px);
}
</style>
