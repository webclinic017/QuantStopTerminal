import { createStore } from "vuex";
import authModule from './modules/auth';
import publicModule from './modules/public';
import adminModule from './modules/admin';

const store = createStore({
  modules:{
    auth: authModule,
    public: publicModule,
    admin: adminModule,
  }
});

export default store;