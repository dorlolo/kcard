import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { VueQueryPlugin } from '@tanstack/vue-query'
import AppShell from './AppShell.vue'
import { router } from './router'
import '../styles/tokens.css'

createApp(AppShell).use(createPinia()).use(VueQueryPlugin).use(router).mount('#app')
