#!/bin/bash
# =========================================
# 百米乐 - aaPanel国际版部署脚本
# =========================================

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/panel_detector.sh"

print_welcome() {
    echo ""
    echo -e "${MAGENTA}========================================${NC}"
    echo -e "${CYAN}  🎵 百米乐音乐平台 - aaPanel部署${NC}"
    echo -e "${MAGENTA}========================================${NC}"
    echo ""
}

check_aapanel() {
    if ! [ -d "/www/server/panel" ] || ! [ -f "/www/server/panel/data/aapanel.pl" ]; then
        log_error "此脚本仅适用于aaPanel！"
        log_info "请使用通用部署脚本: bash deploy.sh"
        exit 1
    fi
    log_success "检测到aaPanel环境"
}

deploy_aapanel() {
    log_header "aaPanel部署"
    
    # 配置站点
    local domain
    read -p "请输入站点域名 (例如: bailemi.com): " domain
    domain=${domain:-bailemi.local}
    
    # 创建目录
    INSTALL_PATH="/www/wwwroot/bailemi"
    mkdir -p "$INSTALL_PATH"
    
    # 复制文件
    local source_dir="$SCRIPT_DIR/.."
    if [ -f "$source_dir/package.json" ]; then
        cp -r "$source_dir"/* "$INSTALL_PATH/" 2>/dev/null || true
    fi
    
    # 设置权限
    chown -R www:www "$INSTALL_PATH"
    chmod -R 755 "$INSTALL_PATH"
    mkdir -p "$INSTALL_PATH"/{uploads,runtime,logs,temp}
    chmod -R 777 "$INSTALL_PATH"/{uploads,runtime,logs,temp}
    
    # 生成配置
    cat > "/www/server/panel/vhost/nginx/${domain}.conf" << 'EOF'
server
{
    listen 80;
    listen [::]:80;
    server_name AAPANEL_DOMAIN www.AAPANEL_DOMAIN;
    index index.php index.html index.htm;
    root /www/wwwroot/bailemi/public;
    
    location / {
        if (!-e $request_filename) {
            rewrite ^(.*)$ /index.php?s=$1 last;
            break;
        }
    }
    
    location ~ \.php$ {
        fastcgi_pass unix:/tmp/php-cgi.sock;
        fastcgi_index index.php;
        include fastcgi.conf;
    }
    
    access_log /www/wwwlogs/AAPANEL_DOMAIN.log;
    error_log /www/wwwlogs/AAPANEL_DOMAIN.error.log;
}
EOF
    
    sed -i "s|AAPANEL_DOMAIN|$domain|g" "/www/server/panel/vhost/nginx/${domain}.conf"
    
    log_success "aaPanel部署完成！"
    echo ""
    echo "请在aaPanel中："
    echo "  1. 添加站点 ${domain}"
    echo "  2. 根目录指向 /www/wwwroot/bailemi/public"
    echo "  3. 配置PHP和数据库"
    echo "  4. 访问 ${domain} 完成安装"
}

main() {
    print_welcome
    check_root
    check_aapanel
    deploy_aapanel
}

main "$@"
