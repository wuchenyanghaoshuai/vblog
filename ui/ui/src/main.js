import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(router)

app.mount('#app')
// 引入ui
import ArcoVue from '@arco-design/web-vue'
import '@arco-design/web-vue/dist/arco.css'
app.use(ArcoVue)
//引入图标库
import ArcoVueIcon from '@arco-design/web-vue/es/icon'
app.use(ArcoVueIcon)
