/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/*.{html,js}", "./views/templates/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        primary: "",
        dark_secondary: "#374151",
        ligth_secondary: "#f9fafb",
        background_color: ""
      }
    },
  },
  plugins: [],
}

