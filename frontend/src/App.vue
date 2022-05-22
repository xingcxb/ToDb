<!--
 * @Author: symbol
 * @Date: 2022-05-18 08:23:08
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-18 18:20:59
 * @FilePath: /todb/frontend/src/App.vue
 * @Description: 
 * 
 * Copyright (c) 2022 by symbol, All Rights Reserved. 
-->
<template>
  <div>
    <!--  <div class="header" data-wails-drag>-->
    <el-container>
      <el-header
        height="69px"
        style="background: #ecedee; border-bottom: 1px solid #bebebe"
      >
        <!--头部 导航-->
        <Header></Header>
      </el-header>
      <!--头部下面的属性-->
      <el-container>
        <el-aside>
          <el-scrollbar style="height: 600px">
            <!--左侧导航-->
            <Aside class="aside-outer-container"></Aside>
          </el-scrollbar>
        </el-aside>
        <el-main>
<!--          <el-scrollbar style="height: 600px">-->
            <!--右侧内容-->
            <div class="connections-list">
              <router-view :key="$route.path + Date.now()" />
            </div>
<!--          </el-scrollbar>-->
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import Header from "./Header.vue";
import Aside from "./Aside.vue";
import { onBeforeMount, ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

let sideWidth = ref(265);

onBeforeMount(() => {
  router.push({
    path: "/rightContent/default",
  });
});
</script>
<style>
body {
  margin: 0 0;
  padding: 0 0;
}
</style>
<style scoped>
/deep/.el-card {
  background-color: rgb(236, 237, 238);
}

body {
  overscroll-behavior: none;
  margin: 0 0;
}

.aside-drag-container {
  position: relative;
  user-select: none;
  /*max-width: 50%;*/
}

.aside-outer-container .connections-list {
  overflow-y: auto;
  width: 98%;
  height: calc(100vh - 54px);
  margin-top: 4px;
  /*border-top: 1px solid #dbdada;*/
}
</style>
