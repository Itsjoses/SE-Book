/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors:{
        'text': '#0F0D11',
        'background': '#FBFAFD',
        'background-second': '#E1E4E9',
        'primary': '#2F27CE',
        'secondary': '#DEDCFF',
        'accent': '#433BFF'

      }
    },
  },
  plugins: [],
}