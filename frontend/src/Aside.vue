<!--
 * @Author: symbol
 * @Date: 2022-05-09 22:20:07
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-14 12:09:56
 * @FilePath: \ToDb\frontend\src\Aside.vue
 * @Description: 
 * 
 * Copyright (c) 2022 by symbol, All Rights Reserved. 
-->
<template>
  <div class="aside-outer-container">
    <!--左侧内容-->
    <el-tree
      :data="listData.data"
      :load="loadNode"
      highlight-current
      accordion
      lazy
    ></el-tree>
  </div>
</template>

<script setup>
import { onBeforeMount, reactive } from "vue";

let listData = reactive({
  data: [],
});

onBeforeMount(() => {
  window.go.main.App.LoadingConnKey().then((resolve) => {
    if (resolve !== "") {
      // 如果返回值中不为空字符串才进行操作
      listData.data = JSON.parse(resolve);
    }
  });
});

// 一次性加载所有的数据库中的数据
function loadNode(node, resolve) {
  console.log("nodeLevel", node.level);
  console.log("nodeData", node.data);
  console.log("resolve", resolve);
  let nodeData = node.data;
  if (node.level == 1) {
    // 表示从基础开始加载
    window.go.main.App.LoadingConnInfo(nodeData.connType,nodeData.label).then((resolve)=>{
      console.log("resolve=====",resolve);
      nodeData.children = JSON.parse(resolve);
      console.log("newListData",listData.data);
    });
  }
}
</script>

<style scoped></style>
