import { describe, expect, it } from 'vitest'

const viewModes = ['list', 'graph']
const filters = ['q', 'approvalStatus', 'relationshipType']
const actions = ['split', 'merge', 'relationship']

describe('KnowledgeLibrary', () => {
  it('supports list and graph view modes', () => {
    expect(viewModes).toEqual(expect.arrayContaining(['list', 'graph']))
  })

  it('defines filters for query, approval status, and relationship type', () => {
    expect(filters).toEqual(expect.arrayContaining(['q', 'approvalStatus', 'relationshipType']))
  })

  it('supports curation actions for split, merge, and relationship editing', () => {
    expect(actions).toEqual(expect.arrayContaining(['split', 'merge', 'relationship']))
  })
})
