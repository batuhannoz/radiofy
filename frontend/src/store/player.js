import {
    getCurrentState,
    setVolume,
    setShuffle,
    play,
    pause,
    seekThePosition,
    repeat
} from "@/services/SpotifyPlayer";
import {createStore} from 'vuex'

const store = createStore({
    state() {
        return{
            Image: null,
            Volume: 50,
            IsShuffle: null,
            IsPlay: null,
            SongName: null,
            ArtistName: null,
            DurationMS: null,
            ProgressMS: null,
            IsRepeat: null,
        }
    },
    getters: {
        GetImage (state) {
            return state.Image;
        },
        GetVolume (state) {
            return state.Volume;
        },
        GetIsShuffle (state) {
            return state.IsShuffle;
        },
        GetIsPlay (state) {
            return state.IsPlay;
        },
        GetSongName (state) {
            return state.SongName;
        },
        GetArtistName (state) {
            return state.ArtistName;
        },
        GetDurationMS (state) {
            return state.DurationMS;
        },
        GetProgressMS (state) {
            return state.ProgressMS;
        },
        GetIsRepeat (state) {
            return state.IsRepeat;
        },
    },
    mutations: {
        ChangeImage(state, Image) {
            state.Image = Image
        },
        ChangeVolume(state, Volume) {
            state.Volume = Volume
        },
        ChangeIsShuffle(state, shuffleState) {
            state.IsShuffle = shuffleState
        },
        ChangeIsPlay(state, playState) {
            state.IsPlay = playState
        },
        ChangeSongName(state, songName) {
            state.SongName = songName;
        },
        ChangeArtistName(state, artistName) {
            state.ArtistName = artistName;
        },
        ChangeDurationMS(state, durationMS) {
            state.DurationMS = durationMS;
        },
        ChangeProgressMS(state, progressMS) {
            state.ProgressMS = progressMS;
        },
        ChangeIsRepeat(state, repeatState) {
            state.IsRepeat = repeatState;
        }
    },
    actions: {
        ChangeImage(context, Image) {
            context.commit('ChangeImage', Image);
        },
        ChangeVolume(context, Volume) {
            setVolume(Volume);
            context.commit('ChangeVolume', Volume);
            console.log(Volume)
        },
        ChangeShuffle(context, IsShuffle) {
            setShuffle(IsShuffle);
            context.commit('ChangeIsShuffle', IsShuffle);
        },
        DisableShuffle(context) {
            setShuffle(false);
            context.commit('ChangeIsShuffle', false);
        },
        EnableShuffle(context) {
            setShuffle(true);
            context.commit('ChangeIsShuffle', true);
        },
        PlaySong(context){
            play();
            context.commit('ChangeIsPlay', true);
        },
        PauseSong(context) {
            pause();
            context.commit('ChangeIsPlay', false);
        },
        ChangeSongName(context, SongName) {
            context.commit('ChangeSongName', SongName);
        },
        ChangeArtistName(context, ArtistName) {
            context.commit('ChangeArtistName', ArtistName);
        },
        ChangeDurationMS(context, DurationMS) {
            context.commit('ChangeDurationMS', DurationMS);
        },
        ChangeProgressMsStore(context, ProgressMS) {
            context.commit('ChangeProgressMS', ProgressMS);
        },
        ChangeProgressMsSpotify(context, ProgressMS) {
            seekThePosition(ProgressMS);
            context.commit('ChangeProgressMS', ProgressMS);
        },
        ChangeRepeat(context, IsRepeat) {
            repeat(IsRepeat);
            context.commit('ChangeIsRepeat', IsRepeat);
        },
        EnableRepeat(context) {
            repeat('track');
            context.commit('ChangeIsRepeat', 'track');
        },
        DisableRepeat(context) {
            repeat('off');
            context.commit('ChangeIsRepeat', 'off');
        },
        RefreshPlayer(context) {
            getCurrentState().then((res) => {
                context.commit("ChangeVolume", res.device.volume_percent);
                context.commit("ChangeIsShuffle", res.shuffle_state);
                context.commit("ChangeIsRepeat", res.repeat_state);
                context.commit("ChangeProgressMS", res.progress_ms);
                context.commit("ChangeArtistName", res.item.artists[0].name);
                context.commit("ChangeSongName", res.item.name);
                context.commit("ChangeImage", res.item.album.images[0].url);
                context.commit("ChangeIsPlay", null);
                context.commit("ChangeIsPlay", res.is_playing);
                context.commit("ChangeDurationMS", res.item.duration_ms);
            })
        }
    }
})
export default store;