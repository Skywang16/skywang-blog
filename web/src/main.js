import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import { router } from './router'
import Vue3Lottie from 'vue3-lottie'
import wLoading from './components/wLoading'

const app = createApp(App)

app.use(router)
app.use(wLoading)
app.use(Vue3Lottie, { name: 'Vue3Lottie' })

app.mount('#app')
