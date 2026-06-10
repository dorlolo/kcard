<script setup lang="ts">
const navItems = [
  { to: '/', label: '首页' },
  { to: '/materials/import', label: '资料导入' },
  { to: '/knowledge', label: '知识库' },
  { to: '/cards', label: '卡片' },
  { to: '/review', label: '复习' },
  { to: '/plans', label: '计划' },
  { to: '/prompts', label: '提示词' }
]
</script>

<template>
  <div class="app-shell">
    <header class="app-shell__header surface-accent-one">
      <h1>KCard 学习卡片</h1>
      <nav aria-label="主导航">
        <RouterLink v-for="item in navItems" :key="item.to" :to="item.to">{{ item.label }}</RouterLink>
      </nav>
    </header>
    <main class="app-shell__main">
      <RouterView />
    </main>
  </div>
</template>
