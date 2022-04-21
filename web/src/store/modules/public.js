import axios from "axios";
import jwtInterceptor from '../../shared/jwt.interceptor'

const state = () => ({
    subsystemStatus: "",
    version: "",
});

const getters = {
    getSubStatus(state) {
        return state.subsystemStatus;
    },
    getVersion(state) {
        return state.version;
    }
};

const actions = {

    async getSubsystemStatus({ commit }) {
        const response = await jwtInterceptor.get("/api/sub-status", {
            withCredentials: true,
            credentials: "include",
            headers: {
                'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
            },
        });
        if (response && response.data) {
            commit("setSubsystemStatus", response.data);
        }
    },

    async getVersion({ commit }) {
        const response = await jwtInterceptor.get("/api/version", {
            withCredentials: true,
            credentials: "include",
            headers: {
                'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
            },
        });
        if (response && response.data) {
            commit("setVersion", response.data);
        }
    },

};

const mutations = {
    setSubsystemStatus(state, data) {
        state.subsystemStatus = data
    },
    setVersion(state, data) {
        state.version = data
    },
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};