import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
    },
  },
  build: {
    outDir: "dist",
    rollupOptions: {
      output: {
        entryFileNames: `src/assets/[name].js`,
        chunkFileNames: `src/assets/[name].js`,
        assetFileNames: `src/assets/[name].[ext]`,
      },
    },
  },
});
