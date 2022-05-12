<template>
  <div class="aside-outer-container">
    <!--左侧内容-->
    <el-tree
        :data="listData.data"
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
       console.log(listData.data);
     }
   });
 });

 // 加载数据库中的数据
 function onSelect(selectedKeys, info) {
   if (selectedKeys.length === 0) {
     //表示是取消选中
     return;
   }
   let parent = info.node.parent;
   if (parent != undefined) {
     //当前为子节点,改变到右侧的页面中显示数据
     //redis,localhost,1
     let parentKey = info.node.parent.key;
     console.log("this is info ", info);
     console.log("this is parentKey ", parentKey);
     let parentKeyArr = parentKey.split(",");
     let connType = parentKeyArr[0];
     let connName = parentKeyArr[1];
     //selectKey是显示具体的节点key，parent显示的是父节点
     window.go.main.App.GetNodeData(
         connType,
         connName,
         selectedKeys[0] + ""
     ).then((resolve) => {
       if (resolve !== "") {
         // 如果返回值中不为空字符串才进行操作
         let data = JSON.parse(resolve);
         let list = info.node.parent.node.children.map((item) => {
           if (item.key == selectedKeys) {
             item.children = data;
           }
           return item;
         });
         info.node.parent.node.children = [...list];
       } else {
         let key = selectedKeys[0];
         console.log("key", key);
         let rootKey = info.node.parent.nodes[0].key;
         let dbKey = info.node.parent.nodes[1].key;
         let rootKeyArr = rootKey.split(",");
         console.log("info", info.node.parent);
         //此处是属于根节点，右侧显示具体的数据
         if (key.indexOf("*") === -1) {
           // 如果不包含*，则表示右侧的页面要渲染
           router.push({
             path: "/rightContent/value",
             query: {
               connType: rootKeyArr[0],
               connName: rootKeyArr[1],
               dbKey: dbKey,
               key: selectedKeys[0],
             },
           });
         }
       }
     });
   } else {
     //当前为根节点
     let key = selectedKeys[0];
     let keys = key.split(",");
     //表示选中一个根节点，通常是选中了一个数据库
     router.push({
       path: "/rightContent/status",
       query: {
         key: keys[1],
       },
     });
   }
 }

</script>

<style scoped>

</style>