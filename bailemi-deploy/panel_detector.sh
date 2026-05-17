#!/bin/bash
# =========================================
# 百米乐音乐平台 - 通用运维面板检测与部署
# 支持：宝塔、小皮、aaPanel、CyberPanel、DirectAdmin、CWP
# 作者：Bailemi Team
# =========================================

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
BOLD='\033[1m'
NC='\033[0m'

# 日志函数
log_info() {
    echo -e "${CYAN}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_header() {
    echo -e "${BOLD}${MAGENTA}========== $1 ==========${NC}"
}

# =========================================
# 面板检测函数
# =========================================

# 检测宝塔面板
detect_baota() {
    if [ -d "/www/server/panel" ] && [ -f "/www/server/panel/data/default.pl" ]; then
        echo "baota"
        return 0
    fi
    return 1
}

# 检测小皮面板
detect_xp() {
    if [ -d "/www/server" ] && [ -f "/www/server/panel/data/plugin.json" ] && [ ! -f "/www/server/panel/data/default.pl" ]; then
        echo "xp"
        return 0
    fi
    if command -v xp >/dev/null 2>&1; then
        echo "xp"
        return 0
    fi
    return 1
}

# 检测aaPanel
detect_aapanel() {
    if [ -d "/www/server/aapanel" ] || [ -d "/www/server/panel" ] && [ -f "/www/server/panel/data/aapanel.pl" ]; then
        echo "aapanel"
        return 0
    fi
    return 1
}

# 检测CyberPanel
detect_cyberpanel() {
    if [ -f "/usr/local/CyberCP/bin/cyberpanel" ] || [ -d "/usr/local/CyberCP" ]; then
        echo "cyberpanel"
        return 0
    fi
    return 1
}

# 检测DirectAdmin
detect_directadmin() {
    if [ -f "/usr/local/directadmin/directadmin" ] && [ -d "/usr/local/directadmin" ]; then
        echo "directadmin"
        return 0
    fi
    return 1
}

# 检测CWP (CentOS Web Panel)
detect_cwp() {
    if [ -f "/usr/local/cwp/php71/bin/php" ] && [ -d "/usr/local/cwp" ]; then
        echo "cwp"
        return 0
    fi
    if [ -f "/usr/local/cwpsrv/htdocs/index.php" ]; then
        echo "cwp"
        return 0
    fi
    return 1
}

# 检测Plesk
detect_plesk() {
    if [ -f "/usr/local/psa/bin/plesk" ] && [ -d "/usr/local/psa" ]; then
        echo "plesk"
        return 0
    fi
    return 1
}

# 检测cPanel/WHM
detect_cpanel() {
    if [ -f "/usr/local/cpanel/version" ] && [ -d "/usr/local/cpanel" ]; then
        echo "cpanel"
        return 0
    fi
    return 1
}

# =========================================
# 主检测函数
# =========================================

detect_panel() {
    log_header "开始检测运维面板"
    
    # 按优先级检测
    local panels=("baota" "xp" "aapanel" "cyberpanel" "directadmin" "cwp" "plesk" "cpanel")
    
    for panel in "${panels[@]}"; do
        log_info "检测 ${panel} 面板..."
        case $panel in
            baota)
                result=$(detect_baota)
                ;;
            xp)
                result=$(detect_xp)
                ;;
            aapanel)
                result=$(detect_aapanel)
                ;;
            cyberpanel)
                result=$(detect_cyberpanel)
                ;;
            directadmin)
                result=$(detect_directadmin)
                ;;
            cwp)
                result=$(detect_cwp)
                ;;
            plesk)
                result=$(detect_plesk)
                ;;
            cpanel)
                result=$(detect_cpanel)
                ;;
        esac
        
        if [ -n "$result" ]; then
            log_success "检测到 ${result} 面板"
            echo "$result"
            return 0
        fi
    done
    
    log_warning "未检测到主流运维面板，使用通用部署模式"
    echo "generic"
    return 0
}

# =========================================
# 获取面板配置
# =========================================

