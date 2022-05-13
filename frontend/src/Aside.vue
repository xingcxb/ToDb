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

 // 加载数据库中的数据,使用懒加载的方案
 function loadNode(node, resolve) {
   console.log("nodeLevel", node.level);
   console.log("selectedKeys",node)
   console.log("resolve",resolve)
   if(node.level == 1){
     // 表示从基础开始加载
   }
 }


</script>

<style scoped>

</style>