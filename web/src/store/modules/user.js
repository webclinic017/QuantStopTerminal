import axios from "axios";
import jwtInterceptor from '../../shared/jwt.interceptor'

const API_URL = 'https://localhost/api/';

const state = () => ({
    publicContent: "",
});

const getters = {
    getPublicAll(state) {
        return state.publicContent;
    },
};

const actions = {

    async getPublicContent({ commit }) {
        const response = await jwtInterceptor.get("https://localhost/api/sub-status", {
            withCredentials: true,
            credentials: "include",
            headers: {
                'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
            },
        });
        if (response && response.data) {
            commit("setPublicContent", response.data);
        }
    },

};

const mutations = {
    setPublicContent(state, data) {
        state.publicContent = data
    },
};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};