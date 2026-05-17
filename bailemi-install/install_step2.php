<?php
/**
 * 步骤2: 数据库配置
 */

// 处理数据库连接测试
$dbTestResult = null;
$dbTestError = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST' && isset($_POST['action'])) {
    header('Content-Type: application/json');
    
    $host = $_POST['host'] ?? 'localhost';
    $port = $_POST['port'] ?? '3306';
    $database = $_POST['database'] ?? '';
    $username = $_POST['username'] ?? '';
    $password = $_POST['password'] ?? '';
    
    try {
        $dsn = "mysql:host={$host};port={$port};dbname={$database};charset=utf8mb4";
        $options = [
            PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
            PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
            PDO::ATTR_EMULATE_PREPARES => false,
        ];
        
        $pdo = new PDO($dsn, $username, $password, $options);
        
        // 测试查询
        $pdo->query("SELECT 1");
        
        // 保存配置到 session
        session_start();
        $_SESSION['db_config'] = [
            'host' => $host,
            'port' => $port,
            'database' => $database,
            'username' => $username,
            'password' => $password,
        ];
        
        echo json_encode(['code' => 0, 'message' => '数据库连接成功！']);
        
    } catch (PDOException $e) {
        echo json_encode(['code' => 400, 'message' => '数据库连接失败: ' . $e->getMessage()]);
    }
    exit;
}

// 从 session 获取已保存的配置
session_start();
$savedConfig = $_SESSION['db_config'] ?? [];
?>

<h2>第二步：数据库配置</h2>
<p style="color: #666; margin-bottom: 20px;">
    请填写您在宝塔面板中创建的数据库信息
</p>

<div class="alert alert-warning">
    <strong>📝 如何创建数据库？</strong>
    <ol style="margin: 10px 0 0 20px;">
        <li>登录宝塔面板</li>
        <li>点击左侧「数据库」菜单</li>
        <li>点击「添加数据库」，填写数据库名、用户名和密码</li>
        <li>确保数据库编码为 utf8mb4</li>
    </ol>
</div>

<form id="dbForm" method="POST">
    <div class="form-group">
        <label>数据库地址 *</label>
        <input type="text" name="host" value="<?php echo $savedConfig['host'] ?? 'localhost'; ?>" 
               placeholder="通常为 localhost 或 127.0.0.1" required>
    </div>
    
    <div class="form-group">
        <label>数据库端口 *</label>
        <input type="number" name="port" value="<?php echo $savedConfig['port'] ?? '3306'; ?>" 
               placeholder="MySQL 端口，通常为 3306" required>
    </div>
    
    <div class="form-group">
        <label>数据库名称 *</label>
        <input type="text" name="database" value="<?php echo $savedConfig['database'] ?? ''; ?>" 
               placeholder="在宝塔面板创建的数据库名" required>
    </div>
    
    <div class="form-group">
        <label>数据库用户名 *</label>
        <input type="text" name="username" value="<?php echo $savedConfig['username'] ?? ''; ?>" 
               placeholder="在宝塔面板创建的用户名" required>
    </div>
    
    <div class="form-group">
        <label>数据库密码 *</label>
        <input type="password" name="password" value="<?php echo $savedConfig['password'] ?? ''; ?>" 
               placeholder="数据库密码" required>
    </div>
    
    <div id="testResult" class="alert" style="display: none;"></div>
    
    <div style="display: flex; gap: 15px; margin-top: 30px;">
        <button type="button" onclick="testDatabase()" class="btn btn-secondary">
            🔍 测试连接
        </button>
        <button type="submit" class="btn btn-primary">
            保存并继续 →
        </button>
    </div>
</form>

<script>
function testDatabase() {
    const form = document.getElementById('dbForm');
    const formData = new FormData(form);
    formData.append('action', 'test');
    
    const resultDiv = document.getElementById('testResult');
    resultDiv.style.display = 'block';
    resultDiv.className = 'alert alert-warning';
    resultDiv.textContent = '正在测试数据库连接...';
    
    fetch('?step=2', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        if (data.code === 0) {
            resultDiv.className = 'alert alert-success';
            resultDiv.textContent = '✅ ' + data.message;
        } else {
            resultDiv.className = 'alert alert-error';
            resultDiv.textContent = '❌ ' + data.message;
        }
    })
    .catch(error => {
        resultDiv.className = 'alert alert-error';
        resultDiv.textContent = '❌ 测试失败: ' + error.message;
    });
}
</script>
