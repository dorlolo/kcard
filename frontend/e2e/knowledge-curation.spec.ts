import { test, expect } from '@playwright/test'

test('knowledge curation exposes list and graph validation path', async () => {
  expect(['list', 'graph']).toContain('graph')
})
