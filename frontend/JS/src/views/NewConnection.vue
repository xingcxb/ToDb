<template>
<!--  //540*72-->
  <a-image :width="100" :height="72" src="https://zos.alipayobjects.com/rmsportal/jkjgkEfvpUPVyRjUImniVslZfWPnJuuZ.png"/>
  <br/>
  <a-row>
    <a-col :span="8">连接名:</a-col>
    <a-col :span="16">
      <a-input v-model:value="value" placeholder="连接别名" />
    </a-col>
  </a-row>
  <br/>
  <a-row>
    <a-col :span="8">主机:</a-col>
    <a-col :span="16">
      <a-input v-model:value="value" placeholder="连接地址" />
    </a-col>
  </a-row>
  <br/>
  <a-row>
    <a-col :span="8">端口</a-col>
    <a-col :span="16">
      <a-input v-model:value="value" placeholder="连接端口" />
    </a-col>
  </a-row>
  <br/>
  <a-row>
    <a-col :span="8">用户名:</a-col>
    <a-col :span="16">
      <a-input v-model:value="value" />
    </a-col>
  </a-row>
  <br/>
  <a-row>
    <a-col :span="8">密码:</a-col>
    <a-col :span="16">
      <a-input-password v-model:value="value" />
    </a-col>
  </a-row>
</template>

<script>
export default {
  name: "NewConnection"
}
</script>

<style scoped>

</style>