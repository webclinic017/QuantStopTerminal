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