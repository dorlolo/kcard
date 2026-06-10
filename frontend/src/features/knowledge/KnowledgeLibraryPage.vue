<script setup lang="ts">
import { computed, ref } from 'vue'
import KnowledgeGraphView from './KnowledgeGraphView.vue'
import type { KnowledgeGraph, KnowledgePoint } from '../../services/knowledge'

const view = ref<'list' | 'graph'>('list')
const query = ref('')
const points = ref<KnowledgePoint[]>([])
const graph = ref<KnowledgeGraph>({ nodes: [], edges: [] })
const filtered = computed(() => points.value.filter((point) => point.content.toLowerCase().includes(query.value.toLowerCase())))
</script>

<template>
  <section class="knowledge-library surface-card">
    <header>
      <h2>知识库</h2>
      <input v-model="query" placeholder="搜索知识点" aria-label="搜索知识点" />
      <button :aria-pressed="view === 'list'" @click="view = 'list'">列表视图</button>
      <button :aria-pressed="view === 'graph'" @click="view = 'graph'">图谱视图</button>
    </header>
    <ul v-if="view === 'list'">
      <li v-for="point in filtered" :key="point.id">{{ point.content }} <small>{{ point.approvalStatus }}</small></li>
    </ul>
    <KnowledgeGraphView v-else :graph="graph" />
  </section>
</template>
