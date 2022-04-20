<template>
  <a-row :align="middle" style="margin-top:10px">
    <a-col :offset="10" :span="4">
      <img src="../../public/info/key.png" alt="key" style="vertical-align: middle;" />
      <span style="font-size: 26px;vertical-align: middle;">{{ key }}</span>
    </a-col>
    <a-col :offset="8" :span="1">
      <!--关闭图片-->
      <img src="../../public/info/close.png" alt="close" style="vertical-align: middle;cursor: pointer;" @click="close"/>
    </a-col>
  </a-row>
  <a-row style="margin-top:20px">
    <a-col :offset="1" :span="10">
      <a-input-group compact>
        <a-input :addon-before="value.data.type" v-model="key" style="width: calc(100% - 30px);">
        </a-input>
        <a-button style="width: 30px;" @click="rename">
          <!--设置新的key-->
          <template #icon>
            <check-outlined />
          </template>
        </a-button>
      </a-input-group>
    </a-col>
    <a-col :offset="1" :span="6">
      <a-input-group compact>
        <a-input addon-before="TTL" v-model:value="value.data.ttl" style="width: calc(100% - 60px);">
        </a-input>
        <a-button style="width: 30px;">
          <!--设置为持久化数据-->
          <template #icon>
            <close-outlined />
          </template>
        </a-button>
        <a-button style="width: 30px;">
          <!--设置数据过期时间-->
          <template #icon>
            <check-outlined />
          </template>
        </a-button>
      </a-input-group>
    </a-col>
    <a-col :offset="2" :span="1">
      <a-button type="primary" danger :size="size">
        <template #icon>
          <delete-outlined />
        </template>
      </a-button>
    </a-col>
    <a-col :span="1">
      <a-button type="primary" :size="size" style="background: #ffb33a;border: none;">
        <template #icon>
          <redo-outlined />
        </template>
      </a-button>
    </a-col>
    <a-col :span="1">
      <a-button type="primary" :size="size" style="background: #07c245;border: none;">
        <template #icon>
          <code-outlined />
        </template>
      </a-button>
    </a-col>
  </a-row>
  <a-row style="margin-top:20px;">
    <a-col :offset="1">
      <a-select 
      ref="select" 
      v-model:value="formatType" 
      style="width:120px"
      :options="formatTypeList" 
      @focus="focus" 
      @change="handleChange">
      </a-select>
    </a-col>
    <a-col>
      <a-button>Size：{{contentSize}}</a-button>
    </a-col>
    <a-col>
      <a-button style="border: none;">
        <copy-outlined />复制
      </a-button>
    </a-col>
  </a-row>
  <a-row style="margin-top:10px;">
    <a-col :offset="1" :span="22">
      <a-textarea v-model:value="content" :auto-size="{minRows:6,maxRows:6}"></a-textarea>
    </a-col>
  </a-row>
</template>

<script setup>
import { CopyOutlined,CheckOutlined, CloseOutlined, DeleteOutlined, RedoOutlined, CodeOutlined } from '@ant-design/icons-vue';
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
let key = ref("");
// redis old key
let oldKey = ref("");
// redis value
let value = reactive({
  data: ""
})
// redis content data size
let contentSize = ref(0);
// redis content data
let content = ref("");
// 内容展示类型
let formatType = ref("Text");

const formatTypeList = ref([{
  value: "Text",
  label: "Text"
}, {
  value: "HEX",
  label: "HEX"
}, {
  value: "Json",
  label: "Json"
}, {
  value: "Msgpack",
  label: "Msgpack"
}, {
  value: "Binary",
  label: "Binary"
}, {
  value: "Unserialize",
  label: "Unserialize"
}, {
  value: "Brotli",
  label: "Brotli"
}, {
  value: "Gzip",
  label: "Gzip"
}, {
  value: "Deflate",
  label: "Deflate"
}]);

// 关闭页面
function close(){
  router.push({
    path: "/rightContent/default"
  });
}

// 修改redis的键
function rename(){
  window.go.main.App.RedisReName(connType, connName,nodeId, oldKey, key)
}

// 改变现实格式
function handleChange(){
  console.log(formatType)
}

// 初始化挂载前的函数
onBeforeMount(() => {
  // 获取路由传递的参数
  // redis键
  key = router.currentRoute.value.query.key
  // redis键
  oldKey = router.currentRoute.value.query.key
  // redis db
  nodeId = router.currentRoute.value.query.dbKey
  // 类型
  connType = router.currentRoute.value.query.connType
  // 连接文件名
  connName = router.currentRoute.value.query.connName
  console.log("这是info页面")
  console.log("key", key, "nodeId", nodeId, "connType", connType, "connName", connName)
  if (connType == "redis") {
    // 从redis获取数据
    window.go.main.App.RedisGetData(connType, connName, nodeId, key).then(res => {
      console.log("==============", res)
      // 此处如果是空值，则应该是该键没有填充值
      value.data = JSON.parse(res);
      content.value = value.data.value
      console.log("++++++++++++======== ",content)
    })
  }
})

</script>

<style scoped>
.key-detail-type {
  text-transform: capitalize;
  text-align: center;
  min-width: 34px;
  display: inline-block;
}
</style>