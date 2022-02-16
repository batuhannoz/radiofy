import { createApp } from 'vue';

import router from './router/router.js';
import App from './App.vue';


import { VuesticPlugin } from "vuestic-ui"
import 'vuestic-ui/dist/vuestic-ui.css'
import './assets/tailwind.css'


const app = createApp(App);

app.use(router);
app.use(VuesticPlugin);

app.mount('#app');

