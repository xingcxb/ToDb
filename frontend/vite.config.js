/*
 * @Author: symbol
 * @Date: 2022-05-08 08:49:54
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-18 11:20:09
 * @FilePath: /todb/frontend/vite.config.js
 * @Description:
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
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
