<template>
  <router-view/>
</template>

<style lang="scss">
@font-face {
  font-family: 'iconfont';
  src: url('assets/iconfont.woff2') format('woff2'),
  url('assets/iconfont.woff') format('woff'),
  url('assets/iconfont.ttf') format('truetype');
}

.iconfont {
  font-family: "iconfont", sans-serif !important;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

html, body {
  background: #f0f3f7;
  margin: 0;
  padding: 0;
  overflow: hidden;
}

#app {
  font-family: Helvetica, 'Hiragino Sans GB', 'Microsoft YaHei UI', 'Microsoft Yahei', '微软雅黑', Arial, sans-serif;

  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}
</style>