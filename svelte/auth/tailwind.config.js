/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			fontFamily: {
				sans: ['Roboto', 'sans-serif'],
				mono: ['Ubuntu Mono', 'monospace']
			}
		}
	},
	plugins: [require('daisyui')],
	daisyui: {
		themes: [
			{
				emerald: {
					...require('daisyui/src/theming/themes')['emerald'],
					'primary-content': '#fff'
				}
			},
			{
				forest: {
					...require('daisyui/src/theming/themes')['forest'],
					'--rounded-btn': '.5rem'
				}
			}
		]
	}
}
