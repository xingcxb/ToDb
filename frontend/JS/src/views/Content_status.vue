<template>
  <br/>
  <!--自动刷新-->
  <a-row style="width: 100%;margin-right: 2px">
    <a-col :span="24">
      <div style="float:right;">
        <a-tag type="card">
          <sync-outlined spin/>
          自动刷新
        </a-tag>
        <a-tooltip placement="bottom">
          <template #title>
            <span>自动刷新开关，每5秒刷新一次</span>
          </template>
          <a-switch @click="changeAutoRefresh" v-model:checked="autoRefresh" checked-children="开"
                    un-checked-children="关"></a-switch>
        </a-tooltip>
      </div>
    </a-col>
  </a-row>
  <br/>
  <!--基础信息-->
  <a-row :gutter="16" justify="space-around" style="margin-left: 2px">
    <!-- server status row -->
    <a-col :span="8" class="gutter-row">
      <a-card>
        <template #title>
          <img src="../../public/status/service.png" alt="service" style="width: 40px; height: 40px;">
          <span>服务器</span>
        </template>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info">
            Redis版本：
            <span class="server-status-text">{{serviceInfo.data.version}}</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            OS：
            <span class="server-status-text">{{serviceInfo.data.os}}</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            进程ID：
            <span class="server-status-text">{{serviceInfo.data.process}}</span>
          </a-tag>
        </p>
      </a-card>
    </a-col>
    <!-- memory row -->
    <a-col :span="8" class="gutter-row">
      <a-card>
        <template #title>
          <img src="../../public/status/memory.png" alt="service" style="width: 40px; height: 40px;">
          <span>内存</span>
        </template>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            已用内存：
            <span class="server-status-text">{{memoryInfo.data.usedMemory}}</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            内存占用峰值：
            <span class="server-status-text">{{memoryInfo.data.usedBigMemory}}</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            Lua占用内存：
            <span class="server-status-text">{{memoryInfo.data.luaMemory}}</span>
          </a-tag>
        </p>
      </a-card>
    </a-col>
    <a-col :span="8" class="gutter-row">
      <a-card>
        <template #title>
          <img src="../../public/status/history.png" alt="service" style="width: 40px; height: 40px;">
          <span>历史</span>
        </template>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            客户端连接数：
            <span class="server-status-text">{{historyInfo.data.connectCount}}</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            历史连接数：
            <span class="server-status-text">{{historyInfo.data.historyCount}}</span>
          </a-tag>
        </p>
        <p class="server-status-tag-p">
          <a-tag class='server-status-container' type="info" size="big">
            历史命令数：
            <span class="server-status-text">{{historyInfo.data.historyInstructions}}</span>
          </a-tag>
        </p>
      </a-card>
    </a-col>
  </a-row>
  <br/>
  <!--键值信息-->
  <a-row :gutter="16" style="margin-left: 2px">
    <a-col>
      <a-card>
        <template #title>
          <img src="../../public/status/kv.png" alt="service" style="width: 40px; height: 40px;">
          <span>键值统计</span>
        </template>
        <a-table :dataSource="kvInfo" :columns="columns" :loading="loading">

        </a-table>
      </a-card>
    </a-col>
  </a-row>
</template>

<script setup>
import {SyncOutlined} from '@ant-design/icons-vue';
import {onBeforeMount, reactive, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();
//是否刷新
let autoRefresh = ref(false);
//定时刷新时间 5s
const refreshTimer = ref(5);
// 获取redis所在服务器信息
let serviceInfo = reactive({
  data:[]
});
// 获取redis占用内存信息
let memoryInfo = reactive({
  data:[]
})
// 获取redis历史信息
let historyInfo = reactive({
  data:[]
})
// 键值信息
let kvInfo = "";

onBeforeMount(() => {
  console.log(router.currentRoute.value.query.data)
  // 获取基本数据
  window.go.main.App.LoadingDbResource(router.currentRoute.value.query.data).then((resolve) => {
    if (resolve !== "") {
      // 如果返回值中不为空字符串才进行操作
      console.log(resolve)
      let _json = JSON.parse(resolve)
      // 获取服务信息
      serviceInfo.data = _json.server;
      memoryInfo.data = _json.memory;
      historyInfo.data = _json.start;
    }
  });
})

function changeAutoRefresh() {
  console.log("----------",autoRefresh)
  autoRefresh = !autoRefresh
  console.log("----------",autoRefresh)
  if (autoRefresh) {
    // 开启自动刷新
    // 开启定时器
    setInterval(() => {
      // 刷新
      window.go.main.App.LoadingDbResource(router.currentRoute.value.query.data).then((resolve) => {
        if (resolve !== "") {
          // 如果返回值中不为空字符串才进行操作
          console.log(resolve)
          let _json = JSON.parse(resolve)
          // 获取服务信息
          serviceInfo.data = _json.server;
          memoryInfo.data = _json.memory;
          historyInfo.data = _json.start;
        }
      });
    }, refreshTimer.value * 1000);
  }
}
</script>

<style scoped>
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
</style>