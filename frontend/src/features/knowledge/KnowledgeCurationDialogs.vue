<script setup lang="ts">
defineProps<{ mode: 'split' | 'merge'; open: boolean }>()
defineEmits<{ confirm: [payload: string[]]; cancel: [] }>()
</script>
<template><dialog :open="open"><h3>{{ mode === 'split' ? '拆分知识点' : '合并知识点' }}</h3><slot /><button @click="$emit('confirm', [])">确认</button><button @click="$emit('cancel')">取消</button></dialog></template>
