<template>
  <div>
    <!--  <div class="header" data-wails-drag>-->
    <el-container>
      <el-header height="69px" style="background: #f5f4f2">
        <!--头部 导航-->
        <Header></Header>
      </el-header>
      <el-container>
        <div class="aside-drag-container" :style="{width: sideWidth + 'px'}">
          <el-aside class="aside-connection">
            <!--左侧导航-->
            <Aside></Aside>
          </el-aside>
          <!-- drag area -->
          <div id="drag-resize-container">
            <div id="drag-resize-pointer"></div>
          </div>
        </div>
        <el-main>
          <!--右侧内容-->
          <router-view :key="$route.path + Date.now()"/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import Header from './Header.vue';
import Aside from "./Aside.vue";
import {onBeforeMount, ref} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();

let sideWidth = ref(265);

onBeforeMount(() => {
  router.push({
    path: "/rightContent/default",
  });
});

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

.aside-drag-container {
  position: relative;
  user-select: none;
  /*max-width: 50%;*/
}

.aside-connection {
  height: 100%;
  border-right: 1px solid #e4e0e0;
  overflow: hidden;
}

#drag-resize-container {
  position: absolute;
  /*height: 100%;*/
  width: 10px;
  right: -12px;
  top: 0px;
}
#drag-resize-pointer {
  position: fixed;
  height: 100%;
  width: 10px;
  cursor: col-resize;
}
#drag-resize-pointer::after {
  content: "";
  display: inline-block;
  width: 2px;
  height: 20px;
  border-left: 1px solid #adabab;
  border-right: 1px solid #adabab;

  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  margin: auto;
}
.dark-mode #drag-resize-pointer::after {
  border-left: 1px solid #b9b8b8;
  border-right: 1px solid #b9b8b8;
}

</style>