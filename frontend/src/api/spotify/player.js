import SpotifyWebApi from "spotify-web-api-js";
const spotifyApi = new SpotifyWebApi();

export const setAccessToken = (token,) => {
    return spotifyApi.setAccessToken(token, );
};

export const setVolume = (volume) => {
    return spotifyApi.setVolume(volume,);
};

export const repeat = (state) => {
    return spotifyApi.setRepeat(state);
};

export const shuffle = (state) => {
    return spotifyApi.setShuffle(state);
};

export const skipToNext = () => {
    return spotifyApi.skipToNext();
};

export const skipToPrevious = () => {
    return spotifyApi.skipToPrevious();
};

export const play = () => {
    return spotifyApi.play();
};

export const pause = () => {
    return spotifyApi.pause();
};

export const getMyCurrentPlayingTrack = () => {
    return spotifyApi.getMyCurrentPlayingTrack();
};

export const getMe = () => {
    return spotifyApi.getMe();
};

export const getCurrentState = () => {
    return spotifyApi.getMyCurrentPlaybackState();
};

export const seekThePosition = (position) => {
    return spotifyApi.seek(position);
};

export const searchItem = (query, limit) => {
    return spotifyApi.searchTracks(query, {limit: limit})
}
