<template>
  <a-row :align="middle" style="margin-top: 10px">
    <a-col :offset="10" :span="4">
      <img
        src="../../../public/info/key.png"
        alt="nowKey"
        style="vertical-align: middle"
      />
      <span style="font-size: 26px; vertical-align: middle">{{ nowKey }}</span>
    </a-col>
    <a-col :offset="8" :span="1">
      <!--关闭图片-->
      <img
        src="../../../public/info/close.png"
        alt="close"
        style="vertical-align: middle; cursor: pointer"
        @click="close"
      />
    </a-col>
  </a-row>
  <a-row style="margin-top: 20px">
    <a-col :offset="1" :span="10">
      <a-input-group compact>
        <a-input
          :addon-before="value.data.type"
          v-model:value="nowKey"
          style="width: calc(100% - 30px)"
        >
        </a-input>
        <a-button style="width: 30px" @click="rename">
          <!--设置新的key-->
          <template #icon>
            <check-outlined />
          </template>
        </a-button>
      </a-input-group>
    </a-col>
    <a-col :offset="1" :span="6">
      <a-input-group compact>
        <a-input
          addon-before="TTL"
          v-model:value="ttl"
          style="width: calc(100% - 60px)"
        >
        </a-input>
        <a-button style="width: 30px">
          <!--设置为持久化数据 无效废弃-->
          <template #icon>
            <close-outlined />
          </template>
        </a-button>
        <a-button style="width: 30px" @click="updateTtl">
          <!--设置数据过期时间-->
          <template #icon>
            <check-outlined />
          </template>
        </a-button>
      </a-input-group>
    </a-col>
    <a-col :offset="2" :span="1">
      <a-button type="primary" danger :size="size" @click="del">
        <template #icon>
          <!--删除-->
          <delete-outlined />
        </template>
      </a-button>
    </a-col>
    <a-col :span="1">
      <a-button
        type="primary"
        @click="getInfo"
        :size="size"
        style="background: #ffb33a; border: none"
      >
        <template #icon>
          <!--刷新-->
          <redo-outlined />
        </template>
      </a-button>
    </a-col>
    <a-col :span="1">
      <a-button
        type="primary"
        :size="size"
        style="background: #07c245; border: none"
        v-clipboard:copy="commandStr"
      >
        <template #icon>
          <!--获取命令-->
          <code-outlined />
        </template>
      </a-button>
    </a-col>
  </a-row>
  <a-row style="margin-top: 20px">
    <a-col :offset="1">
      <a-select
        ref="select"
        v-model:value="formatType"
        style="width: 120px"
        :options="formatTypeList"
        @focus="focus"
        @change="handleChange"
      >
      </a-select>
    </a-col>
    <a-col>
      <a-button>Size：{{ contentSize }}B</a-button>
    </a-col>
    <a-col>
      <a-button
        style="border: none; padding-left: 5px"
        v-clipboard:copy="content"
      >
        <copy-outlined />复制
      </a-button>
    </a-col>
  </a-row>
  <a-row style="margin-top: 10px">
    <a-col :offset="1" :span="22">
      <a-textarea v-model:value="content" :size="large" :rows="4" />
    </a-col>
  </a-row>
  <a-row style="margin-top: 10px">
    <a-col :offset="1" :span="1">
      <a-button type="primary" @click="saveValue">
        <template #icon>
          <!--保存value-->
          <save-outlined />
        </template>
      </a-button>
    </a-col>
  </a-row>
</template>

<script setup>
import {
  CopyOutlined,
  CheckOutlined,
  CloseOutlined,
  DeleteOutlined,
  RedoOutlined,
  CodeOutlined,
  SaveOutlined,
} from "@ant-design/icons-vue";
import { onBeforeMount, reactive, ref } from "vue";
import { useRouter } from "vue-router";

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
  window.go.main.App.RedisSaveStringValue(connType.value,connName.value,nodeId.value,nowKey.value,content.value,ttl.value)
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
  nodeId.value = router.currentRoute.value.query.dbKey;
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
