#!/bin/bash

#=======================================
# 百米乐前端项目 - 宝塔面板一键部署脚本
# 适用系统：CentOS 7+ / Ubuntu 18+ / 宝塔面板
# 作者：百米乐开发团队
#=======================================

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 项目配置
PROJECT_NAME="bailemi"
PROJECT_DIR="/www/wwwroot/bailemi"
DIST_DIR="$PROJECT_DIR/dist"
NGINX_CONFIG_DIR="/www/server/nginx/conf/vhost"
SITE_DOMAIN="bailemi.yourdomain.com"

# 打印函数
print_step() {
    echo -e "${BLUE}[步骤]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[成功]${NC} $1"
}

print_error() {
    echo -e "${RED}[错误]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[警告]${NC} $1"
}

print_info() {
    echo -e "${BLUE}[信息]${NC} $1"
}

# 检测系统
detect_os() {
    if [ -f /etc/redhat-release ]; then
        OS="CentOS"
        PKG_MANAGER="yum"
    elif [ -f /etc/lsb-release ]; then
        OS="Ubuntu"
        PKG_MANAGER="apt"
    elif [ -f /etc/debian_version ]; then
        OS="Debian"
        PKG_MANAGER="apt"
    else
        OS="Unknown"
        PKG_MANAGER="apt"
    fi
    
    print_info "检测到操作系统: $OS"
    print_info "包管理器: $PKG_MANAGER"
}

# 检查是否以 root 权限运行
check_root() {
    if [ "$EUID" -ne 0 ]; then 
        print_error "请使用 root 权限运行此脚本！"
        print_info "运行命令: sudo bash $0"
        exit 1
    fi
}

# 检查 Node.js
check_node() {
    print_step "检查 Node.js 环境..."
    
    if command -v node > /dev/null 2>&1; then
        NODE_VERSION=$(node -v | cut -d'v' -f2 | cut -d'.' -f1)
        if [ "$NODE_VERSION" -ge 18 ]; then
            print_success "Node.js 已安装: $(node -v)"
        else
            print_warning "Node.js 版本过低: $(node -v)，需要 v18+"
            install_nodejs
        fi
    else
        print_warning "Node.js 未安装"
        install_nodejs
    fi
    
    # 检查 npm
    if command -v npm > /dev/null 2>&1; then
        print_success "npm 已安装: $(npm -v)"
    fi
}

# 安装 Node.js
install_nodejs() {
    print_step "开始安装 Node.js 18+..."
    
    # 使用 NodeSource 安装指定版本
    if [ "$PKG_MANAGER" = "yum" ]; then
        # CentOS/RHEL
        curl -fsSL https://rpm.nodesource.com/setup_18.x | bash -
        yum install -y nodejs
    elif [ "$PKG_MANAGER" = "apt" ]; then
        # Ubuntu/Debian
        curl -fsSL https://deb.nodesource.com/setup_18.x | bash -
        apt-get install -y nodejs
    fi
    
    # 验证安装
    if command -v node > /dev/null 2>&1; then
        print_success "Node.js 安装成功: $(node -v)"
    else
        print_error "Node.js 安装失败"
        exit 1
    fi
}

# 检查 pnpm
check_pnpm() {
    print_step "检查 pnpm..."
    
    if command -v pnpm > /dev/null 2>&1; then
        PNPM_VERSION=$(pnpm -v)
        print_success "pnpm 已安装: v$PNPM_VERSION"
    else
        print_warning "pnpm 未安装"
        install_pnpm
    fi
}

# 安装 pnpm
install_pnpm() {
    print_step "开始安装 pnpm..."
    
    npm install -g pnpm
    
    if command -v pnpm > /dev/null 2>&1; then
        print_success "pnpm 安装成功: $(pnpm -v)"
    else
        print_error "pnpm 安装失败"
        exit 1
    fi
}

