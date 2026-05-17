<?php
/**
 * 步骤1: 环境检测
 */

// 获取环境信息
$checks = [];

// PHP版本检测
$phpVersion = PHP_VERSION;
$phpRequired = '7.4.0';
$phpCheck = version_compare($phpVersion, $phpRequired, '>=');
$checks['php'] = [
    'name' => 'PHP 版本',
    'required' => ">= {$phpRequired}",
    'current' => $phpVersion,
    'pass' => $phpCheck,
    'suggestion' => $phpCheck ? '' : '请升级 PHP 版本到 7.4 或更高'
];

// Node.js 版本检测（可选）
$nodeVersion = shell_exec('node -v 2>/dev/null');
$nodeVersion = trim($nodeVersion ?: '');
$nodeCheck = !empty($nodeVersion) && version_compare(str_replace('v', '', $nodeVersion), '16.0', '>=');
$checks['node'] = [
    'name' => 'Node.js 版本',
    'required' => '>= 16.0 (可选)',
    'current' => $nodeVersion ?: '未安装',
    'pass' => $nodeCheck,
    'suggestion' => empty($nodeVersion) ? 'Node.js 未安装，前端构建功能不可用' : ''
];

// 目录读写权限
$dirs = [
    INSTALL_PATH => '网站根目录',
    INSTALL_PATH . '/config' => '配置目录',
    INSTALL_PATH . '/uploads' => '上传目录',
    INSTALL_PATH . '/runtime' => '缓存目录',
];

$permissions = [];
$allWritable = true;
foreach ($dirs as $dir => $name) {
    if (!is_dir($dir)) {
        @mkdir($dir, 0755, true);
    }
    $writable = is_writable($dir);
    $permissions[$dir] = [
        'name' => $name,
        'path' => $dir,
        'writable' => $writable,
    ];
    if (!$writable) $allWritable = false;
}
$checks['permissions'] = [
    'name' => '目录权限',
    'required' => '可读写',
    'current' => $allWritable ? '全部正常' : '存在问题',
    'pass' => $allWritable,
    'suggestion' => $allWritable ? '' : '请设置目录权限: chmod -R 755 目录名'
];

// PHP扩展检测
$extensions = [
    'pdo' => 'PDO',
    'pdo_mysql' => 'MySQLi/PDO',
    'curl' => 'cURL',
    'gd' => 'GD',
    'mbstring' => 'mbstring',
    'openssl' => 'OpenSSL',
    'json' => 'JSON',
];

$extResults = [];
$extAllPass = true;
foreach ($extensions as $ext => $name) {
    $loaded = extension_loaded($ext);
    $extResults[$ext] = [
        'name' => $name,
        'loaded' => $loaded,
    ];
    if (!$loaded) $extAllPass = false;
}
$checks['extensions'] = [
    'name' => 'PHP 扩展',
    'required' => '全部必需',
    'current' => $extAllPass ? '全部正常' : '部分缺失',
    'pass' => $extAllPass,
    'suggestion' => $extAllPass ? '' : '请在宝塔面板中安装缺失的 PHP 扩展'
];

// 函数检测
$functions = [
    'file_get_contents' => 'file_get_contents',
    'file_put_contents' => 'file_put_contents',
    'curl_init' => 'curl_init',
];

$funcResults = [];
$funcAllPass = true;
foreach ($functions as $func => $name) {
    $exists = function_exists($func);
    $funcResults[$func] = [
        'name' => $name,
        'exists' => $exists,
    ];
    if (!$exists) $funcAllPass = false;
}
$checks['functions'] = [
    'name' => 'PHP 函数',
    'required' => '全部可用',
    'current' => $funcAllPass ? '全部正常' : '部分禁用',
    'pass' => $funcAllPass,
    'suggestion' => $funcAllPass ? '' : '部分函数被禁用，可能影响功能'
];

// 环境信息
$serverInfo = [
    '服务器软件' => $_SERVER['SERVER_SOFTWARE'] ?? '未知',
    '服务器系统' => PHP_OS,
    'PHP 运行方式' => PHP_SAPI,
    '最大上传大小' => ini_get('upload_max_filesize'),
    '最大执行时间' => ini_get('max_execution_time') . '秒',
    '内存限制' => ini_get('memory_limit'),
    '时区设置' => date_default_timezone_get(),
];
?>

<h2>第一步：服务器环境检测</h2>
<p style="color: #666; margin-bottom: 20px;">请确保服务器环境满足以下要求后再继续安装</p>

