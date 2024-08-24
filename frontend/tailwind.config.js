/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/**/*.{html,ts,tsx}"],
    theme: {
        container: {
            center: true,
        },
        aspectRatio: {
            auto: "auto",
        },
        extend: {
            colors: {
                'primary': '#ff49db',
            },
        },
    },
    plugins: [],
}