get_panel_config() {
    local panel=$1
    
    case $panel in
        baota|aapanel)
            echo "www_root=/www/wwwroot"
            echo "nginx_conf=/www/server/panel/vhost/nginx"
            echo "php_bin=/www/server/php"
            echo "mysql_bin=/www/server/mysql"
            echo "panel_data=/www/server/panel/data"
            ;;
        xp)
            echo "www_root=/www/wwwroot"
            echo "nginx_conf=/www/server/panel/vhost/nginx"
            echo "php_bin=/www/server/php"
            echo "mysql_bin=/www/server/mysql"
            echo "panel_data=/www/server/panel/data"
            ;;
        cyberpanel)
            echo "www_root=/home"
            echo "nginx_conf=/usr/local/lsws/conf/vhosts"
            echo "php_bin=/usr/local/lsphp"
            echo "mysql_bin=/usr/local/bin/mysql"
            echo "panel_data=/usr/local/CyberCP"
            ;;
        directadmin)
            echo "www_root=/home"
            echo "nginx_conf=/usr/local/directadmin/data/users"
            echo "php_bin=/usr/local/php"
            echo "mysql_bin=/usr/bin/mysql"
            echo "panel_data=/usr/local/directadmin"
            ;;
        cwp)
            echo "www_root=/home"
            echo "nginx_conf=/usr/local/nginx/conf/vhosts"
            echo "php_bin=/usr/local/php"
            echo "mysql_bin=/usr/bin/mysql"
            echo "panel_data=/usr/local/cwp"
            ;;
        plesk)
            echo "www_root=/var/www/vhosts"
            echo "nginx_conf=/var/www/vhosts/system"
            echo "php_bin=/opt/plesk/php"
            echo "mysql_bin=/usr/bin/mysql"
            echo "panel_data=/usr/local/psa"
            ;;
        cpanel)
            echo "www_root=/home"
            echo "nginx_conf=/var/cpanel/userdata"
            echo "php_bin=/usr/local/bin/php"
            echo "mysql_bin=/usr/bin/mysql"
            echo "panel_data=/usr/local/cpanel"
            ;;
        generic)
            echo "www_root=/var/www/html"
            echo "nginx_conf=/etc/nginx/sites-available"
            echo "php_bin=/usr/bin/php"
            echo "mysql_bin=/usr/bin/mysql"
            echo "panel_data=/etc"
            ;;
    esac
}

# =========================================
# 获取面板名称
# =========================================

get_panel_name() {
    local panel=$1
    
    case $panel in
        baota) echo "宝塔面板 (BT.cn)" ;;
        xp) echo "小皮面板 (Xp.cn)" ;;
        aapanel) echo "aaPanel (国际版宝塔)" ;;
        cyberpanel) echo "CyberPanel" ;;
        directadmin) echo "DirectAdmin" ;;
        cwp) echo "CentOS Web Panel (CWP)" ;;
        plesk) echo "Plesk" ;;
        cpanel) echo "cPanel/WHM" ;;
        generic) echo "通用Linux环境" ;;
    esac
}

# =========================================
# 使用示例
# =========================================

if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    echo "==================================================="
    echo "  百米乐运维面板检测器"
    echo "==================================================="
    
    # 检测面板
    panel_type=$(detect_panel)
    panel_name=$(get_panel_name "$panel_type")
    
    echo ""
    log_success "检测结果"
    echo -e "  面板类型: ${panel_type}"
    echo -e "  面板名称: ${panel_name}"
    echo ""
    
    # 获取配置
    log_header "获取面板配置"
    get_panel_config "$panel_type" | while read -r config_line; do
        key=$(echo "$config_line" | cut -d'=' -f1)
        value=$(echo "$config_line" | cut -d'=' -f2-)
        echo -e "  ${CYAN}${key}${NC}=${value}"
    done
    
    echo ""
    log_info "检测完成！"
    echo ""
    echo "==================================================="
    echo "  下一步："
    echo "  1. 如果你想快速部署，运行: bash deploy.sh"
    echo "  2. 如果你想使用特定面板，运行: bash deploy_${panel_type}.sh"
    echo "==================================================="
fi
