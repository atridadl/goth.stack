import { defineConfig } from "astro/config";
import mdx from "@astrojs/mdx";
import sitemap from "@astrojs/sitemap";
import tailwind from "@astrojs/tailwind";
import node from "@astrojs/node";

import react from "@astrojs/react";

// https://astro.build/config
export default defineConfig({
  site: "https://atri.dad/",
  output: "server",
  adapter: node({
    mode: "standalone",
  }),
  server: {
    host: "0.0.0.0",
    port: process.env.PORT || 3000,
  },
  integrations: [mdx(), sitemap(), tailwind(), react()],
  vite: {
    ssr: {
      noExternal: ["react-icons"],
    },
  },
});
