<script setup lang="ts">
defineProps<{ state: 'loading' | 'empty' | 'error' | 'partial' | 'disabled' | 'undo' | 'success'; title?: string; message?: string }>()
</script>

<template>
  <section class="state-view" :data-state="state" role="status" aria-live="polite">
    <strong>{{ title ?? state }}</strong>
    <p>{{ message }}</p>
    <slot />
  </section>
</template>
