import { createStore } from "vuex";

import authModule from "./modules/auth/index.js";
import playerModule from './modules/player/index.js';

const store = createStore({
    modules: {
        player: playerModule,
        auth: authModule
    },
});

export default store;

