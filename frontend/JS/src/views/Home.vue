<template>
  <div class="box" ref="box">
    <div class="left">
      <!--左侧div内容-->
      <a-tree
          :tree-data="treeData"
      >
        <template #title="{ title, key }">
          <span v-if="key === '0-0-1-0'" style="color: #1890ff">{{ title }}</span>
          <template v-else>{{ title }}</template>
        </template>
      </a-tree>
    </div>
    <div class="resize" @mousedown="handleMouseMoveLine">
    </div>
    <div class="mid">
      <!--右侧div内容-->
    </div>
  </div>
</template>

<script setup>
import { defineComponent, ref, watch } from 'vue';



const treeData = [{
  title: 'parent 1',
  key: '0-0',
  children: [{
    title: 'parent 1-0',
    key: '0-0-0',
    disabled: true,
    children: [{
      title: 'leaf',
      key: '0-0-0-0',
      disableCheckbox: true,
    }, {
      title: 'leaf',
      key: '0-0-0-1',
    }],
  }, {
    title: 'parent 1-1',
    key: '0-0-1',
    children: [{
      key: '0-0-1-0',
      title: 'sss',
    }],
  }],
}];


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

<style scoped>
.box{
  min-height: 500px;
  height: 100%;
  width: 100%;
  overflow: hidden;
}

.left {
  width:calc(25% - 10px);
  min-height: 661px;
  height:100%;
  float: left;
  background: rgba(224, 225, 225, 0);
}

.resize{
  min-height: 661px;
  height:100%;
  width: 2px;
  float: left;
  border-right: 1px solid #d5d6d6;
}
.resize:hover{
  cursor:col-resize;
}
.mid{
  float: left;
  width: 68%;   /*右侧初始化宽度*/
  height: 100%;
  background: #fff;
}
</style>