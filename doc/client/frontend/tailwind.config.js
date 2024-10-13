/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/**/*.{js,jsx,ts,tsx}",
    "../main.js", // Include Electron's main process if needed
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
