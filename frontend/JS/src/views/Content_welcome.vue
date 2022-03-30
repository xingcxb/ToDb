<template>
  <a-row style="min-height: 661px;height: 100%;" justify="space-around" align="middle" type="flex">
    <a-col :span="12" :push="8">
      <img src="../../public/logo.png" alt="ToDB"/>
    </a-col>
    <a-col :span="12">
      <span class="app-name">ToDB </span><span class="tail">.app</span>
      <br/>
      <span class="other">Version 0.1.0</span><br/>
      <span class="other">Powered by
          <span class="open-source" @click='openBrower("https://github.com/wailsapp/wails")'>wails</span> and
          <span class="open-source" @click='openBrower("https://icons8.com")'>Icons8</span>
        </span>
    </a-col>
  </a-row>
</template>

<script setup>
function openBrower(v) {
  runtime.BrowserOpenURL(v);
}
</script>

<style scoped>
.app-name {
  font-size: 20px
}

.tail {
  font-size: 6px;
}

.other {
  font-size: 12px;
}

.open-source {
  font-size: 12px;
  color: #1565c0;
  cursor: pointer;
}
</style>