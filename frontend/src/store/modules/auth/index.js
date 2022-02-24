import authGetters from './getters'
import authActions from './actions.js'
import authMutations from './mutations.js'

export default {
    namespaced: true,
    state() {
        return{
            accessToken: "",
            refreshToken: "",
        };
    },
    getters: authGetters,
    mutations: authMutations,
    actions: authActions
};