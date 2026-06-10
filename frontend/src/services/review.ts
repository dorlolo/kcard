import { apiClient } from './apiClient'
export function listReview() { return apiClient.request('/review') }
