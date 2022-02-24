export default {
    setDeviceID(state, deviceID) {
        state.deviceID = deviceID;
    },
    setImage(state, image) {
        state.image = image
    },
    setVolume(state, volume) {
        state.volume = volume
    },
    setShuffle(state, shuffleState) {
        state.isShuffle = shuffleState
    },
    setPlayState(state, playState) {
        state.isPlay = playState
    },
    setSongName(state, songName) {
        state.songName = songName;
    },
    setArtistName(state, artistName) {
        state.artistName = artistName;
    },
    setDurationMS(state, durationMS) {
        state.durationMS = durationMS;
    },
    setProgressMS(state, progressMS) {
        state.progressMS = progressMS;
    },
    setRepeat(state, repeatState) {
        state.isRepeat = repeatState;
    }
};