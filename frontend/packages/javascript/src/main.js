/*
 * @Author: symbol
 * @Date: 2022-04-30 11:18:03
 * @LastEditors: symbol
 * @LastEditTime: 2022-04-30 16:34:05
 * @FilePath: /todb/frontend/JS/src/main.js
 * @Description:
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router";
import i18n from "@/i18n";
import Antd from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
import { VueClipboard } from "@soerenmartius/vue3-clipboard";

// Register global common components
// 注册全局通用组件
import publicComponents from "@/components/public";

createApp(App)
  .use(router)
  .use(i18n)
  .use(publicComponents)
  .use(Antd)
  .use(VueClipboard)
  .mount("#app");
