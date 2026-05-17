# 百米乐 (Bailemi) - 现代开源音乐平台

![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)

一个现代化、美观、功能完整的开源音乐平台，包含完整的前后端实现。

## 技术栈

### 前端
- Vue 3 + Vite
- Tailwind CSS
- Pinia (状态管理)
- Vue Router
- Axios

### 后端（Go + Gin）
- Gin Web Framework
- GORM ORM
- JWT 认证
- MySQL 数据库

### 后端（Spring Boot 可选）
- Spring Boot 3.2
- Spring Security
- JWT
- Spring Data JPA
- MySQL

## 功能特性

### 用户认证
- 用户注册（用户名/邮箱/手机号
- 用户登录
- JWT Token 认证
- 路由守卫
- Token 刷新

### 个人中心
- 用户资料展示
- 头像上传和编辑
- 个人简介编辑
- VIP 等级展示
- 粉丝/关注统计

### 音乐功能
- 热门歌曲推荐
- 推荐歌单
- 音乐播放
- 播放历史
- 我的收藏

### 用户管理
- 用户查询（管理员）
- 用户封禁/解封（管理员）
- 按用户名、注册时间筛选

## 项目结构

```
bailemi-frontend/
├── src/
│   ├── components/
│   │   ├── auth/
│   │   ├── layout/
│   │   └── music/
│   ├── views/
│   │   ├── auth/
│   │   │   ├── Login.vue
│   │   │   └── Register.vue
│   │   ├── music/
│   │   │   └── Home.vue
│   │   └── profile/
│   │       └── Profile.vue
│   ├── stores/
│   │   ├── auth.js
│   │   └── music.js
│   ├── router/
│   │   └── index.js
│   ├── utils/
│   │   └── api.js
│   ├── App.vue
│   ├── main.js
│   └── style.css
├── index.html
├── vite.config.js
├── tailwind.config.js
├── postcss.config.js
└── package.json

bailemi-spring/
├── src/main/java/com/bailemi/
│   ├── config/
│   ├── controller/
│   ├── dto/
│   ├── entity/
│   ├── repository/
│   ├── security/
│   └── service/
└── pom.xml
```

## 快速开始

### 前端开发

```bash
cd bailemi-frontend
npm install
npm run dev
```

访问 http://localhost:5173

### Go 后端开发

```bash
cd bailemi/gateway
go mod tidy
go run cmd/demo/demo.go
```

访问 http://localhost:8080

### Spring Boot 后端开发

```bash
cd bailemi-spring
mvn spring-boot:run
```

访问 http://localhost:8080

## 开源协议

本项目采用 [Apache License 2.0](./LICENSE) 协议开源。

## LICENSE

Copyright 2024 Bailemi

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
