import { describe, it, expect, vi, beforeEach } from 'vitest'
import api from '@/utils/api'

describe('API Utility', () => {
  beforeEach(() => {
    vi.clearAllMocks()
  })

  it('should have correct baseURL configured', () => {
    expect(api.defaults.baseURL).toBe('http://localhost:8080/api/v1')
  })

  it('should have correct timeout configured', () => {
    expect(api.defaults.timeout).toBe(10000)
  })

  it('should export default api instance', () => {
    expect(api).toBeDefined()
    expect(typeof api.get).toBe('function')
    expect(typeof api.post).toBe('function')
    expect(typeof api.put).toBe('function')
    expect(typeof api.delete).toBe('function')
  })

  it('should have request interceptor configured', () => {
    expect(api.interceptors.request.handlers.length).toBeGreaterThan(0)
  })

  it('should have response interceptor configured', () => {
    expect(api.interceptors.response.handlers.length).toBeGreaterThan(0)
  })
})