import {getMe} from '@/api/spotify/player.js';

export default {
    setDisplayName(context, displayName) {
      context.commit("setDisplayName", displayName)
    },
    setUserImage(context, image) {
        context.commit("setUserImage", image)
    },
    setSpotifyID(context, spotifyID) {
        context.commit("setSpotifyID", spotifyID)
    },
    setCountry(context, country) {
        context.commit("setCountry", country)
    },
    refreshUser(context) {
        getMe().then((res) => {
            context.commit("setDisplayName", res.display_name)
            context.commit("setUserImage", res.images[0].url)
            context.commit("setSpotifyID", res.id)
            context.commit("setCountry", res.country)
        })
    }
};