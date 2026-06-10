export type MaterialState = 'draft' | 'queued' | 'processing' | 'needs_review' | 'processed' | 'failed'

export function learnerMessageForMaterialState(state: MaterialState, failureReason?: string): string {
  switch (state) {
    case 'queued': return '资料已进入 AI 分析队列。'
    case 'processing': return 'AI 正在提取知识点草稿，你可以稍后回来查看。'
    case 'needs_review': return '请审阅、编辑、批准或拒绝提取出的知识点。'
    case 'processed': return '已批准的知识点已保存到知识库。'
    case 'failed': return failureReason || '分析失败。你可以重试、替换来源，或手动粘贴文本。'
    default: return '添加资料后开始。'
  }
}
