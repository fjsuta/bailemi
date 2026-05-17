<?php
/**
 * 步骤4: 风格与演示数据选择
 */

session_start();

$dbConfig = $_SESSION['db_config'] ?? null;
$adminCreated = $_SESSION['admin_created'] ?? false;

if (!$dbConfig || !$adminCreated) {
    header('Location: ?step=3');
    exit;
}

// 处理数据导入
if ($_SERVER['REQUEST_METHOD'] === 'POST' && isset($_POST['action'])) {
    header('Content-Type: application/json');
    
    try {
        $dsn = "mysql:host={$dbConfig['host']};port={$dbConfig['port']};dbname={$dbConfig['database']};charset=utf8mb4";
        $pdo = new PDO($dsn, $dbConfig['username'], $dbConfig['password'], [
            PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
        ]);
        
        $selectedDemos = $_POST['demos'] ?? [];
        $selectedStyles = $_POST['styles'] ?? [];
        
        // 插入演示数据
        if (in_array('songs', $selectedDemos)) {
            insertDemoSongs($pdo);
        }
        
        if (in_array('playlists', $selectedDemos)) {
            insertDemoPlaylists($pdo);
        }
        
        if (in_array('comments', $selectedDemos)) {
            insertDemoComments($pdo);
        }
        
        // 更新配置
        updateConfig($selectedStyles);
        
        // 创建配置文件
        createConfigFile($dbConfig);
        
        echo json_encode([
            'code' => 0,
            'message' => '数据导入成功！正在跳转到完成页面...'
        ]);
        
    } catch (Exception $e) {
        echo json_encode(['code' => 400, 'message' => '导入失败: ' . $e->getMessage()]);
    }
    exit;
}

