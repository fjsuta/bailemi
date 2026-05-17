#!/bin/bash
# =========================================
# 百米乐 - 纯Linux环境部署脚本
# 无需任何运维面板，适用于纯净Linux系统
# =========================================

set -e

# 颜色
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

log() { echo -e "${BLUE}[INFO]${NC} $1"; }
success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
warning() { echo -e "${YELLOW}[WARNING]${NC} $1"; }
error() { echo -e "${RED}[ERROR]${NC} $1"; }

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INSTALL_PATH="/var/www/bailemi"

print_welcome() {
    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${CYAN}  🎵 百米乐音乐平台 - 纯Linux部署${NC}"
    echo -e "${GREEN}========================================${NC}"
    echo ""
}

detect_package_manager() {
    log "检测系统包管理器..."
    
    if command -v apt &> /dev/null; then
        echo "apt"
    elif command -v yum &> /dev/null; then
        echo "yum"
    elif command -v dnf &> /dev/null; then
        echo "dnf"
    elif command -v apk &> /dev/null; then
        echo "apk"
    else
        echo "unknown"
    fi
}

install_php() {
    local pm=$1
    log "安装PHP及扩展..."
    
    case $pm in
        apt)
            apt-get update -y
            apt-get install -y php8.2 php8.2-fpm php8.2-mysql php8.2-curl php8.2-gd php8.2-mbstring php8.2-xml php8.2-zip php8.2-redis php8.2-imagick
            ;;
        yum|dnf)
            if command -v dnf &> /dev/null; then
                dnf install -y epel-release
                dnf install -y php php-fpm php-mysqlnd php-curl php-gd php-mbstring php-xml php-zip php-pecl-redis php-pecl-imagick
            else
                yum install -y epel-release
                yum install -y php php-fpm php-mysqlnd php-curl php-gd php-mbstring php-xml php-zip php-pecl-redis php-pecl-imagick
            fi
            ;;
        apk)
            apk add --no-cache php82 php82-fpm php82-mysqli php82-curl php82-gd php82-mbstring php82-xml php82-zip php82-pecl-redis php82-pecl-imagick
            ;;
    esac
    success "PHP安装完成"
}

install_nginx() {
    local pm=$1
    log "安装Nginx..."
    
    case $pm in
        apt)
            apt-get install -y nginx
            ;;
        yum|dnf)
            if command -v dnf &> /dev/null; then
                dnf install -y nginx
            else
                yum install -y nginx
            fi
            ;;
        apk)
            apk add --no-cache nginx
            ;;
    esac
    success "Nginx安装完成"
}

install_mysql() {
    local pm=$1
    log "安装MariaDB/MySQL..."
    
    case $pm in
        apt)
            apt-get install -y mariadb-server mariadb-client
            ;;
        yum|dnf)
            if command -v dnf &> /dev/null; then
                dnf install -y mariadb-server mariadb
            else
                yum install -y mariadb-server mariadb
            fi
            ;;
        apk)
            apk add --no-cache mariadb mariadb-client
            ;;
    esac
    success "MariaDB安装完成"
}

install_node() {
    local pm=$1
    log "安装Node.js..."
    
    if command -v node &> /dev/null; then
        local node_ver=$(node -v | cut -d'v' -f2 | cut -d'.' -f1)
        if [ "$node_ver" -ge 16 ]; then
            log "Node.js ${node_ver} 已满足要求"
            return
        fi
    fi
    
    # 使用NodeSource安装
    curl -fsSL https://deb.nodesource.com/setup_18.x | bash -
    case $pm in
        apt)
            apt-get install -y nodejs
            ;;
        yum|dnf)
            yum install -y nodejs
            ;;
    esac
    npm install -g npm
    success "Node.js安装完成"
}

install_redis() {
    local pm=$1
    log "安装Redis..."
    
    case $pm in
        apt)
            apt-get install -y redis-server
            ;;
        yum|dnf)
            if command -v dnf &> /dev/null; then
                dnf install -y redis
            else
                yum install -y redis
            fi
            ;;
        apk)
            apk add --no-cache redis
            ;;
    esac
    success "Redis安装完成"
}

