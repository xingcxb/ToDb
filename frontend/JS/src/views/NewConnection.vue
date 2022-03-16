<template>
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
      <a-button>确定</a-button>
    </a-col>
    <a-col :span="3">
      <a-button>取消</a-button>
    </a-col>
  </a-row>
</template>

<script setup>
import { reactive } from 'vue';
import { Modal } from 'ant-design-vue';

let connectionInfo = reactive({
  alias: "",
  hostURL: "",
  port: "",
  username: "",
  password: "",
  savePassword: "",
})

// //连接别名
// const alias = ref('')
// //连接地址
// const hostURL = ref('')
// //端口
// const port = ref('')
// //用户名
// const username = ref('')
// //密码
// const password = ref('')
// //是否保存密码
// const savePassword = ref(false)

// 测试连接
function testConnection() {
  window.go.main.App.TestConnection(JSON.stringify(connectionInfo)).then((resolve) => {
    console.log(resolve)
    // var result = JSON.parse(resolve)
    // if (result.code === 200){
    //   let secondsToGo = 5;
    //   const modal = Modal.success({
    //     title:"",
    //     content:"111111"
    //   })
    // }
  });
}


</script>

<style scoped>
.imageRow {
  height: 72px;
  background: #7cb305;
}
</style>