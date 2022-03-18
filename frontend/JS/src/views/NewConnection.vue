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
})

// 测试连接
function testConnection() {
  window.go.main.App.TestConnection(JSON.stringify(connectionInfo)).then((resolve) => {
    console.log(resolve)
    var result = JSON.parse(resolve)
    let secondsToGo = 5;
    if (result.code === 200){
      const modal = Modal.success({
        title:"ToDo",
        content:result.message
      })
    }else{
      const modal = Modal.warning({
        title:"ToDo",
        content:result.message
      })
    }
  });

  // 确定按钮
  function ok(){

  }
}


</script>

<style scoped>
.imageRow {
  height: 72px;
  background: #7cb305;
}
</style>