# 创建项目目录
create_project_dir() {
    print_step "创建项目目录..."
    
    if [ ! -d "$PROJECT_DIR" ]; then
        mkdir -p "$PROJECT_DIR"
        print_success "项目目录已创建: $PROJECT_DIR"
    else
        print_info "项目目录已存在: $PROJECT_DIR"
    fi
}

# 复制项目文件
copy_project_files() {
    print_step "复制项目文件到 $PROJECT_DIR..."
    
    # 获取当前脚本所在目录
    SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
    
    if [ -d "$SCRIPT_DIR/src" ] || [ -f "$SCRIPT_DIR/package.json" ]; then
        # 从脚本目录复制
        rsync -av --exclude='node_modules' --exclude='.git' --exclude='dist' \
              "$SCRIPT_DIR/" "$PROJECT_DIR/"
        print_success "项目文件已复制"
    else
        print_warning "未在脚本目录找到项目文件"
        print_info "请手动将项目文件复制到: $PROJECT_DIR"
    fi
}

# 安装项目依赖
install_dependencies() {
    print_step "安装项目依赖..."
    
    cd "$PROJECT_DIR"
    
    if [ ! -f "package.json" ]; then
        print_error "未找到 package.json 文件"
        exit 1
    fi
    
    # 设置 npm 镜像（国内加速）
    if grep -q "npm.taobao.org\|npmmirror.com" /etc/npmrc 2>/dev/null || [ ! -z "$NP_MIRROR" ]; then
        print_info "检测到使用国内镜像源"
    else
        print_info "配置 npm 镜像为淘宝镜像..."
        npm config set registry https://registry.npmmirror.com
    fi
    
    # 安装依赖
    pnpm install
    
    if [ $? -eq 0 ]; then
        print_success "依赖安装完成"
    else
        print_error "依赖安装失败"
        exit 1
    fi
}

# 构建项目
build_project() {
    print_step "构建生产环境..."
    
    cd "$PROJECT_DIR"
    
    # 检查环境变量文件
    if [ ! -f ".env.production" ] && [ -f ".env.example" ]; then
        print_info "创建生产环境配置文件..."
        cp .env.example .env.production
        
        # 配置 API 地址
        sed -i 's|VITE_API_BASE_URL=.*|VITE_API_BASE_URL=https://api.yourdomain.com|g' .env.production
        print_info "请编辑 .env.production 配置 API 地址"
    fi
    
    # 执行构建
    pnpm run build
    
    if [ $? -eq 0 ]; then
        print_success "构建完成"
    else
        print_error "构建失败"
        exit 1
    fi
}

