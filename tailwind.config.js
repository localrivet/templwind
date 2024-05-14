const defaultTheme = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./**/*.{html,js,templ,txt}"
  ],
  theme: {
    defaultTheme
  },
  plugins: [
    require("@tailwindcss/typography")
  ]
}
