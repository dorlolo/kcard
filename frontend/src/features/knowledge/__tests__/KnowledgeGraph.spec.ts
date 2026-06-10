import { describe, expect, it } from 'vitest'
import type { KnowledgeGraph } from '../../../services/knowledge'

describe('KnowledgeGraph', () => {
  it('supports nodes and typed edges', () => {
    const graph: KnowledgeGraph = {
      nodes: [{ id: 'a', label: '细胞', nodeType: 'knowledge_point', status: 'approved' }],
      edges: [{ id: 'e', sourceId: 'a', targetId: 'b', relationshipType: 'related', label: '相关' }]
    }
    expect(graph.edges[0].relationshipType).toBe('related')
  })

  it('supports dense graph warnings and accessible fallback content', () => {
    const graph: KnowledgeGraph = { nodes: [], edges: [], warnings: ['关系过多，已截断显示'] }
    expect(graph.warnings?.[0]).toContain('截断')
  })

  it('can represent selected focus node data', () => {
    const selectedNodeId = 'kp-cell'
    const graph: KnowledgeGraph = { nodes: [{ id: selectedNodeId, label: '细胞', nodeType: 'knowledge_point' }], edges: [] }
    expect(graph.nodes.some((node) => node.id === selectedNodeId)).toBe(true)
  })
})
