<template>
    <div class="common-layout">
<!--  <div class="header" data-wails-drag>-->
    <el-container>
      <el-header>
        <!--头部 导航-->
        <el-row style="height: 66px">
          <el-col :span="4">
            <el-dropdown trigger="click">
              <span class="el-dropdown-link">
                <img src="./assets/images/quick/conn.png"
                     class="quickIcon"
                     alt=""
                />
                <br/>
                <span style="font-size: 10px">连接</span>
              </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item command="Redis">Redis
<!--                    <a @click="toView('redis')">Redis</a>-->
                  </el-dropdown-item>
                  <el-dropdown-item command="MySQL">
<!--                    <a @click="toView('MySQL')">MySQL</a>-->
                  </el-dropdown-item>
                  <el-divider/>
                  <el-dropdown-item command="other">other</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </el-col>
          <el-col :span="4">
            <a href="" class="header-a">
              <img src="./assets/images/quick/table.png" class="quickIcon" alt=""/>
              <span style="font-size: 10px; text-align: center">表</span>
            </a>
          </el-col>
          <el-col :span="4">
            <a href="" class="header-a">
              <img
                  src="./assets/images/quick/select.png"
                  class="quickIcon"
                  alt=""
              />
              <span style="font-size: 10px; text-align: center">查询</span>
            </a>
          </el-col>
          <el-col :span="4">
            <a @click="importFile" class="header-a">
              <img
                  src="./assets/images/quick/import.png"
                  class="quickIcon"
                  alt=""
              />
              <span style="font-size: 10px; text-align: center">导入</span>
            </a>
          </el-col>
          <el-col :span="4">
            <a @click="exportConn" class="header-a">
              <img
                  src="./assets/images/quick/export.png"
                  class="quickIcon"
                  alt=""
              />
              <span style="font-size: 10px; text-align: center">导出</span>
            </a>
          </el-col>
          <span-col :span="4">
            <!--未想到-->
          </span-col>
        </el-row>
        <Connection
            :visible="visible"
            :connType="connType"
            @ChangeVisible="visible = $event.visible"
            @ChangeConnType="connType = $event.connType"
        ></Connection>
      </el-header>
      <el-container>
        <el-aside width="200px">
          <!--左侧内容-->
          <el-tree-v2 :data="listData.data"></el-tree-v2>
        </el-aside>
        <el-main>
          <!--右侧内容-->
          <router-view :key="$route.path + Date.now()"/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import Connection from "./views/NewConnection.vue";
import {onBeforeMount, reactive, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();

let listData = reactive({
  data: [],
});

let visible = ref(false);
let connType = ref("");

onBeforeMount(() => {
  router.push({
    path: "/rightContent/default",
  });
  window.go.main.App.LoadingConnKey().then((resolve) => {
    if (resolve !== "") {
      // 如果返回值中不为空字符串才进行操作
      listData.data = JSON.parse(resolve);
    }
  });
});

// 创建连接的方法暴露给go
window.runtime.EventsOn("createConn", function (data) {
  console.log("createConn", data);
  toView(data)
})
// 导入的方法暴露给go
window.runtime.EventsOn("importConn", function () {
  importFile()
})

// 导出方法暴露给go
window.runtime.EventsOn("exportConn", function () {
  exportConn()
})

// 打开弹窗
function toView(v) {
  console.log("这是V ", v);
  visible.value = true;
  connType.value = v;
}

// 导入连接
function importFile() {
  window.go.main.App.ImportConn()
}

// 导出连接
function exportConn() {
  window.go.main.App.ExportConn()
}

// 选中文字也可以进行操作
function onSelect(selectedKeys, info) {
  if (selectedKeys.length === 0) {
    //表示是取消选中
    return;
  }
  let parent = info.node.parent;
  if (parent != undefined) {
    //当前为子节点,改变到右侧的页面中显示数据
    //redis,localhost,1
    let parentKey = info.node.parent.key;
    console.log("this is info ", info);
    console.log("this is parentKey ", parentKey);
    let parentKeyArr = parentKey.split(",");
    let connType = parentKeyArr[0];
    let connName = parentKeyArr[1];
    //selectKey是显示具体的节点key，parent显示的是父节点
    window.go.main.App.GetNodeData(
        connType,
        connName,
        selectedKeys[0] + ""
    ).then((resolve) => {
      if (resolve !== "") {
        // 如果返回值中不为空字符串才进行操作
        let data = JSON.parse(resolve);
        let list = info.node.parent.node.children.map((item) => {
          if (item.key == selectedKeys) {
            item.children = data;
          }
          return item;
        });
        info.node.parent.node.children = [...list];
      } else {
        let key = selectedKeys[0];
        console.log("key", key);
        let rootKey = info.node.parent.nodes[0].key;
        let dbKey = info.node.parent.nodes[1].key;
        let rootKeyArr = rootKey.split(",");
        console.log("info", info.node.parent);
        //此处是属于根节点，右侧显示具体的数据
        if (key.indexOf("*") === -1) {
          // 如果不包含*，则表示右侧的页面要渲染
          router.push({
            path: "/rightContent/value",
            query: {
              connType: rootKeyArr[0],
              connName: rootKeyArr[1],
              dbKey: dbKey,
              key: selectedKeys[0],
            },
          });
        }
      }
    });
  } else {
    //当前为根节点
    let key = selectedKeys[0];
    let keys = key.split(",");
    //表示选中一个根节点，通常是选中了一个数据库
    router.push({
      path: "/rightContent/status",
      query: {
        key: keys[1],
      },
    });
  }
}

// 获取本地的连接信息
let onLoadData = (treeNode) => {
  let key = treeNode.dataRef.key.split(",");
  if (key.length == 1) {
    //如果只有一个key，说明不是根节点
    return;
  }
  return new Promise((resolve) => {
    window.go.main.App.LoadingConnInfo(key[1]).then((resolve) => {
      if (resolve !== "") {
        // 如果返回值中不为空字符串才进行操作
        treeNode.dataRef.children = JSON.parse(resolve);
        listData.data = [...listData.data];
      }
    });
    resolve();
  });
};
</script>

<style scoped>
body {
  overscroll-behavior: none;
}

.header {
  height: 64px;
  width: 100%;
  background: #f0efee;
}

.quickIcon {
  padding-top: 3px;
  text-align: center;
  height: 44px;
  width: auto;
}
</style>