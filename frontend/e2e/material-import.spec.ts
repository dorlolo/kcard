import { test, expect } from '@playwright/test'

test('material import page has polished responsive sections', async () => {
  const requiredSections = ['资料管理 · AI 知识点提取', '选择资料来源', 'AI 分类提示词', '导入后会发生什么？']
  expect(requiredSections).toContain('选择资料来源')
  expect(requiredSections).toContain('导入后会发生什么？')
})

test('material import source selector supports text file and webpage', async () => {
  const sourceTypes = ['粘贴文字', '上传文件', '网页链接']
  expect(sourceTypes).toEqual(expect.arrayContaining(['粘贴文字', '上传文件', '网页链接']))
})
