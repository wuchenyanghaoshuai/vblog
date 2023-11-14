import { createRouter, createWebHistory } from 'vue-router'
// import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    //登陆页
    {
      path: '/login',
      name: 'Login',
      component: () => import('../views/login/LoginView.vue')
    },
    //前端页面
    {
      path: '/frontend',
      name: 'FrontendLayout',
      component: () => import('../views/frontend/LayoutView.vue'),
       //重定向到blogs
      children: [
        {
          path: 'blogs',
          name: 'FrontendBlogs',
          component: () => import('../views/frontend/blog/ListView.vue')
        }
      ]
    },
    //后端页面
    {
      path: '/backend',
      name: 'BackendLayout',
      component: () => import('../views/backend/LayoutView.vue'),
      redirect: {name: 'BackendBlogs'},
      children: [
        {
          path: 'blogs',
          name: 'BackendBlogs',
          component: () => import('../views/backend/blog/ListView.vue')
        },
        {
          path: 'blogs_edit',
          name: 'BlogEdit',
          component: () => import('../views/backend/blog/EditView.vue')
        },
        {
          path: 'comments',
          name: 'CommentList',
          component: () => import('../views/backend/comment/ListPage.vue')
        },
      ]
    }
  ]
})

export default router
