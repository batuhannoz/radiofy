import {createStore} from 'vuex'

const store = createStore({
    state() {
        return{
            CurrentSongImage: 'https://i.scdn.co/image/ab67616d000048511624590458126fc8b8c64c2f',
            Volume: 50
        }
    },
    getters: {
        CurrentVolume (state) {
            return state.Volume
        },
        CurrentSongImage (state) {
            return state.CurrentSongImage
        }
    },
    mutations: {
        ChangeCurrentImage(state, img) {
            state.CurrentSongImage = img
        }
    },
    actions: {
        ChangeCurrentImage(context, img) {
            context.commit('ChangeCurrentImage', img)
        }
    }
})
export default store;