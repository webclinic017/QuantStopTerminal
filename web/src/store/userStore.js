
import { reactive } from 'vue'
import axios from "axios";
import jwtInterceptor from "../shared/jwt.interceptor";

export const userStore = reactive({
    loginStatus: "",
    registerExchangeStatus: "",
    userProfile: {
        id: 0,
        username: "",
        roles: "",
    },
    logOut: false,

    /* Getters */
    getLoginStatus() {
        return this.loginStatus;
    },
    getRegisterExchangeStatus() {
        return this.registerExchangeStatus;
    },
    getUserProfile() {
        console.log('getting userProfile from localStorage ...');
        if (localStorage.getItem('userProfile')) {
            try {
                let p = JSON.parse(localStorage.getItem('userProfile'))
                this.userProfile = p;
            } catch(e) {
                console.log('error getting chartStyle: ' + e);
                localStorage.removeItem('userProfile');
            }
        } else {
            // first time no local storage saved yet, save the defaults
            console.log('no userProfile has been saved, saving defaults ...');
            this.setUserProfile(this.userProfile)
        }
        return this.userProfile;
    },
    getLogout() {
        return this.logOut;
    },

    /* Setters */
    setLoginStatus(data) {
        this.loginStatus = data;
    },
    setRegisterExchangeStatus(data) {
        this.registerExchangeStatus = data;
    },
    setUserProfile(data) {
        this.userProfile = {
            id: data.id,
            username: data.username,
            roles: data.roles
        };
        localStorage.setItem("userProfile", JSON.stringify(this.userProfile));
    },
    setLogout(data) {
        this.logOut = data;
    },

    /* Actions */
    async actionLogin(payload) {
        const response = await axios.post("/api/session", payload, {
            withCredentials: true,
            credentials: "include",
            headers: {
                'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
            },
        });
        if (response.status === 200 && response.data) {
            localStorage.setItem("isAuthenticated", "true");
            this.setUserProfile(response.data)
            this.setLoginStatus("success")
        } else {
            this.setLoginStatus("failed")
        }
    },
    async actionRegisterExchange({ commit }, payload) {
        const response = await axios.post("/api/exchange", payload, {
            withCredentials: true,
            credentials: "include",
            headers: {
                'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
            },
        });
        if (response && response.data) {
            this.setRegisterExchangeStatus("success")
        } else {
            this.setRegisterExchangeStatus("failed")
        }
    },
    async actionGetUserProfile({ commit }) {
        const response = await jwtInterceptor.get("/api/user", {
            withCredentials: true,
            credentials: "include",
            headers: {
                'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
            },
        });

        if (response && response.data) {
            this.setUserProfile(response.data)
        }
    },
    async actionLogout() {
        const response = await axios.delete("/api/session", {
            withCredentials: true,
            credentials: "include",
            headers: {
                'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
            },
        });

        if (response && response.data) {
            this.setLogout(true)
            localStorage.removeItem("isAuthenticated");
            let p = {
                id: 0,
                username: "",
                roles: "",
            }
            this.setUserProfile(p)
        } else {
            this.setLogout(false)
        }
    },



})