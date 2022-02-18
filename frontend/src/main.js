import { createApp } from 'vue';

import router from './router/router.js';
import App from './App.vue';


import './assets/tailwind.css'


const app = createApp(App);

app.use(router);

app.mount('#app');

