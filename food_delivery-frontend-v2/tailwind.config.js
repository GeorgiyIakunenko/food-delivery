/** @type {import('tailwindcss').Config} */

module.exports = {
    purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
    theme: {
        colors: {
            'card-bg': '#F4F5F7',
            'primary': {
                500: '#FF991F',
                400: '#FF991F',
                300: '#FFAB00',
                200: '#FFC400',
                100: '#FFE380',
                75: '#FFF0B3',
                50: '#FFFAE5',
            },
            'neutral': {
                400: '#505F79',
                300: '#5E6C84',
                200: '#6B778C',
                100: '#7A869A',
                90: '#8993A4',
                80: '#97A0AF',
                70: '#A5ADBA',
                60 : '#B3BAC5',
                50 : '#C1C7D0',
                40 : '#DFE1E6',
                30 : '#EBECF0',
                0o1: '#FAFBFC',
                0o0: '#FFFFFF',

            },
            'teal' : {
                500 : '#008DA6',
                400 : '#00A3BF',
                300 : '#00B8D9',
                200 : '#00C7E6',
                100 : '#79E2F2',
                75 : '#B3F5FF',
                50 : '#E6FCFF',
            },
            'purple' : {
                500 : '#403294',
                400 : '#5243AA',
                300 : '#6554C0',
                200 : '#8777D9',
                100 : '#998DD9',
                75 : '#C0B6F2',
                50 : '#EAE6FF',
            },
            'green' : {
                500 : '#064',
                400 : '#00875A',
                300 : '#36B37E',
                200 : '#57D9A3',
                100 : '#79F2C0',
                75 : '#ABF5D1',
                50 : '#E3FCEF',
            },
            'blue' : {
                600 : '#063F94',
                500 : '#0747A6',
                400 : '#0052CC',
                300 : '#0065FF',
                200 : '#2684FF',
                100 : '#4C9AFF',
                75 : '#B3D4FF',
                50 : '#DEEBFF',
            }


        },
        extend: {
            fontFamily: {
                'inter': ['Inter', 'sans-serif'],
                'roboto': ['Roboto', 'sans-serif'],
                'sans': ['DM Sans','sans-serif']
            }
        },
    },
    variants: {
        extend: {},
    },
    plugins: [],
}
