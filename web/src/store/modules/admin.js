import axios from "axios";
import jwtInterceptor from '../../shared/jwt.interceptor'

const state = () => ({
    users: "",

});

const getters = {
    getUsers(state) {
        return state.users;
    },

};

const actions = {
    async getAllUsers({ commit }) {
        const response = await axios.get("https://localhost/api/get-users", {
            withCredentials: true,
            credentials: "include",
            headers: {
                'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
            },
        });

        if (response.status === 200 && response.data) {
            commit("setUsers", response.data);
        } else {
            commit("setUsers", "failed");
        }
    },



};

const mutations = {
    setUsers(state, data) {
        state.users = data;
    },

};

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations,
};
