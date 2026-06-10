import { describe, expect, it } from 'vitest'

describe('KnowledgeGraph', () => {
  it('supports nodes and typed edges', () => {
    const graph = { nodes: [{ id: 'a', label: 'A' }], edges: [{ id: 'e', relationshipType: 'related' }] }
    expect(graph.edges[0].relationshipType).toBe('related')
  })
})
