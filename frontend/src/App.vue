<template>
  <div>
    <!--  <div class="header" data-wails-drag>-->
    <el-container>
      <el-header height="69px">
        <!--头部 导航-->
        <Header></Header>
      </el-header>
      <!--头部下面的属性-->
      <el-container>
        <el-aside style="background: #e0e1e1;">
          <el-scrollbar style="height: 659px">
            <!--左侧导航-->
            <Aside class="aside-outer-container"></Aside>
          </el-scrollbar>
        </el-aside>
        <el-main>
          <el-scrollbar style="height: 659px">
            <!--右侧内容-->
            <div class="connections-list">
              <router-view :key="$route.path + Date.now()"/>
            </div>
          </el-scrollbar>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import Header from './Header.vue';
import Aside from "./Aside.vue";
import {onBeforeMount, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();

let sideWidth = ref(265);

onBeforeMount(() => {
  router.push({
    path: "/rightContent/default",
  });
});


</script>

<style scoped>
body {
  overscroll-behavior: none;
}

.aside-drag-container {
  position: relative;
  user-select: none;
  /*max-width: 50%;*/
}

.aside-outer-container .connections-list {
  overflow-y: auto;
  height: calc(100vh - 54px);
  margin-top: 4px;
  /*border-top: 1px solid #dbdada;*/
}
</style>