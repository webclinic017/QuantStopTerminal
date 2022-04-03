import Home from "./components/Home.vue";
import Login from "./components/Login.vue";
import Register from "./components/Register.vue";
import Profile from "./components/Profile.vue";
import Verify from "./components/Verify";
import BoardAdmin from "./components/BoardAdmin.vue";
import BoardModerator from "./components/BoardModerator.vue";
import BoardUser from "./components/BoardUser.vue";
import { createRouter, createWebHistory } from "vue-router";
import store from "./store/index";

const routes = [
    { path: "/", component: Home, meta: { requiredAuth: false } },
    { path: "/home", component: Home, meta: { requiredAuth: false } },
    { path: "/login", component: Login, meta: { requiredAuth: false } },
    { path: "/register", component: Register, meta: { requiredAuth: false } },
    { path: "/profile", component: Profile, meta: { requiredAuth: true } },
    { path: "/verify/:code", component: Verify, meta: { requiredAuth: false } },
    { path: "/admin", component: BoardAdmin, meta: { requiredAuth: true } },
    { path: "/mod", component: BoardModerator, meta: { requiredAuth: true } },
    { path: "/user", component: BoardUser, meta: { requiredAuth: true } },
];

export const routeConfig = createRouter({
    history: createWebHistory(),
    routes: routes,
    //props: route => ({ query: route.query })
});

routeConfig.beforeEach(async (to, from, next) => {
    let userProfile = store.getters["auth/getUserProfile"];
    let isAuthenticated = localStorage.getItem("isAuthenticated");
    if (userProfile.id !== 0 && isAuthenticated) {
        await store.dispatch("auth/userProfile");
        userProfile = store.getters["auth/getUserProfile"];
    }

    if (to.meta.requiredAuth) {
        if (userProfile.id === 0) {
            return next({ path: "/login" });
        }
    } /*else {
        if (userProfile.id !== 0) {
            return next({ path: "/profile" });
        }
    }*/
    return next();
});
