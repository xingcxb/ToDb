<template>
<a-modal v-model:visible="props.visible" title="Basic Modal" @ok="handleOk">
  <a-row>
    <a-col :span="24" class="imageRow">
      <!--  //540*72-->
      <a-image :width="100" :height="72"
               src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"/>
    </a-col>
  </a-row>
  <br/>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="4">连接名:</a-col>
    <a-col :span="18">
      <a-input v-model:value="connectionInfo.alias" placeholder="连接别名"/>
    </a-col>
    <a-col :span="1"></a-col>
  </a-row>
  <br/>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="4">主机:</a-col>
    <a-col :span="18">
      <a-input v-model:value="connectionInfo.hostURL" placeholder="连接地址"/>
    </a-col>
    <a-col :span="1"></a-col>
  </a-row>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="4">端口</a-col>
    <a-col :span="5">
      <a-input v-model:value="connectionInfo.port" placeholder="连接端口"/>
    </a-col>
    <a-col :span="14"></a-col>
  </a-row>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="4">用户名:</a-col>
    <a-col :span="10">
      <a-input v-model:value="connectionInfo.username" placeholder="用户名"/>
    </a-col>
    <a-col :span="9"></a-col>
  </a-row>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="4">密码:</a-col>
    <a-col :span="10">
      <a-input-password v-model:value="connectionInfo.password" placeholder="密码"/>
    </a-col>
    <a-col :span="9"></a-col>
  </a-row>
  <a-row>
    <a-col :span="5"></a-col>
    <a-col :span="18">
      <a-checkbox v-model:checked="connectionInfo.savePassword">保存密码</a-checkbox>
    </a-col>
    <a-col :span="1"></a-col>
  </a-row>
  <br/>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="3">
      <a-button type="primary" @click="testConnection">测试连接</a-button>
    </a-col>
    <a-col :span="13"></a-col>
    <a-col :span="3">
      <a-button @click="ok">确定</a-button>
    </a-col>
    <a-col :span="3">
      <a-button @click="cancel">取消</a-button>
    </a-col>
  </a-row>
</a-modal>
</template>

<script setup>
import {reactive, ref} from 'vue';
import {useRouter} from "vue-router";

// const router = useRouter()

const props = defineProps(['visible'])
console.log("++++++++++ ",props)

let connectionInfo = reactive({
  //连接别名
  alias: "",
  //连接地址
  hostURL: "",
  //端口
  port: "",
  //用户名
  username: "",
  //密码
  password: "",
  //是否保存密码
  savePassword: false,
  //类型
  // connType: router.currentRoute.value.query.connType,
})

// console.log("=========" + router.currentRoute.value.query.type)

function handleOk(){
  console.log("[[[[[[[[[[[[",props.visible)
  props.visible = false
  console.log("]]]]]]]]",props.visible)
}

// 测试连接
function testConnection() {
  window.go.main.App.TestConnection(JSON.stringify(connectionInfo)).then(() => {
  });
}

// // 确定按钮
// function ok() {
//   window.go.main.App.Ok(JSON.stringify(connectionInfo)).then((resolve) => {
//     var result = JSON.parse(resolve)
//     if (result.code === 200) {
//       //跳转到主页面
//       router.push({
//         name: 'Home'
//       })
//     }
//   });
// }

// // 取消按钮
// function cancel() {
//   router.push({
//     name: 'Home'
//   })
// }


</script>

<style scoped>
.imageRow {
  height: 72px;
  background: #7cb305;
}
</style>