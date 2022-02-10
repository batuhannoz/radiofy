import { createRouter, createWebHistory } from 'vue-router';


import Login from './pages/Login.vue'
import Token from './pages/token.vue'
import Radiofy from './pages/radiofy.vue'


const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', component: Login},
        {path:'/token', component: Token},
        {path: '/home', component: Radiofy}
        //{path: '/:id', component: Home},
        //{path: '/room/:id',component: Room},
        //{path: '/:notFound(.*)', component: NotFound},
        //{path: '/home'}
    ]
});

export default router;
