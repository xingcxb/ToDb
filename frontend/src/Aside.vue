<template>
  <div class="aside-outer-container">
    <!--左侧内容-->
    <el-tree-v2 :data="listData.data"></el-tree-v2>
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


</script>

<style scoped>

</style>