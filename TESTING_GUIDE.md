# 测试指南

## 📋 测试概览

本项目包含前后端两套完整的测试体系，采用**系统化调试方法（Systematic Debugging）**建立，确保代码质量和功能正确性。

---

## 🖥️ 后端测试

### 技术栈
- **JUnit 5** - 单元测试框架
- **Mockito** - Mock 对象框架
- **MockMvc** - Web 层集成测试
- **H2 Database** - 内存数据库（测试用）

### 测试目录结构
```
bailemi-spring/
└── src/
    └── test/
        ├── java/com/bailemi/
        │   ├── controller/
        │   │   ├── AuthControllerTest.java
        │   │   └── AdminControllerTest.java
        │   ├── service/
        │   │   └── UserServiceTest.java
        │   └── repository/
        └── resources/
            └── application.yml  # 测试配置
```

### 运行测试

#### 运行所有测试
```bash
cd bailemi-spring
mvn test
```

#### 运行特定测试类
```bash
mvn test -Dtest=UserServiceTest
mvn test -Dtest=AuthControllerTest
mvn test -Dtest=AdminControllerTest
```

#### 运行测试并生成覆盖率报告
```bash
mvn test jacoco:report
```

### 测试覆盖范围

#### 1. UserServiceTest - 业务逻辑测试
- ✅ 用户注册（正常流程、重复用户名、重复邮箱）
- ✅ 用户登录（正常流程、封禁用户）
- ✅ 更新用户状态
- ✅ 获取用户详情
- ✅ 删除用户（软删除）

#### 2. AuthControllerTest - 认证接口测试
- ✅ 注册接口（成功、失败、参数验证）
- ✅ 登录接口（成功、失败、参数验证）

#### 3. AdminControllerTest - 管理员接口测试
- ✅ 查询用户列表（权限验证）
- ✅ 获取用户详情
- ✅ 封禁/解封用户
- ✅ 删除用户
- ✅ 权限控制测试（ADMIN vs 普通用户）

---

## 🎨 前端测试

### 技术栈
- **Vitest** - 测试运行器
- **@testing-library/vue** - Vue 组件测试
- **@testing-library/jest-dom** - DOM 断言
- **jsdom** - DOM 环境模拟

### 测试目录结构
```
bailemi-frontend/
├── src/
│   └── test/
│       ├── setup.js           # 测试环境配置
│       ├── stores/
│       │   ├── auth.test.js
│       │   └── notification.test.js
│       ├── utils/
│       │   └── api.test.js
│       └── components/
├── vitest.config.js           # Vitest 配置
└── package.json               # 测试依赖
```

### 安装测试依赖
```bash
cd bailemi-frontend
npm install
```

### 运行测试

#### 运行所有测试
```bash
npm test
```

#### 运行测试（监视模式）
```bash
npm run test:watch
```

#### 运行测试并打开 UI
```bash
npm run test:ui
```

#### 运行测试并生成覆盖率报告
```bash
npm run test:coverage
```

### 测试覆盖范围

#### 1. auth.test.js - 认证 Store 测试
- ✅ 初始化状态
- ✅ 设置用户信息
- ✅ 设置和清除 Token
- ✅ 登出功能
- ✅ 认证状态计算

#### 2. notification.test.js - 通知 Store 测试
- ✅ 添加通知
- ✅ 移除通知
- ✅ 标记已读
- ✅ 全部标记已读
- ✅ 清空所有通知
- ✅ 快捷方法（success, error, warning, info）
- ✅ 自动移除定时器

#### 3. api.test.js - API 工具测试
- ✅ 配置验证
- ✅ 实例导出
- ✅ 拦截器配置

---

## 📊 测试覆盖率目标

| 模块 | 目标覆盖率 |
|------|-----------|
| **后端 Service 层** | ≥ 80% |
| **后端 Controller 层** | ≥ 70% |
| **前端 Store/Pinia** | ≥ 80% |
| **前端 Utils** | ≥ 90% |

---

## 🔧 最佳实践

### 后端测试
1. **单元隔离** - 使用 Mockito 隔离外部依赖
2. **测试命名** - 使用清晰的测试方法命名
3. **AAA 模式** - Arrange（准备）→ Act（执行）→ Assert（断言）
4. **单一职责** - 每个测试只验证一个功能点

### 前端测试
1. **以用户为中心** - 测试用户行为而非实现细节
2. **就近原则** - 测试文件与源文件放在同一目录的 test 文件夹
3. **组件测试** - 测试组件的渲染和交互，而非内部状态
4. **Mock API** - 使用 Mock 避免真实网络请求

---

## 🚀 持续集成

建议在 CI/CD 中配置自动测试：

### GitHub Actions 示例
```yaml
name: Test
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Backend Tests
        run: |
          cd bailemi-spring
          mvn test
      - name: Frontend Tests
        run: |
          cd bailemi-frontend
          npm install
          npm run test:coverage
```

---

## 📚 扩展测试

### 添加新测试

#### 后端 - 添加 Service 测试
1. 在 `src/test/java/com/bailemi/service/` 下创建 `XxxServiceTest.java`
2. 使用 `@ExtendWith(MockitoExtension.class)` 注解
3. 使用 `@Mock` 创建 Mock 对象
4. 使用 `@InjectMocks` 注入被测试对象

#### 前端 - 添加 Store 测试
1. 在 `src/test/stores/` 下创建 `xxx.test.js`
2. 使用 `setActivePinia(createPinia())` 初始化 Pinia
3. 使用 `describe` 和 `it` 组织测试
4. 使用 `expect` 进行断言

---

## 🐛 调试测试失败的技巧

### 使用系统化调试方法

#### Phase 1: 理解错误信息
- 仔细阅读断言错误
- 查看完整的堆栈跟踪
- 注意行号和文件名

#### Phase 2: 隔离问题
- 运行单个测试
- 检查 Mock 对象是否正确配置
- 验证测试数据是否正确

#### Phase 3: 验证假设
- 使用 `console.log` 或断点
- 简化测试到最小可复现案例
- 对比成功和失败的测试

#### Phase 4: 修复并验证
- 只修改必要部分
- 运行所有测试确保没有破坏其他功能
- 添加新的测试用例覆盖边界情况

---

## 📞 获取帮助

- 查看 [Vitest 文档](https://vitest.dev/)
- 查看 [Testing Library Vue 文档](https://testing-library.com/docs/vue-testing-library/intro/)
- 查看 [Spring Boot 测试文档](https://docs.spring.io/spring-boot/docs/current/reference/html/features.html#features.testing)
- 查看 [Mockito 文档](https://site.mockito.org/)

---

**记住：系统化调试是解决问题的最佳方式！** 🔍