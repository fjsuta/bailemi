# 百米乐安装向导 - 宝塔面板一键部署

## 📦 包含文件

```
bailemi-install/
├── install.php          # 主安装向导入口
├── install_step1.php     # 环境检测
├── install_step2.php     # 数据库配置
├── install_step3.php     # 数据表与管理员初始化
├── install_step4.php     # 风格与演示数据选择
├── install_step5.php     # 完成安装
├── install.sql          # 数据库结构
└── README.md           # 使用说明
```

## 🚀 快速开始

### 方式一：通过浏览器访问安装向导

1. 将整个 `bailemi-install` 文件夹上传到宝塔面板的网站根目录
2. 在浏览器中访问：`http://yourdomain.com/install/install.php`
3. 按照向导提示完成安装

### 方式二：通过命令行一键安装

```bash
cd /www/wwwroot/bailemi
sudo bash setup.sh
```

## 📋 安装步骤

### 步骤 1：环境检测

系统自动检测以下环境：

- ✅ **PHP 版本**：需要 PHP 7.4+
- ✅ **Node.js 版本**：可选，用于前端构建
- ✅ **目录权限**：网站根目录、配置目录、上传目录、缓存目录
- ✅ **PHP 扩展**：PDO、MySQLi、cURL、GD、mbstring、OpenSSL、JSON
- ✅ **PHP 函数**：必需函数可用性

#### 如何修复环境问题

**PHP 版本过低**
```bash
# 宝塔面板操作
# 软件商店 → PHP设置 → 安装版本 → 选择 PHP 8.0+
```

**扩展缺失**
```bash
# 宝塔面板操作
# 软件商店 → PHP设置 → 安装扩展 → 勾选所需扩展
```

**目录权限问题**
```bash
# SSH 登录服务器执行
chmod -R 755 /www/wwwroot/bailemi
chown -R www:www /www/wwwroot/bailemi
```

### 步骤 2：数据库配置

#### 在宝塔面板中创建数据库

1. 登录宝塔面板
2. 点击左侧菜单「数据库」
3. 点击「添加数据库」
4. 填写信息：
   - 数据库名：`bailemi`
   - 用户名：`bailemi`
   - 密码：（自动生成或手动设置）
   - 编码：`utf8mb4`
5. 点击「提交」

#### 填写安装向导表单

- **数据库地址**：通常填写 `localhost` 或 `127.0.0.1`
- **数据库端口**：MySQL 端口，默认 `3306`
- **数据库名称**：在宝塔面板创建的数据库名
- **用户名**：在宝塔面板创建的用户名
- **密码**：在宝塔面板设置的密码

点击「测试连接」验证连接是否成功。

### 步骤 3：创建数据表和管理员

#### 自动创建数据表

点击「创建数据表」按钮，系统将自动创建以下表：

- ✅ users - 用户表
- ✅ songs - 歌曲表
- ✅ albums - 专辑表
- ✅ artists - 歌手表
- ✅ playlists - 歌单表
- ✅ playlist_songs - 歌单歌曲关联表
- ✅ comments - 评论表
- ✅ favorites - 收藏表
- ✅ play_history - 播放记录表
- ✅ roles - 角色表
- ✅ permissions - 权限表
- ✅ follows - 关注表
- ✅ operation_logs - 操作日志表

#### 创建管理员账号

- **用户名**：管理员登录账号（至少3个字符）
- **密码**：管理员密码（至少6个字符）
- **邮箱**：用于找回密码（可选）
- **手机**：用于找回密码（可选）

### 步骤 4：风格与演示数据选择

#### 选择主题风格

- 🎨 紫色主题（默认）
- 🌊 蓝色主题
- 🌿 绿色主题
- 🌙 深色主题

#### 选择演示数据

- 🎵 演示歌曲（10首热门歌曲）
- 📋 演示歌单（4个推荐歌单）
- 💬 演示评论（示例评论）
- 👥 演示用户（示例用户）

### 步骤 5：完成安装

安装完成后，系统会：

1. ✅ 生成 `install.lock` 文件
2. ✅ 创建 `config.php` 配置文件
3. ✅ 导入选定的演示数据
4. ✅ 显示访问地址和管理后台地址

## 🔐 安全提示

### 立即执行的安全操作

1. **删除安装文件**
   ```bash
   rm -rf /www/wwwroot/bailemi/install*
   ```

2. **修改管理员密码**
   登录后台后立即修改默认管理员密码

3. **设置目录权限**
   ```bash
   chmod -R 644 /www/wwwroot/bailemi/config
   chmod 755 /www/wwwroot/bailemi/uploads
   chmod 755 /www/wwwroot/bailemi/runtime
   ```

4. **配置 SSL 证书**（强烈推荐）
   - 宝塔面板 → 网站 → 设置 → SSL
   - 选择 Let's Encrypt 或上传商业证书
   - 开启强制 HTTPS

## 📁 重要文件说明

| 文件 | 说明 | 权限建议 |
|------|------|----------|
| `install.lock` | 安装锁定文件，防止重复安装 | 644 |
| `config.php` | 数据库和应用配置 | 644 |
| `.env` | 环境变量配置 | 644 |
| `uploads/` | 用户上传文件目录 | 755 |
| `runtime/` | 缓存文件目录 | 755 |

## 🛠️ 常见问题

### Q1: 安装向导显示空白页面

**原因**：PHP 版本过低或内存限制不足

**解决**：
```bash
# 检查 PHP 版本
php -v

# 修改 php.ini
memory_limit = 256M
max_execution_time = 300
```

### Q2: 数据库连接失败

**检查项**：
1. 数据库服务是否启动
2. 用户名密码是否正确
3. 数据库是否已创建
4. 用户是否有权限访问数据库

```bash
# 测试 MySQL 连接
mysql -u bailemi -p -h localhost
```

### Q3: 权限不足

**解决**：
```bash
# 设置目录所有者
chown -R www:www /www/wwwroot/bailemi

# 设置目录权限
chmod -R 755 /www/wwwroot/bailemi
```

### Q4: 如何重新安装？

如果需要重新安装，删除 `install.lock` 文件即可：

```bash
rm -f /www/wwwroot/bailemi/install.lock
```

然后重新访问安装向导。

## 📞 获取帮助

- 📖 官方文档：https://docs.bailemi.com
- 💬 技术支持：https://support.bailemi.com
- 🐛 问题反馈：https://github.com/bailemi/bailemi/issues
- 📧 邮箱：support@bailemi.com

## 📄 开源协议

本项目采用 Apache 2.0 协议开源。

---

**版本**: 1.0.0  
**最后更新**: 2024  
**支持**: CentOS 7+ / Ubuntu 18+ / 宝塔面板 7.0+
