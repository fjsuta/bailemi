#!/bin/bash
# =========================================
# 百米乐 - CyberPanel部署脚本
# =========================================

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/panel_detector.sh"

print_welcome() {
    echo ""
    echo -e "${MAGENTA}========================================${NC}"
    echo -e "${CYAN}  🎵 百米乐音乐平台 - CyberPanel部署${NC}"
    echo -e "${MAGENTA}========================================${NC}"
    echo ""
}

check_cyberpanel() {
    if ! [ -f "/usr/local/CyberCP/bin/cyberpanel" ] && ! [ -d "/usr/local/CyberCP" ]; then
        log_error "此脚本仅适用于CyberPanel！"
        log_info "请使用通用部署脚本: bash deploy.sh"
        exit 1
    fi
    log_success "检测到CyberPanel环境"
}

deploy_cyberpanel() {
    log_header "CyberPanel部署"
    
    local domain
    read -p "请输入站点域名 (例如: bailemi.com): " domain
    domain=${domain:-bailemi.local}
    
    # 获取用户名（从配置或输入）
    local username
    read -p "请输入CyberPanel用户名 (例如: admin): " username
    username=${username:-admin}
    
    # 创建站点目录
    local user_home="/home/${username}"
    local site_path="${user_home}/public_html"
    
    if [ -d "$site_path" ]; then
        INSTALL_PATH="${user_home}/public_html/bailemi"
    else
        INSTALL_PATH="/home/${domain}/public_html"
    fi
    
    mkdir -p "$INSTALL_PATH"
    log_info "站点目录: ${INSTALL_PATH}"
    
    # 复制文件
    local source_dir="$SCRIPT_DIR/.."
    if [ -f "$source_dir/package.json" ]; then
        cp -r "$source_dir"/* "$INSTALL_PATH/" 2>/dev/null || true
    fi
    
    # 设置权限
    if [ -d "/home/${username}" ]; then
        chown -R "${username}:nobody" "$INSTALL_PATH"
    else
        chown -R nobody:nobody "$INSTALL_PATH"
    fi
    chmod -R 755 "$INSTALL_PATH"
    mkdir -p "$INSTALL_PATH"/{uploads,runtime,logs,temp}
    chmod -R 777 "$INSTALL_PATH"/{uploads,runtime,logs,temp}
    
    # 生成OpenLiteSpeed配置
    local litespeed_conf="/usr/local/lsws/conf/vhosts/${domain}.conf"
    
    if [ ! -f "$litespeed_conf" ]; then
        cat > "$litespeed_conf" << 'EOF'
docRoot                   /home/CP_USER/public_html/bailemi/public
vhDomain                  CP_DOMAIN
vhAliases                 www.CP_DOMAIN
adminEmails               admin@CP_DOMAIN
enableGzip                1
enableIpGeo               0
enableRLimit              0
errorlog                  /usr/local/lsws/logs/CP_DOMAIN.error.log
accesslog                 /usr/local/lsws/logs/CP_DOMAIN.access.log

context / {
  location                /home/CP_USER/public_html/bailemi/public
  allowBrowse             1
  indexFiles              index.php index.html index.htm

  rewrite  {
    enable                1
    autoLoadHtaccess      1
  }
}

context /php/ {
  location                /usr/local/lsws/fcgi-bin/
  allowBrowse             0
}

scriptHandler /php/ {
  lsapi                   /usr/local/lsws/fcgi-bin/lsphp
}
EOF
        
        sed -i "s|CP_DOMAIN|$domain|g" "$litespeed_conf"
        sed -i "s|CP_USER|$username|g" "$litespeed_conf"
    fi
    
    log_success "CyberPanel部署完成！"
    echo ""
    echo "请在CyberPanel中："
    echo "  1. 创建站点 ${domain}"
    echo "  2. 网站根目录指向 ${INSTALL_PATH}/public"
    echo "  3. 创建数据库"
    echo "  4. 重启OpenLiteSpeed"
    echo "  5. 访问 ${domain} 完成安装"
}

main() {
    print_welcome
    check_root
    check_cyberpanel
    deploy_cyberpanel
}

main "$@"
