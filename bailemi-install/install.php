<?php
/**
 * 百米乐音乐平台 - 一键安装向导
 * 版本: 1.0.0
 * 作者: Bailemi Development Team
 */

// 错误报告
error_reporting(E_ALL);
ini_set('display_errors', 0);

// 定义常量
define('INSTALL_PATH', dirname(__FILE__));
define('CONFIG_FILE', INSTALL_PATH . '/config.php');
define('ENV_FILE', INSTALL_PATH . '/.env');
define('LOCK_FILE', INSTALL_PATH . '/install.lock');

// 检测是否已安装
if (file_exists(LOCK_FILE) && !isset($_GET['force'])) {
    die(json_encode([
        'code' => 403,
        'message' => '安装已完成，如需重新安装请删除 install.lock 文件'
    ]));
}

// 获取当前步骤
$step = isset($_GET['step']) ? intval($_GET['step']) : 1;
$action = isset($_GET['action']) ? $_GET['action'] : '';

// 路由处理
header('Content-Type: text/html; charset=utf-8');
?>
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>百米乐安装向导</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }
        .container {
            max-width: 900px;
            margin: 0 auto;
            background: white;
            border-radius: 20px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            overflow: hidden;
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 40px;
            text-align: center;
        }
        .header h1 { font-size: 2.5rem; margin-bottom: 10px; }
        .header p { opacity: 0.9; font-size: 1.1rem; }
        .steps {
            display: flex;
            justify-content: space-between;
            padding: 30px 40px;
            background: #f8f9fa;
            border-bottom: 1px solid #e9ecef;
        }
        .step {
            display: flex;
            align-items: center;
            gap: 10px;
            color: #adb5bd;
        }
        .step.active { color: #667eea; font-weight: bold; }
        .step.completed { color: #28a745; }
        .step-number {
            width: 36px;
            height: 36px;
            border-radius: 50%;
            background: #e9ecef;
            display: flex;
            align-items: center;
            justify-content: center;
            font-weight: bold;
        }
        .step.active .step-number { background: #667eea; color: white; }
        .step.completed .step-number { background: #28a745; color: white; }
        .content {
            padding: 40px;
            min-height: 400px;
        }
        .btn {
            padding: 12px 30px;
            border: none;
            border-radius: 8px;
            font-size: 1rem;
            cursor: pointer;
            transition: all 0.3s;
            text-decoration: none;
            display: inline-block;
        }
        .btn-primary {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
        }
        .btn-primary:hover { transform: translateY(-2px); box-shadow: 0 5px 20px rgba(102, 126, 234, 0.4); }
        .btn-success { background: #28a745; color: white; }
        .btn-secondary { background: #6c757d; color: white; }
        .form-group {
            margin-bottom: 20px;
        }
        .form-group label {
            display: block;
            margin-bottom: 8px;
            font-weight: 600;
            color: #333;
        }
        .form-group input, .form-group select, .form-group textarea {
            width: 100%;
            padding: 12px;
            border: 2px solid #e9ecef;
            border-radius: 8px;
            font-size: 1rem;
            transition: border-color 0.3s;
        }
        .form-group input:focus, .form-group select:focus, .form-group textarea:focus {
            outline: none;
            border-color: #667eea;
        }
        .checkbox-group {
            display: grid;
            grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
            gap: 15px;
        }
        .checkbox-item {
            padding: 15px;
            border: 2px solid #e9ecef;
            border-radius: 10px;
            cursor: pointer;
            transition: all 0.3s;
        }
        .checkbox-item:hover { border-color: #667eea; }
        .checkbox-item.checked { border-color: #667eea; background: #f8f9ff; }
        .checkbox-item input { display: none; }
        .checkbox-item h4 { margin-bottom: 5px; color: #333; }
        .checkbox-item p { color: #666; font-size: 0.9rem; }
        .alert {
            padding: 15px;
            border-radius: 8px;
            margin-bottom: 20px;
        }
        .alert-success { background: #d4edda; color: #155724; border: 1px solid #c3e6cb; }
        .alert-error { background: #f8d7da; color: #721c24; border: 1px solid #f5c6cb; }
        .alert-warning { background: #fff3cd; color: #856404; border: 1px solid #ffeeba; }
        .check-list { list-style: none; padding: 0; }
        .check-list li {
            padding: 12px;
            margin-bottom: 10px;
            border-radius: 8px;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .check-pass { background: #d4edda; color: #155724; }
        .check-fail { background: #f8d7da; color: #721c24; }
        .check-item-name { font-weight: 600; }
        .check-item-status { font-weight: bold; }
        .footer {
            padding: 30px 40px;
            background: #f8f9fa;
            display: flex;
            justify-content: space-between;
            align-items: center;
        }
        .progress-bar {
            height: 6px;
            background: #e9ecef;
            border-radius: 3px;
            margin-top: 20px;
            overflow: hidden;
        }
        .progress-fill {
            height: 100%;
            background: linear-gradient(90deg, #667eea, #764ba2);
            transition: width 0.5s;
        }
        .admin-avatar {
            width: 100px;
            height: 100px;
            border-radius: 50%;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            display: flex;
            align-items: center;
            justify-content: center;
            margin: 0 auto 20px;
            font-size: 3rem;
            color: white;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🎵 百米乐安装向导</h1>
            <p>现代化开源音乐平台 - 简单、快速、强大</p>
            <div class="progress-bar">
                <div class="progress-fill" style="width: <?php echo ($step / 5) * 100; ?>%"></div>
            </div>
        </div>

        <div class="steps">
            <div class="step <?php echo $step >= 1 ? ($step > 1 ? 'completed' : 'active') : ''; ?>">
                <div class="step-number"><?php echo $step > 1 ? '✓' : '1'; ?></div>
                <span>环境检测</span>
            </div>
            <div class="step <?php echo $step >= 2 ? ($step > 2 ? 'completed' : 'active') : ''; ?>">
                <div class="step-number"><?php echo $step > 2 ? '✓' : '2'; ?></div>
                <span>数据库配置</span>
            </div>
            <div class="step <?php echo $step >= 3 ? ($step > 3 ? 'completed' : 'active') : ''; ?>">
                <div class="step-number"><?php echo $step > 3 ? '✓' : '3'; ?></div>
                <span>管理员设置</span>
            </div>
            <div class="step <?php echo $step >= 4 ? ($step > 4 ? 'completed' : 'active') : ''; ?>">
                <div class="step-number"><?php echo $step > 4 ? '✓' : '4'; ?></div>
                <span>风格选择</span>
            </div>
            <div class="step <?php echo $step >= 5 ? 'completed active' : ''; ?>">
                <div class="step-number"><?php echo $step >= 5 ? '✓' : '5'; ?></div>
                <span>完成安装</span>
            </div>
        </div>

        <div class="content">
            <?php
            switch ($step) {
                case 1:
                    include 'install_step1.php';
                    break;
                case 2:
                    include 'install_step2.php';
                    break;
                case 3:
                    include 'install_step3.php';
                    break;
                case 4:
                    include 'install_step4.php';
                    break;
                case 5:
                    include 'install_step5.php';
                    break;
                default:
                    include 'install_step1.php';
            }
            ?>
        </div>

        <div class="footer">
            <span style="color: #666;">百米乐 v1.0.0</span>
            <span style="color: #999;">Powered by Bailemi Team</span>
        </div>
    </div>
</body>
</html>
