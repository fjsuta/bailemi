#!/bin/bash

echo "🧪 测试百米乐音乐平台 API"
echo "===================================="
echo ""

# 测试健康检查
echo "1. 测试健康检查..."
curl -s http://localhost:8080/health
echo -e "\n"

# 测试登录
echo "2. 测试登录..."
curl -s -X POST http://localhost:8080/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"login_type":"username","account":"demo","password":"demo"}'
echo -e "\n"

# 测试获取热门歌曲
echo "3. 测试获取热门歌曲..."
curl -s http://localhost:8080/v1/song/hot | head -c 200
echo -e "\n"

# 测试注册
echo "4. 测试注册..."
curl -s -X POST http://localhost:8080/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","email":"test@example.com","password":"password123","verify_code":"123456"}'
echo -e "\n"

# 测试验证码发送
echo "5. 测试发送验证码..."
curl -s -X POST http://localhost:8080/v1/auth/send-code \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com"}'
echo -e "\n"

echo "===================================="
echo "✅ 所有测试完成！"
echo ""
echo "前端访问地址: http://localhost:3000"
echo "API 地址: http://localhost:8080"
