# 🎵 百米乐 - 现代开源音乐平台

![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)
![Version](https://img.shields.io/badge/version-1.0.0-green.svg)
![Vue](https://img.shields.io/badge/Vue-3.4-blue.svg)
![Node](https://img.shields.io/badge/Node-18%2B-green.svg)

> 🎉 百米乐是一款现代化的开源音乐平台，支持多种部署方式，开箱即用！

## ✨ 功能特性

### 🎨 用户界面
- **🌙 黑暗模式**：支持浅色/深色模式切换，自动记忆用户偏好
- **📱 响应式设计**：完美适配桌面端、移动端和平板设备
- **✨ 玻璃拟态UI**：现代化设计，流畅的动画效果
- **🔔 系统通知**：实时的操作反馈通知

### 👤 用户系统
- **📝 用户注册**：支持用户名、邮箱、手机号注册
- **🔐 安全登录**：JWT Token 认证，支持 Token 刷新
- **👤 个人中心**：完善的个人资料管理
- **📸 头像上传**：支持自定义头像
- **⭐ 我的收藏**：收藏喜欢的歌曲
- **📜 最近播放**：记录播放历史

### 🎵 音乐功能
- **🎶 歌曲播放**：流畅的音乐播放体验
- **📋 歌单管理**：创建和管理个人歌单
- **🔍 智能搜索**：快速搜索歌曲、歌手、歌单
- **🏆 排行榜**：热门歌曲排行榜
- **💬 歌曲评论**：与其他用户互动

### 🛡️ 安全特性
- **🔒 RBAC权限控制**：基于角色的权限管理
- **🛡️ SQL注入防护**：参数化查询，防止SQL注入
- **🔐 密码加密**：bcrypt加密存储
- **⏰ JWT过期机制**：安全的Token管理
- **🚫 CSRF防护**：防止跨站请求伪造

### 🚀 部署特性
- **🌐 多面板支持**：宝塔、小皮、aaPanel、CyberPanel
- **🐳 Docker部署**：一键容器化部署
- **🖥️ 纯环境部署**：支持无面板纯净Linux环境
- **🔄 自动检测**：智能识别服务器环境
- **📦 一键安装**：提供Web安装向导

---

## 🏗️ 技术栈

### 前端技术
- **Vue 3** - 渐进式JavaScript框架
- **Vite** - 下一代前端构建工具
- **Tailwind CSS** - 实用优先的CSS框架
- **Pinia** - Vue状态管理
- **Vue Router** - Vue官方路由管理
- **Axios** - HTTP请求库

### 后端技术
- **Go + Gin** - 高性能Web框架
- **Spring Boot** - Java企业级框架（可选）
- **JWT** - 身份认证
- **bcrypt** - 密码加密
- **GORM** - Go ORM库
- **Spring Security** - Java安全框架（可选）

### 数据库
- **MySQL 8.0** - 关系型数据库
- **Redis** - 缓存和会话存储

### 部署
- **Nginx** - Web服务器
- **Docker** - 容器化平台
- **Docker Compose** - 容器编排

---

## 📦 项目结构

```
bailemi/
├── bailemi-frontend/          # Vue 3 前端项目
│   ├── src/
│   │   ├── components/         # 组件
│   │   │   ├── music/          # 音乐组件
│   │   │   └── *.vue           # UI组件
│   │   ├── views/              # 页面视图
│   │   │   ├── auth/           # 认证页面
│   │   │   ├── music/          # 音乐页面
│   │   │   └── profile/        # 个人中心
│   │   ├── stores/             # Pinia状态管理
│   │   ├── router/             # 路由配置
│   │   └── utils/              # 工具函数
│   ├── tailwind.config.js      # Tailwind配置
│   ├── vite.config.js          # Vite配置
│   └── package.json            # 依赖配置
│
├── bailemi-gateway/           # Go Gin API网关
│   ├── cmd/                    # 入口文件
│   │   ├── server/             # 服务入口
│   │   └── demo/               # 示例代码
│   ├── internal/               # 内部包
│   │   ├── handler/            # 控制器层
│   │   ├── service/            # 业务逻辑层
│   │   ├── repository/         # 数据访问层
│   │   ├── model/              # 数据模型
│   │   └── middleware/         # 中间件
│   ├── pkg/                    # 公共包
│   │   ├── jwt/                # JWT工具
│   │   ├── errors/             # 错误处理
│   │   └── response/           # 响应封装
│   └── configs/                # 配置文件
│
├── bailemi-spring/             # Spring Boot 后端（可选）
│   └── src/main/java/
│       └── com/bailemi/
│           ├── config/         # 配置类
│           ├── controller/      # 控制器
│           ├── service/         # 服务层
│           ├── repository/      # 数据访问
│           ├── entity/          # 实体类
│           └── security/        # 安全配置
│
├── bailemi-install/            # PHP安装向导
│   ├── install.php             # 安装入口
│   ├── install_step*.php       # 安装步骤
│   └── install.sql             # 数据库结构
│
├── bailemi-deploy/             # 部署脚本
│   ├── deploy.sh               # 通用部署
│   ├── deploy_baota.sh         # 宝塔部署
│   ├── deploy_docker.sh        # Docker部署
│   ├── panel_detector.sh       # 面板检测
│   └── docker/                  # Docker配置
│
├── bailemi-infra/              # 基础设施
│   └── docker/                 # Docker编排
│
└── README.md                   # 项目文档
```

---

## 🚀 快速开始

### 方式一：使用安装向导

1. 下载源码到Web目录
2. 访问 `http://your-domain/install.php`
3. 按照向导提示完成安装
4. 开始使用！

### 方式二：Docker部署

```bash
cd bailemi-deploy/docker
bash deploy_docker.sh
```

### 方式三：宝塔面板部署

```bash
cd bailemi-deploy
bash deploy_baota.sh
```

### 方式四：通用部署

```bash
cd bailemi-deploy
bash deploy.sh
```

---

## 🎨 界面预览

### 登录页面
- 玻璃拟态设计
- 渐变背景动画
- 表单验证
- 响应式布局

### 首页
- 热门歌曲展示
- 推荐歌单
- 快速搜索
- 音乐播放栏

### 个人中心
- 用户资料展示
- 头像上传
- 我的收藏
- 最近播放
- 主题切换

---

## 🔧 配置说明

### 环境变量

创建 `.env.production` 文件：

```env
VITE_API_BASE_URL=https://api.yourdomain.com
VITE_APP_TITLE=百米乐
```

### 数据库配置

安装向导会自动配置，或手动修改：

```php
// config.php
return [
    'db' => [
        'host' => 'localhost',
        'port' => 3306,
        'database' => 'bailemi',
        'username' => 'bailemi',
        'password' => 'your_password',
    ],
];
```

### API 端点

- `POST /api/v1/auth/register` - 用户注册
- `POST /api/v1/auth/login` - 用户登录
- `GET /api/v1/song/hot` - 热门歌曲
- `GET /api/v1/playlist/recommended` - 推荐歌单
- `GET /api/v1/user/profile` - 用户资料
- `PUT /api/v1/user/profile` - 更新资料

---

## 🛡️ 安全建议

1. **修改默认密码**：安装后立即修改管理员密码
2. **启用HTTPS**：生产环境务必启用SSL证书
3. **配置防火墙**：仅开放必要端口
4. **定期备份**：定期备份数据库和文件
5. **删除安装文件**：安装完成后删除 `install.php`

---

## 🤝 贡献指南

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

---

## 📄 开源协议

本项目采用 Apache License 2.0 协议开源。

```
Apache License
Version 2.0, January 2004
http://www.apache.org/licenses/

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
```

---

## 📞 技术支持

- 📖 官方文档：https://docs.bailemi.com
- 💬 问题反馈：https://github.com/bailemi/bailemi/issues
- 📧 邮箱：support@bailemi.com

---

## 🙏 致谢

- [Vue.js](https://vuejs.org/) - 渐进式JavaScript框架
- [Tailwind CSS](https://tailwindcss.com/) - 实用优先CSS框架
- [Gin](https://github.com/gin-gonic/gin) - Go Web框架
- [Spring Boot](https://spring.io/projects/spring-boot) - Java框架
- 所有开源贡献者！

---

## ⭐ Star History

如果这个项目对您有帮助，请给我们一个 Star！

[![Star History Chart](https://api.star-history.com/svg?repos=bailemi/bailemi&type=Timeline)](https://star-history.com/#bailemi/bailemi&Timeline)

---

<p align="center">
  <strong>Made with ❤️ by Bailemi Team</strong>
  <br>
  <strong>Released under the Apache-2.0 License</strong>
</p>
