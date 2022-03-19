module.exports = {
  mode: 'jit',
  purge: ['./public/**/*.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      boxShadow: {
        '4xl': '0 8px 24px rgb(0 0 0 / 50%)',
      },
      colors: {
        SpotifyGreen: '#1db954',
        SpotifyBackround: '#121212',
        SpotifyPlayer: '#181818',
        SpotifyBlack: '#000000'
      }
    },
  },
  variants: {
    extend: {
      translate: ({after}) => after(['group-hover'])
    },
  },
  plugins: [require("@tailwindcss/line-clamp")],
}
