<!--
 * @Author: symbol
 * @Date: 2022-05-09 22:20:07
 * @LastEditors: symbol
 * @LastEditTime: 2022-05-22 14:39:26
 * @FilePath: /todb/frontend/src/Aside.vue
 * @Description: 
 * 
 * Copyright (c) 2022 by symbol, All Rights Reserved. 
-->
<template>
  <div>
    <!--左侧内容-->
    <el-tree
      :data="listData.data"
      :load="loadNode"
      highlight-current
      accordion
      lazy
      style="background: #e0e1e1"
    ></el-tree>
  </div>
</template>

<script setup>
import { onBeforeMount, reactive } from "vue";
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
  let nodeData = node.data;
  if (node.level == 1) {
    // 表示只处理一级节点从基础开始加载
    router.push({
      // 当选中连接数据库时候需要将状态信息加载出来
      path: "/rightContent/status",
      query: {
        fileName: nodeData.label,
      },
    });
    window.go.main.App.LoadingConnInfo(nodeData.connType, nodeData.label).then(
      (resp) => {
        if (resp !== "") {
          setTimeout(resolve(JSON.parse(resp)), 500);
        } else {
          // 如果返回的数据不存在值，标记为没有子节点
          setTimeout(resolve([]), 500);
        }
      }
    );
  } else if (node.level == 2) {
    // console.log("node:", node);
    // 表示进入具体的库，需要加载key
    // 获取父节点的数据
    let parentNode = node.parent.data;
    window.go.main.App.GetNodeData(
      parentNode.connType,
      parentNode.label,
      node.data.key
    ).then((resp) => {
      // 当数据不存在的时候返回的是
      if (resp != "") {
        setTimeout(resolve(JSON.parse(resp)), 500);
      } else {
        // 如果返回的数据不存在值，标记为没有子节点
        setTimeout(resolve([]), 500);
      }
    });
  } else {
    if (node.level > 2) {
      let arr = [];
      if (nodeData.children && nodeData.children.length > 0) {
        arr = nodeData.children;
      } else {
        // 这里是不存在子节点，将右边进行改变
        // 获取到选中的顶级父类节点
        let topParentNode = node.parent;
        let nextParentNode = node.parent;
        for (let i = 1; i < node.level - 1; i++) {
          topParentNode = topParentNode.parent;
          if (i === node.level - 3) {
            nextParentNode = topParentNode;
          }
        }
        window.go.main.App.ChangeRightWindowStyle(
          JSON.stringify(topParentNode.data),
          JSON.stringify(nextParentNode.data),
          JSON.stringify(nodeData)
        ).then((resp) => {
          let fullStr = nodeData.fullStr;
          let dbId = nextParentNode.data.key;
          let connType = topParentNode.data.connType;
          let connName = topParentNode.data.title;
          console.log("类型为：", resp);
          let path = "";
          switch (resp) {
            case "string":
              // 字符串类型
              path = "/rightContent/value_string";
              break;
            case "list":
              // list类型
              path = "/rightContent/value_list";
              break;
            case "hash":
              // hash类型
              path = "/rightContent/value_hash";
              break;
            case "set":
              // set类型
              path = "/rightContent/value_set";
              break;
            case "stream":
              // stream类型
              path = "/rightContent/value_stream";
              break;
            case "zset":
              // zset类型
              path = "/rightContent/value_zset";
              break;
            default:
              // 其他的返回到默认的页面
              path = "/rightContent/default";
              break;
          }
          router.push({
            path: path,
            query: {
              key: fullStr,
              dbId: dbId,
              connType: connType,
              connName: connName,
            },
          });
          setTimeout(resolve([]), 500);
        });
      }
      setTimeout(resolve([...arr]), 500);
    }
  }
}
</script>

<style scoped>
.interval_row {
  margin-top: 10px;
}
</style>
