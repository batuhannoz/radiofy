import {
    getCurrentState,
    setVolume,
    shuffle,
    play,
    pause,
    repeat,
    seekThePosition,
    skipToNext,
    playSong,
    skipToPrevious} from '@/api/spotify/player.js';

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
    enableShuffle(context) {
        shuffle(true).then(() => {
            context.commit('setShuffle', true);
        })
    },
    disableShuffle(context) {
        shuffle(false).then(() => {
            context.commit('setShuffle', false);
        })
    },
    play(context) {
        play().then(() => {
            context.commit('setPlayState', true);
        })
    },
    pauseSong(context) {
        pause().then(() => {
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
    skipToNext({dispatch}) {
        skipToNext().then(() => {
            setTimeout(() => {
                dispatch("refreshPlayer")
            }, 750)
        })
    },
    skipToPrevious({dispatch}) {
        skipToPrevious().then(() => {
            setTimeout(() => {
                dispatch("refreshPlayer")
            }, 750)
        })
    },
    playSong({dispatch}, songInfo) {
        playSong(songInfo.albumID, songInfo.position, songInfo.deviceID).then(() => {
            setTimeout(() => {
                dispatch("refreshPlayer")
            }, 750)
        })
    },
    refreshPlayer(context) {
        getCurrentState().then((res) => {
            context.commit("setPosition", res.item.track_number)
            context.commit("setAlbumID", res.item.album.uri)
            context.commit("setVolume", res.device.volume_percent);
            context.commit("setShuffle", res.shuffle_state);
            context.commit("setRepeat", res.repeat_state);
            context.commit("setProgressMS", res.progress_ms);
            context.commit("setArtistName", res.item.artists[0].name);
            context.commit("setSongName", res.item.name);
            context.commit("setImage", res.item.album.images[0].url);
            context.commit("setPlayState", res.is_playing);
            context.commit("setDurationMS", res.item.duration_ms);
            context.commit("setSongID", res.item.id)
        })
    }
}