<script setup lang="ts">
defineProps<{ status: string; progressPercent?: number; failureReason?: string }>()
defineEmits<{ retry: [] }>()
</script>

<template>
  <section class="surface-accent-two state-view" aria-live="polite">
    <strong>处理状态：{{ status }}</strong>
    <progress v-if="progressPercent !== undefined" :value="progressPercent" max="100" />
    <p v-if="failureReason">{{ failureReason }}</p>
    <button v-if="status === 'failed'" @click="$emit('retry')">重新分析</button>
  </section>
</template>
