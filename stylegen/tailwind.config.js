/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["../pages/**/*.{html,go}"],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["night"],
  },
  plugins: [require("daisyui"), require('@tailwindcss/typography')],
}

