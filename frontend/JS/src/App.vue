<template>
  <!--header-->
  <!--头部-->
  <div class="header" data-wails-drag>
    <!--导航-->
    <a-row style="height: 66px">
      <a-col :span="4">
        <a-dropdown :trigger="['click']" :placement="top">
          <a class="ant-dropdown-link" style="text-align: center;display: block" @click.prevent>
            <img src="./assets/images/quick/conn.png" class="quickIcon" alt=""/>
            <p style="font-size:10px">连接</p>
          </a>
          <template #overlay>
            <a-menu>
              <a-menu-item key="0">
<!--                <a href="http://www.alipay.com/">Redis</a>-->
                <router-link to="/newConnection">Redis</router-link>
              </a-menu-item>
              <a-menu-item key="1">
                <a href="http://www.taobao.com/">MySQL</a>
              </a-menu-item>
              <a-menu-divider/>
              <a-menu-item key="3">other</a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </a-col>
      <a-col :span="4">
        <a href="" style="text-align: center;display: block">
          <img src="./assets/images/quick/table.png" class="quickIcon" alt=""/>
          <p style="font-size:10px;text-align: center">表</p>
        </a>
      </a-col>
      <a-col :span="4">
        <a href="" style="text-align: center;display: block">
          <img src="./assets/images/quick/select.png" class="quickIcon" alt=""/>
          <p style="font-size:10px;text-align: center">查询</p>
        </a>
      </a-col>
      <a-col :span="4">
        <a href="" style="text-align: center;display: block">
          <img src="./assets/images/quick/import.png" class="quickIcon" alt=""/>
          <p style="font-size:10px;text-align: center">导入</p>
        </a>
      </a-col>
      <a-col :span="4">
        <a href="" style="text-align: center;display: block">
          <img src="./assets/images/quick/export.png" class="quickIcon" alt=""/>
          <p style="font-size:10px;text-align: center">导出</p>
        </a>
      </a-col>
      <a-col :span="4">
        <!--广告-->
      </a-col>
    </a-row>
  </div>
  <router-view/>
</template>

<script setup>

</script>
<style>
.header {
  height: 64px;
  width: 100%;
  background: #f0efee;
}

.quickIcon {
  padding: 0;
  text-align: center;
  height: 44px;
  width: auto
}

</style>
