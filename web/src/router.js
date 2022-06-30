import { createRouter, createWebHistory } from 'vue-router'
import AccountView from './views/Account/index.vue'
import ClientView from './views/Client/index.vue'
import TokenView from './views/Token/index.vue'
import HomeView from './App.vue'



const routes = [
    { path: '/', redirect: '/client' },
    { path: '/client', component: ClientView },
    { path: '/token', component: TokenView },
    { path: '/account', component: AccountView },

]


const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router