import { createApp } from 'vue';

import App from './App.vue';
import './assets/tailwind.css';

import router from './router/router.js';
import store from './store/index.js';

const app = createApp(App);

app.use(store)
app.use(router);

app.mount('#app');

