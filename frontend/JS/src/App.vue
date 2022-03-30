<template>
  <!--header-->
  <!--头部-->
  <div class="header" data-wails-drag>
    <!--导航-->
    <a-row style="height: 66px">
      <a-col :span="4">
        <a-dropdown :trigger="['click']" :placement="top">
          <a class="ant-dropdown-link" style="text-align: center;display: block" @click.prevent>
            <img src="./assets/images/quick/conn.png" class="quickIcon" alt=""/>
            <p style="font-size:10px">连接</p>
          </a>
          <template #overlay>
            <a-menu>
              <a-menu-item key="0">
                <a @click="toView('redis')">Redis</a>
                <!--                <span :click="toView('redis')">Redis</span>-->
              </a-menu-item>
              <a-menu-item key="1">
                <a href="http://www.taobao.com/">MySQL</a>
              </a-menu-item>
              <a-menu-divider/>
              <a-menu-item key="3">other</a-menu-item>
            </a-menu>
          </template>
        </a-dropdown>
      </a-col>
      <a-col :span="4">
        <a href="" style="text-align: center;display: block">
          <img src="./assets/images/quick/table.png" class="quickIcon" alt=""/>
          <p style="font-size:10px;text-align: center">表</p>
        </a>
      </a-col>
      <a-col :span="4">
        <a href="" style="text-align: center;display: block">
          <img src="./assets/images/quick/select.png" class="quickIcon" alt=""/>
          <p style="font-size:10px;text-align: center">查询</p>
        </a>
      </a-col>
      <a-col :span="4">
        <a href="" style="text-align: center;display: block">
          <img src="./assets/images/quick/import.png" class="quickIcon" alt=""/>
          <p style="font-size:10px;text-align: center">导入</p>
        </a>
      </a-col>
      <a-col :span="4">
        <a href="" style="text-align: center;display: block">
          <img src="./assets/images/quick/export.png" class="quickIcon" alt=""/>
          <p style="font-size:10px;text-align: center">导出</p>
        </a>
      </a-col>
      <a-col :span="4">
        <!--广告-->
      </a-col>
    </a-row>
  </div>
    <div class="box" ref="box">
      <div class="left">
        <!--左侧div内容-->
        <a-tree
            v-model:expandedKeys="expandedKeys"
            v-model:selectedKeys="selectedKeys"
            :load-data="onLoadData"
            :tree-data="listData.data"
            @select="onSelect"
            style="width: calc(20% - 10px); background: rgba(224, 225, 225, 0.1)"
        />
      </div>
      <div class="resize" @mousedown="handleMouseMoveLine">
      </div>
      <div class="mid">
        <!--右侧div内容-->
        <router-view/>
      </div>
    </div>
</template>

<script setup>
import {onBeforeMount,onMounted, reactive, ref} from 'vue';
import {useRouter} from "vue-router";

const router = useRouter();

let listData = reactive({
  data: []
})
const expandedKeys = ref([]);
const selectedKeys = ref([]);

onBeforeMount(() => {
  window.go.main.App.LoadingConnKey().then((resolve) => {
    if (resolve !== "") {
      // 如果返回值中不为空字符串才进行操作
      listData.data = JSON.parse(resolve)
    }
  });
  router.push({
    path:"/rightContent/default",
  })
})

function toView(v) {
  console.log(v)
  router.push({
    path: "/newConnection",
    query: {
      connType: v
    }
  })
}

// onMounted(()=>{
//   router.go({
//     path: '/home/welcome',
//     replace: true
//   })
// })

// 选中文字也可以进行操作
// function onSelect(selectedKeys,{node}) {
//   node.onExpand()
// }

// 获取本地的连接信息
let onLoadData = treeNode => {
  return new Promise(resolve => {
    window.go.main.App.LoadingConnInfo(treeNode.dataRef.key).then((resolve) => {
      if (resolve !== "") {
        // 如果返回值中不为空字符串才进行操作
        treeNode.dataRef.children = JSON.parse(resolve)
        listData.data =[...listData.data]
      }
    })
    resolve();
  });
}

//鼠标移动到边线
function handleMouseMoveLine() {
  var resize = document.getElementsByClassName('resize');
  var left = document.getElementsByClassName('left');
  var mid = document.getElementsByClassName('mid');
  var box = document.getElementsByClassName('box');
  for (let i = 0; i < resize.length; i++) {
    // 鼠标按下事件
    resize[i].onmousedown = function (e) {
      //颜色改变提醒
      resize[i].style.background = '#818181';
      var startX = e.clientX;
      resize[i].left = resize[i].offsetLeft;
      // 鼠标拖动事件
      document.onmousemove = function (e) {
        var endX = e.clientX;
        // （endx-startx）=移动的距离。resize[i].left+移动的距离=左边区域最后的宽度
        var moveLen = resize[i].left + (endX - startX);
        // 容器宽度 - 左边区域的宽度 = 右边区域的宽度
        var maxT = box[i].clientWidth - resize[i].offsetWidth;

        // 左边区域的最小宽度为32px
        if (moveLen < 32) moveLen = 32;
        //右边区域最小宽度为150px
        if (moveLen > maxT - 150) moveLen = maxT - 150;

        // 设置左侧区域的宽度
        resize[i].style.left = moveLen;

        for (let j = 0; j < left.length; j++) {
          left[j].style.width = moveLen + 'px';
          mid[j].style.width = (box[i].clientWidth - moveLen - 10) + 'px';
        }
      };
      // 鼠标松开事件
      document.onmouseup = function (evt) {
        //颜色恢复
        resize[i].style.background = '#d6d6d6';
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
  width: calc(20% - 10px);
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
  width: 68%; /*右侧初始化宽度*/
  min-height: 661px;
  height: 100%;
}
.quickIcon {
  padding: 0;
  text-align: center;
  height: 44px;
  width: auto
}

</style>
