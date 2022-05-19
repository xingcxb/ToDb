/*
 * @Author: symbol
 * @Date: 2022-04-30 11:18:03
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-19 14:36:07
 * @FilePath: /todb/frontend/src/router/index.js
 * @Description:路由
 *
 * Copyright (c) 2022 by symbol, All Rights Reserved.
 */
import Default from "../Welcome.vue";
import NewConnect from "../views/NewConnection.vue";
import Status from "../views/redis/Content_status.vue";
import StringInfo from "../views/redis/Content_string_info.vue";
import ListInfo from "../views/redis/Content_list_info.vue";
import HashInfo from "../views/redis/Content_hash_info.vue";
import SetInfo from "../views/redis/Content_set_info.vue";
import StreamInfo from "../views/redis/Content_stream_info.vue";
import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    // 新建连接
    path: "/newConnection",
    name: "newConnection",
    component: NewConnect,
  },
  {
    // 右侧默认页面
    path: "/rightContent/default",
    name: "default",
    component: Default,
  },
  {
    // 右侧状态页面
    path: "/rightContent/status",
    name: "status",
    component: Status,
  },
  {
    // 右侧详情页面string类型
    path: "/rightContent/value_string",
    name: "stringInfo",
    component: StringInfo,
  },
  {
    // 右侧详情页面list类型
    path: "/rightContent/value_list",
    name: "listInfo",
    component: ListInfo,
  },
  {
    // 右侧详情页面hash类型
    path: "/rightContent/value_hash",
    name: "hashInfo",
    component: HashInfo,
  },
  {
    // 右侧详情页面set类型
    path: "/rightContent/value_set",
    name: "setInfo",
    component: SetInfo,
  },
  {
    // 右侧详情页面stream类型
    path: "/rightContent/value_stream",
    name: "streamInfo",
    component: StreamInfo,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
