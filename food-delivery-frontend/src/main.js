import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router/router'
import components from '@/components/UI/index'

const app = createApp(App)
const pinia = createPinia()

components.forEach(component => {
    app.component(component.name, component)
})

//app.component('Button', Button)

app.use(pinia)
app.use(router)




app.mount('#app')
