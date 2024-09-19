import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import {setAuth, axios} from './services/axios.js';
import utils from './services/utils.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/dashboard.css'
import './assets/main.css'


const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$utils = utils;
app.config.globalProperties.$setAuth = setAuth;

// Components
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
//app.component("PostModal", PostModal);


app.use(router)
app.mount('#app')
