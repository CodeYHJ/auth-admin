import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import 'ant-design-vue/dist/antd.css';
// window.config = { BASE_URL: import.meta.env.BASE_URL }

const app = createApp(App)

app.use(router)

app.mount('#app')

