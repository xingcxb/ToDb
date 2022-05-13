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
 import {onBeforeMount,reactive} from "vue";

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
   console.log("nodeData",node.data)
   console.log("resolve",resolve)
   let nodeData = node.data;
   if(node.level == 1){
     // 表示从基础开始加载
     // window.go.main.App.GetNodeData(nodeData.connType,nodeData.)
   }
 }


</script>

<style scoped>

</style>