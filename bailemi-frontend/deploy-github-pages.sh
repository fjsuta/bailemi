#!/bin/bash

# GitHub Pages 部署脚本
# 用于将前端构建产物部署到 gh-pages 分支

set -e

REPO_NAME="bailemi"
REPO_OWNER="fjsuta"
SOURCE_BRANCH="master"
DEPLOY_BRANCH="gh-pages"

echo "🚀 开始部署 GitHub Pages..."

# 克隆仓库到临时目录
TEMP_DIR=$(mktemp -d)
echo "📦 克隆仓库到临时目录: $TEMP_DIR"

git clone --depth 1 --branch $SOURCE_BRANCH https://github.com/$REPO_OWNER/$REPO_NAME.git $TEMP_DIR/$REPO_NAME

cd $TEMP_DIR/$REPO_NAME/bailemi-frontend

# 安装依赖
echo "📦 安装依赖..."
npm install

# 构建生产版本
echo "🔨 构建生产版本..."
npm run build

# 进入构建目录
cd dist

# 初始化临时仓库
echo "📦 初始化部署仓库..."
git init
git config user.name "GitHub Actions"
git config user.email "actions@github.com"

# 添加所有文件
git add .

# 提交
git commit -m "Deploy to GitHub Pages - $(date +'%Y-%m-%d %H:%M:%S')"

# 推送到 gh-pages 分支
echo "🚀 推送到 gh-pages 分支..."
git push -f https://github.com/$REPO_OWNER/$REPO_NAME.git master:$DEPLOY_BRANCH

# 清理
rm -rf $TEMP_DIR

echo "✅ 部署完成！"
echo "🌐 访问地址: https://$REPO_OWNER.github.io/$REPO_NAME/"
