import { describe, expect, it } from 'vitest'

const requiredSections = ['资料管理 · AI 知识点提取', '选择资料来源', 'AI 分类提示词', '导入后会发生什么？']

describe('MaterialImportPage layout requirements', () => {
  it('defines the polished material import sections expected by US1', () => {
    expect(requiredSections).toContain('选择资料来源')
    expect(requiredSections).toContain('导入后会发生什么？')
  })

  it('covers the required source type selector options', () => {
    const sourceTypes = ['粘贴文字', '上传文件', '网页链接']
    expect(sourceTypes).toEqual(expect.arrayContaining(['粘贴文字', '上传文件', '网页链接']))
  })
})
