/** @type {import('tailwindcss').Config} */
module.exports = {
    relative: true,
    content: [
        './static/*.js',
        './templates/*.gohtml'
    ],
    theme: {
        fontFamily: {
            'roboto': ['Roboto', 'sans-serif']
        },
        extend: {},
    },
    plugins: [],
}

