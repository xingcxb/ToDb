<template>
  <!--header-->
  <!--头部-->
  <div class="header" data-wails-drag>
    <!--导航-->
    <a-row style="height: 66px">
      <a-col :span="4">
        <a-dropdown :trigger="['click']" :placement="top">
          <a
            class="ant-dropdown-link header-a"
            @click.prevent
          >
            <img
              src="./assets/images/quick/conn.png"
              class="quickIcon"
              alt=""
            />
            <p style="font-size: 10px">连接</p>
          </a>
          <template #overlay>
            <a-menu>
              <a-menu-item key="0">
                <a @click="toView('redis')">Redis</a>
                <!--                <span :click="toView('redis')">Redis</span>-->
              </a-menu-item>
              <a-menu-item key="1">
                <a @click="toView('MySQL')">MySQL</a>
              </a-menu-item>
              <a-menu-divider />
              <a-menu-item key="3">other</a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </a-col>
      <a-col :span="4">
        <a href="" class="header-a">
          <img src="./assets/images/quick/table.png" class="quickIcon" alt="" />
          <p style="font-size: 10px; text-align: center">表</p>
        </a>
      </a-col>
      <a-col :span="4">
        <a href="" class="header-a">
          <img
            src="./assets/images/quick/select.png"
            class="quickIcon"
            alt=""
          />
          <p style="font-size: 10px; text-align: center">查询</p>
        </a>
      </a-col>
      <a-col :span="4">
        <a @click="importFile" class="header-a">
          <img
            src="./assets/images/quick/import.png"
            class="quickIcon"
            alt=""
          />
          <p style="font-size: 10px; text-align: center">导入</p>
        </a>
      </a-col>
      <a-col :span="4">
        <a @click="exportConn" class="header-a">
          <img
            src="./assets/images/quick/export.png"
            class="quickIcon"
            alt=""
          />
          <p style="font-size: 10px; text-align: center">导出</p>
        </a>
      </a-col>
      <a-col :span="4">
        <!--广告-->
      </a-col>
    </a-row>
    <Connection
      :visible="visible"
      :connType="connType"
      @ChangeVisible="visible = $event.visible"
      @ChangeConnType="connType = $event.connType"
    ></Connection>
  </div>
  <div class="box" ref="box">
    <div class="left">
      <!--左侧div内容-->
      <a-directory-tree
        v-model:expandedKeys="expandedKeys"
        v-model:selectedKeys="selectedKeys"
        :tree-data="listData.data"
        :load-data="onLoadData"
        @select="onSelect"
        style="
          min-width: 210px;
          /**width: calc(20% - 10px);*/
          width: 406px;
          background: rgba(224, 225, 225, 0.1);
        "
      >
      </a-directory-tree>
    </div>
    <div class="resize" @mousedown="handleMouseMoveLine"></div>
    <div class="mid">
      <!--右侧div内容-->
      <router-view :key="$route.path + Date.now()" />
    </div>
  </div>
</template>

<script setup>
import Connection from "./views/NewConnection.vue";
import { onBeforeMount, reactive, ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

let listData = reactive({
  data: [],
});

let visible = ref(false);
let connType = ref("");
const expandedKeys = ref([]);
const selectedKeys = ref([]);

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

// 打开弹窗
function toView(v) {
  visible.value = true;
  connType.value = v;
}

// 导入连接
function importFile() {
  console.log("点击了导入连接")
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

//鼠标移动到边线
function handleMouseMoveLine() {
  var resize = document.getElementsByClassName("resize");
  var left = document.getElementsByClassName("left");
  var mid = document.getElementsByClassName("mid");
  var box = document.getElementsByClassName("box");
  console.log(document)
  console.log(resize[0].left+"xxxxx")
  console.log("resize:", resize);
  console.log("left:",left);
  console.log("mid:",mid);
  console.log("box:", box)

  for (let i = 0; i < resize.length; i++) {
    console.log("resize:", resize[i]);
    // 鼠标按下事件
    resize[i].onmousedown = function (e) {
      //颜色改变提醒
      resize[i].style.background = "#818181";
      var startX = e.clientX;
      resize[i].left = resize[i].offsetLeft;
      // 鼠标拖动事件
      document.getElementById("app").onmousemove = function (e) {
        console.log("this is e    ",e)
        var endX = e.clientX;
        // （endx-startx）=移动的距离。resize[i].left+移动的距离=左边区域最后的宽度
        var moveLen = resize[i].left + (endX - startX);
        // 容器宽度 - 左边区域的宽度 = 右边区域的宽度
        var maxT = box[i].clientWidth - resize[i].offsetWidth;

        // 左边区域的最小宽度为32px
        if (moveLen < 406) moveLen = 406;
        //右边区域最小宽度为150px
        if (moveLen > maxT - 655) moveLen = maxT - 655;

        // 设置左侧区域的宽度
        resize[i].style.left = moveLen;

        for (let j = 0; j < left.length; j++) {
          console.log("this is moveLen ", moveLen)
          left[j].style.width = moveLen + "px";
          mid[j].style.width = box[i].clientWidth - moveLen - 10 + "px";
        }
      };
      // 鼠标松开事件
      document.onmouseup = function (evt) {
        //颜色恢复
        resize[i].style.background = "#d6d6d6";
        document.onmousemove = null;
        document.onmouseup = null;
        //当你不在需要继续获得鼠标消息就要应该调用ReleaseCapture()释放掉
        resize[i].releaseCapture && resize[i].releaseCapture();
      };
      //该函数在属于当前线程的指定窗口里设置鼠标捕获
      resize[i].setCapture && resize[i].setCapture();
      return false;
    };
  }
}
</script>
<style>
body {
  overscroll-behavior: none;
}

.header {
  height: 64px;
  width: 100%;
  background: #f0efee;
}

.header-a{
  text-align: center;
  display: block;
  cursor: default;
}

.box {
  min-height: 500px;
  height: 100%;
  width: 100%;
  overflow: hidden;
  /*border-top: 1px solid black;*/
}

/**
 * 左侧区域
 */
.left {
  /*width: calc(20% - 10px);*/
  width: 406px;
  min-height: 661px;
  height: 100%;
  float: left;
  background: rgba(224, 225, 225, 0.1);
}

.resize {
  min-height: 661px;
  height: 100%;
  width: 2px;
  float: left;
  border-right: 2px solid #d5d6d6;
  background-image: url("./assets/line.png");
}

.resize:hover {
  cursor: col-resize;
}

.mid {
  float: left;
  width: 80%; /*右侧初始化宽度*/
  min-height: 661px;
  height: 100%;
}

.quickIcon {
  padding: 0;
  text-align: center;
  height: 44px;
  width: auto;
}
</style>