setup_project() {
    log "设置项目目录..."
    
    mkdir -p "$INSTALL_PATH"
    
    # 复制项目文件
    local source_dir="$SCRIPT_DIR/.."
    if [ -f "$source_dir/package.json" ]; then
        cp -r "$source_dir"/* "$INSTALL_PATH/" 2>/dev/null || true
    fi
    
    # 创建必要目录
    mkdir -p "$INSTALL_PATH"/{public,uploads,runtime,logs,config,temp}
    
    # 设置权限
    local www_user="www-data"
    if id -u nginx &>/dev/null 2>&1; then
        www_user="nginx"
    elif id -u apache &>/dev/null 2>&1; then
        www_user="apache"
    elif ! id -u www-data &>/dev/null 2>&1; then
        useradd -r -s /sbin/nologin www-data 2>/dev/null || true
    fi
    
    chown -R "$www_user:$www_user" "$INSTALL_PATH"
    chmod -R 755 "$INSTALL_PATH"
    chmod -R 777 "$INSTALL_PATH"/{uploads,runtime,logs,temp}
    
    success "项目目录设置完成"
}

configure_nginx() {
    log "配置Nginx..."
    
    local nginx_conf_dir
    if [ -d "/etc/nginx/sites-available" ]; then
        nginx_conf_dir="/etc/nginx/sites-available"
    elif [ -d "/etc/nginx/conf.d" ]; then
        nginx_conf_dir="/etc/nginx/conf.d"
    else
        mkdir -p /etc/nginx/conf.d
        nginx_conf_dir="/etc/nginx/conf.d"
    fi
    
    cat > "${nginx_conf_dir}/bailemi.conf" << 'EOF'
server {
    listen 80;
    listen [::]:80;
    server_name _;
    
    root /var/www/bailemi/public;
    index index.php index.html index.htm;
    
    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }
    
    location ~ \.php$ {
        fastcgi_pass unix:/var/run/php/php8.2-fpm.sock;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
    }
    
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2|ttf|eot|mp3)$ {
        expires 30d;
        access_log off;
    }
    
    location ~ /\. {
        deny all;
    }
    
    access_log /var/log/nginx/bailemi_access.log;
    error_log /var/log/nginx/bailemi_error.log;
}
EOF
    
    # 启用站点
    if [ -d "/etc/nginx/sites-enabled" ]; then
        ln -sf "${nginx_conf_dir}/bailemi.conf" "/etc/nginx/sites-enabled/bailemi.conf"
    fi
    
    success "Nginx配置完成"
}

configure_mysql() {
    log "配置MySQL..."
    
    # 启动服务
    if command -v systemctl &> /dev/null; then
        systemctl start mariadb 2>/dev/null || systemctl start mysql
    elif command -v service &> /dev/null; then
        service mariadb start 2>/dev/null || service mysql start
    fi
    
    # 创建数据库和用户（安全起见，交互式）
    echo ""
    warning "请手动完成数据库配置："
    echo "  mysql -u root -p"
    echo "  CREATE DATABASE bailemi CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
    echo "  CREATE USER 'bailemi'@'localhost' IDENTIFIED BY 'your_secure_password';"
    echo "  GRANT ALL PRIVILEGES ON bailemi.* TO 'bailemi'@'localhost';"
    echo "  FLUSH PRIVILEGES;"
    echo ""
}

start_services() {
    log "启动服务..."
    
    if command -v systemctl &> /dev/null; then
        systemctl enable nginx php8.2-fpm redis mariadb 2>/dev/null || systemctl enable nginx php-fpm redis mysql
        systemctl start nginx php8.2-fpm redis mariadb 2>/dev/null || systemctl start nginx php-fpm redis mysql
    elif command -v service &> /dev/null; then
        service nginx start
        service php8.2-fpm start 2>/dev/null || service php-fpm start
        service redis start
        service mariadb start 2>/dev/null || service mysql start
    fi
    
    success "服务启动完成"
}

setup_firewall() {
    log "配置防火墙..."
    
    if command -v ufw &> /dev/null; then
        ufw allow 80/tcp
        ufw allow 443/tcp
        ufw reload
        log "UFW防火墙已配置"
    elif command -v firewall-cmd &> /dev/null; then
        firewall-cmd --permanent --add-service=http
        firewall-cmd --permanent --add-service=https
        firewall-cmd --reload
        log "firewalld已配置"
    else
        warning "未检测到防火墙，请手动配置"
    fi
}

show_completed() {
    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${CYAN}  🎵 部署完成！${NC}"
    echo -e "${GREEN}========================================${NC}"
    echo ""
    echo -e "  项目路径: ${INSTALL_PATH}"
    echo -e "  访问地址: http://$(hostname -I | awk '{print $1}')"
    echo ""
    echo -e "  ${CYAN}下一步：${NC}"
    echo -e "  1. 配置数据库"
    echo -e "  2. 运行安装向导"
    echo -e "  3. 配置SSL（推荐）"
    echo ""
}

main() {
    print_welcome
    
    if [ "$EUID" -ne 0 ]; then
        error "请以root用户运行此脚本！"
        exit 1
    fi
    
    local pm=$(detect_package_manager)
    
    if [ "$pm" = "unknown" ]; then
        error "无法检测到包管理器！"
        exit 1
    fi
    
    log "检测到系统: ${pm}"
    echo ""
    
    read -p "是否继续安装所有组件? (y/n): " confirm
    if [[ ! "$confirm" =~ ^[Yy]$ ]]; then
        log "安装已取消"
        exit 0
    fi
    
    install_php "$pm"
    install_nginx "$pm"
    install_mysql "$pm"
    install_node "$pm"
    install_redis "$pm"
    setup_project
    configure_nginx
    configure_mysql
    start_services
    setup_firewall
    show_completed
}

main "$@"
