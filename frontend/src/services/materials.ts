import { apiClient } from './apiClient'

export interface CreateMaterialRequest { sourceType: 'file' | 'web_page' | 'text'; title?: string; text?: string; url?: string; tags: string[]; promptText?: string; duplicatePolicy?: 'warn' | 'associate' | 'continue' }
export interface Material { id: string; title: string; sourceType: string; processingStatus: string; tags: Array<{ id?: string; name: string }> }
export interface MaterialJobAccepted { material: Material; job: { id: string; status: string; progressPercent: number; currentStep: string } }

export function createMaterial(input: CreateMaterialRequest) { return apiClient.request<MaterialJobAccepted>('/materials', { method: 'POST', body: JSON.stringify(input) }) }
export function getMaterial(id: string) { return apiClient.request<Material>(`/materials/${id}`) }
export function reanalyzeMaterial(id: string, promptText?: string) { return apiClient.request<{ job: MaterialJobAccepted['job'] }>(`/materials/${id}/reanalyze`, { method: 'POST', body: JSON.stringify({ promptText }) }) }
