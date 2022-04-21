import Home from "./pages/Home.vue";
import Login from "./pages/Login.vue";
import Profile from "./pages/Profile.vue";
import Settings from "./pages/settings/Settings";
import UserManager from "./pages/admin/UserManager";
import FinancialChart from "./pages/FinancialChart";
import Exchanges from "./pages/Exchanges";
import { createRouter, createWebHistory } from "vue-router";
import store from "./store/index";

const routes = [
    { path: "/", component: Home, meta: { requiredAuth: true } },
    { path: "/home", component: Home, meta: { requiredAuth: true } },
    { path: "/login", component: Login, meta: { requiredAuth: false } },
    { path: "/profile", component: Profile, meta: { requiredAuth: true } },
    { path: "/settings", component: Settings, meta: { requiredAuth: true } },
    { path: "/users", component: UserManager, meta: { requiredAuth: true } },
    { path: "/chart", component: FinancialChart, meta: { requiredAuth: true } },
    { path: "/exchanges", component: Exchanges, meta: { requiredAuth: true } },
];

export const routeConfig = createRouter({
    history: createWebHistory(),
    routes: routes,
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
    }
    return next();
});
