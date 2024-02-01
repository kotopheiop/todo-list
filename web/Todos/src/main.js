import {createApp} from 'vue'
import App from './App.vue'
import './index.css'
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

const app = createApp(App)


const options = {}

app.use(Toast, options)

app.mount('#app')
