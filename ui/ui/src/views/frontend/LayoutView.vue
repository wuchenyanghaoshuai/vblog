<template>
  <div>
    <TopBar></TopBar>
    Frontend Layout
    <div>博客前台布局页面</div>
    <!-- 需要显示子路由的页面必须要添加RouteView-->
  </div>
  <router-view></router-view>
</template>

<script setup>
import TopBar from "@/components/TopBar.vue";

</script>

<style lang="scss" scoped></style>
