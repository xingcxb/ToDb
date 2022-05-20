/*
 * @Author: symbol
 * @Date: 2022-04-30 11:18:03
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-20 16:09:11
 * @FilePath: /todb/frontend/src/main.js
 * @Description:
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import i18n from "./i18n";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import { VueClipboard } from "@soerenmartius/vue3-clipboard";

createApp(App)
  .use(router)
  .use(i18n)
  .use(ElementPlus)
  .use(VueClipboard)
  .mount("#app");
