-- =========================================
-- 百米乐音乐平台数据库安装脚本
-- 版本: 1.0.0
-- 创建时间: 2024
-- =========================================

-- 1. 用户表
CREATE TABLE IF NOT EXISTS `users` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
    `username` VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名',
    `password` VARCHAR(255) NOT NULL COMMENT '密码哈希',
    `email` VARCHAR(100) UNIQUE COMMENT '邮箱',
    `phone` VARCHAR(20) UNIQUE COMMENT '手机号',
    `avatar_url` VARCHAR(500) DEFAULT NULL COMMENT '头像URL',
    `nickname` VARCHAR(50) DEFAULT NULL COMMENT '昵称',
    `bio` VARCHAR(500) DEFAULT NULL COMMENT '个人简介',
    `role` ENUM('user', 'admin', 'vip') DEFAULT 'user' COMMENT '用户角色',
    `status` TINYINT DEFAULT 1 COMMENT '状态: 1正常, 0封禁',
    `last_login_ip` VARCHAR(45) DEFAULT NULL COMMENT '最后登录IP',
    `last_login_at` DATETIME DEFAULT NULL COMMENT '最后登录时间',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
    INDEX `idx_username` (`username`),
    INDEX `idx_email` (`email`),
    INDEX `idx_phone` (`phone`),
    INDEX `idx_role` (`role`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 2. 角色表
CREATE TABLE IF NOT EXISTS `roles` (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '角色ID',
    `name` VARCHAR(50) NOT NULL UNIQUE COMMENT '角色名称',
    `description` VARCHAR(200) DEFAULT NULL COMMENT '角色描述',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 3. 权限表
CREATE TABLE IF NOT EXISTS `permissions` (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '权限ID',
    `name` VARCHAR(100) NOT NULL UNIQUE COMMENT '权限名称',
    `description` VARCHAR(200) DEFAULT NULL COMMENT '权限描述',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- 4. 用户角色关联表
CREATE TABLE IF NOT EXISTS `user_roles` (
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `role_id` INT UNSIGNED NOT NULL COMMENT '角色ID',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`user_id`, `role_id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';

-- 5. 歌手表
CREATE TABLE IF NOT EXISTS `artists` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '歌手ID',
    `name` VARCHAR(100) NOT NULL COMMENT '歌手姓名',
    `avatar_url` VARCHAR(500) DEFAULT NULL COMMENT '头像URL',
    `cover_url` VARCHAR(500) DEFAULT NULL COMMENT '封面URL',
    `bio` TEXT DEFAULT NULL COMMENT '歌手简介',
    `country` VARCHAR(50) DEFAULT NULL COMMENT '国家/地区',
    `genre` VARCHAR(50) DEFAULT NULL COMMENT '音乐风格',
    `fans_count` INT UNSIGNED DEFAULT 0 COMMENT '粉丝数量',
    `songs_count` INT UNSIGNED DEFAULT 0 COMMENT '歌曲数量',
    `albums_count` INT UNSIGNED DEFAULT 0 COMMENT '专辑数量',
    `status` TINYINT DEFAULT 1 COMMENT '状态: 1正常, 0下架',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX `idx_name` (`name`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='歌手表';

-- 6. 专辑表
CREATE TABLE IF NOT EXISTS `albums` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '专辑ID',
    `title` VARCHAR(200) NOT NULL COMMENT '专辑标题',
    `artist_id` BIGINT UNSIGNED NOT NULL COMMENT '歌手ID',
    `cover_url` VARCHAR(500) DEFAULT NULL COMMENT '专辑封面',
    `description` TEXT DEFAULT NULL COMMENT '专辑描述',
    `genre` VARCHAR(50) DEFAULT NULL COMMENT '音乐风格',
    `release_date` DATE DEFAULT NULL COMMENT '发行日期',
    `company` VARCHAR(100) DEFAULT NULL COMMENT '发行公司',
    `songs_count` INT UNSIGNED DEFAULT 0 COMMENT '歌曲数量',
    `play_count` BIGINT UNSIGNED DEFAULT 0 COMMENT '播放次数',
    `status` TINYINT DEFAULT 1 COMMENT '状态: 1正常, 0下架',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (`artist_id`) REFERENCES `artists`(`id`) ON DELETE CASCADE,
    INDEX `idx_title` (`title`),
    INDEX `idx_artist_id` (`artist_id`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='专辑表';

-- 7. 歌曲表
CREATE TABLE IF NOT EXISTS `songs` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '歌曲ID',
    `title` VARCHAR(200) NOT NULL COMMENT '歌曲标题',
    `artist_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '歌手ID',
    `artist` VARCHAR(100) DEFAULT NULL COMMENT '歌手名称（冗余字段）',
    `album_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '专辑ID',
    `album` VARCHAR(200) DEFAULT NULL COMMENT '专辑名称（冗余字段）',
    `cover_url` VARCHAR(500) DEFAULT NULL COMMENT '封面URL',
    `play_url` VARCHAR(500) DEFAULT NULL COMMENT '播放URL',
    `duration` INT UNSIGNED DEFAULT 0 COMMENT '时长（秒）',
    `lyric_url` VARCHAR(500) DEFAULT NULL COMMENT '歌词URL',
    `genre` VARCHAR(50) DEFAULT NULL COMMENT '音乐风格',
    `language` VARCHAR(20) DEFAULT '国语' COMMENT '语言',
    `quality` VARCHAR(20) DEFAULT 'standard' COMMENT '音质: standard,HQ,SQ',
    `play_count` BIGINT UNSIGNED DEFAULT 0 COMMENT '播放次数',
    `like_count` INT UNSIGNED DEFAULT 0 COMMENT '点赞次数',
    `download_count` INT UNSIGNED DEFAULT 0 COMMENT '下载次数',
    `status` TINYINT DEFAULT 1 COMMENT '状态: 1正常, 0下架',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (`artist_id`) REFERENCES `artists`(`id`) ON DELETE SET NULL,
    FOREIGN KEY (`album_id`) REFERENCES `albums`(`id`) ON DELETE SET NULL,
    INDEX `idx_title` (`title`),
    INDEX `idx_artist_id` (`artist_id`),
    INDEX `idx_album_id` (`album_id`),
    INDEX `idx_status` (`status`),
    INDEX `idx_play_count` (`play_count`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='歌曲表';

-- 8. 歌单表
CREATE TABLE IF NOT EXISTS `playlists` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '歌单ID',
    `title` VARCHAR(200) NOT NULL COMMENT '歌单标题',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '创建用户ID',
    `cover_url` VARCHAR(500) DEFAULT NULL COMMENT '封面URL',
    `description` TEXT DEFAULT NULL COMMENT '歌单描述',
    `tags` VARCHAR(200) DEFAULT NULL COMMENT '标签，多个用逗号分隔',
    `song_count` INT UNSIGNED DEFAULT 0 COMMENT '歌曲数量',
    `play_count` BIGINT UNSIGNED DEFAULT 0 COMMENT '播放次数',
    `favorite_count` INT UNSIGNED DEFAULT 0 COMMENT '收藏次数',
    `share_count` INT UNSIGNED DEFAULT 0 COMMENT '分享次数',
    `is_public` TINYINT DEFAULT 1 COMMENT '是否公开: 1公开, 0私密',
    `status` TINYINT DEFAULT 1 COMMENT '状态: 1正常, 0下架',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    INDEX `idx_title` (`title`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_is_public` (`is_public`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='歌单表';

-- 9. 歌单歌曲关联表
CREATE TABLE IF NOT EXISTS `playlist_songs` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    `playlist_id` BIGINT UNSIGNED NOT NULL COMMENT '歌单ID',
    `song_id` BIGINT UNSIGNED NOT NULL COMMENT '歌曲ID',
    `position` INT UNSIGNED DEFAULT 0 COMMENT '排序位置',
    `added_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    FOREIGN KEY (`playlist_id`) REFERENCES `playlists`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`song_id`) REFERENCES `songs`(`id`) ON DELETE CASCADE,
    UNIQUE KEY `uk_playlist_song` (`playlist_id`, `song_id`),
    INDEX `idx_playlist_id` (`playlist_id`),
    INDEX `idx_song_id` (`song_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='歌单歌曲关联表';

-- 10. 收藏表
CREATE TABLE IF NOT EXISTS `favorites` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `song_id` BIGINT UNSIGNED NOT NULL COMMENT '歌曲ID',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '收藏时间',
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`song_id`) REFERENCES `songs`(`id`) ON DELETE CASCADE,
    UNIQUE KEY `uk_user_song` (`user_id`, `song_id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_song_id` (`song_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='收藏表';

-- 11. 评论表
CREATE TABLE IF NOT EXISTS `comments` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '评论ID',
    `content` TEXT NOT NULL COMMENT '评论内容',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `song_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '歌曲ID',
    `playlist_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '歌单ID',
    `parent_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '父评论ID（回复）',
    `like_count` INT UNSIGNED DEFAULT 0 COMMENT '点赞数',
    `status` TINYINT DEFAULT 1 COMMENT '状态: 1正常, 0删除',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`song_id`) REFERENCES `songs`(`id`) ON DELETE SET NULL,
    FOREIGN KEY (`playlist_id`) REFERENCES `playlists`(`id`) ON DELETE SET NULL,
    FOREIGN KEY (`parent_id`) REFERENCES `comments`(`id`) ON DELETE CASCADE,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_song_id` (`song_id`),
    INDEX `idx_playlist_id` (`playlist_id`),
    INDEX `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='评论表';

-- 12. 播放记录表
CREATE TABLE IF NOT EXISTS `play_history` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `song_id` BIGINT UNSIGNED NOT NULL COMMENT '歌曲ID',
    `play_duration` INT UNSIGNED DEFAULT 0 COMMENT '播放时长（秒）',
    `total_duration` INT UNSIGNED DEFAULT 0 COMMENT '歌曲总时长',
    `quality` VARCHAR(20) DEFAULT 'standard' COMMENT '播放音质',
    `source` VARCHAR(50) DEFAULT 'web' COMMENT '播放来源: web, app, mini',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '播放时间',
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`song_id`) REFERENCES `songs`(`id`) ON DELETE CASCADE,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_song_id` (`song_id`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='播放记录表';

-- 13. 关注表
CREATE TABLE IF NOT EXISTS `follows` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    `follower_id` BIGINT UNSIGNED NOT NULL COMMENT '关注者ID',
    `following_id` BIGINT UNSIGNED NOT NULL COMMENT '被关注者ID',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '关注时间',
    FOREIGN KEY (`follower_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`following_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    UNIQUE KEY `uk_follow_pair` (`follower_id`, `following_id`),
    INDEX `idx_follower_id` (`follower_id`),
    INDEX `idx_following_id` (`following_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='关注表';

-- 14. 歌手收藏表
CREATE TABLE IF NOT EXISTS `artist_favorites` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
    `artist_id` BIGINT UNSIGNED NOT NULL COMMENT '歌手ID',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '收藏时间',
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`artist_id`) REFERENCES `artists`(`id`) ON DELETE CASCADE,
    UNIQUE KEY `uk_user_artist` (`user_id`, `artist_id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_artist_id` (`artist_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='歌手收藏表';

-- 15. 操作日志表
CREATE TABLE IF NOT EXISTS `operation_logs` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
    `user_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '操作用户ID',
    `username` VARCHAR(50) DEFAULT NULL COMMENT '用户名',
    `action` VARCHAR(100) NOT NULL COMMENT '操作类型',
    `target_type` VARCHAR(50) DEFAULT NULL COMMENT '目标类型',
    `target_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '目标ID',
    `ip` VARCHAR(45) DEFAULT NULL COMMENT 'IP地址',
    `user_agent` VARCHAR(500) DEFAULT NULL COMMENT 'User Agent',
    `description` TEXT DEFAULT NULL COMMENT '操作描述',
    `status` TINYINT DEFAULT 1 COMMENT '状态: 1成功, 0失败',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_action` (`action`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志表';

-- =========================================
-- 初始化数据
-- =========================================

-- 插入默认角色
INSERT INTO `roles` (`name`, `description`) VALUES
('user', '普通用户'),
('vip', 'VIP会员'),
('admin', '系统管理员');

-- 插入默认权限
INSERT INTO `permissions` (`name`, `description`) VALUES
('user.view', '查看用户信息'),
('user.edit', '编辑用户信息'),
('user.delete', '删除用户'),
('song.add', '添加歌曲'),
('song.edit', '编辑歌曲'),
('song.delete', '删除歌曲'),
('admin.panel', '访问后台'),
('admin.users', '用户管理'),
('admin.system', '系统管理');
