<template>
  <a-row style="width: 100%">
    <a-col :span="24">
      <div style="float:right;">
        <a-tag type="card">
          <sync-outlined spin/>
          自动刷新
        </a-tag>
        <a-tooltip placement="bottom">
          <template #title>
            <span>自动刷新开关，每2秒刷新一次</span>
          </template>
          <a-switch @click="changeAutoRefresh" v-model:checked="autoRefresh" checked-children="开" un-checked-children="关"></a-switch>
        </a-tooltip>
      </div>
    </a-col>
  </a-row>
  <a-row :guttes="10" class="status-container status-card">
    <!-- server status row -->
  <a-col :span="8">
    <a-card class="box-card">
      <template #title>
        <img src="../../public/status/service.png" alt="service" style="width: 40px; height: 40px;">
        <span>服务器</span>
      </template>
      <p class="server-status-tag-p">
        <a-tag class='server-status-container' type="info" size="big">
          Redis版本：
          <span class="server-status-text">6.0.1</span>
        </a-tag>
      </p>
      <p class="server-status-tag-p">
        <a-tag class='server-status-container' type="info" size="big">
          OS：
          <span class="server-status-text" >macOs</span>
        </a-tag>
      </p>
      <p class="server-status-tag-p">
        <a-tag class='server-status-container' type="info" size="big">
          进程ID：
          <span class="server-status-text">12203</span>
        </a-tag>
      </p>
    </a-card>
  </a-col>
    <!-- memory row -->
    <a-col :span="8">
      <a-card class="box-card">
        <template #title>
          <img src="../../public/status/memory.png" alt="service" style="width: 40px; height: 40px;">
          <span>内存</span>
        </template>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            Redis版本：
            <span class="server-status-text">6.0.1</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            OS：
            <span class="server-status-text" >macOs</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            进程ID：
            <span class="server-status-text">12203</span>
          </a-tag>
        </p>
      </a-card>
    </a-col>
    <a-col :span="8">
      <a-card class="box-card">
        <template #title>
          <img src="../../public/status/history.png" alt="service" style="width: 40px; height: 40px;">
          <span>服务器</span>
        </template>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            Redis版本：
            <span class="server-status-text">6.0.1</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            OS：
            <span class="server-status-text" >macOs</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            进程ID：
            <span class="server-status-text">12203</span>
          </a-tag>
        </p>
      </a-card>
    </a-col>
  </a-row>
</template>

<script setup>
import {SyncOutlined} from '@ant-design/icons-vue';
import {onBeforeMount, reactive, ref, toRef} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();
//是否刷新
let autoRefresh = ref(false);
//定时刷新时间 5s
const refreshTimer = ref(5);
// 连接状态
let connectionStatus = "";
// 状态连接
let statusConnection = "";
//所有信息过滤
let allInfoFilter = "";

onBeforeMount(() => {
  console.log(router.currentRoute.value.query.data)
  window.go.main.App.LoadingDbResource(router.currentRoute.value.query.data).then((resolve) => {
    if (resolve !== "") {
      // 如果返回值中不为空字符串才进行操作
      console.log(resolve)
    }
  });
})

function changeAutoRefresh() {
  autoRefresh.value = !autoRefresh.value
  console.log(autoRefresh.value)
  if (autoRefresh) {
    // 开启自动刷新
    // 开启定时器
    setInterval(() => {
      // 刷新
      window.go.main.App.LoadingDbResource(router.currentRoute.value.query.data).then((resolve) => {
        if (resolve !== "") {
          // 如果返回值中不为空字符串才进行操作
          console.log(resolve)
        }
      });
    }, refreshTimer.value * 1000);
  }
}
</script>

<style scoped>
.status-card{
  margin-top: 20px;
  margin-left: 20px;
}
.server-status-tag-p {
  height: 32px;
}
.server-status-container{
  width: 100%;
  overflow-x: hidden;
  text-overflow: ellipsis;
}
.server-status-text{
  color: #43b50b;
}
.status-filter-input {
  float: right;
  width: 100px;
}
/*fix table height changes[scrollTop changes] when tab toggled*/
.status-card {
  height: 50px;
}
.status-card{
  /*height: calc(100% - 50px) !important;*/
}
</style>