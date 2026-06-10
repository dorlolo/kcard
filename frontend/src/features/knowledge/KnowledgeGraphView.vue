<script setup lang="ts">
import { computed } from 'vue'
import type { KnowledgeGraph } from '../../services/knowledge'

const props = defineProps<{ graph: KnowledgeGraph; selectedNodeId?: string }>()
const emit = defineEmits<{ selectNode: [id: string] }>()

const positionedNodes = computed(() => {
  const radius = 145
  const centerX = 380
  const centerY = 210
  const total = Math.max(props.graph.nodes.length, 1)
  return props.graph.nodes.map((node, index) => {
    const angle = (Math.PI * 2 * index) / total - Math.PI / 2
    return { ...node, x: centerX + Math.cos(angle) * radius, y: centerY + Math.sin(angle) * radius }
  })
})

const nodeMap = computed(() => new Map(positionedNodes.value.map((node) => [node.id, node])))
const visibleEdges = computed(() => props.graph.edges.map((edge) => ({ ...edge, source: nodeMap.value.get(edge.sourceId), target: nodeMap.value.get(edge.targetId) })).filter((edge) => edge.source && edge.target))
</script>

<template>
  <section class="knowledge-graph surface-accent-two" aria-label="知识关系图谱">
    <p v-if="graph.warnings?.length" class="graph-warning" role="status">{{ graph.warnings.join(' ') }}</p>
    <svg viewBox="0 0 760 430" role="img" aria-label="类 Obsidian 知识图谱">
      <g v-for="edge in visibleEdges" :key="edge.id">
        <line :x1="edge.source!.x" :y1="edge.source!.y" :x2="edge.target!.x" :y2="edge.target!.y" stroke="var(--color-muted)" stroke-width="1.5" />
        <text :x="(edge.source!.x + edge.target!.x) / 2" :y="(edge.source!.y + edge.target!.y) / 2 - 6" text-anchor="middle" class="graph-edge-label">
          {{ edge.label || edge.relationshipType }}
        </text>
      </g>
      <g
        v-for="node in positionedNodes"
        :key="node.id"
        :transform="`translate(${node.x}, ${node.y})`"
        tabindex="0"
        role="button"
        :aria-label="`打开知识点 ${node.label}`"
        class="graph-node"
        :class="{ 'is-selected': selectedNodeId === node.id }"
        @click="emit('selectNode', node.id)"
        @keydown.enter.prevent="emit('selectNode', node.id)"
        @keydown.space.prevent="emit('selectNode', node.id)"
      >
        <circle r="28" />
        <text y="44" text-anchor="middle">{{ node.label.slice(0, 18) }}</text>
      </g>
    </svg>
    <details>
      <summary>可访问的图谱列表</summary>
      <ul>
        <li v-for="node in graph.nodes" :key="node.id">
          <button type="button" @click="emit('selectNode', node.id)">{{ node.label }}</button>
        </li>
      </ul>
    </details>
  </section>
</template>
