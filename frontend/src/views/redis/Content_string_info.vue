<template>
  <el-row type="flex" justify="center" align="middle" style="margin-top: 10px">
    <el-col :span="14">
      <img
        src="../../public/info/key.png"
        alt="nowKey"
        style="vertical-align: middle; height: 28px; width: 28px"
      />
      <span style="font-size: 20px; vertical-align: middle">{{ nowKey }}</span>
    </el-col>
    <el-col :offset="8" :span="1">
      <!--关闭图片-->
      <img
        src="../../public/info/close.png"
        alt="close"
        style="vertical-align: middle; cursor: pointer"
        @click="close"
      />
    </el-col>
  </el-row>
  <el-row style="margin-top: 20px" :gutter="5">
    <el-col :span="10">
      <!--key-->
      <el-input v-model:value="nowKey" style="width: calc(100% - 30px)">
        <template #prepend>{{ value.data.type }}</template>
        <template #append>
          <el-button @click="rename">
            <template #icon>
              <CheckmarkOutline />
            </template>
          </el-button>
        </template>
      </el-input>
    </el-col>
    <el-col :offset="1" :span="6">
      <!--到期时间-->
      <el-input v-model:value="ttl" style="width: calc(100% - 60px)">
        <template #prepend>TTL</template>
        <template #append>
          <el-button @click="updateTtl">
            <template #icon>
              <CheckmarkOutline />
            </template>
          </el-button>
        </template>
      </el-input>
    </el-col>
    <el-col :offset="1" :span="2">
      <el-button type="primary" danger :size="small" @click="del">
        <template #icon>
          <!--删除-->
          <Delete28Regular />
        </template>
      </el-button>
    </el-col>
    <el-col :span="2">
      <el-button
        type="primary"
        :size="small"
        @click="getInfo"
        style="background: #ffb33a; border: none"
      >
        <template #icon>
          <!--刷新-->
          <Refresh />
        </template>
      </el-button>
    </el-col>
    <el-col :span="2">
      <el-button
        type="primary"
        :size="small"
        style="background: #07c245; border: none"
        v-clipboard:copy="commandStr"
      >
        <template #icon>
          <!--获取命令-->
          <Code />
        </template>
      </el-button>
    </el-col>
  </el-row>
  <el-row style="margin-top: 20px; height: 24px" :gutter="5">
    <el-col :span="2">
      <el-select size="small" v-model="formatType" change="handleChange">
        <el-option
          v-for="item in formatTypeList"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        />
      </el-select>
    </el-col>
    <el-col :span="2">
      <el-tag>
        <span style="font-size: 10px"> Size：{{ contentSize }}B </span>
      </el-tag>
    </el-col>
    <el-col :span="1">
      <el-button
        :size="small"
        style="border: none; padding-left: 5px; height: 24px"
        v-clipboard:copy="content"
      >
        <template #icon>
          <CopyOutline />
        </template>
        复制
      </el-button>
    </el-col>
  </el-row>
  <el-row style="margin-top: 20px">
    <el-col :span="23">
      <el-input v-model="content" type="textarea" size="large" :rows="10" />
    </el-col>
  </el-row>
  <el-row style="margin-top: 10px">
    <el-col :span="1">
      <el-button type="primary" @click="saveValue"> 保存</el-button>
    </el-col>
  </el-row>
</template>

<script setup>
import { onBeforeMount, reactive, ref } from "vue";
import { useRouter } from "vue-router";
import {
  CheckmarkOutline,
  Refresh,
  Code,
  CopyOutline,
} from "@vicons/ionicons5";
import { Delete28Regular } from "@vicons/fluent";

const router = useRouter();

// 节点id
let nodeId = ref("");
// 连接类型
let connType = ref("");
// 连接文件名
let connName = ref("");
// redis key
let nowKey = ref("");
// redis old key
let oldKey = ref("");
// redis value
let value = reactive({
  data: "",
});
// redis 剩余时间
let ttl = ref("");
// redis content data size
let contentSize = ref(0);
// redis content data
let content = ref("");
// 内容展示类型
let formatType = ref("Text");
// 命令
let commandStr = ref("");

const formatTypeList = ref([
  {
    value: "Text",
    label: "Text",
  },
  {
    value: "HEX",
    label: "HEX",
  },
  {
    value: "Json",
    label: "Json",
  },
  {
    value: "Msgpack",
    label: "Msgpack",
  },
  {
    value: "Binary",
    label: "Binary",
  },
  {
    value: "Unserialize",
    label: "Unserialize",
  },
  {
    value: "Brotli",
    label: "Brotli",
  },
  {
    value: "Gzip",
    label: "Gzip",
  },
  {
    value: "Deflate",
    label: "Deflate",
  },
]);

// 关闭页面
function close() {
  router.push({
    path: "/rightContent/default",
  });
}

// 删除
function del() {
  window.go.main.App.RedisDelKey(
    connType.value,
    connName.value,
    nodeId.value,
    nowKey.value
  );
}

// 更新剩余时间
function updateTtl() {
  window.go.main.App.RedisUpTtl(
    connType.value,
    connName.value,
    nodeId.value,
    nowKey.value,
    ttl.value
  );
}

// 修改redis的键
function rename() {
  window.go.main.App.RedisReName(
    connType.value,
    connName.value,
    nodeId.value,
    oldKey.value,
    nowKey.value
  );
}

// 获取基础信息
function getInfo() {
  if (connType.value == "redis") {
    // 从redis获取数据
    window.go.main.App.RedisGetData(
      connType.value,
      connName.value,
      nodeId.value,
      nowKey.value
    ).then((res) => {
      // 此处如果是空值，则应该是该键没有填充值
      value.data = JSON.parse(res);
      content.value = value.data.value;
      ttl.value = value.data.ttl;
      contentSize.value = value.data.size;
      commandStr.value = value.data.commandStr;
    });
  }
}

// 保存值
function saveValue() {
  window.go.main.App.RedisSaveStringValue(
    connType.value,
    connName.value,
    nodeId.value,
    nowKey.value,
    content.value,
    ttl.value
  );
}

// 改变现实格式
function handleChange() {
  console.log(formatType);
}

// 初始化挂载前的函数
onBeforeMount(() => {
  // 获取路由传递的参数
  // redis键
  nowKey.value = router.currentRoute.value.query.key;
  // redis键
  oldKey.value = router.currentRoute.value.query.key;
  // redis db
  nodeId.value = router.currentRoute.value.query.dbId;
  // 类型
  connType.value = router.currentRoute.value.query.connType;
  // 连接文件名
  connName.value = router.currentRoute.value.query.connName;
  getInfo();
});
</script>

<style scoped>
.key-detail-type {
  text-transform: capitalize;
  text-align: center;
  min-width: 34px;
  display: inline-block;
}
</style>
