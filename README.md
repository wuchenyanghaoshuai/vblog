# Vblog博客平台

## 项目简介
Vblog 是一个基于Go语言开发的简洁博客平台，支持用户注册、登录、文章发布、令牌管理等功能，适合个人或小型团队搭建自己的博客系统。

---

## 主要功能模块

### 1. 用户模块（apps/user）
- 用户注册、登录、信息查询
- 密码加密存储，支持标签和角色管理

### 2. 博客模块（apps/blog）
- 文章的增删改查（CRUD）
- 支持标签、摘要、状态（草稿/已发布）管理
- 文章发布与状态变更

### 3. Token模块（apps/token）
- 用户登录后颁发访问令牌和刷新令牌
- 令牌校验、撤销，支持过期时间管理

---

## 数据库结构

- users：用户表，包含用户名、密码、标签、角色等
- tokens：令牌表，包含access_token、refresh_token、过期时间等
- blogs：博客文章表，包含标题、作者、内容、标签、状态等

详细建表SQL见 `doc/table/table.sql`

---

## 启动和使用方法

1. 安装依赖：
   ```
   go mod tidy
   ```
2. 启动服务：
   ```
   go run main.go start
   ```
3. 配置文件：
   - 配置文件位于 `conf/application.yml`，可根据实际情况修改数据库等参数

---

## 目录结构说明

- apps/         业务模块（blog、user、token）
- cmd/          启动和初始化命令
- common/       通用工具和基础库
- conf/         配置文件及加载逻辑
- doc/          项目文档和数据库表结构
- etc/          其他配置
- exception/    异常处理
- ioc/          依赖注入相关
- middleware/   中间件
- response/     统一响应结构
- test/         测试用例
- main.go       项目入口

---

## API接口说明（简要）

### 用户模块
- 创建用户：CreateUser
- 查询用户：QueryUser

### 博客模块
- 查询文章列表：QueryBlog
- 查询文章详情：DescribeBlog
- 创建文章：CreateBlog
- 更新文章：UpdateBlog
- 发布/变更状态：UpdateBlogStatus
- 删除文章：DeleteBlog

### Token模块
- 颁发令牌：IssueToken
- 撤销令牌：RevolkToken
- 校验令牌：ValidateToken

详细参数和返回值请参考各模块 `interface.go` 文件。

---

## 后续建议与改进
- 增加前端页面，支持Web可视化管理
- 增加更多文章分类、评论、点赞等功能
- 完善接口文档和单元测试
- 支持Docker一键部署

---

如有问题或建议，欢迎提issue或联系作者。
```
go run main.go start
```