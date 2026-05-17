#!/bin/bash
# =========================================
# 百米乐 - Docker一键部署脚本
# =========================================

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# 颜色
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log() { echo -e "${BLUE}[INFO]${NC} $1"; }
success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
warning() { echo -e "${YELLOW}[WARNING]${NC} $1"; }
error() { echo -e "${RED}[ERROR]${NC} $1"; }

print_welcome() {
    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${BLUE}  🎵 百米乐音乐平台 - Docker部署${NC}"
    echo -e "${GREEN}========================================${NC}"
    echo ""
}

check_docker() {
    log "检查Docker环境..."
    if ! command -v docker &> /dev/null; then
        error "Docker未安装！"
        echo ""
        echo "请先安装Docker和Docker Compose："
        echo "CentOS/RHEL: curl -fsSL https://get.docker.com | sh"
        echo "Ubuntu/Debian: apt-get install docker-ce docker-compose"
        echo ""
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
        error "Docker Compose未安装！"
        exit 1
    fi
    
    success "Docker环境检测通过"
}

copy_files() {
    log "复制Docker配置文件..."
    
    # 创建目录
    mkdir -p "${SCRIPT_DIR}/logs"
    mkdir -p "${SCRIPT_DIR}/ssl"
    mkdir -p "${SCRIPT_DIR}/mysql/init"
    
    # 复制源码
    if [ -d "$SCRIPT_DIR/.." ] && [ -f "$SCRIPT_DIR/../package.json" ]; then
        log "复制项目文件..."
        cp -r "$SCRIPT_DIR/../"* . 2>/dev/null || true
    fi
    
    success "文件复制完成"
}

create_env() {
    log "创建环境配置文件..."
    
    if [ ! -f .env ]; then
        cat > .env << 'EOF'
# =========================================
# 百米乐 Docker 环境配置
# =========================================

# 应用配置
APP_NAME=Bailemi
APP_ENV=production
APP_KEY=base64:$(openssl rand -base64 32)
APP_URL=http://localhost
APP_DEBUG=false

# 数据库配置
DB_CONNECTION=mysql
DB_HOST=mysql
DB_PORT=3306
DB_DATABASE=bailemi
DB_USERNAME=bailemi
DB_PASSWORD=bailemi_2024

# Redis配置
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=bailemi_redis_2024

# 文件上传
UPLOAD_PATH=/var/www/html/storage/uploads
MAX_UPLOAD_SIZE=100M

# 安全
JWT_SECRET=$(openssl rand -hex 32)
JWT_TTL=86400
EOF
        success ".env文件创建完成"
    else
        warning ".env文件已存在，跳过创建"
    fi
}

build_containers() {
    log "构建Docker容器..."
    
    if docker compose version &> /dev/null; then
        docker compose build
    else
        docker-compose build
    fi
    
    success "容器构建完成"
}

start_containers() {
    log "启动Docker容器..."
    
    if docker compose version &> /dev/null; then
        docker compose up -d
    else
        docker-compose up -d
    fi
    
    success "容器启动完成"
}

wait_for_services() {
    log "等待服务就绪..."
    sleep 10
    
    # 等待MySQL
    log "检查MySQL连接..."
    local max_tries=30
    local try=0
    
    while [ $try -lt $max_tries ]; do
        if docker exec bailemi-mysql mysqladmin ping -h localhost -u bailemi -pbailemi_2024 &> /dev/null; then
            success "MySQL已就绪"
            break
        fi
        try=$((try + 1))
        log "等待MySQL启动... (${try}/${max_tries})"
        sleep 3
    done
}

init_database() {
    log "初始化数据库..."
    
    # 检查是否有SQL文件
    if [ -f "$SCRIPT_DIR/../bailemi-install/install.sql" ]; then
        log "执行数据库初始化..."
        docker cp "$SCRIPT_DIR/../bailemi-install/install.sql" bailemi-mysql:/docker-entrypoint-initdb.d/
        
        # 执行SQL
        docker exec bailemi-mysql mysql -u bailemi -pbailemi_2024 bailemi -e "source /docker-entrypoint-initdb.d/install.sql" 2>/dev/null || true
    fi
    
    success "数据库初始化完成"
}

show_access_info() {
    echo ""
    echo -e "${GREEN}========================================${NC}"
    echo -e "${BLUE}  🎵 部署成功！${NC}"
    echo -e "${GREEN}========================================${NC}"
    echo ""
    echo -e "  ${CYAN}访问地址：${NC}"
    echo "    前端: http://localhost"
    echo "    phpMyAdmin: http://localhost:8080"
    echo ""
    echo -e "  ${CYAN}默认登录信息：${NC}"
    echo "    MySQL用户: bailemi"
    echo "    MySQL密码: bailemi_2024"
    echo "    Redis密码: bailemi_redis_2024"
    echo ""
    echo -e "  ${CYAN}常用命令：${NC}"
    echo "    查看容器状态: docker compose ps"
    echo "    查看日志: docker compose logs -f"
    echo "    重启服务: docker compose restart"
    echo "    停止服务: docker compose down"
    echo "    进入容器: docker exec -it bailemi-php sh"
    echo ""
    echo -e "${YELLOW}注意：首次部署后请运行安装向导！${NC}"
    echo ""
}

main() {
    print_welcome
    
    check_docker
    
    read -p "是否继续Docker部署? (y/n): " confirm
    if [[ ! "$confirm" =~ ^[Yy]$ ]]; then
        log "部署已取消"
        exit 0
    fi
    
    copy_files
    create_env
    build_containers
    start_containers
    wait_for_services
    init_database
    show_access_info
}

# 颜色定义
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'

main "$@"
