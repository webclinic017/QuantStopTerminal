import axios from "axios";
import jwtInterceptor from '../../shared/jwt.interceptor'

const state = () => ({
  loginApiStatus: "",
  registerApiStatus: "",
  userProfile: {
    id: 0,
    username: "",
    roles: "",
  },
  logOut: false
});

const getters = {
  getLoginApiStatus(state) {
    return state.loginApiStatus;
  },
  getRegisterApiStatus(state) {
    return state.registerApiStatus;
  },
  getUserProfile(state) {
    return state.userProfile;
  },
  getLogout(state) {
    return state.logOut;
  }
};

const actions = {
  async loginApi({ commit }, payload) {
    const response = await axios.post("https://localhost/api/session", payload, {
      withCredentials: true,
      credentials: "include",
      headers: {
        'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
      },
    });

    if (response.status === 200 && response.data) {
      localStorage.setItem("isAuthenticated", "true");
      commit("setUserProfile", response.data);
      commit("setLoginApiStatus", "success");
    } else {
      commit("setLoginApiStatus", "failed");
    }
  },

  async registerApi({ commit }, payload) {
    const response = await axios.post("https://localhost/api/signup", payload, {
      withCredentials: true,
      credentials: "include",
      headers: {
        'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
      },
    });

    if (response && response.data) {
      localStorage.setItem("isAuthenticated", "true");
      commit("setRegisterApiStatus", "success");
    } else {
      commit("setRegisterApiStatus", "failed");
    }
  },

  async userProfile({ commit }) {
    const response = await jwtInterceptor.get("https://localhost/api/user", {
      withCredentials: true,
      credentials: "include",
      headers: {
        'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
      },
    });

    if (response && response.data) {
      commit("setUserProfile", response.data);
    }
  },

  async userLogout({ commit }) {
    const response = await axios.delete("https://localhost/api/session", {
      withCredentials: true,
      credentials: "include",
      headers: {
        'X-Requested-With': 'XMLHttpRequest' // CSRF prevention
      },
    });

    if (response && response.data) {
      commit("setLogout", true);
      localStorage.removeItem("isAuthenticated");
    } else {
      commit("setLogout", false);
    }
  },
};

const mutations = {
  setLoginApiStatus(state, data) {
    state.loginApiStatus = data;
  },

  setRegisterApiStatus(state, data) {
    state.registerApiStatus = data;
  },

  setUserProfile(state, data) {
    state.userProfile = {
      id: data.id,
      username: data.username,
      roles: data.roles
    };
  },

  setLogout(state, payload) {
    state.logOut = payload;
  }
};

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations,
};
