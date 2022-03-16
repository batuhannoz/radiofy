import playerGetters from './getters.js'
import playerActions from './actions.js'
import playerMutations from './mutations.js'

export default {
    namespaced: true,
    state() {
        return{
            deviceID: null,
            image: null,
            volume: 30,
            isShuffle: null,
            isPlay: null,
            songName: null,
            artistName: null,
            durationMS: null,
            progressMS: null,
            isRepeat: null,
            songID: null
        };
    },
    getters: playerGetters,
    mutations: playerMutations,
    actions: playerActions
};