import { createRouter, createWebHistory } from 'vue-router';
import store from "@/store";
import Main from '../pages/Main.vue';
import Login from '../pages/Login';
import Callback from '../pages/Callback';
import CreateClub from '../pages/CreateClub';
//import VueCookies from "vue-cookies";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', component: Main},
        {path: '/login', component: Login},
        {path: '/callback', component: Callback},
        {path: '/create', component: CreateClub},
    ]
});

router.beforeEach(function(to, from, next,) {
    if(to.path === "/login" || to.path === "/callback") {
        next();
    } else if(!store.getters["auth/getAccessToken"]) {
        router.push("/login");
    }
    next();
})

export default router;
