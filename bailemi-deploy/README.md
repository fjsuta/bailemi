# 🎵 百米乐音乐平台 - 全能部署系统

支持多种部署方式的现代化开源音乐平台

---

## 📋 目录

1. [支持的部署方式](#-支持的部署方式)
2. [快速开始](#-快速开始)
3. [部署方式详细说明](#-部署方式详细说明)
4. [常见问题](#-常见问题)
5. [技术支持](#-技术支持)

---

## 🚀 支持的部署方式

| 方式 | 状态 | 适用环境 | 难度 |
|------|------|----------|------|
| 通用部署 | ✅ | 任意Linux | ⭐ |
| 宝塔面板 | ✅ | CentOS/Ubuntu | ⭐ |
| 小皮面板 | ✅ | Windows/Linux | ⭐ |
| aaPanel | ✅ | CentOS/Ubuntu | ⭐ |
| CyberPanel | ✅ | CentOS/Ubuntu | ⭐⭐ |
| DirectAdmin | ✅ | CentOS/CloudLinux | ⭐⭐ |
| Docker部署 | ✅ | 任意系统 | ⭐⭐ |
| 纯环境部署 | ✅ | 纯净Linux | ⭐⭐⭐ |

---

## ⚡ 快速开始

### 方式一：自动检测部署（推荐）

```bash
# 1. 进入部署目录
cd /workspace/bailemi-deploy

# 2. 赋予执行权限
chmod +x *.sh docker/*.sh

# 3. 自动检测并部署
bash panel_detector.sh
bash deploy.sh
```

### 方式二：指定面板部署

```bash
# 宝塔面板
bash deploy_baota.sh

# 小皮面板
bash deploy_xp.sh

# aaPanel
bash deploy_aapanel.sh

# CyberPanel
bash deploy_cyberpanel.sh

# Docker部署
cd docker && bash deploy_docker.sh

# 纯Linux环境
bash deploy_native.sh
```

---

## 📖 部署方式详细说明

### 1️⃣ 通用部署

**适用场景**：任意Linux环境，不限制面板

**命令**：
```bash
bash deploy.sh
```

**功能**：
- ✅ 自动检测系统环境
- ✅ 自动检测运维面板
- ✅ 智能配置安装路径
- ✅ 自动设置权限
- ✅ 生成Nginx配置
- ✅ 提供安装指引

---

### 2️⃣ 宝塔面板

**适用场景**：宝塔面板环境

**命令**：
```bash
bash deploy_baota.sh
```

**功能**：
- ✅ 检测宝塔面板
- ✅ 获取可用PHP版本
- ✅ 检查PHP扩展
- ✅ 自动创建站点配置
- ✅ 生成宝塔兼容Nginx配置
- ✅ 设置正确目录权限

**宝塔面板操作步骤**：
1. 在宝塔中添加站点
2. 设置根目录为 `/www/wwwroot/bailemi/public`
3. 配置PHP和数据库
4. 访问站点完成安装向导

---

### 3️⃣ 小皮面板

**适用场景**：小皮面板(Xp.cn)环境

**命令**：
```bash
bash deploy_xp.sh
```

**功能**：
- ✅ 检测小皮面板
- ✅ 自动配置站点
- ✅ 生成Nginx配置
- ✅ 权限自动设置

---

### 4️⃣ aaPanel (国际版宝塔)

**适用场景**：aaPanel国际版面板

**命令**：
```bash
bash deploy_aapanel.sh
```

**功能**：
- ✅ 检测aaPanel
- ✅ 国际化部署配置
- ✅ 符合国际版路径规范

---

### 5️⃣ CyberPanel

**适用场景**：CyberPanel + OpenLiteSpeed

**命令**：
```bash
bash deploy_cyberpanel.sh
```

**功能**：
- ✅ 检测CyberPanel
- ✅ OpenLiteSpeed配置
- ✅ LiteSpeed重写规则
- ✅ 正确的用户权限

---

### 6️⃣ Docker部署

**适用场景**：任意系统，追求快速部署

**命令**：
```bash
cd docker
bash deploy_docker.sh
```

**包含服务**：
- 🐳 Nginx (Web服务器)
- 🐳 PHP 8.2-FPM (应用服务)
- 🐳 MySQL 8.0 (数据库)
- 🐳 Redis 7 (缓存)
- 🐳 Node.js 18 (前端构建)
- 🐳 phpMyAdmin (数据库管理)

**管理命令**：
```bash
# 启动服务
docker compose up -d

# 查看状态
docker compose ps

# 查看日志
docker compose logs -f

# 重启服务
docker compose restart

# 停止服务
docker compose down

# 进入容器
docker exec -it bailemi-php sh
```

---

### 7️⃣ 纯环境部署

**适用场景**：纯净Linux系统，无任何面板

**命令**：
```bash
bash deploy_native.sh
```

**自动安装**：
- ✅ Nginx
- ✅ PHP 8.2+ (含扩展)
- ✅ MariaDB/MySQL
- ✅ Redis
- ✅ Node.js 18+

**支持系统**：
- Ubuntu 20.04+
- Debian 11+
- CentOS 7+
- AlmaLinux 8+
- RockyLinux 8+
- Alpine Linux

---

## 📁 目录结构

```
bailemi-deploy/
├── panel_detector.sh          # 面板自动检测
├── deploy.sh                  # 通用部署脚本
├── deploy_baota.sh            # 宝塔专用
├── deploy_xp.sh               # 小皮专用
├── deploy_aapanel.sh          # aaPanel专用
├── deploy_cyberpanel.sh       # CyberPanel专用
├── deploy_native.sh           # 纯Linux部署
├── README.md                  # 本文档
└── docker/
    ├── deploy_docker.sh       # Docker部署脚本
    ├── docker-compose.yml     # 服务编排配置
    ├── nginx/
    │   └── conf.d/
    │       └── bailemi.conf   # Nginx配置
    └── php/
        └── Dockerfile         # PHP容器镜像
```

---

## 🔧 常见问题

### Q1：权限不足怎么办？

```bash
# 赋予所有脚本执行权限
chmod +x *.sh docker/*.sh

# 以root运行
sudo bash deploy.sh
```

### Q2：部署后访问502错误？

检查PHP-FPM是否在运行：
```bash
systemctl status php-fpm
systemctl restart php-fpm
```

### Q3：如何重置安装？

```bash
# 删除安装锁文件
rm /path/to/bailemi/install.lock

# 重新访问安装向导
```

### Q4：Docker部署MySQL连不上？

等待MySQL完全启动：
```bash
docker compose logs -f mysql
```

默认连接信息：
- 主机: `mysql` (容器间通信)
- 端口: `3306`
- 用户: `bailemi`
- 密码: `bailemi_2024`

### Q5：如何配置HTTPS？

**面板方式**：在面板中配置Let's Encrypt证书

**Docker方式**：
```yaml
# 在docker-compose.yml中配置SSL证书挂载
volumes:
  - ./ssl:/etc/nginx/ssl
```

---

## 📊 环境要求

### 最低配置
- CPU: 1核
- 内存: 512MB
- 磁盘: 10GB
- 系统: Linux任意发行版

### 推荐配置
- CPU: 2核+
- 内存: 2GB+
- 磁盘: 40GB+
- 系统: Ubuntu 20.04+ / CentOS 7+

---

## 🤝 贡献

欢迎提交Issue和Pull Request！

---

## 📜 许可证

Apache License 2.0

---

## 📞 技术支持

- 文档: https://docs.bailemi.com
- 支持: https://support.bailemi.com
- GitHub: https://github.com/bailemi
- Email: support@bailemi.com

---

## ⭐ Star History

如果本项目对您有帮助，请给我们一个Star！

---

**🎉 开始您的音乐平台之旅！**
