import { apiClient } from './apiClient'

export interface KnowledgePoint { id: string; content: string; summary?: string; approvalStatus: string; tags?: Array<{ id: string; name: string }> }
export interface KnowledgeGraphNode { id: string; label: string; nodeType: 'knowledge_point' | 'source_material' | 'tag' | 'card'; status?: string }
export interface KnowledgeGraphEdge { id: string; sourceId: string; targetId: string; relationshipType: string; label?: string; weight?: number }
export interface KnowledgeGraph { nodes: KnowledgeGraphNode[]; edges: KnowledgeGraphEdge[]; warnings?: string[] }

export function listKnowledgePoints(params = new URLSearchParams()) { return apiClient.request<{ items: KnowledgePoint[] }>(`/knowledge-points?${params}`) }
export function getKnowledgeGraph(params = new URLSearchParams()) { return apiClient.request<KnowledgeGraph>(`/knowledge-graph?${params}`) }
export function createKnowledgeRelationship(input: { sourceKnowledgePointId: string; targetKnowledgePointId: string; relationshipType: string; label?: string }) { return apiClient.request<KnowledgeGraphEdge>('/knowledge-relationships', { method: 'POST', body: JSON.stringify(input) }) }
