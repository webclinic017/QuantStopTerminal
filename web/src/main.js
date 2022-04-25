import { createApp } from 'vue'
import App from './App.vue'
import * as appRouter from './router'
import store from './store/index'
import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import { FontAwesomeIcon } from './plugins/font-awesome'

const app = createApp(App)
app.use(appRouter.routeConfig);
app.use(store);
app.component("font-awesome-icon", FontAwesomeIcon)
app.mount('body')
