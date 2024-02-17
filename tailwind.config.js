/** @type {import('tailwind').Config} **/
module.exports = {
  content: ["internal/templates/**/*.{html,js,templ,go}"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Inter", "sans-serif"],
      },
      colors: {
        primary: {
          100: "#E6F6FF",
          200: "#B3E0FF",
          300: "#80CAFF",
          400: "#4DB4FF",
          500: "#1A9EFF",
          600: "#007ACC",
          700: "#005C99",
          800: "#004066",
          900: "#001C33",
        },
        secondary: {
          100: "#FFF5E6",
          200: "#FFE0B3",
          300: "#FFCA80",
          400: "#FFB94D",
          500: "#FFA61A",
          600: "#CC8400",
          700: "#996300",
          800: "#664200",
          900: "#332100",
        },
      },
      plugins: [
        require("@tailwindcss/forms"),
        require("@tailwindcss/typography"),
        require("autoprefixer"),
      ],
    },
  },
};
