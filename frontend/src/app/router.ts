import { defineComponent, h } from 'vue'
import { createRouter, createWebHistory, useRoute } from 'vue-router'

const Placeholder = defineComponent({
  name: 'RoutePlaceholder',
  setup() {
    const route = useRoute()
    return () =>
      h('section', { class: 'surface-card' }, [
        h('h2', String(route.meta.title ?? '工作区')),
        h('p', '该工作区功能正在搭建中。')
      ])
  }
})

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', component: Placeholder, meta: { title: '首页' } },
    { path: '/materials/import', component: () => import('../features/materials/MaterialImportPage.vue'), meta: { title: '导入资料' } },
    { path: '/knowledge', component: Placeholder, meta: { title: '知识库' } },
    { path: '/cards', component: Placeholder, meta: { title: '卡片与卡组' } },
    { path: '/review', component: Placeholder, meta: { title: '复习会话' } },
    { path: '/plans', component: Placeholder, meta: { title: '复习计划' } },
    { path: '/prompts', component: Placeholder, meta: { title: '提示词预设' } }
  ]
})
