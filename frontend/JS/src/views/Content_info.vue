<template>
123
</template>

<script setup>

import {onBeforeMount} from "vue";
import {useRouter} from "vue-router";

const router = useRouter();

// 初始化挂载前的函数
onBeforeMount(()=>{
  // 获取路由传递的参数
  let parameter = router.currentRoute.value.query.key
  console.log("这是info页面")
  console.log(parameter)
  // 从redis获取数据
  
})

</script>

<style scoped>

</style>