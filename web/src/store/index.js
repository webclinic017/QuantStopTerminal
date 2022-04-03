/*
import { createStore } from "vuex";
import { auth } from "./auth.module";

const store = createStore({
  modules: {
    auth,
  },
});

export default store;
*/

import { createStore } from "vuex";
import authModule from './modules/auth';
import userModule from './modules/user';

const store = createStore({
  modules:{
    auth: authModule,
    user: userModule
  }
});

export default store;