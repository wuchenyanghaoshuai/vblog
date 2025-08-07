import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(createPinia())
app.use(router)

//引入 Arco Design Vue组件
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
app.use(ArcoVue);
app.mount('#app')
