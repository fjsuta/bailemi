<?php
/**
 * 步骤5: 完成安装
 */

// 完成安装，生成锁文件
session_start();
session_destroy();

// 生成安装锁文件
$lockContent = json_encode([
    'version' => '1.0.0',
    'installed_at' => date('Y-m-d H:i:s'),
    'installed_by' => 'Bailemi Install Wizard',
    'php_version' => PHP_VERSION,
], JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE);

file_put_contents(LOCK_FILE, $lockContent);
?>

<h2>🎉 安装完成！</h2>
<p style="color: #666; margin-bottom: 30px;">
    恭喜！百米乐音乐平台已成功安装到您的服务器
</p>

<div class="alert alert-success" style="margin-bottom: 30px;">
    <h3 style="margin-bottom: 15px;">✅ 安装检查清单</h3>
    <ul style="list-style: none; padding: 0;">
        <li style="padding: 5px 0;">✅ 环境检测通过</li>
        <li style="padding: 5px 0;">✅ 数据库配置完成</li>
        <li style="padding: 5px 0;">✅ 数据表创建成功</li>
        <li style="padding: 5px 0;">✅ 管理员账号已创建</li>
        <li style="padding: 5px 0;">✅ 演示数据已导入</li>
        <li style="padding: 5px 0;">✅ 安装锁定文件已生成</li>
    </ul>
</div>

<div style="background: #f8f9fa; padding: 30px; border-radius: 10px; margin: 30px 0;">
    <h3 style="margin-bottom: 20px; color: #333;">🚀 访问您的平台</h3>
    
    <div style="margin-bottom: 20px;">
        <strong>前台地址：</strong>
        <code style="background: #e9ecef; padding: 5px 10px; border-radius: 5px; margin-left: 10px;">
            <?php echo (isset($_SERVER['HTTPS']) && $_SERVER['HTTPS'] === 'on' ? 'https' : 'http') . '://' . $_SERVER['HTTP_HOST'] . dirname($_SERVER['SCRIPT_NAME']); ?>
        </code>
    </div>
    
    <div style="margin-bottom: 20px;">
        <strong>后台管理：</strong>
        <code style="background: #e9ecef; padding: 5px 10px; border-radius: 5px; margin-left: 10px;">
            <?php echo (isset($_SERVER['HTTPS']) && $_SERVER['HTTPS'] === 'on' ? 'https' : 'http') . '://' . $_SERVER['HTTP_HOST'] . dirname($_SERVER['SCRIPT_NAME']) . '/admin'; ?>
        </code>
    </div>
</div>

<div style="background: #fff3cd; padding: 20px; border-radius: 10px; border-left: 4px solid #ffc107; margin: 30px 0;">
    <h4 style="margin-bottom: 10px; color: #856404;">⚠️ 重要安全提示</h4>
    <ul style="color: #856404; line-height: 1.8; padding-left: 20px;">
        <li>请立即登录管理后台，修改管理员密码</li>
        <li>建议删除 <code>install.php</code> 和 <code>install_*</code> 文件</li>
        <li>定期备份数据库和配置文件</li>
        <li>启用 HTTPS 以保障数据安全</li>
    </ul>
</div>

<div style="margin: 40px 0; text-align: center;">
    <a href="../" class="btn btn-primary btn-lg" style="margin: 0 10px;">
        🚀 访问前台
    </a>
    <a href="../admin" class="btn btn-success btn-lg" style="margin: 0 10px;">
        ⚙️ 管理后台
    </a>
</div>

<div style="margin-top: 40px; padding-top: 30px; border-top: 1px solid #e9ecef;">
    <h4 style="margin-bottom: 15px; color: #333;">📚 后续操作</h4>
    <ol style="color: #666; line-height: 2;">
        <li>配置 SSL 证书（宝塔面板 → 网站 → 设置 → SSL）</li>
        <li>配置邮件服务（如需要发送邮件）</li>
        <li>上传真实的音乐文件到服务器</li>
        <li>配置 CDN 加速（如使用）</li>
        <li>设置定时备份任务</li>
    </ol>
</div>

<div style="margin-top: 40px; text-align: center; color: #999; font-size: 0.9rem;">
    <p>百米乐音乐平台 v1.0.0</p>
    <p>Powered by Bailemi Development Team</p>
    <p style="margin-top: 10px;">
        <a href="https://github.com/bailemi" target="_blank" style="color: #667eea;">GitHub</a> |
        <a href="https://bailemi.com" target="_blank" style="color: #667eea;">官网</a> |
        <a href="https://docs.bailemi.com" target="_blank" style="color: #667eea;">文档</a>
    </p>
</div>
