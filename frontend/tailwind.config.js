module.exports = {
  mode: 'jit',
  purge: ['./public/**/*.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      colors: {
        SpotifyGreen: '#1db954',
        SpotifyBackround: '#121212',
        SpotifyPlayer: '#181818',
        SpotifyBlack: '#000000'
      }
    },
  },
  variants: {
    extend: {},
  },
  plugins: [require("@tailwindcss/line-clamp")],
}
