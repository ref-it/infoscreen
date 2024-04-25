/** @type {import('tailwindcss').Config} */

const defaultTheme = require('tailwindcss/defaultTheme');

module.exports = {
    content: ["./src/**/*.vue", "./index.html"],
    theme: {
        fontSize: {
            sm: '1rem',
            base: '1.2rem',
            xl: '1.45rem',
            '2xl': '1.763rem',
            '3xl': '2.153rem',
            '4xl': '2.641rem',
            '5xl': '3.252rem',
        },
        extend: {
            fontFamily: {
                sans: ["PT Sans", ...defaultTheme.fontFamily.sans],
            },
        },
    },
    plugins: [
        require('@tailwindcss/forms'),
    ],
}

