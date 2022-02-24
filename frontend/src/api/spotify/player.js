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

export const shuffle = (state, deviceID) => {
    return spotifyApi.setShuffle(state, {device_id: deviceID});
};

export const skipToNext = () => {
    return spotifyApi.skipToNext();
};

export const skipToPrevious = () => {
    return spotifyApi.skipToPrevious();
};

export const play = (deviceID) => {
    return spotifyApi.play({device_id: deviceID});
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
