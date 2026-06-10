import { apiClient } from './apiClient'

export interface Tag { id?: string; name: string; color?: string }
export interface KnowledgePoint { id: string; content: string; summary?: string; notes?: string; approvalStatus: 'draft' | 'approved' | 'rejected' | 'needs_review'; creationSource?: string; tags?: Tag[]; duplicateGroupId?: string }
export interface KnowledgeGraphNode { id: string; label: string; nodeType: 'knowledge_point' | 'source_material' | 'tag' | 'card'; status?: string; weight?: number }
export interface KnowledgeGraphEdge { id: string; sourceId: string; targetId: string; relationshipType: string; label?: string; weight?: number; sourceType?: string }
export interface KnowledgeGraph { nodes: KnowledgeGraphNode[]; edges: KnowledgeGraphEdge[]; warnings?: string[] }

export interface KnowledgeListParams { q?: string; approvalStatus?: string; tag?: string; includeRejected?: boolean }
export interface KnowledgeGraphParams extends KnowledgeListParams { focusKnowledgePointId?: string; depth?: number; relationshipType?: string[]; includeArchived?: boolean }

export function listKnowledgePoints(params: KnowledgeListParams = {}) {
  return apiClient.request<{ items: KnowledgePoint[]; meta: { total: number } }>(`/knowledge-points?${toSearchParams(params)}`)
}

export function updateKnowledgePoint(id: string, input: Partial<Pick<KnowledgePoint, 'content' | 'summary' | 'notes' | 'approvalStatus'>>) {
  return apiClient.request<KnowledgePoint>(`/knowledge-points/${id}`, { method: 'PATCH', body: JSON.stringify(input) })
}

export function splitKnowledgePoint(id: string, items: Array<{ content: string; summary?: string; tags?: string[] }>) {
  return apiClient.request<{ items: KnowledgePoint[] }>(`/knowledge-points/${id}/split`, { method: 'POST', body: JSON.stringify({ items }) })
}

export function mergeKnowledgePoints(knowledgePointIds: string[], content: string) {
  return apiClient.request<KnowledgePoint>('/knowledge-points/merge', { method: 'POST', body: JSON.stringify({ knowledgePointIds, content }) })
}

export function getKnowledgeGraph(params: KnowledgeGraphParams = {}) {
  return apiClient.request<KnowledgeGraph>(`/knowledge-graph?${toSearchParams(params)}`)
}

export function createKnowledgeRelationship(input: { sourceKnowledgePointId: string; targetKnowledgePointId: string; relationshipType: string; label?: string; weight?: number }) {
  return apiClient.request<KnowledgeGraphEdge>('/knowledge-relationships', { method: 'POST', body: JSON.stringify(input) })
}

export function archiveKnowledgeRelationship(id: string) {
  return apiClient.request<KnowledgeGraphEdge>(`/knowledge-relationships/${id}`, { method: 'PATCH', body: JSON.stringify({ archived: true }) })
}

function toSearchParams(params: object) {
  const search = new URLSearchParams()
  Object.entries(params).forEach(([key, value]) => {
    if (value === undefined || value === '' || value === false) return
    if (Array.isArray(value)) value.forEach((item) => search.append(key, String(item)))
    else search.set(key, String(value))
  })
  return search
}
