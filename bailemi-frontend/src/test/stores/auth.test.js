import { setActivePinia, createPinia } from 'pinia'
import { describe, it, expect, beforeEach } from 'vitest'
import { useAuthStore } from '@/stores/auth'

describe('useAuthStore', () => {
  beforeEach(() => {
    localStorage.clear()
    setActivePinia(createPinia())
  })

  it('should initialize with empty user and tokens', () => {
    const authStore = useAuthStore()
    
    expect(authStore.user).toBe(null)
    expect(authStore.token).toBe('')
    expect(authStore.refreshToken).toBe('')
    expect(authStore.isAuthenticated).toBe(false)
  })

  it('should set user correctly', () => {
    const authStore = useAuthStore()
    const userData = { id: 1, username: 'testuser', email: 'test@example.com' }
    
    authStore.setUser(userData)
    
    expect(authStore.user).toEqual(userData)
  })

  it('should set tokens and save to localStorage', () => {
    const authStore = useAuthStore()
    
    authStore.setTokens('access-token-123', 'refresh-token-456')
    
    expect(authStore.token).toBe('access-token-123')
    expect(authStore.refreshToken).toBe('refresh-token-456')
    expect(localStorage.getItem('access_token')).toBe('access-token-123')
    expect(localStorage.getItem('refresh_token')).toBe('refresh-token-456')
  })

  it('should clear localStorage when setting empty tokens', () => {
    const authStore = useAuthStore()
    
    authStore.setTokens('', '')
    
    expect(localStorage.getItem('access_token')).toBe(null)
    expect(localStorage.getItem('refresh_token')).toBe(null)
  })

  it('should logout and clear all data', () => {
    const authStore = useAuthStore()
    
    authStore.setTokens('access-token-123', 'refresh-token-456')
    authStore.setUser({ id: 1, username: 'testuser' })
    authStore.logout()
    
    expect(authStore.user).toBe(null)
    expect(authStore.token).toBe('')
    expect(authStore.refreshToken).toBe('')
    expect(localStorage.getItem('access_token')).toBe(null)
    expect(localStorage.getItem('refresh_token')).toBe(null)
    expect(authStore.isAuthenticated).toBe(false)
  })

  it('should be authenticated when has token and user', () => {
    const authStore = useAuthStore()
    
    authStore.setTokens('access-token-123', 'refresh-token-456')
    authStore.setUser({ id: 1, username: 'testuser' })
    
    expect(authStore.isAuthenticated).toBe(true)
  })

  it('should not be authenticated when only has token', () => {
    const authStore = useAuthStore()
    
    authStore.setTokens('access-token-123', 'refresh-token-456')
    
    expect(authStore.isAuthenticated).toBe(false)
  })

  it('should not be authenticated when only has user', () => {
    const authStore = useAuthStore()
    
    authStore.setUser({ id: 1, username: 'testuser' })
    
    expect(authStore.isAuthenticated).toBe(false)
  })
})