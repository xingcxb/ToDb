<!--
 * @Author: symbol
 * @Date: 2022-04-30 11:18:03
 * @LastEditors: symbol
 * @LastEditTime: 2022-04-30 17:53:49
 * @FilePath: /todb/frontend/JS/src/views/redis/Content_welcome.vue
 * @Description: 
 * 
 * Copyright (c) 2022 by symbol, All Rights Reserved. 
-->
<template>
  <el-row
    style="min-height: 600px; height: 100%"
    justify="space-around"
    align="middle"
    type="flex"
  >
    <el-col :span="12" :push="8">
      <img src="../../assets/logos/logo.png" alt="ToDB" />
    </el-col>
    <el-col :span="12">
      <span class="app_name">ToDB </span><span class="tail">.app</span>
      <br />
      <span class="other">Version 0.1.0</span><br />
      <span class="other"
        >Powered by
        <span
          class="open_source"
          @click="openBrowser('https://github.com/wailsapp/wails')"
          >wails</span
        >
        and
        <span class="open_source" @click="openBrowser('https://icons8.com')"
          >Icons8</span
        >
      </span>
    </el-col>
  </el-row>
</template>

<script setup>
function openBrowser(v) {
  runtime.BrowserOpenURL(v);
}
</script>

<style scoped>
.app_name {
  font-size: 20px;
}

.tail {
  font-size: 6px;
}

.other {
  font-size: 12px;
}

.open_source {
  font-size: 12px;
  color: #1565c0;
  cursor: pointer;
}
</style>
