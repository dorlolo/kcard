<script setup lang="ts">
import { computed, ref } from 'vue'

const relationshipTypes = [
  { value: 'related', label: '相关概念' },
  { value: 'prerequisite', label: '前置知识' },
  { value: 'supports', label: '支持说明' },
  { value: 'contradicts', label: '可能矛盾' },
  { value: 'duplicate', label: '重复/相似' }
]

const props = defineProps<{ sourceId?: string; targetId?: string; disabled?: boolean }>()
const emit = defineEmits<{ save: [relationship: { relationshipType: string; label: string }] }>()

const relationshipType = ref('related')
const label = ref('相关概念')
const canSave = computed(() => Boolean(props.sourceId && props.targetId && props.sourceId !== props.targetId && !props.disabled))

function save() {
  if (!canSave.value) return
  emit('save', { relationshipType: relationshipType.value, label: label.value })
}
</script>

<template>
  <form class="relationship-editor surface-card" @submit.prevent="save">
    <h3>添加知识关系</h3>
    <p class="muted-text">选择两个不同知识点后，可以手动补充它们之间的关系。</p>
    <label class="form-field">
      关系类型
      <select v-model="relationshipType" :disabled="!canSave">
        <option v-for="type in relationshipTypes" :key="type.value" :value="type.value">{{ type.label }}</option>
      </select>
    </label>
    <label class="form-field">
      关系标签
      <input v-model="label" :disabled="!canSave" />
    </label>
    <button class="primary-action" type="submit" :disabled="!canSave">保存关系</button>
  </form>
</template>
