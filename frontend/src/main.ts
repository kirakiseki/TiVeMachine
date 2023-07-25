import { createApp } from 'vue'
import ArcoVue from '@arco-design/web-vue'
import './style.css'
import '@arco-design/web-vue/dist/arco.css'
import App from './App.vue'

createApp(App).use(ArcoVue).mount('#app')
