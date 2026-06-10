<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import KnowledgeCurationDialogs from './KnowledgeCurationDialogs.vue'
import KnowledgeGraphView from './KnowledgeGraphView.vue'
import KnowledgeRelationshipEditor from './KnowledgeRelationshipEditor.vue'
import {
  createKnowledgeRelationship,
  getKnowledgeGraph,
  listKnowledgePoints,
  mergeKnowledgePoints,
  splitKnowledgePoint,
  type KnowledgeGraph,
  type KnowledgePoint
} from '../../services/knowledge'

const view = ref<'list' | 'graph'>('list')
const query = ref('')
const approvalStatus = ref('')
const relationshipType = ref('')
const points = ref<KnowledgePoint[]>([])
const graph = ref<KnowledgeGraph>({ nodes: [], edges: [] })
const selectedIds = ref<string[]>([])
const selectedNodeId = ref('')
const targetNodeId = ref('')
const dialogMode = ref<'split' | 'merge' | null>(null)
const loading = ref(false)
const errorMessage = ref('')

const filtered = computed(() => points.value)
const selectedPoint = computed(() => points.value.find((point) => point.id === selectedNodeId.value || point.id === selectedIds.value[0]))

async function loadKnowledge() {
  loading.value = true
  errorMessage.value = ''
  try {
    const params = { q: query.value, approvalStatus: approvalStatus.value, includeRejected: true }
    const [list, graphResult] = await Promise.all([
      listKnowledgePoints(params),
      getKnowledgeGraph({ ...params, focusKnowledgePointId: selectedNodeId.value, relationshipType: relationshipType.value ? [relationshipType.value] : undefined, depth: 2, includeRejected: true })
    ])
    points.value = list.items
    graph.value = graphResult
  } catch (error) {
    errorMessage.value = error instanceof Error ? error.message : '知识库加载失败'
  } finally {
    loading.value = false
  }
}

function toggleSelection(id: string) {
  selectedIds.value = selectedIds.value.includes(id) ? selectedIds.value.filter((item) => item !== id) : [...selectedIds.value, id]
}

async function confirmDialog(payload: string[]) {
  if (dialogMode.value === 'split' && selectedPoint.value) await splitKnowledgePoint(selectedPoint.value.id, payload.map((content) => ({ content })))
  if (dialogMode.value === 'merge' && selectedIds.value.length >= 2) await mergeKnowledgePoints(selectedIds.value, payload[0])
  dialogMode.value = null
  selectedIds.value = []
  await loadKnowledge()
}

async function saveRelationship(input: { relationshipType: string; label: string }) {
  if (!selectedNodeId.value || !targetNodeId.value) return
  await createKnowledgeRelationship({ sourceKnowledgePointId: selectedNodeId.value, targetKnowledgePointId: targetNodeId.value, ...input })
  targetNodeId.value = ''
  await loadKnowledge()
}

function selectGraphNode(id: string) {
  if (!selectedNodeId.value || selectedNodeId.value === id) selectedNodeId.value = id
  else targetNodeId.value = id
  view.value = 'graph'
  void loadKnowledge()
}

onMounted(loadKnowledge)
</script>

<template>
  <section class="knowledge-library surface-card">
    <header class="knowledge-library-header">
      <div>
        <p class="eyebrow">知识整理</p>
        <h2>知识库</h2>
        <p class="muted-text">在列表中筛选、拆分、合并知识点，也可以切换到图谱视图查看关系。</p>
      </div>
      <div class="view-toggle" role="group" aria-label="知识库视图切换">
        <button :aria-pressed="view === 'list'" @click="view = 'list'">列表视图</button>
        <button :aria-pressed="view === 'graph'" @click="view = 'graph'">图谱视图</button>
      </div>
    </header>

    <div class="knowledge-toolbar surface-accent-one">
      <label class="form-field">搜索<input v-model="query" placeholder="搜索知识点" aria-label="搜索知识点" @change="loadKnowledge" /></label>
      <label class="form-field">状态<select v-model="approvalStatus" @change="loadKnowledge"><option value="">全部</option><option value="draft">草稿</option><option value="approved">已批准</option><option value="needs_review">待复核</option><option value="rejected">已拒绝</option></select></label>
      <label class="form-field">关系<select v-model="relationshipType" @change="loadKnowledge"><option value="">全部关系</option><option value="related">相关</option><option value="prerequisite">前置</option><option value="duplicate">重复</option><option value="supports">支持</option><option value="contradicts">矛盾</option></select></label>
      <button type="button" @click="loadKnowledge">刷新</button>
    </div>

    <p v-if="loading" class="state-view">正在加载知识库……</p>
    <p v-if="errorMessage" class="state-view" data-state="error">{{ errorMessage }}</p>

    <div v-if="view === 'list'" class="knowledge-list">
      <article v-for="point in filtered" :key="point.id" class="knowledge-item" :class="{ 'is-selected': selectedIds.includes(point.id) }">
        <label class="knowledge-select"><input type="checkbox" :checked="selectedIds.includes(point.id)" @change="toggleSelection(point.id)" /> 选择</label>
        <h3>{{ point.summary || point.content }}</h3>
        <p>{{ point.content }}</p>
        <div class="tag-chip-row"><span class="tag-chip">{{ point.approvalStatus }}</span><span v-for="tag in point.tags" :key="tag.id || tag.name" class="tag-chip">{{ tag.name }}</span><span v-if="point.duplicateGroupId" class="tag-chip">可能重复</span></div>
        <div class="item-actions"><button type="button" @click="selectedIds = [point.id]; dialogMode = 'split'">拆分</button><button type="button" :disabled="selectedIds.length < 2" @click="dialogMode = 'merge'">合并所选</button><button type="button" @click="selectGraphNode(point.id)">查看图谱</button></div>
      </article>
      <p v-if="!filtered.length && !loading" class="state-view" data-state="empty">暂无匹配知识点，请调整筛选条件。</p>
    </div>

    <div v-else class="knowledge-graph-layout">
      <KnowledgeGraphView :graph="graph" :selected-node-id="selectedNodeId" @select-node="selectGraphNode" />
      <KnowledgeRelationshipEditor :source-id="selectedNodeId" :target-id="targetNodeId" @save="saveRelationship" />
    </div>

    <KnowledgeCurationDialogs :open="dialogMode !== null" :mode="dialogMode || 'split'" :selected-count="selectedIds.length" @confirm="confirmDialog" @cancel="dialogMode = null" />
  </section>
</template>
