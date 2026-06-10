import { apiClient } from './apiClient'
export function listDashboard() { return apiClient.request('/dashboard') }