# 创建 Nginx 配置文件
create_nginx_config() {
    print_step "生成 Nginx 配置文件..."
    
    # 宝塔面板 Nginx 配置文件路径
    NGINX_CONF="$NGINX_CONFIG_DIR/$SITE_DOMAIN.conf"
    
    # 创建配置内容
    cat > "$NGINX_CONF" << 'EOF'
#==========================================
# 百米乐音乐平台 - Nginx 配置文件
# 适用：Vue/React 单页应用（SPA）
# 宝塔面板自动生成
#==========================================

server {
    listen 80;
    listen [::]:80;
    server_name bailemi.yourdomain.com;
    
    # SSL 配置（如需 HTTPS，请取消注释并配置证书）
    # listen 443 ssl http2;
    # listen [::]:443 ssl http2;
    # ssl_certificate /www/server/nginx/conf/ssl/bailemi.yourdomain.com.pem;
    # ssl_certificate_key /www/server/nginx/conf/ssl/bailemi.yourdomain.com.key;
    # ssl_protocols TLSv1.2 TLSv1.3;
    # ssl_ciphers ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384;
    # ssl_prefer_server_ciphers on;
    # ssl_session_cache shared:SSL:10m;
    # ssl_session_timeout 10d;
    
    # 字符编码
    charset utf-8;
    
    # 根目录
    root /www/wwwroot/bailemi/dist;
    
    # 访问日志
    access_log /www/wwwlogs/bailemi_access.log;
    error_log /www/wwwlogs/bailemi_error.log;
    
    # 默认索引文件
    index index.html;
    
    # Gzip 压缩配置
    gzip on;
    gzip_vary on;
    gzip_proxied any;
    gzip_comp_level 6;
    gzip_types text/plain text/css text/xml application/json application/javascript application/rss+xml application/atom+xml image/svg+xml application/font-woff2 application/octet-stream;
    gzip_min_length 1000;
    
    # 安全头配置
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;
    
    # 静态资源缓存配置
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2|ttf|eot)$ {
        expires 30d;
        add_header Cache-Control "public, immutable";
        access_log off;
    }
    
    # HTML 文件不缓存（确保更新及时生效）
    location ~* \.html$ {
        expires -1h;
        add_header Cache-Control "no-cache, no-store, must-revalidate";
    }
    
    # API 代理配置（如需代理到后端）
    location /api/ {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        # 超时配置
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
        
        # WebSocket 支持（如需要）
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
    }
    
    # ========================================
    # Vue/React SPA 路由支持 - 关键配置！
    # ========================================
    location / {
        try_files $uri $uri/ /index.html;
    }
    
    # 防止访问隐藏文件
    location ~ /\. {
        deny all;
        access_log off;
        log_not_found off;
    }
    
    # 禁止访问敏感文件
    location ~* \.(env|git|htaccess|htpasswd|ini|log|sh|sql|conf|bak)$ {
        deny all;
        access_log off;
        log_not_found off;
    }
    
    # 禁止访问 node_modules
    location ~ /node_modules/ {
        deny all;
        access_log off;
        log_not_found off;
    }
    
    # 禁止访问 .git 目录
    location ~ /\.git {
        deny all;
        access_log off;
        log_not_found off;
    }
    
    # 禁止访问服务器内部文件
    location ~ /\.(?!well-known) {
        deny all;
    }
    
    # 限制请求大小
    client_max_body_size 10M;
    
    # 超时配置
    client_body_timeout 65;
    send_timeout 65;
    
    # 关闭日志（静态资源）
    location ~* \.(png|jpe?g|gif|ico|css|js|svg|woff2?|ttf|eot)$ {
        access_log off;
    }
}

# HTTP 重定向到 HTTPS（如启用 HTTPS）
# server {
#     listen 80;
#     listen [::]:80;
#     server_name bailemi.yourdomain.com;
#     return 301 https://$server_name$request_uri;
# }
EOF
    
    print_success "Nginx 配置文件已创建: $NGINX_CONF"
    print_info "配置文件内容预览："
    echo "----------------------------------------"
    head -n 30 "$NGINX_CONF"
    echo "----------------------------------------"
}

# 测试 Nginx 配置
test_nginx_config() {
    print_step "测试 Nginx 配置..."
    
    if nginx -t 2>&1 | grep -q "syntax is ok"; then
        print_success "Nginx 配置语法正确"
    else
        print_error "Nginx 配置存在语法错误"
        nginx -t
        exit 1
    fi
}

# 重载 Nginx
reload_nginx() {
    print_step "重载 Nginx 服务..."
    
    # 宝塔面板使用 service 命令
    if command -v bt > /dev/null 2>&1; then
        print_info "检测到宝塔面板，重载 Nginx..."
        bt reload
    else
        systemctl reload nginx 2>/dev/null || service nginx reload 2>/dev/null || nginx -s reload
    fi
    
    if [ $? -eq 0 ]; then
        print_success "Nginx 已重载"
    else
        print_warning "Nginx 重载失败，请手动重载"
        print_info "命令: systemctl reload nginx"
    fi
}

