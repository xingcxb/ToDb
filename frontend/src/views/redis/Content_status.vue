<template>
  <br/>
  <!--自动刷新-->
  <el-row style="width: 100%; margin-right: 2px">
    <el-col :span="24">
      <div style="float: right">
        <el-tag type="card">
          <sync-outlined spin/>
          自动刷新
        </el-tag>
        <el-tooltip placement="bottom" content="自动刷新开关，每5秒刷新一次">
          <el-switch
              @change="changeAutoRefresh"
              v-model="autoRefresh"
              inline-prompt
              active-text="开"
              inactive-text="关"
          >{{ autoRefresh }}
          </el-switch>
        </el-tooltip>
      </div>
    </el-col>
  </el-row>
  <br/>
  <!--基础信息-->
  <el-row :gutter="16" justify="space-around" style="margin-left: 2px">
    <!-- server status row -->
    <el-col :span="8" class="gutter-row">
      <el-card shadow="hover">
        <template #header>
          <img
              src="../../public/status/service.png"
              alt="service"
              style="width: 40px; height: 40px"
          />
          <span>服务器</span>
        </template>
        <p class="server-status-tag-p">
          <el-tag class="server-status-container" type="info">
            Redis版本：
            <span class="server-status-text">{{
                serviceInfo.data.version
              }}</span>
          </el-tag>
        </p>
        <p class="server-status-tag-p">
          <el-tag class="server-status-container" type="info" size="big">
            OS：
            <span class="server-status-text">{{ serviceInfo.data.os }}</span>
          </el-tag>
        </p>
        <p class="server-status-tag-p">
          <el-tag class="server-status-container" type="info" size="big">
            进程ID：
            <span class="server-status-text">{{
                serviceInfo.data.process
              }}</span>
          </el-tag>
        </p>
      </el-card>
    </el-col>
    <!-- memory row -->
    <el-col :span="8" class="gutter-row">
      <el-card shadow="hover">
        <template #header>
          <img
              src="../../public/status/memory.png"
              alt="service"
              style="width: 40px; height: 40px"
          />
          <span>内存</span>
        </template>
        <p class="server-status-tag-p">
          <el-tag class="server-status-container" type="info" size="big">
            已用内存：
            <span class="server-status-text">{{
                memoryInfo.data.usedMemory
              }}</span>
          </el-tag>
        </p>
        <p class="server-status-tag-p">
          <el-tag class="server-status-container" type="info" size="big">
            内存占用峰值：
            <span class="server-status-text">{{
                memoryInfo.data.usedBigMemory
              }}</span>
          </el-tag>
        </p>
        <p class="server-status-tag-p">
          <el-tag class="server-status-container" type="info" size="big">
            Lua占用内存：
            <span class="server-status-text">{{
                memoryInfo.data.luaMemory
              }}</span>
          </el-tag>
        </p>
      </el-card>
    </el-col>
    <el-col :span="8" class="gutter-row">
      <el-card shadow="hover">
        <template #header>
          <img
              src="../../public/status/history.png"
              alt="service"
              style="width: 40px; height: 40px"
          />
          <span>历史</span>
        </template>
        <p class="server-status-tag-p">
          <el-tag class="server-status-container" type="info" size="big">
            客户端连接数：
            <span class="server-status-text">{{
                historyInfo.data.connectCount
              }}</span>
          </el-tag>
        </p>
        <p class="server-status-tag-p">
          <el-tag class="server-status-container" type="info" size="big">
            历史连接数：
            <span class="server-status-text">{{
                historyInfo.data.historyCount
              }}</span>
          </el-tag>
        </p>
        <p class="server-status-tag-p">
          <el-tag class="server-status-container" type="info" size="big">
            历史命令数：
            <span class="server-status-text">{{
                historyInfo.data.historyInstructions
              }}</span>
          </el-tag>
        </p>
      </el-card>
    </el-col>
  </el-row>
  <br/>
  <!--键值信息-->
  <el-row style="width: 100%; margin-right: 2px">
    <el-col :span="24">
      <el-card shadow="hover">
        <template #title>
          <img
              src="../../public/status/kv.png"
              alt="service"
              style="width: 40px; height: 40px"
          />
          <span>键值统计</span>
        </template>
        <el-table
            :data="kvInfo.data"
            style="width: 100%"
            border
        >
          <el-table-column prop="db" label="DB" />
          <el-table-column prop="keys" label="Keys" />
          <el-table-column prop="expires" label="Expires" />
          <el-table-column prop="avgTtl" label="Avg TTL" />
        </el-table>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import {onBeforeMount, onBeforeUnmount, reactive, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();
//是否刷新
let autoRefresh = ref(false);
//定时刷新时间 5s
const refreshTimer = ref(5);
let timer = null;
// 获取redis所在服务器信息
let serviceInfo = reactive({
  data: [],
});
// 获取redis占用内存信息
let memoryInfo = reactive({
  data: [],
});
// 获取redis历史信息
let historyInfo = reactive({
  data: [],
});
// 键值信息
let kvInfo = reactive({
  data: [],
});

// 页面加载时同步加载redis数据
onBeforeMount(() => {
  // console.log("status页面", router.currentRoute.value.query.key)
  let fileName = router.currentRoute.value.query.fileName;
  if (fileName === "") {
    // 如果返回空值返回到默认界面
    router.push({
      path: "/rightContent/default",
    });
    return;
  }
  // 获取基本数据
  window.go.main.App.LoadingDbResource(fileName).then((resolve) => {
    if (resolve !== "") {
      // 如果返回值中不为空字符串才进行操作
      // console.log(resolve)
      let _json = JSON.parse(resolve);
      // 获取服务信息
      serviceInfo.data = _json.server;
      memoryInfo.data = _json.memory;
      historyInfo.data = _json.start;
      // 获取键值信息
      let _kvInfo = _json.dbkv;
      for (let i = 0; i < 16; i++) {
        let _temp = _kvInfo["db" + i];
        if (_temp === "") {
          continue;
        }
        let splitArry = _temp.split(",");
        let __info = "";
        for (let j = 0; j < splitArry.length; j++) {
          let _temp2 = splitArry[j].split("=");
          if (__info.length === 0) {
            __info = "{";
          }
          __info += '"' + _temp2[0] + '":"' + _temp2[1] + '"';
          if (j < splitArry.length - 1) {
            __info += ",";
          }
        }
        __info += "}";
        let _info = JSON.parse(__info);
        kvInfo.data.push({
          key: i,
          db: "db" + i,
          keys: _info.keys,
          expires: _info.expires,
          avgTtl: _info.avg_ttl,
        });
      }
    } else {
      // 如果返回空值返回到默认界面
      router.push({
        path: "/rightContent/default",
      });
    }
  });
});

// 卸载时关闭定时器
onBeforeUnmount(() => {
  clearInterval(timer);
  timer = null;
});

// 获取节点数据
function getNodeData() {
  window.go.main.App.GetNodeData("redis", "localhost", 13).then((resolve) => {
    let val = ["1:2:3", "1111", "234234", "12312", "1:2:4"];
    console.log(resolve);
  });
}

// 自动刷新按钮
function changeAutoRefresh() {
  console.log("自动刷新按钮",autoRefresh.value);
  if (autoRefresh.value) {
    // 开启定时器
    timer = setInterval(() => {
      // 刷新
      console.log("value:",router.currentRoute.value.query.fileName);
      window.go.main.App.LoadingDbResource(
          router.currentRoute.value.query.fileName
      ).then((resolve) => {
        if (resolve !== "") {
          // 如果返回值中不为空字符串才进行操作
          console.log(resolve);
          let _json = JSON.parse(resolve);
          // 获取服务信息
          serviceInfo.data = _json.server;
          memoryInfo.data = _json.memory;
          historyInfo.data = _json.start;
        }
      });
    }, refreshTimer.value * 1000);
  } else {
    // 关闭定时器
    clearInterval(timer);
    timer = null;
  }
}
</script>

<style scoped>
.server-status-tag-p {
  height: 32px;
}

.server-status-container {
  width: 100%;
  overflow-x: hidden;
  text-overflow: ellipsis;
}

.server-status-text {
  color: #43b50b;
}
</style>
