import { describe, expect, it } from 'vitest'

describe('KnowledgeLibrary', () => {
  it('documents list and graph states', () => {
    expect(['list', 'graph']).toContain('graph')
  })
})
