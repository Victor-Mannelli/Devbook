/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/*.{html,js}", "./views/templates/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        dark_bg: "#374151",
        dark_secondary_bg: "#6b7280",
        light_bg: "#fff",
        light_secondary_bg: "#c3ccdb",
      }
    },
  },
  plugins: [],
}

