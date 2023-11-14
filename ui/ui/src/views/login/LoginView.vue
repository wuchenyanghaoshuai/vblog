<template>
  <div class="login-page">
    <div class="login-form">
      <a-form :model="loginForm" @submit="handleSubmit">
        <a-form-item hide-label>
          <span class="login-title">欢迎登陆xxx博客系统</span>
        </a-form-item>

        <a-form-item hide-label field="username" :rules="[{required:true,message:'请输入用户名'}]">
          <a-input v-model="loginForm.username" placeholder="请输入用户名" allow-clear>
            <template #prefix>
              <icon-user />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item  hide-label field="password" :rules="[{required:true,message:'请输入密码'},{minLength:5,message:'must be greater than 5 characters'}]">
          <a-input-password v-model="loginForm.password" placeholder="请输入密码" allow-clear>
            <template #prefix>
              <icon-lock />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item hide-label>
          <a-button html-type="submit" style="width: 100%" type="primary">登陆</a-button>
        </a-form-item>

      </a-form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import {LOGIN} from "@/api/token";
import { Message } from '@arco-design/web-vue';
import {useRouter} from "vue-router";
import {state} from "@/stores/app";

//路由对象
const router = useRouter()
//定义表单数据
//表单对应的响应式数据，提交给后段的数据
// 按照vblog login api 来设计
const loginForm = ref({
  username: '',
  password: ''
})

//表单数据提交函数
const handleSubmit = async (data) => {

  if (data.errors === undefined){
    try {
      const resp = await LOGIN(data.values)

       state.value.is_login = true
      state.value.token = resp.data
      router.push({name: 'BackendBlogs'})
    }catch (error){
      console.log(error)
    }
  }
}
</script>

<style lang="css" scoped>
.login-page {
  height: 100vh;
  width: 100vw;
  display: flex;
  justify-content: center;
  align-items: center;
  background-position: left bottom;
  background-image: url("../../image/login_bg.svg");
  background-repeat: no-repeat;
}

.login-form {
  display: flex;
  align-content: center;
  flex-direction: column;
  height: 400px;
  width: 400px;
}

.login-title {
  display: flex;
  width: 100%;
  justify-content: center;
  font-size: 16px;
  font-weight: 550;

}
</style>
