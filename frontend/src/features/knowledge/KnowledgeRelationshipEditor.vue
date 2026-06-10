<script setup lang="ts">
const relationshipTypes = ['related', 'prerequisite', 'supports', 'contradicts']
defineProps<{ sourceId?: string; targetId?: string }>()
defineEmits<{ save: [relationship: { relationshipType: string; label: string }] }>()
</script>
<template><form class="surface-card" @submit.prevent="$emit('save', { relationshipType: 'related', label: '相关概念' })"><label>关系 <select><option v-for="type in relationshipTypes" :key="type">{{ type }}</option></select></label><label>标签 <input value="相关概念" /></label><button>保存关系</button></form></template>
