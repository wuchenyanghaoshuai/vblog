<template>
  <div class="top-bar">
    <div style="margin-left: 12px">
      老鱼的博客系统
    </div>

    <div>
      <a-button type="text" @click=" router.push({name: 'FrontendLayout'})"> 前台</a-button>
      <a-button type="text" @click=" router.push({name: 'BackendLayout'})"> 后台</a-button>
      <a-button type="text" v-if="!state.is_login" @click="login"> <icon-import/>登陆</a-button>


      <a-button type="text">
        <span style="margin-right: 12px" @click="logout" v-if="state.is_login"> <icon-export/>退出登陆</span>
      </a-button>


    </div>
  </div>
</template>

<script setup>
import {useRouter} from "vue-router";
console.log("打印是否登陆",state.value.is_login)
const router = useRouter()
import { state } from '../stores/app.js'
console.log("state",state.is_login)

const logout = () => {
  state.value.is_login = false
  router.push({name: 'Login'})
  //或者是调用后段的token销毁接口 当作作业来做
}
const login = () => {
  router.push({name: 'Login'})
  state.value.is_login = true

}
</script>

<style lang="css" scoped>
.top-bar{
  height: 60px;
  border-bottom: 1px solid rgb(229,230,235);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 12px;
}
</style>