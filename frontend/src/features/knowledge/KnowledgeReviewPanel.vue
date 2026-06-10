<script setup lang="ts">
export interface KnowledgePoint { id: string; content: string; summary?: string; approvalStatus: 'draft' | 'approved' | 'rejected' | 'needs_review'; notes?: string }
defineProps<{ points: KnowledgePoint[] }>()
defineEmits<{ approve: [id: string]; reject: [id: string]; edit: [point: KnowledgePoint] }>()
</script>

<template>
  <section class="knowledge-review">
    <article v-for="point in points" :key="point.id" class="surface-card">
      <h3>{{ point.summary || '知识点' }}</h3>
      <p>{{ point.content }}</p>
      <small>状态：{{ point.approvalStatus }}</small>
      <div>
        <button @click="$emit('approve', point.id)">批准</button>
        <button @click="$emit('reject', point.id)">拒绝</button>
        <button @click="$emit('edit', point)">编辑</button>
      </div>
    </article>
  </section>
</template>
