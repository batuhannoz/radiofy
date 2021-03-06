import { createRouter, createWebHistory } from 'vue-router';
import store from "@/store";
import ClubList from '../pages/ClubList.vue';
import Login from '../pages/Login';
import Callback from '../pages/Callback';
import CreateClub from '../pages/CreateClub';
import ClubLeader from "@/pages/ClubLeader";
import PageNotFound from '../pages/PageNotFound';
import ClubListener from "@/pages/ClubListener";
import VueCookies from "vue-cookies";



const router = createRouter({
    history: createWebHistory(),
    routes: [
        {path: '/', component: ClubList},
        {path: '/login', component: Login},
        {path: '/callback', component: Callback},
        {path: '/create', component: CreateClub},
        {path: '/club/:id', component: ClubListener, props: true},
        {path: '/club/leader/:id', component: ClubLeader, props: true},
        {path: '/:notFound(.*)', component: PageNotFound}
    ]
});

router.beforeEach(function(to, from, next,) {
    if (store.getters["auth/getAccessToken"] && store.getters["auth/getRadiofyToken"]) {
        next();
    }
    else if (VueCookies.get("access_token") && VueCookies.get("radiofy_token")) {
        store.dispatch("auth/setRadiofyToken", VueCookies.get("radiofy_token"))
        store.dispatch("auth/setAccessToken", VueCookies.get("access_token"))
        router.push("/callback")
    } else if (to.path === "/login" || to.path === "/callback") {
        next()
    } else {
        router.push("/login")
    }
})

export default router;
