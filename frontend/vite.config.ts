import path from "node:path";
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
import { tanstackRouter } from "@tanstack/router-plugin/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    tanstackRouter({
      target: "react",
      autoCodeSplitting: true,
    }),
    tailwindcss(),
    react({
      babel: {
        plugins: [["babel-plugin-react-compiler"]],
      },
    }),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@wails": path.resolve(__dirname, "./wailsjs"),
      "@modules": path.resolve(__dirname, "./src/modules"),
      "@employees": path.resolve(__dirname, "./src/modules/employees"),
      "@teams": path.resolve(__dirname, "./src/modules/teams"),
      "@iam": path.resolve(__dirname, "./src/modules/iam"),
      "@catalog": path.resolve(__dirname, "./src/modules/catalog"),
      "@shared": path.resolve(__dirname, "./src/modules/shared"),
      "@organizations": path.resolve(__dirname, "./src/modules/organizations"),
    },
  },
});
