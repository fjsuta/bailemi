#!/bin/bash
# =========================================
# 百米乐 - 宝塔面板专用部署脚本
# =========================================

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/panel_detector.sh"

print_welcome() {
    echo ""
    echo -e "${MAGENTA}========================================${NC}"
    echo -e "${CYAN}  🎵 百米乐音乐平台 - 宝塔版部署${NC}"
    echo -e "${MAGENTA}========================================${NC}"
    echo ""
}

# 检查是否是宝塔
check_baota() {
    if ! [ -d "/www/server/panel" ] || ! [ -f "/www/server/panel/data/default.pl" ]; then
        log_error "此脚本仅适用于宝塔面板！"
        log_info "请使用通用部署脚本: bash deploy.sh"
        exit 1
    fi
    log_success "检测到宝塔面板环境"
}

# 获取可用PHP版本
get_php_version() {
    local php_dir="/www/server/php"
    local php_versions=()
    
    if [ -d "$php_dir" ]; then
        for dir in "$php_dir"/*/; do
            if [ -d "$dir" ] && [ -f "$dir/bin/php" ]; then
                local ver=$(basename "$dir")
                php_versions+=("$ver")
            fi
        done
    fi
    
    if [ ${#php_versions[@]} -eq 0 ]; then
        log_error "未找到PHP版本，请在宝塔中安装PHP"
        exit 1
    fi
    
    # 按版本排序，返回最新
    local sorted=$(printf '%s\n' "${php_versions[@]}" | sort -rV)
    echo "$sorted" | head -n1
}

# 创建宝塔站点
create_baota_site() {
    log_header "创建宝塔站点"
    
    local domain
    read -p "请输入站点域名 (例如: bailemi.com): " domain
    
    # 默认域名
    if [ -z "$domain" ]; then
        domain="bailemi.local"
        log_warning "使用默认域名: ${domain}"
    fi
    
    # 获取PHP版本
    local php_ver=$(get_php_version)
    log_info "使用PHP版本: ${php_ver}"
    
    # 检查是否已存在
    local site_path="/www/wwwroot/$PROJECT_NAME"
    if [ -d "$site_path" ]; then
        log_warning "站点目录已存在，将覆盖内容"
    fi
    
    # 使用宝塔CLI创建站点（如果存在）
    if command -v bt >/dev/null 2>&1; then
        log_info "尝试使用宝塔CLI创建站点..."
        # 注意：这里是演示代码，实际宝塔需要通过API调用
        log_info "提示：请在宝塔面板中手动添加站点 ${domain}"
    fi
    
    # 创建目录
    mkdir -p "$site_path"
    INSTALL_PATH="$site_path"
    log_success "站点目录: ${site_path}"
    SITE_DOMAIN="$domain"
}

# 安装PHP扩展
install_php_extensions() {
    log_header "检查PHP扩展"
    
    local php_ver=$(get_php_version)
    local php_cli="/www/server/php/${php_ver}/bin/php"
    local php_ini="/www/server/php/${php_ver}/etc/php.ini"
    
    # 检查必要的扩展
    local required=("pdo" "mysql" "curl" "gd" "mbstring" "openssl" "json" "fileinfo")
    local missing=()
    
    for ext in "${required[@]}"; do
        if ! "$php_cli" -m | grep -q "$ext" 2>/dev/null; then
            missing+=("$ext")
        fi
    done
    
    if [ ${#missing[@]} -eq 0 ]; then
        log_success "所有必要扩展已安装"
    else
        log_warning "缺少扩展: ${missing[*]}"
        echo ""
        echo -e "${YELLOW}请在宝塔面板安装以下扩展：${NC}"
        for ext in "${missing[@]}"; do
            echo "  - ${ext}"
        done
        echo ""
        read -p "是否继续部署? (y/n): " cont
        if [[ ! "$cont" =~ ^[Yy]$ ]]; then
            log_info "已取消部署"
            exit 0
        fi
    fi
}

# 配置宝塔站点
configure_baota_site() {
    log_header "配置宝塔站点"
    
    # 复制文件
    log_info "复制项目文件..."
    local source_dir="$SCRIPT_DIR/.."
    if [ -f "$source_dir/package.json" ]; then
        cp -r "$source_dir"/* "$INSTALL_PATH/" 2>/dev/null || true
    fi
    
    # 设置权限
    log_info "设置目录权限..."
    chown -R www:www "$INSTALL_PATH"
    chmod -R 755 "$INSTALL_PATH"
    chmod -R 777 "$INSTALL_PATH/uploads" 2>/dev/null || true
    chmod -R 777 "$INSTALL_PATH/runtime" 2>/dev/null || true
    chmod -R 777 "$INSTALL_PATH/temp" 2>/dev/null || true
    chmod -R 777 "$INSTALL_PATH/logs" 2>/dev/null || true
    
    # 生成宝塔可用的Nginx配置
    local nginx_conf="/www/server/panel/vhost/nginx/$SITE_DOMAIN.conf"
    
    if [ ! -f "$nginx_conf" ]; then
        log_info "生成Nginx配置文件..."
        cat > "$nginx_conf" << 'EOF'
server
{
    listen 80;
    listen [::]:80;
    server_name SITE_DOMAIN_PLACEHOLDER www.SITE_DOMAIN_PLACEHOLDER;
    index index.php index.html index.htm;
    root /www/wwwroot/bailemi/public;
    
    location / {
        if (!-e $request_filename) {
            rewrite ^(.*)$ /index.php?s=$1 last;
            break;
        }
    }
    
    location ~ \.php$ {
        fastcgi_pass unix:/tmp/php-cgi-PHP_VER_PLACEHOLDER.sock;
        fastcgi_index index.php;
        include fastcgi.conf;
    }
    
    # 静态资源缓存
    location ~ .*\.(gif|jpg|jpeg|png|bmp|swf|js|css|woff|woff2|ttf|eot|svg|mp3)$ {
        expires 30d;
        add_header Cache-Control "public, immutable";
        access_log off;
    }
    
    # 日志
    access_log /www/wwwlogs/SITE_DOMAIN_PLACEHOLDER.log;
    error_log /www/wwwlogs/SITE_DOMAIN_PLACEHOLDER.error.log;
}
EOF
        
        # 替换变量
        sed -i "s|SITE_DOMAIN_PLACEHOLDER|$SITE_DOMAIN|g" "$nginx_conf"
        sed -i "s|PHP_VER_PLACEHOLDER|${php_ver}|g" "$nginx_conf"
        
        log_success "Nginx配置已生成: ${nginx_conf}"
    else
        log_warning "站点配置已存在: ${nginx_conf}"
    fi
}

# 宝塔部署主函数
main() {
    print_welcome
    check_root
    check_baota
    
    echo ""
    read -p "按Enter继续部署..."
    
    create_baota_site
    install_php_extensions
    configure_baota_site
    
    echo ""
    log_header "部署完成！"
    echo ""
    log_success "站点已部署到: ${INSTALL_PATH}"
    log_success "站点域名: ${SITE_DOMAIN}"
    echo ""
    log_info "请在宝塔面板中完成以下配置："
    echo ""
    echo -e "  ${CYAN}1.${NC} 添加站点 ${SITE_DOMAIN}"
    echo -e "  ${CYAN}2.${NC} 选择网站根目录为 ${INSTALL_PATH}/public"
    echo -e "  ${CYAN}3.${NC} 配置PHP版本并安装必要扩展"
    echo -e "  ${CYAN}4.${NC} 配置数据库"
    echo -e "  ${CYAN}5.${NC} 设置SSL（推荐）"
    echo -e "  ${CYAN}6.${NC} 访问 ${SITE_DOMAIN} 完成安装向导"
    echo ""
    
    log_warning "安全提示："
    echo "  部署完成后，请删除安装脚本文件！"
}

PROJECT_NAME="bailemi"

main "$@"
