export default {
    getDeviceID(state) {
        return state.deviceID
    },
    getImage (state) {
        return state.image;
    },
    getVolume (state) {
        return state.volume;
    },
    getIsShuffle (state) {
        return state.isShuffle;
    },
    getIsPlay (state) {
        return state.isPlay;
    },
    getSongName (state) {
        return state.songName;
    },
    getArtistName (state) {
        return state.artistName;
    },
    getDurationMS (state) {
        return state.durationMS;
    },
    getProgressMS (state) {
        return state.progressMS;
    },
    getIsRepeat (state) {
        return state.isRepeat;
    },
}
