import { createApp } from 'vue';
import App from './App.vue';
import { VaSlider } from 'vuestic-ui'

const app = createApp(App);

app.component('va-slider',VaSlider)

app.mount('#app');

