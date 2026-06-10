<script setup lang="ts">
import type { KnowledgeGraph } from '../../services/knowledge'

defineProps<{ graph: KnowledgeGraph }>()
defineEmits<{ selectNode: [id: string] }>()
</script>

<template>
  <section class="knowledge-graph surface-accent-two" aria-label="知识关系图谱">
    <p v-if="graph.warnings?.length" role="status">{{ graph.warnings.join(' ') }}</p>
    <svg viewBox="0 0 800 420" role="img" aria-label="类 Obsidian 知识图谱">
      <line v-for="(edge, index) in graph.edges" :key="edge.id" :x1="80 + index * 40" y1="120" :x2="160 + index * 40" y2="220" stroke="currentColor" />
      <g v-for="(node, index) in graph.nodes" :key="node.id" :transform="`translate(${80 + (index % 8) * 90}, ${100 + Math.floor(index / 8) * 90})`" tabindex="0" role="button" @click="$emit('selectNode', node.id)">
        <circle r="24" fill="var(--color-accent-lilac)" stroke="var(--color-border)" />
        <text y="44" text-anchor="middle">{{ node.label.slice(0, 18) }}</text>
      </g>
    </svg>
    <details>
      <summary>可访问的图谱列表</summary>
      <ul><li v-for="node in graph.nodes" :key="node.id">{{ node.label }}</li></ul>
    </details>
  </section>
</template>
