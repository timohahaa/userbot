/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{vue,js,ts,jsx,tsx}', './index.html'],
	theme: {
		extend: {
			colors: {
				'dark-main': '#17181c',
				'light-main': '#e4e4e4',
				'text-color': '#fafafa',
				// border: '#aaaaaa',
				'table-bg-color': '#23232c',
				'button-text-color': '#1ad4cb',
				'button-color': '#22333b',
				'checkbox-color': '#383841',
			},
			text: {
				
			}
		},
		container: {
			center: true,
		},
	},
	plugins: [],
}