# 创建 SSL 配置（可选）
create_ssl_config() {
    print_info "是否需要配置 SSL 证书？"
    read -p "请输入域名证书路径（或按回车跳过）: " CERT_PATH
    read -p "请输入域名证书密钥路径（或按回车跳过）: " CERT_KEY_PATH
    
    if [ ! -z "$CERT_PATH" ] && [ ! -z "$CERT_KEY_PATH" ]; then
        if [ -f "$CERT_PATH" ] && [ -f "$CERT_KEY_PATH" ]; then
            print_info "检测到证书文件，更新 Nginx 配置..."
            
            # 取消 SSL 注释并配置证书路径
            sed -i "s|# ssl_certificate /www/server/nginx/conf/ssl/bailemi.yourdomain.com.pem;|ssl_certificate $CERT_PATH;|" "$NGINX_CONF"
            sed -i "s|# ssl_certificate_key /www/server/nginx/conf/ssl/bailemi.yourdomain.com.key;|ssl_certificate_key $CERT_KEY_PATH;|" "$NGINX_CONF"
            sed -i "s|# listen 443 ssl http2;|listen 443 ssl http2;|" "$NGINX_CONF"
            sed -i "s|# listen \[::\]:443 ssl http2;|listen [::]:443 ssl http2;|" "$NGINX_CONF"
            sed -i "s|# return 301 https://|return 301 https://|" "$NGINX_CONF"
            
            print_success "SSL 配置已更新"
        else
            print_warning "证书文件不存在，跳过 SSL 配置"
        fi
    fi
}

# 配置防火墙
config_firewall() {
    print_step "配置防火墙..."
    
    # CentOS 7+ 使用 firewall-cmd
    if command -v firewall-cmd > /dev/null 2>&1; then
        if systemctl is-active --quiet firewalld; then
            print_info "开放 80 和 443 端口..."
            firewall-cmd --permanent --add-service=http
            firewall-cmd --permanent --add-service=https
            firewall-cmd --reload
            print_success "防火墙规则已配置"
        fi
    # Ubuntu/Debian 使用 ufw
    elif command -v ufw > /dev/null 2>&1; then
        if systemctl is-active --quiet ufw; then
            print_info "开放 80 和 443 端口..."
            ufw allow 80/tcp
            ufw allow 443/tcp
            print_success "防火墙规则已配置"
        fi
    else
        print_info "未检测到防火墙或防火墙未运行"
    fi
}

