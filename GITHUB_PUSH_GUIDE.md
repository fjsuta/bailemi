# 🚀 GitHub 推送指南

## 当前状态

项目代码已准备就绪，但需要 GitHub 认证才能推送。

## 解决方案

### 方法一：使用 GitHub Token（推荐用于自动化）

1. **创建 Personal Access Token**
   - 访问 https://github.com/settings/tokens
   - 点击 "Generate new token (classic)"
   - 设置 Token 名称（如：bailemi-deploy）
   - 选择权限：
     - ✅ repo (Full control of private repositories)
   - 点击 "Generate token"
   - **立即复制 Token（只会显示一次）**

2. **设置环境变量**
   ```bash
   export GH_TOKEN="ghp_xxxxxxxxxxxxxxxxxxxx"
   ```

3. **运行推送脚本**
   ```bash
   cd /workspace/bailemi
   bash ../push_to_github.sh
   ```

### 方法二：交互式登录（本地开发环境）

```bash
cd /workspace/bailemi

# 交互式登录
gh auth login

# 设置 Git 凭证助手
gh auth setup-git

# 创建仓库
gh repo create bailemi --public --description "百米乐开源音乐平台"

# 添加远程仓库
git remote add origin https://github.com/YOUR_USERNAME/bailemi.git

# 推送代码
git push -u origin master
```

### 方法三：手动推送

1. **在 GitHub 上创建仓库**
   - 访问 https://github.com/new
   - 仓库名称：`bailemi`
   - 描述：`🎵 百米乐 - 现代开源音乐平台`
   - 选择 Public（公开）
   - 不要勾选 "Initialize this repository with a README"
   - 点击 "Create repository"

2. **运行以下命令**（替换 `YOUR_USERNAME` 为您的 GitHub 用户名）
   ```bash
   cd /workspace/bailemi
   git remote add origin https://github.com/YOUR_USERNAME/bailemi.git
   git branch -M main
   git push -u origin main
   ```

## 推送脚本

我已经为您创建了自动推送脚本：

```bash
# 脚本位置
/workspace/push_to_github.sh
```

## 验证推送

推送成功后，您可以访问：
- 仓库地址：https://github.com/YOUR_USERNAME/bailemi

## 注意事项

1. 确保您有 GitHub 账号
2. 如果使用 Token，确保 Token 有 `repo` 权限
3. 仓库将公开（Public），所有人都可以访问
4. 代码采用 Apache 2.0 开源协议

## 下一步

1. 创建 GitHub 账号（如果没有）
2. 获取 Personal Access Token
3. 运行推送脚本

祝您使用愉快！🎵
