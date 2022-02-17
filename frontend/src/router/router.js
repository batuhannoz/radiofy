import { createRouter, createWebHistory } from 'vue-router';

import Token from '../pages/Token.vue'
import MainPage from '../pages/Clubs'
import Login from "@/pages/Login";
const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', component: MainPage},
        {path: "/login", component: Login},
        {path: '/token', component: Token},
    ]
});

export default router;
