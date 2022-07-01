export default {
    setAccessToken(context, accessToken) {
        context.commit("setAccessToken", accessToken)
    },
    setRefreshToken(context, refreshToken) {
        context.commit("setRefreshToken", refreshToken)
    },
    setRadiofyToken(context, refreshToken) {
        context.commit("setRadiofyToken", refreshToken)
    }
}