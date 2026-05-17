import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useNotificationStore = defineStore('notification', () => {
  const notifications = ref([])
  const notificationCount = ref(0)

  const addNotification = (notification) => {
    const id = Date.now() + Math.random()
    const newNotification = {
      id,
      type: notification.type || 'info', // success, error, warning, info
      title: notification.title || '',
      message: notification.message || '',
      duration: notification.duration || 5000, // 默认5秒
      read: false,
      timestamp: new Date().toISOString()
    }

    notifications.value.unshift(newNotification)
    notificationCount.value++

    // 自动移除
    if (newNotification.duration > 0) {
      setTimeout(() => {
        removeNotification(id)
      }, newNotification.duration)
    }

    return id
  }

  const removeNotification = (id) => {
    const index = notifications.value.findIndex(n => n.id === id)
    if (index > -1) {
      notifications.value.splice(index, 1)
    }
  }

  const markAsRead = (id) => {
    const notification = notifications.value.find(n => n.id === id)
    if (notification) {
      notification.read = true
      notificationCount.value = Math.max(0, notificationCount.value - 1)
    }
  }

  const markAllAsRead = () => {
    notifications.value.forEach(n => n.read = true)
    notificationCount.value = 0
  }

  const clearAll = () => {
    notifications.value = []
    notificationCount.value = 0
  }

  // 快捷方法
  const success = (title, message) => addNotification({ type: 'success', title, message })
  const error = (title, message) => addNotification({ type: 'error', title, message })
  const warning = (title, message) => addNotification({ type: 'warning', title, message })
  const info = (title, message) => addNotification({ type: 'info', title, message })

  return {
    notifications,
    notificationCount,
    addNotification,
    removeNotification,
    markAsRead,
    markAllAsRead,
    clearAll,
    success,
    error,
    warning,
    info
  }
})
