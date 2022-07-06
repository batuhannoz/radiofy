import userGetters from './getters.js'
import userActions from './actions.js'
import userMutations from './mutations.js'

export default {
    namespaced: true,
    state() {
        return {
            displayName: null,
            image: null,
            spotifyID: null,
            country: null
        };
    },
    getters: userGetters,
    mutations: userMutations,
    actions: userActions
};