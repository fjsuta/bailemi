#!/bin/bash
# =========================================
# 百米乐音乐平台 - 通用源码初始化部署脚本
# 支持所有主流Linux系统和运维面板
# =========================================

set -e

# 引入面板检测器
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/panel_detector.sh"

# =========================================
# 配置变量
# =========================================

PROJECT_NAME="bailemi"
VERSION="1.0.0"
INSTALL_PATH=""
PANEL_TYPE=""
WWW_ROOT=""
NGINX_CONF_DIR=""
DETECTED_CONFIG=""

# =========================================
# 函数定义
# =========================================

# 打印欢迎信息
print_welcome() {
    log_header "百米乐音乐平台 - 通用部署工具"
    echo ""
    echo -e "  🎵 版本: ${VERSION}"
    echo -e "  📦 开源协议: Apache 2.0"
    echo ""
    log_info "此工具支持所有主流Linux系统和运维面板"
    echo ""
}

# 检查root权限
check_root() {
    if [ "$EUID" -ne 0 ]; then
        log_error "请使用root权限运行此脚本！"
        echo -e "运行命令: ${YELLOW}sudo bash $0${NC}"
        exit 1
    fi
}

# 系统检测
detect_system() {
    log_header "系统信息检测"
    
    # 系统类型
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        OS_NAME=$PRETTY_NAME
    elif [ -f /etc/redhat-release ]; then
        OS_NAME=$(cat /etc/redhat-release)
    elif [ -f /etc/debian_version ]; then
        OS_NAME="Debian $(cat /etc/debian_version)"
    else
        OS_NAME=$(uname -s)
    fi
    
    # 系统架构
    ARCH=$(uname -m)
    KERNEL=$(uname -r)
    
    log_info "操作系统: ${OS_NAME}"
    log_info "系统架构: ${ARCH}"
    log_info "内核版本: ${KERNEL}"
}

# 检测安装路径
detect_install_path() {
    log_header "检测安装路径"
    
    # 从面板配置获取默认路径
    local configs=($DETECTED_CONFIG)
    for config in "${configs[@]}"; do
        if [[ "$config" =~ ^www_root= ]]; then
            WWW_ROOT=$(echo "$config" | cut -d'=' -f2-)
        elif [[ "$config" =~ ^nginx_conf= ]]; then
            NGINX_CONF_DIR=$(echo "$config" | cut -d'=' -f2-)
        fi
    done
    
    if [ -n "$WWW_ROOT" ] && [ -d "$WWW_ROOT" ]; then
        INSTALL_PATH="$WWW_ROOT/$PROJECT_NAME"
        log_info "检测到Web根目录: ${WWW_ROOT}"
    else
        if [ -d "/var/www/html" ]; then
            INSTALL_PATH="/var/www/html/$PROJECT_NAME"
        elif [ -d "/www/wwwroot" ]; then
            INSTALL_PATH="/www/wwwroot/$PROJECT_NAME"
        else
            INSTALL_PATH="/usr/share/nginx/html/$PROJECT_NAME"
        fi
        log_warning "未检测到面板Web目录，使用默认路径"
    fi
    
    log_info "项目安装路径: ${INSTALL_PATH}"
}

# 检查环境
check_environment() {
    log_header "检查部署环境"
    
    # 检查PHP
    if command -v php >/dev/null 2>&1; then
        PHP_VER=$(php -v | head -n1 | cut -d' ' -f2 | cut -d'.' -f1-2)
        log_success "PHP ${PHP_VER} 已安装"
    else
        log_warning "PHP 未安装，部分功能不可用"
    fi
    
    # 检查Node.js
    if command -v node >/dev/null 2>&1; then
        NODE_VER=$(node -v | cut -d'v' -f2)
        log_success "Node.js ${NODE_VER} 已安装"
    else
        log_warning "Node.js 未安装，前端构建功能不可用"
    fi
    
    # 检查Git
    if command -v git >/dev/null 2>&1; then
        log_success "Git 已安装"
    else
        log_warning "Git 未安装"
    fi
    
    # 检查数据库
    if command -v mysql >/dev/null 2>&1; then
        log_success "MySQL/MariaDB 已安装"
    else
        log_warning "MySQL/MariaDB 未安装"
    fi
    
    # 检查Nginx/Apache
    if command -v nginx >/dev/null 2>&1; then
        log_success "Nginx 已安装"
    elif command -v apache2 >/dev/null 2>&1 || command -v httpd >/dev/null 2>&1; then
        log_success "Apache 已安装"
    else
        log_warning "Web服务器未检测到"
    fi
}

