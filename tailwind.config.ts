import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/**/*.{astro,html,js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        "primary": "#1F2937",
        "secondary": "#F9FAFB",
        "accent": "#34D399",
        "accent-secondary": "#FB923C",
        "font-color": "#F9FAFB",
        "font-color-dark": "#1F2937",
      },
    },
  },
} satisfies Config;
