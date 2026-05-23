import { setActivePinia, createPinia } from 'pinia'
import { describe, it, expect, beforeEach, vi } from 'vitest'
import { useNotificationStore } from '@/stores/notification'

describe('useNotificationStore', () => {
  beforeEach(() => {
    localStorage.clear()
    vi.useFakeTimers()
    setActivePinia(createPinia())
  })

  afterEach(() => {
    vi.useRealTimers()
  })

  it('should initialize with empty notifications', () => {
    const notificationStore = useNotificationStore()
    
    expect(notificationStore.notifications).toEqual([])
    expect(notificationStore.notificationCount).toBe(0)
  })

  it('should add notification with default values', () => {
    const notificationStore = useNotificationStore()
    
    const id = notificationStore.addNotification({
      title: 'Test Title',
      message: 'Test Message'
    })
    
    expect(notificationStore.notifications.length).toBe(1)
    expect(notificationStore.notificationCount).toBe(1)
    expect(notificationStore.notifications[0].title).toBe('Test Title')
    expect(notificationStore.notifications[0].message).toBe('Test Message')
    expect(notificationStore.notifications[0].type).toBe('info')
    expect(notificationStore.notifications[0].duration).toBe(5000)
    expect(notificationStore.notifications[0].read).toBe(false)
  })

  it('should add notification with custom type', () => {
    const notificationStore = useNotificationStore()
    
    notificationStore.addNotification({
      type: 'success',
      title: 'Success',
      message: 'Operation successful'
    })
    
    expect(notificationStore.notifications[0].type).toBe('success')
  })

  it('should add notification with custom duration', () => {
    const notificationStore = useNotificationStore()
    
    notificationStore.addNotification({
      title: 'Test',
      message: 'Test message',
      duration: 10000
    })
    
    expect(notificationStore.notifications[0].duration).toBe(10000)
  })

  it('should remove notification by id', () => {
    const notificationStore = useNotificationStore()
    
    const id = notificationStore.addNotification({
      title: 'Test',
      message: 'Test message'
    })
    
    expect(notificationStore.notifications.length).toBe(1)
    
    notificationStore.removeNotification(id)
    
    expect(notificationStore.notifications.length).toBe(0)
  })

  it('should mark notification as read', () => {
    const notificationStore = useNotificationStore()
    
    const id = notificationStore.addNotification({
      title: 'Test',
      message: 'Test message'
    })
    
    expect(notificationStore.notificationCount).toBe(1)
    
    notificationStore.markAsRead(id)
    
    expect(notificationStore.notifications[0].read).toBe(true)
    expect(notificationStore.notificationCount).toBe(0)
  })

  it('should mark all notifications as read', () => {
    const notificationStore = useNotificationStore()
    
    notificationStore.addNotification({ title: 'Test 1', message: 'Message 1' })
    notificationStore.addNotification({ title: 'Test 2', message: 'Message 2' })
    
    expect(notificationStore.notificationCount).toBe(2)
    
    notificationStore.markAllAsRead()
    
    expect(notificationStore.notificationCount).toBe(0)
    expect(notificationStore.notifications.every(n => n.read)).toBe(true)
  })

  it('should clear all notifications', () => {
    const notificationStore = useNotificationStore()
    
    notificationStore.addNotification({ title: 'Test 1', message: 'Message 1' })
    notificationStore.addNotification({ title: 'Test 2', message: 'Message 2' })
    
    expect(notificationStore.notifications.length).toBe(2)
    
    notificationStore.clearAll()
    
    expect(notificationStore.notifications.length).toBe(0)
    expect(notificationStore.notificationCount).toBe(0)
  })

  it('should use shortcut methods correctly', () => {
    const notificationStore = useNotificationStore()
    
    notificationStore.success('Success Title', 'Success Message')
    expect(notificationStore.notifications[0].type).toBe('success')
    
    notificationStore.error('Error Title', 'Error Message')
    expect(notificationStore.notifications[0].type).toBe('error')
    
    notificationStore.warning('Warning Title', 'Warning Message')
    expect(notificationStore.notifications[0].type).toBe('warning')
    
    notificationStore.info('Info Title', 'Info Message')
    expect(notificationStore.notifications[0].type).toBe('info')
  })

  it('should auto-remove notification after duration', () => {
    const notificationStore = useNotificationStore()
    
    notificationStore.addNotification({
      title: 'Test',
      message: 'Test message',
      duration: 5000
    })
    
    expect(notificationStore.notifications.length).toBe(1)
    
    vi.advanceTimersByTime(6000)
    
    expect(notificationStore.notifications.length).toBe(0)
  })

  it('should not auto-remove notification with duration 0', () => {
    const notificationStore = useNotificationStore()
    
    notificationStore.addNotification({
      title: 'Test',
      message: 'Test message',
      duration: 0
    })
    
    expect(notificationStore.notifications.length).toBe(1)
    
    vi.advanceTimersByTime(10000)
    
    expect(notificationStore.notifications.length).toBe(1)
  })
})