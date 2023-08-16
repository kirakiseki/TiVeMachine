import { createApp } from 'vue'
import ArcoVue from '@arco-design/web-vue'
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
import '@arco-design/web-vue/dist/arco.css'
import generatedRoutes from 'virtual:generated-pages'
import { setupLayouts } from 'virtual:generated-layouts'
import { createRouter, createWebHistory } from 'vue-router'
import App from './App.vue'
import setupLoginGuard from './router/setup/loginGuard'
import pinia from './store'
import '~/api/middleware'

export const routes = setupLayouts(generatedRoutes)
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    ...routes,
  ],
})

setupLoginGuard(router)

createApp(App).use(router).use(ArcoVue).use(ArcoVueIcon).use(pinia).mount('#app')
