import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import {setAuth, axios} from './services/axios.js';
import utils from './services/utils.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import PhotoGrid from './components/PhotoGrid.vue';
import UserInfo from './components/UserInfo.vue';
import ChangeUsernameForm from './components/ChangeUsernameModal.vue';
import UploadPhotoModal from './components/UploadPhotoModal.vue'
import PhotoModal from './components/PhotoModal.vue';


const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.config.globalProperties.$utils = utils;
app.config.globalProperties.$setAuth = setAuth;

// Components
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("PhotoGrid", PhotoGrid);
app.component("UserInfo", UserInfo);
app.component("ChangeUsernameForm", ChangeUsernameForm);
app.component("UploadPhotoModal", UploadPhotoModal)
app.component("PhotoModal", PhotoModal)


app.use(router)
app.mount('#app')