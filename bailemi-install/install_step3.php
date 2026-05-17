<?php
/**
 * 步骤3: 数据表创建与管理员设置
 */

session_start();

// 处理数据表创建
if ($_SERVER['REQUEST_METHOD'] === 'POST' && isset($_POST['action'])) {
    header('Content-Type: application/json');
    
    if ($_POST['action'] === 'create_tables') {
        try {
            // 从 session 获取数据库配置
            $dbConfig = $_SESSION['db_config'] ?? null;
            if (!$dbConfig) {
                throw new Exception('数据库配置不存在，请返回上一步重新配置');
            }
            
            $dsn = "mysql:host={$dbConfig['host']};port={$dbConfig['port']};dbname={$dbConfig['database']};charset=utf8mb4";
            $pdo = new PDO($dsn, $dbConfig['username'], $dbConfig['password'], [
                PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
            ]);
            
            // 读取 SQL 文件
            $sqlFile = INSTALL_PATH . '/install.sql';
            if (!file_exists($sqlFile)) {
                throw new Exception('SQL 安装文件不存在');
            }
            
            $sql = file_get_contents($sqlFile);
            
            // 分割 SQL 语句
            $statements = array_filter(array_map('trim', explode(';', $sql)));
            
            // 执行 SQL
            $pdo->beginTransaction();
            foreach ($statements as $statement) {
                if (!empty($statement)) {
                    $pdo->exec($statement);
                }
            }
            $pdo->commit();
            
            echo json_encode(['code' => 0, 'message' => '数据表创建成功！']);
            
        } catch (Exception $e) {
            if (isset($pdo) && $pdo->inTransaction()) {
                $pdo->rollBack();
            }
            echo json_encode(['code' => 400, 'message' => '创建失败: ' . $e->getMessage()]);
        }
        exit;
    }
    
    if ($_POST['action'] === 'create_admin') {
        try {
            $username = $_POST['username'] ?? '';
            $password = $_POST['password'] ?? '';
            $email = $_POST['email'] ?? '';
            $phone = $_POST['phone'] ?? '';
            
            if (strlen($username) < 3) {
                throw new Exception('用户名至少需要3个字符');
            }
            if (strlen($password) < 6) {
                throw new Exception('密码至少需要6个字符');
            }
            
            $dbConfig = $_SESSION['db_config'] ?? null;
            if (!$dbConfig) {
                throw new Exception('数据库配置不存在');
            }
            
            $dsn = "mysql:host={$dbConfig['host']};port={$dbConfig['port']};dbname={$dbConfig['database']};charset=utf8mb4";
            $pdo = new PDO($dsn, $dbConfig['username'], $dbConfig['password'], [
                PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
            ]);
            
            // 检查管理员是否已存在
            $stmt = $pdo->prepare("SELECT COUNT(*) FROM users WHERE role = 'admin'");
            $stmt->execute();
            if ($stmt->fetchColumn() > 0) {
                throw new Exception('管理员已存在，如需重置请清空 users 表');
            }
            
            // 密码哈希
            $passwordHash = password_hash($password, PASSWORD_DEFAULT);
            
            // 创建管理员
            $stmt = $pdo->prepare("
                INSERT INTO users (username, password, email, phone, role, status, created_at) 
                VALUES (?, ?, ?, ?, 'admin', 1, NOW())
            ");
            $stmt->execute([$username, $passwordHash, $email, $phone]);
            
            // 保存到 session
            $_SESSION['admin_created'] = true;
            
            echo json_encode([
                'code' => 0, 
                'message' => '管理员创建成功！',
                'data' => [
                    'username' => $username,
                    'password' => $password
                ]
            ]);
            
        } catch (Exception $e) {
            echo json_encode(['code' => 400, 'message' => $e->getMessage()]);
        }
        exit;
    }
}

$dbConfig = $_SESSION['db_config'] ?? null;
$tablesCreated = isset($_SESSION['tables_created']);
$adminCreated = isset($_SESSION['admin_created']);
?>

<h2>第三步：数据表创建与管理员设置</h2>
<p style="color: #666; margin-bottom: 20px;">
    创建数据库表并设置第一个管理员账号
</p>

<?php if (!$tablesCreated): ?>
<div class="alert alert-warning">
    <strong>📦 步骤 3.1: 创建数据表</strong>
    <p style="margin-top: 10px;">点击下方按钮自动创建所有必需的数据表</p>
</div>

<div id="tableResult" class="alert" style="display: none;"></div>

<div style="margin: 30px 0;">
    <button type="button" onclick="createTables()" class="btn btn-primary btn-lg">
        🚀 创建数据表
    </button>
</div>

<div style="margin-top: 30px;">
    <h4 style="margin-bottom: 15px;">将创建以下数据表：</h4>
    <ul style="color: #666; line-height: 2;">
        <li>✅ users - 用户表</li>
        <li>✅ songs - 歌曲表</li>
        <li>✅ albums - 专辑表</li>
        <li>✅ artists - 歌手表</li>
        <li>✅ playlists - 歌单表</li>
        <li>✅ playlist_songs - 歌单歌曲关联表</li>
        <li>✅ comments - 评论表</li>
        <li>✅ favorites - 收藏表</li>
        <li>✅ play_history - 播放记录表</li>
        <li>✅ roles - 角色表</li>
        <li>✅ permissions - 权限表</li>
    </ul>
</div>
<?php else: ?>
<div class="alert alert-success">
    ✅ 数据表已创建成功！
</div>
<?php endif; ?>

<?php if ($tablesCreated && !$adminCreated): ?>
<div class="alert alert-warning" style="margin-top: 30px;">
    <strong>👤 步骤 3.2: 设置管理员账号</strong>
    <p style="margin-top: 10px;">请设置您的管理员登录凭证</p>
</div>

<form id="adminForm">
    <div class="form-group">
        <label>管理员用户名 *</label>
        <input type="text" name="username" placeholder="请输入管理员用户名（至少3个字符）" 
               required minlength="3">
    </div>
    
    <div class="form-group">
        <label>管理员密码 *</label>
        <input type="password" name="password" placeholder="请输入管理员密码（至少6个字符）" 
               required minlength="6">
    </div>
    
    <div class="form-group">
        <label>管理员邮箱</label>
        <input type="email" name="email" placeholder="用于找回密码（可选）">
    </div>
    
    <div class="form-group">
        <label>管理员手机</label>
        <input type="tel" name="phone" placeholder="用于找回密码（可选）">
    </div>
    
    <div id="adminResult" class="alert" style="display: none;"></div>
    
    <div style="margin-top: 30px;">
        <button type="submit" class="btn btn-primary btn-lg">
            👤 创建管理员账号
        </button>
    </div>
</form>
<?php elseif ($adminCreated): ?>
<div class="alert alert-success" style="margin-top: 30px;">
    ✅ 管理员账号已创建成功！
</div>
<?php endif; ?>

<div style="margin-top: 30px; display: flex; justify-content: space-between;">
    <a href="?step=2" class="btn btn-secondary">← 上一步</a>
    <?php if ($tablesCreated && $adminCreated): ?>
    <a href="?step=4" class="btn btn-primary">继续安装 →</a>
    <?php endif; ?>
</div>

<script>
function createTables() {
    const btn = event.target;
    btn.disabled = true;
    btn.textContent = '创建中...';
    
    const resultDiv = document.getElementById('tableResult');
    resultDiv.style.display = 'block';
    resultDiv.className = 'alert alert-warning';
    resultDiv.textContent = '正在创建数据表，请稍候...';
    
    fetch('?step=3', {
        method: 'POST',
        headers: {'Content-Type': 'application/x-www-form-urlencoded'},
        body: 'action=create_tables'
    })
    .then(response => response.json())
    .then(data => {
        if (data.code === 0) {
            resultDiv.className = 'alert alert-success';
            resultDiv.innerHTML = '✅ ' + data.message + '<br>页面将自动刷新...';
            setTimeout(() => location.reload(), 1500);
        } else {
            resultDiv.className = 'alert alert-error';
            resultDiv.textContent = '❌ ' + data.message;
            btn.disabled = false;
            btn.textContent = '🚀 重新创建数据表';
        }
    })
    .catch(error => {
        resultDiv.className = 'alert alert-error';
        resultDiv.textContent = '❌ 创建失败: ' + error.message;
        btn.disabled = false;
        btn.textContent = '🚀 重新创建数据表';
    });
}

document.getElementById('adminForm')?.addEventListener('submit', function(e) {
    e.preventDefault();
    
    const formData = new FormData(this);
    formData.append('action', 'create_admin');
    
    const btn = this.querySelector('button[type="submit"]');
    btn.disabled = true;
    btn.textContent = '创建中...';
    
    const resultDiv = document.getElementById('adminResult');
    resultDiv.style.display = 'block';
    resultDiv.className = 'alert alert-warning';
    resultDiv.textContent = '正在创建管理员账号...';
    
    fetch('?step=3', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        if (data.code === 0) {
            resultDiv.className = 'alert alert-success';
            resultDiv.innerHTML = `
                ✅ 管理员创建成功！<br><br>
                <strong>登录信息：</strong><br>
                用户名: ${data.data.username}<br>
                密码: ${data.data.password}<br><br>
                <em>请妥善保管您的登录凭证！</em>
            `;
            setTimeout(() => location.href = '?step=4', 2000);
        } else {
            resultDiv.className = 'alert alert-error';
            resultDiv.textContent = '❌ ' + data.message;
            btn.disabled = false;
            btn.textContent = '👤 重新创建管理员';
        }
    })
    .catch(error => {
        resultDiv.className = 'alert alert-error';
        resultDiv.textContent = '❌ 创建失败: ' + error.message;
        btn.disabled = false;
        btn.textContent = '👤 重新创建管理员';
    });
});
</script>
