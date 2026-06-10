<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps<{ mode: 'split' | 'merge'; open: boolean; selectedCount?: number }>()
const emit = defineEmits<{ confirm: [payload: string[]]; cancel: [] }>()

const first = ref('')
const second = ref('')
const merged = ref('')

watch(() => props.open, () => {
  first.value = ''
  second.value = ''
  merged.value = ''
})

function confirm() {
  if (props.mode === 'split') emit('confirm', [first.value, second.value].filter(Boolean))
  else emit('confirm', [merged.value].filter(Boolean))
}
</script>

<template>
  <dialog :open="open" class="curation-dialog surface-card">
    <h3>{{ mode === 'split' ? '拆分知识点' : '合并知识点' }}</h3>
    <p v-if="mode === 'merge'">已选择 {{ selectedCount ?? 0 }} 个知识点，请填写合并后的内容。</p>
    <template v-if="mode === 'split'">
      <label class="form-field">拆分内容一<textarea v-model="first" rows="3" /></label>
      <label class="form-field">拆分内容二<textarea v-model="second" rows="3" /></label>
    </template>
    <label v-else class="form-field">合并后内容<textarea v-model="merged" rows="4" /></label>
    <div class="dialog-actions">
      <button class="primary-action" type="button" @click="confirm">确认</button>
      <button type="button" @click="emit('cancel')">取消</button>
    </div>
  </dialog>
</template>
