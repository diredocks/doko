import { solidStart } from "@solidjs/start/config";
import { nitroV2Plugin as nitro } from "@solidjs/vite-plugin-nitro-2";
import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";

export default defineConfig({
	clearScreen: false,
	plugins: [
		solidStart({
			ssr: false,
		}),
		nitro(),
		tailwindcss(),
	],
});
