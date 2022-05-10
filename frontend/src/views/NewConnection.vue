<template>
  <el-dialog
      v-model="visible"
      title="connType"
      width="600px"
      :before-close="handleClose"
      draggable
      center
  >
    <el-row>
      <el-col :span="24" class="imageRow">
        <!--  //472*72-->
        <img
            :width="100"
            :height="72"
            src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"
        />
      </el-col>
    </el-row>
    <br/>
    <el-row align="middle">
      <el-col :offset="1" :span="4">
        <span>连接名:</span>
      </el-col>
      <el-col :span="18">
        <el-input v-model="connectionInfo.alias" placeholder="连接别名"/>
      </el-col>
      <el-col :span="1"></el-col>
    </el-row>
    <el-row align="middle" class="rowInterval">
      <el-col :offset="1" :span="4">主&nbsp;&nbsp;&nbsp;机:</el-col>
      <el-col :span="18">
        <el-input
            v-model="connectionInfo.hostURL"
            placeholder="连接地址"
        />
      </el-col>
    </el-row>
    <el-row align="middle" class="rowInterval">
      <el-col :offset="1" :span="4">端&nbsp;&nbsp;&nbsp;口:</el-col>
      <el-col :span="5">
        <el-input v-model="connectionInfo.port" placeholder="连接端口"/>
      </el-col>
      <el-col :span="14"></el-col>
    </el-row>
    <el-row align="middle" class="rowInterval">
      <el-col :offset="1" :span="4">用户名:</el-col>
      <el-col :span="10">
        <el-input v-model="connectionInfo.username" placeholder="用户名"/>
      </el-col>
    </el-row>
    <el-row align="middle" class="rowInterval">
      <el-col :offset="1" :span="4">密&nbsp;&nbsp;&nbsp;码:</el-col>
      <el-col :span="10">
        <el-input
            type="password"
            v-model="connectionInfo.password"
            placeholder="密码"
            show-password
        />
      </el-col>
    </el-row>
    <el-row>
      <el-col :offset="5" :span="18">
        <el-checkbox v-model="connectionInfo.savePassword">
          保存密码
        </el-checkbox>
      </el-col>
      <el-col :span="1"></el-col>
    </el-row>
    <br/>
    <el-row>
      <el-col :offset="1" :span="3">
        <el-button type="primary" @click="testConnection">
          测试连接
        </el-button>
      </el-col>
      <el-col :offset="13" :span="3">
        <el-button @click="ok">确定</el-button>
      </el-col>
      <el-col :span="3">
        <el-button @click="handleClose">取消</el-button>
      </el-col>
    </el-row>
  </el-dialog>
</template>

<script setup>
import {reactive, ref, watch} from "vue";

// 接收父组件参数
const props = defineProps(["visible", "connType"]);
// 子组件中声明
const visible = ref(props.visible)
const connType = ref(props.connType)
// 声明提交事件
const emit = defineEmits([`ChangeVisible`, `ChangeConnType`]);
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
  connType: "",
});
// 监听props变化
watch([props], () => {
  visible.value = props.visible;
  connType.value = props.connType;
  connectionInfo.connType = connType.value
});
// 弹窗异步属性
let confirmLoading = ref(false)

// 弹窗关闭的回调
function handleClose() {
  emit("ChangeVisible", visible.value);
  emit("ChangeConnType", connType.value);
}

// 创建连接的方法暴露给go
window.runtime.EventsOn("createConn", function (data) {
  console.log("createConn", data);
  toView(data)
})

// 测试连接
function testConnection() {
  window.go.main.App.TestConnection(JSON.stringify(connectionInfo)).then(
      () => {
      }
  );
}

// 确定按钮
function ok() {
  window.go.main.App.Ok(JSON.stringify(connectionInfo)).then((resolve) => {
    var result = JSON.parse(resolve)
    if (result.code === 200) {
      confirmLoading.value = true
      setTimeout(() => {
        confirmLoading.value = false
        handleClose()
      }, 2000)
    }
  });
}

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
  /*background: #7cb305;*/
}
.rowInterval{
  margin-top: 3px;
}
</style>
