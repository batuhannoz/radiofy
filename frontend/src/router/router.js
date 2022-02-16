import { createRouter, createWebHistory } from 'vue-router';

import Token from '../pages/Token.vue'
import Login from "@/pages/Login";
import MainPage from '../pages/Clubs'
const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', component: MainPage},
        {path: "/login", component: Login},
        {path: '/token', component: Token},
    ]
});

export default router;
