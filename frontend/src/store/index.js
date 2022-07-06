import { createStore } from "vuex";

import userModule from "./modules/user/index"
import authModule from "./modules/auth/index.js";
import playerModule from './modules/player/index.js';

const store = createStore({
    modules: {
        player: playerModule,
        auth: authModule,
        user: userModule
    },
});

export default store;

