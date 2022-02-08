import { createRouter, createWebHistory } from 'vue-router';




const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', component: Login},
        {path: '/home', component: Home},
        {path: '/room/:id',component: Room},
        {path: '/:notFound(.*)', component: NotFound}


    ]
});

export default router;
