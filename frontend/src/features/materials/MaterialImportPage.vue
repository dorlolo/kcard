<script setup lang="ts">
import { computed, ref } from 'vue'
import { createMaterial } from '../../services/materials'
import StateViews from '../../components/state/StateViews.vue'

type SourceType = 'text' | 'file' | 'web_page'

const sourceOptions: Array<{ type: SourceType; title: string; description: string; accent: string }> = [
  { type: 'text', title: '粘贴文字', description: '适合课堂笔记、文章片段和临时整理内容。', accent: 'surface-accent-two' },
  { type: 'file', title: '上传文件', description: '为 PDF、文档和讲义预留入口，后续会进入异步解析。', accent: 'surface-accent-one' },
  { type: 'web_page', title: '网页链接', description: '保存网页地址并准备后续抓取与重新分析。', accent: 'surface-card' }
]

const title = ref('')
const text = ref('')
const url = ref('')
const tags = ref('')
const promptText = ref('请提取适合制作复习卡片的原子化知识点，保留来源语境，并标记可能重复或需要人工确认的内容。')
const sourceType = ref<SourceType>('text')
const status = ref<'empty' | 'loading' | 'success' | 'error'>('empty')
const message = ref('选择资料来源，补充标签和提示词后开始分析。')

const tagChips = computed(() => tags.value.split(/[，,]/).map((tag) => tag.trim()).filter(Boolean))
const canSubmit = computed(() => {
  if (status.value === 'loading') return false
  if (sourceType.value === 'text') return text.value.trim().length > 0
  if (sourceType.value === 'web_page') return url.value.trim().length > 0
  return title.value.trim().length > 0
})

async function submit() {
  status.value = 'loading'
  try {
    const result = await createMaterial({
      sourceType: sourceType.value,
      title: title.value,
      text: sourceType.value === 'text' ? text.value : undefined,
      url: sourceType.value === 'web_page' ? url.value : undefined,
      tags: tagChips.value,
      promptText: promptText.value,
      duplicatePolicy: 'warn'
    })
    status.value = 'success'
    message.value = `已提交分析任务 ${result.job.id}，可以稍后回到知识点审阅。`
  } catch (error) {
    status.value = 'error'
    message.value = error instanceof Error ? error.message : '无法导入资料，请检查内容后重试。'
  }
}
</script>

<template>
  <section class="material-import-page" aria-labelledby="material-import-title">
    <div class="material-import-hero surface-gradient-study">
      <p class="eyebrow">资料管理 · AI 知识点提取</p>
      <div>
        <h2 id="material-import-title">导入资料</h2>
        <p>把文件、网页或文字整理为可审阅的知识点草稿。AI 输出不会直接进入卡片，必须先由你确认。</p>
      </div>
    </div>

    <div class="material-import-grid">
      <form class="surface-card material-import-form" @submit.prevent="submit">
        <fieldset class="source-type-grid">
          <legend>选择资料来源</legend>
          <button
            v-for="option in sourceOptions"
            :key="option.type"
            type="button"
            class="source-type-card"
            :class="[{ 'is-selected': sourceType === option.type }, option.accent]"
            :aria-pressed="sourceType === option.type"
            @click="sourceType = option.type"
          >
            <strong>{{ option.title }}</strong>
            <span>{{ option.description }}</span>
          </button>
        </fieldset>

        <label class="form-field">
          <span>标题</span>
          <input v-model="title" placeholder="例如：细胞结构复习资料" />
        </label>

        <label v-if="sourceType === 'web_page'" class="form-field">
          <span>网页地址</span>
          <input v-model="url" type="url" placeholder="https://example.com/article" />
        </label>

        <label v-if="sourceType === 'text'" class="form-field">
          <span>资料正文</span>
          <textarea v-model="text" rows="10" placeholder="粘贴课堂笔记、教材摘录或网页正文……" />
        </label>

        <div v-if="sourceType === 'file'" class="upload-dropzone surface-accent-two" role="note">
          <strong>文件上传入口</strong>
          <p>当前先记录标题和标签，后续会接入文件 token、进度和失败重试。</p>
        </div>

        <label class="form-field">
          <span>标签</span>
          <input v-model="tags" placeholder="生物，考试，错题" />
        </label>
        <div class="tag-chip-row" aria-label="已输入标签">
          <span v-for="tag in tagChips" :key="tag" class="tag-chip">{{ tag }}</span>
          <span v-if="tagChips.length === 0" class="muted-text">还没有标签</span>
        </div>

        <label class="form-field prompt-field">
          <span>AI 分类提示词</span>
          <textarea v-model="promptText" rows="4" />
        </label>

        <button class="primary-action" :disabled="!canSubmit" type="submit">分析资料</button>
      </form>

      <aside class="material-import-side" aria-label="导入流程说明">
        <section class="surface-card checklist-card">
          <h3>导入后会发生什么？</h3>
          <ol>
            <li>系统检查是否可能重复。</li>
            <li>AI 根据提示词提取知识点草稿。</li>
            <li>你在知识点审阅区编辑、批准或拒绝。</li>
            <li>只有已批准的知识点会进入制卡流程。</li>
          </ol>
        </section>

        <section class="surface-accent-two preview-card">
          <h3>状态预览</h3>
          <StateViews :state="status" title="资料导入" :message="message" />
          <RouterLink v-if="status === 'success'" class="primary-action knowledge-link" to="/knowledge">前往知识库审阅</RouterLink>
        </section>
      </aside>
    </div>
  </section>
</template>
