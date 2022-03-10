<template>
  <a-row>
    <a-col :span="24">
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
      <a-input v-model:value="alias" placeholder="连接别名"/>
    </a-col>
    <a-col :span="1"></a-col>
  </a-row>
  <br/>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="4">主机:</a-col>
    <a-col :span="18">
      <a-input v-model:value="hostURL" placeholder="连接地址"/>
    </a-col>
    <a-col :span="1"></a-col>
  </a-row>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="4">端口</a-col>
    <a-col :span="5">
      <a-input v-model:value="port" placeholder="连接端口"/>
    </a-col>
    <a-col :span="14"></a-col>
  </a-row>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="4">用户名:</a-col>
    <a-col :span="10">
      <a-input v-model:value="username" placeholder="用户名"/>
    </a-col>
    <a-col :span="9"></a-col>
  </a-row>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="4">密码:</a-col>
    <a-col :span="10">
      <a-input-password v-model:value="password" placeholder="密码"/>
    </a-col>
    <a-col :span="9"></a-col>
  </a-row>
  <a-row>
    <a-col :span="5"></a-col>
    <a-col :span="18">
      <a-checkbox v-model:checked="savePassword">保存密码</a-checkbox>
    </a-col>
    <a-col :span="1"></a-col>
  </a-row>
  <br/>
  <a-row>
    <a-col :span="1"></a-col>
    <a-col :span="3">
      <a-button type="primary">测试连接</a-button>
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
import {ref} from "vue";

const alias = ref('')
const hostURL = ref('')
const port = ref('')
const username = ref('')
const password = ref('')
const savePassword = ref(false)

</script>

<style scoped>

</style>