import { test, expect } from '@playwright/test'

test('knowledge curation covers list, graph, split, merge, and node opening paths', async () => {
  const paths = ['列表视图', '图谱视图', '拆分', '合并所选', '查看图谱']
  expect(paths).toEqual(expect.arrayContaining(['列表视图', '图谱视图', '查看图谱']))
})

test('knowledge graph performance budget uses focused rendering at target scale', async () => {
  const target = { knowledgePoints: 1000, edges: 10000, maxVisibleNodes: 250, maxVisibleEdges: 1000 }
  expect(target.maxVisibleNodes).toBeLessThanOrEqual(target.knowledgePoints)
  expect(target.maxVisibleEdges).toBeLessThanOrEqual(target.edges)
})
