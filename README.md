# 🎵 百米乐 (Bailemi) - 开源音乐平台

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](./LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org/)
[![Vue Version](https://img.shields.io/badge/Vue-3.4+-4FC08D.svg)](https://vuejs.org/)
[![Spring Boot](https://img.shields.io/badge/Spring%20Boot-3.2-6DB33F.svg)](https://spring.io/projects/spring-boot)

> 百米乐是一个现代化的开源音乐平台，提供音乐播放、用户社区、第三方登录等完整功能。支持多种服务器面板一键部署，开箱即用。

[🌐 在线演示](https://bailemi.com) | [📖 文档](https://docs.bailemi.com) | [🐛 问题反馈](https://github.com/fjsuta/bailemi/issues)

---

## ✨ 功能特性

### 🔐 用户认证系统
- **本地账号注册/登录**：支持用户名、邮箱、手机号
- **JWT Token 认证**：安全的无状态认证机制
- **第三方 OAuth 登录**：
  - 🔵 Google
  - 🟢 Microsoft
  - 🍎 Apple
  - 💚 微信
  - 🐧 QQ
- **账号安全**：
  - 修改密码
  - 账号注销（需密码确认）
  - 第三方账号绑定/解绑

### 🎵 音乐功能
- **热门歌曲推荐**：智能推荐热门音乐
- **歌曲搜索**：支持关键词搜索歌曲、歌手、歌单
- **音乐播放**：流畅的在线音乐播放体验
- **歌单管理**：创建、编辑个人歌单
- **收藏功能**：收藏喜欢的歌曲
- **播放历史**：记录播放记录

### 👥 社交功能
- **用户关注**：关注喜欢的用户
- **粉丝系统**：查看粉丝和关注列表
- **评论互动**：对歌曲发表评论

### 🎨 界面特色
- **暗色/亮色模式**：支持主题切换
- **毛玻璃效果**：现代 Glassmorphism 设计
- **响应式布局**：适配各种设备

### 🔧 管理功能
- **用户管理**：管理员可查询、封禁用户
- **内容管理**：歌曲、歌单、评论管理
- **系统配置**：第三方登录配置开关
- **自动更新检查**：每天检查仓库更新，支持一键拉取代码
- **图形化公安备案**：内置公安备案管理工具
- **多云对象存储**：支持阿里云、腾讯云、七牛云、又拍云等国内主流云服务商及本地存储

---

## 🏗️ 技术架构

### 前端技术栈
| 技术 | 说明 | 版本 |
|------|------|------|
| Vue 3 | 渐进式 JavaScript 框架 | 3.4+ |
| Vite | 新一代前端构建工具 | 5.0+ |
| Tailwind CSS | 原子化 CSS 框架 | 3.4+ |
| Pinia | Vue 状态管理 | 2.1+ |
| Vue Router | Vue 官方路由 | 4.2+ |

### 后端技术栈（Go + Gin）
| 技术 | 说明 |
|------|------|
| Gin | 高性能 Go Web 框架 |
| GORM | Go 语言 ORM 库 |
| JWT | JSON Web Token 认证 |
| MySQL | 关系型数据库 |
| Redis | 缓存（可选） |

### 对象存储支持
- 📁 **本地存储**：服务器本地文件系统
- ☁️ **阿里云 OSS**：阿里云对象存储
- ☁️ **腾讯云 COS**：腾讯云对象存储
- ☁️ **七牛云**：七牛云对象存储
- ☁️ **又拍云**：又拍云存储
- ☁️ **华为云 OBS**：华为云对象存储
- ☁️ **百度云 BOS**：百度云对象存储

---

## 📁 项目结构

```
bailemi/
├── bailemi-frontend/          # Vue 3 前端项目
│   ├── src/
│   │   ├── components/         # 公共组件
│   │   ├── views/             # 页面视图
│   │   │   ├── auth/          # 认证页面
│   │   │   ├── music/         # 音乐页面
│   │   │   ├── admin/         # 管理后台
│   │   │   └── profile/        # 个人中心
│   │   ├── stores/            # Pinia 状态管理
│   │   ├── router/            # 路由配置
│   │   └── utils/             # 工具函数
│   └── ...
│
├── bailemi/                    # Go Gin 后端
│   ├── gateway/                # API 网关
│   │   ├── cmd/               # 入口文件
│   │   ├── internal/          # 内部包
│   │   │   ├── handler/       # 控制器
│   │   │   ├── service/       # 业务逻辑
│   │   │   ├── repository/    # 数据访问
│   │   │   ├── model/         # 数据模型
│   │   │   ├── middleware/    # 中间件
│   │   │   └── tasks/         # 定时任务
│   │   └── pkg/               # 公共工具包
│   │
│   └── frontend/              # 原生前端（可选）
│
├── bailemi-spring/            # Spring Boot 后端（可选）
│   └── src/main/java/com/bailemi/
│
├── bailemi-deploy/            # 部署脚本
│   ├── deploy_baota.sh        # 宝塔面板部署
│   ├── deploy_xp.sh           # XP 面板部署
│   ├── deploy_aapanel.sh      # aaPanel 部署
│   ├── deploy_cyberpanel.sh   # CyberPanel 部署
│   ├── deploy_native.sh        # 原生 Linux 部署
│   └── docker/                # Docker 部署
│
└── bailemi-install/            # 安装向导
    └── install.php            # PHP 安装向导
```

---

## 🚀 快速开始

### 方式一：Docker 部署（推荐）

```bash
# 克隆项目
git clone https://github.com/fjsuta/bailemi.git
cd bailemi/bailemi-deploy/docker

# 启动服务
docker-compose up -d
```

### 方式二：宝塔/aaPanel/XP 面板

```bash
# 在服务器上执行
cd bailemi/bailemi-deploy
chmod +x deploy_baota.sh
./deploy_baota.sh
```

### 方式三：手动部署

#### 1. 前端部署

```bash
cd bailemi-frontend

# 安装依赖
npm install

# 开发模式
npm run dev

# 生产构建
npm run build
```

#### 2. Go 后端部署

```bash
cd bailemi/gateway

# 安装依赖
go mod tidy

# 运行服务
go run cmd/demo/demo.go
```

#### 3. PHP 安装向导

访问 `http://your-domain/bailemi-install/install.php` 完成数据库配置。

---

## 🔌 API 接口

### 认证接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/v1/auth/register` | 用户注册 |
| POST | `/api/v1/auth/login` | 用户登录 |
| POST | `/api/v1/auth/refresh` | 刷新 Token |
| POST | `/api/v1/auth/logout` | 退出登录 |
| GET | `/api/v1/auth/oauth/config` | 获取 OAuth 配置 |

### OAuth 接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/oauth/:provider/authorize` | OAuth 授权 |
| GET | `/api/v1/oauth/:provider/callback` | OAuth 回调 |
| POST | `/api/v1/oauth/bind` | 绑定账号 |
| DELETE | `/api/v1/oauth/bind/:provider` | 解除绑定 |
| GET | `/api/v1/oauth/bindings` | 获取绑定列表 |

### 用户接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/user/me` | 获取当前用户 |
| PUT | `/api/v1/user/profile` | 更新资料 |
| POST | `/api/v1/user/avatar` | 上传头像 |
| POST | `/api/v1/user/password/change` | 修改密码 |
| POST | `/api/v1/user/delete` | 注销账号 |

### 音乐接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/search` | 搜索歌曲 |
| GET | `/api/v1/song/hot` | 热门歌曲 |
| GET | `/api/v1/playlist` | 歌单列表 |

### 管理接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/v1/admin/update/check` | 检查更新 |
| POST | `/api/v1/admin/update/do` | 执行更新 |
| GET | `/api/v1/admin/icp` | 获取备案信息 |
| POST | `/api/v1/admin/icp` | 设置备案信息 |
| GET | `/api/v1/admin/storage` | 获取存储配置 |
| POST | `/api/v1/admin/storage` | 设置存储配置 |

---

## ⚙️ 配置说明

### OAuth 配置

在 `bailemi/gateway/config/oauth.json` 中配置第三方登录：

```json
{
  "google": {
    "enabled": true,
    "client_id": "your-client-id",
    "client_secret": "your-client-secret"
  },
  "microsoft": {
    "enabled": true,
    "client_id": "your-client-id",
    "client_secret": "your-client-secret"
  },
  "wechat": {
    "enabled": false,
    "app_id": "your-app-id",
    "app_secret": "your-app-secret"
  },
  "qq": {
    "enabled": false,
    "app_id": "your-app-id",
    "app_key": "your-app-key"
  }
}
```

### 存储配置

在 `bailemi/gateway/config/storage.json` 中配置对象存储：

```json
{
  "type": "local",
  "local": {
    "path": "./uploads"
  },
  "aliyun_oss": {
    "enabled": false,
    "access_key_id": "",
    "access_key_secret": "",
    "bucket": "",
    "region": ""
  },
  "tencent_cos": {
    "enabled": false,
    "secret_id": "",
    "secret_key": "",
    "bucket": "",
    "region": ""
  },
  "qiniu": {
    "enabled": false,
    "access_key": "",
    "secret_key": "",
    "bucket": "",
    "domain": ""
  }
}
```

### 环境变量

```env
# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your-password
DB_NAME=bailemi

# JWT 配置
JWT_SECRET=your-secret-key
JWT_EXPIRES_IN=7d

# OAuth 配置
OAUTH_CONFIG_PATH=./config/oauth.json

# 自动更新配置
UPDATE_CHECK_ENABLE=true
UPDATE_REPO_URL=https://github.com/fjsuta/bailemi

# 基础URL
BASE_URL=https://your-domain.com
```

---

## 📊 数据库表结构

| 表名 | 说明 |
|------|------|
| users | 用户表 |
| user_profiles | 用户资料表 |
| songs | 歌曲表 |
| playlists | 歌单表 |
| playlist_songs | 歌单歌曲关联表 |
| favorites | 收藏表 |
| play_history | 播放历史表 |
| follows | 关注关系表 |
| comments | 评论表 |
| oauth_providers | OAuth 绑定表 |
| permissions | 权限表 |
| roles | 角色表 |
| role_permissions | 角色权限关联表 |
| system_config | 系统配置表 |
| icp_info | 公安备案信息表 |
| update_logs | 更新日志表 |

---

## 🛠️ 开发指南

### 前端开发

```bash
# 进入前端目录
cd bailemi-frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建生产版本
npm run build

# 预览生产版本
npm run preview
```

### 后端开发 (Go)

```bash
# 进入后端目录
cd bailemi/gateway

# 安装依赖
go mod tidy

# 运行服务
go run cmd/demo/demo.go

# 运行测试
go test ./...
```

### 后端开发 (Spring Boot)

```bash
# 进入目录
cd bailemi-spring

# 构建项目
mvn clean package

# 运行服务
mvn spring-boot:run
```

---

## 🤝 贡献指南

1. **Fork 本仓库**
2. **创建特性分支** (`git checkout -b feature/amazing-feature`)
3. **提交更改** (`git commit -m 'Add some amazing feature'`)
4. **推送分支** (`git push origin feature/amazing-feature`)
5. **创建 Pull Request**

---

## 📄 开源协议

本项目基于 [Apache License 2.0](./LICENSE) 开源，你可以免费使用、修改和分发本项目。

---

## 🙏 致谢

- [Vue.js](https://vuejs.org/) - 渐进式 JavaScript 框架
- [Tailwind CSS](https://tailwindcss.com/) - 原子化 CSS 框架
- [Gin](https://gin-gonic.com/) - Go Web 框架
- [Spring Boot](https://spring.io/projects/spring-boot) - Spring 初始化框架
- [Vite](https://vitejs.dev/) - 新一代前端构建工具

---

## 📬 联系与反馈

- **GitHub Issues**: [提交问题](https://github.com/fjsuta/bailemi/issues)
- **邮箱**: replab@zohomail.cn
- **微博**: https://weibo.com/u/7799762062
- **软天空**: http://a.ruansky.com/user/8615546/

---

<div align="center">
  <p>如果你觉得这个项目对你有帮助，请给我一个 ⭐️</p>
  <p>Made with ❤️ by <a href="https://github.com/fjsuta">fjsuta</a></p>
</div>
