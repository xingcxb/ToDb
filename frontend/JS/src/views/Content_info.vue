<template>
  <a-row :align="middle" style="margin-top:10px">
    <a-col :offset="10" :span="4">
      <img src="../../public/info/key.png" alt="key" style="vertical-align: middle;" />
      <span style="font-size: 26px;vertical-align: middle;">{{key}}</span>
    </a-col>
    <a-col :offset="9" :span="1">
      <img src="../../public/info/close.png" alt="close" style="vertical-align: middle;" />
    </a-col>
  </a-row>
  <a-row style="margin-top:10px">
    <a-col :offset="1" :span="8">
      <a-input addon-before="String" addon-after="✓"/>
    </a-col>
  </a-row>
  {{key}}
  <br/>
  {{value}}
</template>

<script setup>
import {CloseOutlined} from '@ant-design/icons-vue';
import {onBeforeMount, reactive, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();

let key = ref("");
let value = reactive({
  data:""
})


// 初始化挂载前的函数
onBeforeMount(() => {
  // 获取路由传递的参数
  // redis键
  key = router.currentRoute.value.query.key
  // redis db
  let nodeId = router.currentRoute.value.query.dbKey
  // 类型
  let connType = router.currentRoute.value.query.connType
  // 连接文件名
  let connName = router.currentRoute.value.query.connName
  console.log("这是info页面")
  console.log("key",key,"nodeId",nodeId, "connType", connType, "connName", connName)
  if (connType == "redis") {
    // 从redis获取数据
    window.go.main.App.RedisGetData(connType,connName,nodeId,key).then(res => {
      console.log("==============",res)
      // 此处如果是空值，则应该是该键没有填充值
      value.data = res;
    })
  }
})

</script>

<style scoped>

</style>