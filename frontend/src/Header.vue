<template>
  <!--头部 导航-->
  <el-row style="height: 69px">
    <el-col :offset="1" :span="3" style="background: #f5f4f2">
      <el-dropdown trigger="click">
        <el-card shadow="hover" class="card" :body-style="{padding:'0px',border:'none'}">
          <div class="quickText_div">
            <img src="./assets/images/quick/conn.png" class="quickImg">
            <br/>
            <span class="quickText">连接</span>
          </div>
        </el-card>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click.native="toView('Redis')" command="Redis">Redis</el-dropdown-item>
            <el-dropdown-item @click.native="toView('MySQL')" command="MySQL">MySQL</el-dropdown-item>
            <el-dropdown-item command="Other" divided disabled>Other</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </el-col>
    <el-col :offset="1" :span="3">
      <el-dropdown trigger="contextmenu">
        <el-card shadow="hover" class="card" :body-style="{padding:'0px',border:'none'}">
          <div class="quickText_div">
            <img src="./assets/images/quick/table.png" class="quickImg">
            <br/>
            <span class="quickText">表</span>
          </div>
        </el-card>
      </el-dropdown>
    </el-col>
    <el-col :offset="1" :span="3">
      <el-dropdown trigger="contextmenu">
        <el-card shadow="hover" class="card" :body-style="{padding:'0px',border:'none'}">
          <div class="quickText_div">
            <img src="./assets/images/quick/select.png" class="quickImg">
            <br/>
            <span class="quickText">查询</span>
          </div>
        </el-card>
      </el-dropdown>
    </el-col>
    <el-col :offset="1" :span="3">
      <el-dropdown trigger="contextmenu">
        <el-card @click="importFile" shadow="hover" class="card" :body-style="{padding:'0px',border:'none'}">
          <div class="quickText_div">
            <img src="./assets/images/quick/import.png" class="quickImg">
            <br/>
            <span class="quickText">导入</span>
          </div>
        </el-card>
      </el-dropdown>
    </el-col>
    <el-col :offset="1" :span="3">
      <el-dropdown trigger="contextmenu">
        <el-card @click="exportConn" shadow="hover" class="card" :body-style="{padding:'0px',border:'none'}">
          <div class="quickText_div">
            <img src="./assets/images/quick/export.png" class="quickImg">
            <br/>
            <span class="quickText">导出</span>
          </div>
        </el-card>
      </el-dropdown>
    </el-col>
    <el-col :span="4">
      <!--未想到-->
    </el-col>
  </el-row>
  <Connection
      :visible="visible"
      :connType="connType"
      @ChangeVisible="visible = $event.visible"
      @ChangeConnType="connType = $event.connType"
  ></Connection>
</template>

<script setup>
import Connection from "./views/NewConnection.vue";
import {ref} from "vue";

let visible = ref(false);
let connType = ref("");

// 打开弹窗
function toView(v) {
  visible.value = true;
  connType.value = v;
}

// 导入的方法暴露给go
window.runtime.EventsOn("importConn", function () {
  importFile()
})

// 导出方法暴露给go
window.runtime.EventsOn("exportConn", function () {
  exportConn()
})

// 导入连接
function importFile() {
  console.log("导入连接");
  window.go.main.App.ImportConn()
}

// 导出连接
function exportConn() {
  window.go.main.App.ExportConn()
}

</script>

<style scoped>
.card {
  width: 80px;
  height: 67px;
  border: 0px;
  cursor: default;
}

.quickText_div {
  text-align: center;
  align-items: center;
  background: #f5f4f2;
}

.quickImg {
  width: 49px;
  height: 49px;
}

.quickText {
  font-size: 10px;
}
</style>