<div class="alert alert-warning">
    <strong>⚠️ 注意：</strong>如果有任何检测项未通过，请先在宝塔面板中修复后再继续安装
</div>

<h3 style="margin: 20px 0 15px; color: #333;">📋 环境检测结果</h3>

<ul class="check-list">
    <?php foreach ($checks as $check): ?>
    <li class="<?php echo $check['pass'] ? 'check-pass' : 'check-fail'; ?>">
        <div>
            <strong><?php echo $check['name']; ?></strong>
            <div style="font-size: 0.9em; margin-top: 5px;">
                当前值: <code><?php echo $check['current']; ?></code>
                <?php if ($check['required']): ?>
                | 要求: <code><?php echo $check['required']; ?></code>
                <?php endif; ?>
            </div>
            <?php if ($check['suggestion']): ?>
            <div style="font-size: 0.9em; margin-top: 5px; color: #856404;">
                💡 <?php echo $check['suggestion']; ?>
            </div>
            <?php endif; ?>
        </div>
        <span class="check-item-status">
            <?php echo $check['pass'] ? '✅ 通过' : '❌ 未通过'; ?>
        </span>
    </li>
    <?php endforeach; ?>
</ul>

<h3 style="margin: 30px 0 15px; color: #333;">📊 服务器信息</h3>
<table style="width: 100%; border-collapse: collapse; background: #f8f9fa; border-radius: 8px;">
    <?php foreach ($serverInfo as $key => $value): ?>
    <tr style="border-bottom: 1px solid #e9ecef;">
        <td style="padding: 12px; font-weight: 600; color: #495057; width: 40%;"><?php echo $key; ?></td>
        <td style="padding: 12px; color: #212529;"><?php echo $value; ?></td>
    </tr>
    <?php endforeach; ?>
</table>

<h3 style="margin: 30px 0 15px; color: #333;">📁 目录权限</h3>
<table style="width: 100%; border-collapse: collapse; background: #f8f9fa; border-radius: 8px;">
    <?php foreach ($permissions as $perm): ?>
    <tr style="border-bottom: 1px solid #e9ecef;">
        <td style="padding: 12px; font-weight: 600; color: #495057;"><?php echo $perm['name']; ?></td>
        <td style="padding: 12px; color: #212529; font-family: monospace;"><?php echo $perm['path']; ?></td>
        <td style="padding: 12px; text-align: right;">
            <span style="color: <?php echo $perm['writable'] ? '#28a745' : '#dc3545'; ?>; font-weight: bold;">
                <?php echo $perm['writable'] ? '✅ 可写' : '❌ 不可写'; ?>
            </span>
        </td>
    </tr>
    <?php endforeach; ?>
</table>

<h3 style="margin: 30px 0 15px; color: #333;">🔧 PHP 扩展</h3>
<table style="width: 100%; border-collapse: collapse; background: #f8f9fa; border-radius: 8px;">
    <?php foreach ($extResults as $ext => $result): ?>
    <tr style="border-bottom: 1px solid #e9ecef;">
        <td style="padding: 12px; font-weight: 600; color: #495057;"><?php echo $result['name']; ?></td>
        <td style="padding: 12px; color: #212529; font-family: monospace;">extension=<?php echo $ext; ?></td>
        <td style="padding: 12px; text-align: right;">
            <span style="color: <?php echo $result['loaded'] ? '#28a745' : '#dc3545'; ?>; font-weight: bold;">
                <?php echo $result['loaded'] ? '✅ 已加载' : '❌ 未加载'; ?>
            </span>
        </td>
    </tr>
    <?php endforeach; ?>
</table>

<div style="margin-top: 30px; text-align: center;">
    <?php
    $allPassed = array_reduce($checks, function($carry, $item) {
        return $carry && $item['pass'];
    }, true);
    ?>
    
    <?php if ($allPassed): ?>
    <div class="alert alert-success">
        ✅ 环境检测全部通过！可以继续安装
    </div>
    <a href="?step=2" class="btn btn-primary btn-lg">
        继续安装 →
    </a>
    <?php else: ?>
    <div class="alert alert-error">
        ❌ 环境检测未全部通过，请先修复上述问题后再继续
    </div>
    <p style="color: #666; margin-top: 15px;">
        💡 提示：您可以在宝塔面板的 PHP 设置中安装扩展和调整配置
    </p>
    <?php endif; ?>
</div>
