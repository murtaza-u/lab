/** @type {import('tailwindcss').Config} */
export default {
  content: ["./view/**/*.templ"], // this is where our templates are located
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["emerald", "forest"],
  },
  plugins: [require("daisyui")],
};
