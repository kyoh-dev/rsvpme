// @ts-check
import { defineConfig, envField } from "astro/config";

import node from "@astrojs/node";
import tailwindcss from "@tailwindcss/vite";
import tsconfigPaths from "vite-tsconfig-paths";

import preact from "@astrojs/preact";

// https://astro.build/config
export default defineConfig({
  env: {
    schema: {
      DATABASE_URL: envField.string({
        context: "server",
        access: "public",
        optional: false,
      }),
      DATABASE_AUTH_TOKEN: envField.string({
        context: "server",
        access: "secret",
        optional: false,
      }),
    },
  },

  vite: {
    plugins: [tailwindcss(), tsconfigPaths()],
  },

  adapter: node({
    mode: "standalone",
  }),

  output: "server",

  integrations: [preact({ devtools: true })],
});