# 创建目录
create_directories() {
    log_header "创建项目目录"
    
    if [ ! -d "$INSTALL_PATH" ]; then
        mkdir -p "$INSTALL_PATH"
        log_success "创建目录: ${INSTALL_PATH}"
    fi
    
    # 子目录
    local dirs=("uploads" "runtime" "config" "logs" "temp")
    for dir in "${dirs[@]}"; do
        if [ ! -d "$INSTALL_PATH/$dir" ]; then
            mkdir -p "$INSTALL_PATH/$dir"
            log_success "创建目录: ${dir}"
        fi
    done
}

# 复制项目文件
copy_project_files() {
    log_header "复制项目文件"
    
    local current_dir=$(pwd)
    local source_dir="$SCRIPT_DIR/.."
    
    # 检查源文件
    if [ -f "$source_dir/package.json" ]; then
        log_info "从 ${source_dir} 复制文件"
        cp -r "$source_dir"/* "$INSTALL_PATH/" 2>/dev/null || true
    else
        log_warning "未找到源码文件，创建演示结构"
        create_demo_structure
    fi
    
    log_success "项目文件已复制"
}

# 创建演示结构（如果无源码）
create_demo_structure() {
    cat > "$INSTALL_PATH/index.php" << 'EOF'
<?php
header('Content-Type: text/html; charset=utf-8');
echo '<h1>百米乐音乐平台 v1.0.0</h1>';
echo '<p>这是一个演示页面，请上传完整源码！</p>';
EOF
    
    cat > "$INSTALL_PATH/readme.txt" << 'EOF'
百米乐音乐平台 - 部署说明
============================
请上传完整的项目源码到此目录！
EOF
}

# 设置权限
set_permissions() {
    log_header "设置目录权限"
    
    # 获取Web用户
    local web_user=""
    local web_group=""
    
    if id -u www-data >/dev/null 2>&1; then
        web_user="www-data"
        web_group="www-data"
    elif id -u www >/dev/null 2>&1; then
        web_user="www"
        web_group="www"
    elif id -u apache >/dev/null 2>&1; then
        web_user="apache"
        web_group="apache"
    else
        web_user=$(stat -c '%U' "$INSTALL_PATH" 2>/dev/null || echo "root")
        web_group=$(stat -c '%G' "$INSTALL_PATH" 2>/dev/null || echo "root")
    fi
    
    log_info "使用用户: ${web_user}:${web_group}"
    
    # 设置权限
    chown -R "$web_user:$web_group" "$INSTALL_PATH" || true
    chmod -R 755 "$INSTALL_PATH"
    chmod -R 777 "$INSTALL_PATH/uploads" 2>/dev/null || true
    chmod -R 777 "$INSTALL_PATH/runtime" 2>/dev/null || true
    chmod -R 777 "$INSTALL_PATH/logs" 2>/dev/null || true
    chmod -R 777 "$INSTALL_PATH/temp" 2>/dev/null || true
    
    log_success "目录权限已设置"
}

# 生成Nginx配置
generate_nginx_config() {
    log_header "生成Nginx配置"
    
    local site_domain="bailemi.local"
    local config_file=""
    
    if [ -n "$NGINX_CONF_DIR" ] && [ -d "$NGINX_CONF_DIR" ]; then
        config_file="$NGINX_CONF_DIR/$PROJECT_NAME.conf"
    elif [ -d "/etc/nginx/sites-available" ]; then
        config_file="/etc/nginx/sites-available/$PROJECT_NAME.conf"
    elif [ -d "/etc/nginx/conf.d" ]; then
        config_file="/etc/nginx/conf.d/$PROJECT_NAME.conf"
    else
        config_file="$INSTALL_PATH/nginx.conf"
    fi
    
    log_info "配置文件: ${config_file}"
    
    cat > "$config_file" << 'EOF'
# =========================================
# 百米乐音乐平台 - Nginx配置
# =========================================

server {
    listen 80;
    listen [::]:80;
    server_name bailemi.local www.bailemi.local;
    
    # 网站根目录（根据实际情况修改）
    root /var/www/html/bailemi/public;
    index index.php index.html index.htm;
    
    # 字符编码
    charset utf-8;
    
    # 访问日志
    access_log /var/log/nginx/bailemi_access.log;
    error_log /var/log/nginx/bailemi_error.log;
    
    # Gzip压缩
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_proxied any;
    gzip_comp_level 6;
    gzip_types
        text/plain
        text/css
        text/xml
        text/javascript
        application/json
        application/javascript
        application/x-javascript
        application/xml
        application/rss+xml
        font/truetype
        font/opentype
        image/svg+xml;
    
    # 安全头部
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-XSS-Protection "1; mode=block" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header Referrer-Policy "no-referrer-when-downgrade" always;
    add_header Content-Security-Policy "default-src 'self' http: https: data: blob: 'unsafe-inline'" always;
    
    # 前端路由支持
    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }
    
    # PHP文件处理
    location ~ \.php$ {
        try_files $uri /index.php =404;
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_pass unix:/var/run/php/php7.4-fpm.sock;
        fastcgi_index index.php;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }
    
    # 静态资源缓存
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2|ttf|eot)$ {
        expires 30d;
        add_header Cache-Control "public, immutable";
        access_log off;
    }
    
    # 拒绝访问敏感文件
    location ~ /\. {
        deny all;
        access_log off;
        log_not_found off;
    }
    
    location ~ /(composer|package|package-lock|\.env|\.git|node_modules|storage|vendor) {
        deny all;
    }
    
    # 限制上传大小
    client_max_body_size 100M;
    client_body_buffer_size 128k;
}
EOF
    
    log_success "Nginx配置文件已生成"
}

# 生成安装提示
generate_install_tips() {
    local tips_file="$INSTALL_PATH/INSTALL_GUIDE.txt"
    
    cat > "$tips_file" << 'EOF'
# =========================================
# 百米乐音乐平台 - 部署成功！
# =========================================

项目位置: INSTALL_PATH_PLACEHOLDER
面板类型: PANEL_TYPE_PLACEHOLDER
安装时间: TIME_PLACEHOLDER

=========================================
下一步操作
=========================================

1. 配置数据库
   - 在面板中创建数据库
   - 修改配置文件: config.php 或 .env

2. 配置Web服务器
   - 检查配置文件
   - 根据面板添加站点

3. 运行安装向导
   - 访问域名
   - 按照提示完成安装

4. 安全配置
   - 修改文件权限
   - 配置SSL证书
   - 设置防火墙

=========================================
常用命令
=========================================

# 查看日志
tail -f /var/log/nginx/error.log

# 重启服务
systemctl restart nginx
systemctl restart php-fpm

# 查看权限
ls -la INSTALL_PATH_PLACEHOLDER

=========================================
技术支持
=========================================
- 文档: https://docs.bailemi.com
- 支持: https://support.bailemi.com
- GitHub: https://github.com/bailemi
EOF
    
    # 替换变量
    sed -i "s|INSTALL_PATH_PLACEHOLDER|$INSTALL_PATH|g" "$tips_file"
    sed -i "s|PANEL_TYPE_PLACEHOLDER|$PANEL_TYPE|g" "$tips_file"
    sed -i "s|TIME_PLACEHOLDER|$(date '+%Y-%m-%d %H:%M:%S')|g" "$tips_file"
    
    log_info "安装指南已生成: $tips_file"
}

# =========================================
# 主函数
# =========================================

main() {
    print_welcome
    check_root
    
    # 检测面板
    PANEL_TYPE=$(detect_panel)
    local panel_name=$(get_panel_name "$PANEL_TYPE")
    DETECTED_CONFIG=$(get_panel_config "$PANEL_TYPE")
    
    log_success "检测到环境: ${panel_name}"
    echo ""
    
    # 检测系统
    detect_system
    echo ""
    
    # 检测安装路径
    detect_install_path
    echo ""
    
    # 检查环境
    check_environment
    echo ""
    
    # 交互确认
    read -p "是否继续部署到 ${INSTALL_PATH}? (y/n): " confirm
    if [[ ! "$confirm" =~ ^[Yy]$ ]]; then
        log_info "部署已取消"
        exit 0
    fi
    
    # 执行部署
    create_directories
    copy_project_files
    set_permissions
    generate_nginx_config
    generate_install_tips
    
    echo ""
    log_header "部署完成！"
    log_success "项目已成功部署到: ${INSTALL_PATH}"
    log_success "面板类型: ${panel_name}"
    echo ""
    log_info "请查看: ${INSTALL_PATH}/INSTALL_GUIDE.txt"
    echo ""
    log_warning "重要提示："
    echo -e "  1. 请根据 ${PANEL_TYPE} 面板配置Web站点"
    echo -e "  2. 在面板中创建数据库"
    echo -e "  3. 访问站点完成安装向导"
    echo ""
}

# 执行主函数
main "$@"
