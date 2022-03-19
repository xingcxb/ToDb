<template>
  <a-layout class="layout">
    <a-layout-header :style="{ background: '#ffffff' }">
      <div class="logo"/>
<!--      <a-menu-->
<!--          theme="light"-->
<!--          mode="horizontal"-->
<!--          :style="{ lineHeight: '64px', background: '#eff0f0'}"-->
<!--      >-->
<!--        <a-sub-menu key="conn">-->
<!--&lt;!&ndash;          <image src="assets/header/conn.png"></image>&ndash;&gt;-->
<!--          连接-->
<!--        </a-sub-menu>-->
<!--        <a-menu-item key="2">nav 2</a-menu-item>-->
<!--        <a-menu-item key="3">nav 3</a-menu-item>-->
<!--      </a-menu>-->
      <a-menu v-model:selectedKeys="current" mode="horizontal">
        <a-sub-menu key="">
<!--          <template #icon>-->
<!--            <img src="assets/header/conn.png" alt="">-->
<!--          </template>-->
          <template #title>连接</template>
            <a-menu-item key="setting:1">Option 1</a-menu-item>
            <a-menu-item key="setting:2">Option 2</a-menu-item>
        </a-sub-menu>
        <a-menu-item key="">
          <a href="" target="_blank">
            123123
          </a>
        </a-menu-item>
      </a-menu>
    </a-layout-header>
    <a-layout-content style="padding: 0 50px" :style="{background: '#eff0f0'}">
      <div :style="{ background: '#eff0f0', padding: '24px', minHeight: '280px' }">Content</div>
    </a-layout-content>
    <a-layout-footer style="text-align: center">
      Ant Design ©2018 Created by Ant UED
    </a-layout-footer>
  </a-layout>
</template>

<script setup>

</script>
<style>
/*.site-layout-content {*/
/*  min-height: 280px;*/
/*  padding: 24px;*/
/*  background: #fff;*/
/*}*/
.logo {
  float: left;
  width: 57px;
  height: 57px;
  margin: 4px 0 3px 0;
  background: url("assets/logo.png");
}
[data-theme='light'] .site-layout-content {
  background: #eff0f0;
}
</style>