# 生成使用说明
generate_readme() {
    print_step "生成部署说明..."
    
    README_FILE="$PROJECT_DIR/DEPLOY_GUIDE.md"
    
    cat > "$README_FILE" << 'EOF'
# 百米乐音乐平台 - 宝塔面板部署指南

## 部署完成 ✅

您的百米乐前端项目已成功部署到服务器！

## 访问地址

- **前台地址**: http://bailemi.yourdomain.com
- **API 地址**: http://api.yourdomain.com (已配置代理)

## 重要配置

### 1. 修改域名
编辑 Nginx 配置文件，将 `bailemi.yourdomain.com` 替换为您的实际域名：
```bash
vi /www/server/nginx/conf/vhost/bailemi.yourdomain.com.conf
```

### 2. 配置 SSL 证书
在宝塔面板中为网站配置 SSL 证书：
1. 登录宝塔面板
2. 点击「网站」→「设置」→「SSL」
3. 选择 Let's Encrypt 或上传商业证书
4. 开启强制 HTTPS

### 3. 修改 API 地址
编辑 `.env.production` 文件，配置正确的 API 地址：
```bash
cd /www/wwwroot/bailemi
vi .env.production
```

修改以下配置：
```env
VITE_API_BASE_URL=https://api.yourdomain.com
```

### 4. 重新构建（如修改了配置）
```bash
cd /www/wwwroot/bailemi
pnpm run build
```

## 常用命令

### 查看日志
```bash
# Nginx 访问日志
tail -f /www/wwwlogs/bailemi_access.log

# Nginx 错误日志
tail -f /www/wwwlogs/bailemi_error.log
```

### 重启 Nginx
```bash
# 宝塔面板
bt reload

# 或命令行
systemctl reload nginx
```

### 清理缓存
```bash
# 清理浏览器缓存
Ctrl + Shift + Delete

# 清理 CDN 缓存（如使用）
# 请登录对应的 CDN 控制台清理
```

## 宝塔面板操作

### 1. 添加网站
1. 登录宝塔面板
2. 点击「网站」→「添加站点」
3. 填写域名，选择纯静态
4. 网站目录设置为 `/www/wwwroot/bailemi/dist`

### 2. 配置伪静态
宝塔面板已自动配置，无需手动设置

### 3. 配置 SSL
1. 点击网站「设置」→「SSL」
2. 选择证书类型并配置

### 4. 设置网站目录
确保「网站目录」指向 `/www/wwwroot/bailemi/dist`

## 故障排查

### 页面显示空白
1. 检查浏览器控制台是否有错误
2. 确认 `dist` 目录存在且有文件
3. 检查 Nginx 配置是否正确
4. 确认 API 地址配置正确

### 静态资源 404
1. 检查文件是否在 `dist` 目录
2. 确认 Nginx 配置中的 `root` 路径正确
3. 检查文件权限

### API 请求失败
1. 确认后端服务正在运行
2. 检查 `.env.production` 配置
3. 测试 API 连接：
```bash
curl http://127.0.0.1:8080/health
```

## 性能优化建议

### 1. 开启 CDN 加速
将 `dist` 目录上传至 CDN，七牛云、阿里云 OSS 等

### 2. 配置浏览器缓存
当前配置已优化，可根据需求调整 `expires` 时间

### 3. 开启 HTTP/2
在 Nginx 配置中已启用 `http2`

### 4. 压缩资源
当前配置已启用 Gzip 压缩

## 技术支持

如有问题，请查看：
- Nginx 错误日志: `/www/wwwlogs/bailemi_error.log`
- 前端控制台错误信息
- 宝塔面板日志

---
Generated by Bailemi Setup Script
EOF
    
    print_success "部署说明已生成: $README_FILE"
}

# 主函数
main() {
    echo ""
    echo "========================================"
    echo "  百米乐前端项目 - 宝塔面板一键部署"
    echo "========================================"
    echo ""
    
    # 交互式配置
    echo "请配置以下信息（直接回车使用默认值）："
    echo ""
    
    read -p "网站域名 [$SITE_DOMAIN]: " INPUT_DOMAIN
    SITE_DOMAIN=${INPUT_DOMAIN:-$SITE_DOMAIN}
    
    read -p "项目目录 [$PROJECT_DIR]: " INPUT_DIR
    PROJECT_DIR=${INPUT_DIR:-$PROJECT_DIR}
    
    DIST_DIR="$PROJECT_DIR/dist"
    
    echo ""
    print_info "配置确认："
    echo "  域名: $SITE_DOMAIN"
    echo "  项目目录: $PROJECT_DIR"
    echo "  静态资源目录: $DIST_DIR"
    echo ""
    
    # 执行部署步骤
    check_root
    detect_os
    check_node
    check_pnpm
    create_project_dir
    copy_project_files
    install_dependencies
    build_project
    create_nginx_config
    test_nginx_config
    reload_nginx
    config_firewall
    generate_readme
    
    echo ""
    echo "========================================"
    print_success "部署完成！🎉"
    echo "========================================"
    echo ""
    print_info "访问地址: http://$SITE_DOMAIN"
    print_info "项目目录: $PROJECT_DIR"
    print_info "配置文件: /www/server/nginx/conf/vhost/$SITE_DOMAIN.conf"
    print_info "部署说明: $PROJECT_DIR/DEPLOY_GUIDE.md"
    echo ""
    print_warning "请记得："
    echo "  1. 修改域名为实际域名"
    echo "  2. 配置 SSL 证书（建议）"
    echo "  3. 配置正确的 API 地址"
    echo ""
}

# 运行主函数
main "$@"
