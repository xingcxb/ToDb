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
import {onBeforeMount, reactive} from "vue";
import router from "./router";

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


// 【第二层】加载连接信息下的表/库
function loadNode(node, resolve) {
  console.log("node--------",node);
  let nodeData = node.data;
  if (node.level == 1) {
    // 表示只处理一级节点从基础开始加载
    // 当选中第一个库的时候需要将状态信息加载出来
    router.push({
      path:"/rightCount/status",
      query:{
        key: nodeData.label
      }
    })
    window.go.main.App.LoadingConnInfo(nodeData.connType, nodeData.label).then((resp) => {
      if (resp !== "") {
        setTimeout(resolve(JSON.parse(resp)), 500)
      } else {
        // 如果返回的数据不存在值，标记为没有子节点
        setTimeout(resolve([]), 500)
      }
    });
  } else if (node.level == 2) {
    console.log("node:", node);
    // 表示进入具体的库，需要加载key
    // 获取父节点的数据
    let parentNode = node.parent.data;
    window.go.main.App.GetNodeData(parentNode.connType,parentNode.label,node.key).then((resp) => {
      if (resp !== "") {
        console.log("res+++++:",resp)
        setTimeout(resolve(JSON.parse(resp)), 500)
        console.log("附加后的node:", node);
        console.log("listData:",listData.data)
      } else {
        // 如果返回的数据不存在值，标记为没有子节点
        setTimeout(resolve([]), 500)
      }
    });
  }else{
    // if(node.data.children.length > 0) {
    //   resolve(node.data.children)
    // }
    console.log("level=======", node.level);
  }
}
</script>

<style scoped></style>