// 插入演示歌曲
function insertDemoSongs($pdo) {
    $songs = [
        ['title' => '晴天', 'artist' => '周杰伦', 'album' => '叶惠美', 'duration' => 267, 'play_count' => 1000000],
        ['title' => '七里香', 'artist' => '周杰伦', 'album' => '七里香', 'duration' => 289, 'play_count' => 800000],
        ['title' => '稻香', 'artist' => '周杰伦', 'album' => '魔杰座', 'duration' => 231, 'play_count' => 900000],
        ['title' => '江南', 'artist' => '林俊杰', 'album' => '第二天堂', 'duration' => 245, 'play_count' => 850000],
        ['title' => '修炼爱情', 'artist' => '林俊杰', 'album' => '因你而在', 'duration' => 254, 'play_count' => 750000],
        ['title' => '光年之外', 'artist' => '邓紫棋', 'album' => '光年之外', 'duration' => 235, 'play_count' => 600000],
        ['title' => '泡沫', 'artist' => '邓紫棋', 'album' => 'Xposed', 'duration' => 268, 'play_count' => 550000],
        ['title' => '十年', 'artist' => '陈奕迅', 'album' => '黑白灰', 'duration' => 205, 'play_count' => 1200000],
        ['title' => '浮夸', 'artist' => '陈奕迅', 'album' => '认了吧', 'duration' => 280, 'play_count' => 950000],
        ['title' => '富士山下', 'artist' => '陈奕迅', 'album' => 'What's Going On...?', 'duration' => 242, 'play_count' => 880000],
    ];
    
    $stmt = $pdo->prepare("
        INSERT INTO songs (title, artist, album, duration, play_count, status, created_at) 
        VALUES (?, ?, ?, ?, ?, 1, NOW())
    ");
    
    foreach ($songs as $song) {
        $stmt->execute([
            $song['title'],
            $song['artist'],
            $song['album'],
            $song['duration'],
            $song['play_count']
        ]);
    }
}

// 插入演示歌单
function insertDemoPlaylists($pdo) {
    $playlists = [
        ['title' => '华语流行精选', 'description' => '精选最热门的华语流行歌曲', 'song_count' => 50],
        ['title' => '经典老歌回忆', 'description' => '回味经典老歌，回忆美好时光', 'song_count' => 40],
        ['title' => '深夜治愈系', 'description' => '适合深夜一个人听的歌', 'song_count' => 30],
        ['title' => '跑步健身BGM', 'description' => '动感十足的运动音乐', 'song_count' => 35],
    ];
    
    $stmt = $pdo->prepare("
        INSERT INTO playlists (title, description, song_count, user_id, is_public, status, created_at) 
        VALUES (?, ?, ?, 1, 1, 1, NOW())
    ");
    
    foreach ($playlists as $playlist) {
        $stmt->execute([
            $playlist['title'],
            $playlist['description'],
            $playlist['song_count']
        ]);
    }
}

// 插入演示评论
function insertDemoComments($pdo) {
    $comments = [
        ['content' => '这首歌太好听了！百听不厌！', 'song_id' => 1, 'user_id' => 1],
        ['content' => '周董永远的神！', 'song_id' => 1, 'user_id' => 1],
        ['content' => '这首让我想起了青春', 'song_id' => 2, 'user_id' => 1],
    ];
    
    $stmt = $pdo->prepare("
        INSERT INTO comments (content, song_id, user_id, status, created_at) 
        VALUES (?, ?, ?, 1, NOW())
    ");
    
    foreach ($comments as $comment) {
        $stmt->execute([
            $comment['content'],
            $comment['song_id'],
            $comment['user_id']
        ]);
    }
}

// 更新配置
function updateConfig($styles) {
    // 这里可以根据选择的风格更新配置
    // 例如：主题颜色、布局等
}

// 创建配置文件
function createConfigFile($dbConfig) {
    $config = "<?php\n";
    $config .= "/**\n";
    $config .= " * 百米乐配置文件\n";
    $config .= " * 由安装向导自动生成\n";
    $config .= " */\n\n";
    $config .= "return [\n";
    $config .= "    'db' => [\n";
    $config .= "        'host' => '{$dbConfig['host']}',\n";
    $config .= "        'port' => '{$dbConfig['port']}',\n";
    $config .= "        'database' => '{$dbConfig['database']}',\n";
    $config .= "        'username' => '{$dbConfig['username']}',\n";
    $config .= "        'password' => '{$dbConfig['password']}',\n";
    $config .= "        'charset' => 'utf8mb4',\n";
    $config .= "    ],\n";
    $config .= "    'app' => [\n";
    $config .= "        'name' => '百米乐',\n";
    $config .= "        'version' => '1.0.0',\n";
    $config .= "        'timezone' => 'Asia/Shanghai',\n";
    $config .= "    ],\n";
    $config .= "    'security' => [\n";
    $config .= "        'jwt_secret' => bin2hex(random_bytes(32)),\n";
    $config .= "        'jwt_ttl' => 86400,\n";
    $config .= "    ],\n";
    $config .= "];\n";
    
    file_put_contents(CONFIG_FILE, $config);
}
?>

<h2>第四步：风格与演示数据选择</h2>
<p style="color: #666; margin-bottom: 20px;">
    选择您想要的主题风格和导入的演示数据
</p>

<form id="installForm" method="POST">
    <input type="hidden" name="action" value="import_data">
    
    <h3 style="margin: 30px 0 20px; color: #333;">🎨 主题风格</h3>
    <p style="color: #666; margin-bottom: 15px;">选择平台的主题风格</p>
    
    <div class="checkbox-group">
        <label class="checkbox-item checked">
            <input type="checkbox" name="styles[]" value="purple" checked>
            <h4>🎵 紫色主题（默认）</h4>
            <p>紫色渐变，现代科技感</p>
        </label>
        
        <label class="checkbox-item">
            <input type="checkbox" name="styles[]" value="blue">
            <h4>🌊 蓝色主题</h4>
            <p>蓝色渐变，清新自然</p>
        </label>
        
        <label class="checkbox-item">
            <input type="checkbox" name="styles[]" value="green">
            <h4>🌿 绿色主题</h4>
            <p>绿色渐变，健康活力</p>
        </label>
        
        <label class="checkbox-item">
            <input type="checkbox" name="styles[]" value="dark">
            <h4>🌙 深色主题</h4>
            <p>深色背景，护眼舒适</p>
        </label>
    </div>
    
    <h3 style="margin: 40px 0 20px; color: #333;">📦 演示数据</h3>
    <p style="color: #666; margin-bottom: 15px;">
        导入演示数据可以让您快速了解平台功能（可选）
    </p>
    
    <div class="checkbox-group">
        <label class="checkbox-item checked">
            <input type="checkbox" name="demos[]" value="songs" checked>
            <h4>🎵 演示歌曲</h4>
            <p>导入10首热门歌曲</p>
        </label>
        
        <label class="checkbox-item checked">
            <input type="checkbox" name="demos[]" value="playlists" checked>
            <h4>📋 演示歌单</h4>
            <p>导入4个推荐歌单</p>
        </label>
        
        <label class="checkbox-item">
            <input type="checkbox" name="demos[]" value="comments">
            <h4>💬 演示评论</h4>
            <p>导入示例评论数据</p>
        </label>
        
        <label class="checkbox-item">
            <input type="checkbox" name="demos[]" value="users">
            <h4>👥 演示用户</h4>
            <p>导入示例用户数据</p>
        </label>
    </div>
    
    <div id="importResult" class="alert" style="display: none; margin-top: 30px;"></div>
    
    <div style="margin-top: 40px; display: flex; justify-content: space-between;">
        <a href="?step=3" class="btn btn-secondary">← 上一步</a>
        <button type="submit" class="btn btn-primary btn-lg">
            导入数据并完成安装 →
        </button>
    </div>
</form>

<script>
// 点击选中效果
document.querySelectorAll('.checkbox-item').forEach(item => {
    item.addEventListener('click', function() {
        const checkbox = this.querySelector('input[type="checkbox"]');
        checkbox.checked = !checkbox.checked;
        this.classList.toggle('checked', checkbox.checked);
    });
});

// 表单提交
document.getElementById('installForm').addEventListener('submit', function(e) {
    e.preventDefault();
    
    const form = this;
    const btn = form.querySelector('button[type="submit"]');
    btn.disabled = true;
    btn.textContent = '导入中...';
    
    const resultDiv = document.getElementById('importResult');
    resultDiv.style.display = 'block';
    resultDiv.className = 'alert alert-warning';
    resultDiv.textContent = '正在导入数据，请稍候...';
    
    const formData = new FormData(form);
    
    fetch('?step=4', {
        method: 'POST',
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        if (data.code === 0) {
            resultDiv.className = 'alert alert-success';
            resultDiv.textContent = '✅ ' + data.message;
            setTimeout(() => location.href = '?step=5', 1500);
        } else {
            resultDiv.className = 'alert alert-error';
            resultDiv.textContent = '❌ ' + data.message;
            btn.disabled = false;
            btn.textContent = '重新导入';
        }
    })
    .catch(error => {
        resultDiv.className = 'alert alert-error';
        resultDiv.textContent = '❌ 导入失败: ' + error.message;
        btn.disabled = false;
        btn.textContent = '重新导入';
    });
});
</script>
