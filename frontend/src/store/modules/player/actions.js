import {
    getCurrentState,
    setVolume,
    shuffle,
    play,
    pause,
    repeat,
    seekThePosition} from '@/api/spotify/player.js';

export default {
    setDeviceID(context, deviceID) {
        context.commit("setDeviceID", deviceID)
    },
    setImage(context, image) {
        context.commit("setImage", image)
    },
    setVolume(context, volume) {
        setVolume(volume).then(() => {
            context.commit("setVolume", volume)
        })
    },
    setShuffle({context, state}, shuffleState) {
        shuffle(shuffleState, state.deviceID).then(() => {
            context.commit('setShuffle', shuffleState);
        })
    },
    enableShuffle({context, state}) {
        shuffle(true, state.deviceID).then(() => {
            context.commit('setShuffle', true);
        })
    },
    disableShuffle({context, state}) {
        shuffle('off', state.deviceID).then(() => {
            context.commit('setShuffle', 'off');
        })
    },
    playSong({context, state}) {
        play(state.deviceID).then(() => {
            context.commit('setPlayState', true);
        })
    },
    pauseSong({context, state}) {
        pause(state.deviceID).then(() => {
            context.commit('setPlayState', false);
        })
    },
    setSongName(context, songName) {
        context.commit('setSongName', songName);
    },
    setArtistName(context, artistName) {
        context.commit('setArtistName', artistName);
    },
    setDurationMS(context, durationMS) {
        context.commit('setDurationMS', durationMS);
    },
    setProgress(context, progressMS) {
        context.commit('setProgressMS', progressMS);
    },
    setRepeat(context, isRepeat) {
        repeat(isRepeat).then(() => {
            context.commit('setRepeat', isRepeat);
        })
    },
    enableRepeat(context) {
        repeat('track').then(() => {
            context.commit('setRepeat', 'track');
        });
    },
    disableRepeat(context) {
        repeat('off').then(() => {
            context.commit('setRepeat', 'off');
        });
    },
    seekThePosition(context, position) {
      seekThePosition(position).then(() => {
          context.commit("setProgressMS", position);
      });
    },
    refreshPlayer(context) {
        getCurrentState().then((res) => {
            context.commit("setVolume", res.device.volume_percent);
            context.commit("setShuffle", res.shuffle_state);
            context.commit("setRepeat", res.repeat_state);
            context.commit("setProgressMS", res.progress_ms);
            context.commit("setArtistName", res.item.artists[0].name);
            context.commit("setSongName", res.item.name);
            context.commit("setImage", res.item.album.images[0].url);
            context.commit("setPlayState", res.is_playing);
            context.commit("setDurationMS", res.item.duration_ms);
        })
    }
}