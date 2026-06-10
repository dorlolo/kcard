import { apiClient } from './apiClient'
export function listPrompts() { return apiClient.request('/prompts') }
