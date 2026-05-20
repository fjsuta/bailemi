# 百米乐音乐平台功能开发文档

## 开发时间
2026年5月19日

## 已完成的后端API

### 1. 排行榜API

**新增Handler:** `chart_handler.go`
- `GET /v1/charts` - 获取排行榜（支持type参数：hot/new/rising）
- `GET /v1/charts/hot` - 获取热歌榜
- `GET /v1/charts/new` - 获取新歌榜
- `GET /v1/charts/rising` - 获取飙升榜

**新增Service方法:** `music_service.go`
- `GetChartSongs(ctx, chartType, limit, userID)` - 获取排行榜歌曲

**新增Repository方法:** `music_repo.go`
- `GetRisingSongs(ctx, limit)` - 获取飙升歌曲（按播放量排序）

### 2. 歌手API

**新增Handler:** `artist_handler.go`
- `GET /v1/artist` - 获取歌手列表（支持分页、排序、字母筛选）
- `GET /v1/artist/:artist_id` - 获取歌手详情
- `GET /v1/artist/:artist_id/songs` - 获取歌手歌曲列表
- `GET /v1/artist/:artist_id/albums` - 获取歌手专辑列表

**新增Service方法:** `music_service.go`
- `ListArtists(ctx, page, pageSize, sort, letter, userID)` - 获取歌手列表

**新增Repository方法:** `music_repo.go`
- `ListArtists(ctx, page, pageSize, sort, letter)` - 数据库查询歌手列表

### 3. 歌单API

**新增Handler方法:** `playlist_handler.go`
- `ListPlaylists(c)` - 获取歌单列表（支持标签筛选、排序）

**新增Service方法:** `social_service.go`
- `ListPlaylists(ctx, page, pageSize, tag, sort, userID)` - 获取歌单列表

**新增Repository方法:** `social_repo.go`
- `ListPlaylists(ctx, page, pageSize, tag, sort)` - 数据库查询歌单列表

## 路由注册

已在 `cmd/server/main.go` 中注册以下路由组：

```go
// 排行榜路由
charts := api.Group("/charts")
{
    charts.GET("", authMiddleware.OptionalAuth(), chartHandler.GetCharts)
    charts.GET("/hot", authMiddleware.OptionalAuth(), chartHandler.GetHotSongs)
    charts.GET("/new", authMiddleware.OptionalAuth(), chartHandler.GetNewSongs)
    charts.GET("/rising", authMiddleware.OptionalAuth(), chartHandler.GetRisingSongs)
}

// 歌手路由
artist := api.Group("/artist")
{
    artist.GET("", authMiddleware.OptionalAuth(), artistHandler.ListArtists)
    artist.GET("/:artist_id", authMiddleware.OptionalAuth(), artistHandler.GetArtist)
    artist.GET("/:artist_id/songs", artistHandler.GetArtistSongs)
    artist.GET("/:artist_id/albums", artistHandler.GetArtistAlbums)
}

// 歌单路由
playlist := api.Group("/playlist")
{
    playlist.GET("", authMiddleware.OptionalAuth(), playlistHandler.ListPlaylists)
    // ... 其他路由
}
```

## API参数说明

### 排行榜API

**请求示例:**
```bash
GET /api/v1/charts?type=hot&limit=20
GET /api/v1/charts?type=new&limit=20
GET /api/v1/charts?type=rising&limit=20
```

**响应示例:**
```json
{
  "code": 0,
  "data": {
    "type": "hot",
    "songs": [
      {
        "id": 1,
        "title": "歌曲名",
        "artist": {
          "id": 1,
          "name": "歌手名"
        },
        "play_count": 100000,
        "favorite_count": 5000,
        "comment_count": 100
      }
    ]
  }
}
```

### 歌手API

**请求示例:**
```bash
GET /api/v1/artist?page=1&pageSize=20&sort=hot&letter=A
GET /api/v1/artist?page=1&pageSize=20&sort=new
GET /api/v1/artist/1
```

**响应示例:**
```json
{
  "code": 0,
  "data": {
    "list": [...],
    "total": 100,
    "page": 1,
    "pageSize": 20
  }
}
```

### 歌单API

**请求示例:**
```bash
GET /api/v1/playlist?page=1&pageSize=20&tag=study&sort=hot
GET /api/v1/playlist?page=1&pageSize=20&tag=rock&sort=new
```

## 前端组件（待开发）

### 排行榜页面
- Tab切换：热歌榜/新歌榜/飙升榜
- 列表展示：排名、封面、歌名、歌手、播放量
- 前三名特殊样式：金银铜色
- 点击播放功能
- 加入播放队列功能

### 歌手页面
- 歌手列表：网格布局展示
- 字母索引条（A-Z）
- 歌手详情页：大头像、关注按钮、粉丝数
- 热门歌曲列表
- 专辑封面墙

### 歌单页面
- 标签筛选栏（华语、欧美、摇滚、民谣等）
- 瀑布流布局展示歌单
- 歌单详情页
- 歌曲列表
- 播放全部、收藏歌单功能

## 数据库索引优化建议

为了提升查询性能，建议在数据库中添加以下索引：

```sql
-- songs表
CREATE INDEX idx_songs_play_count ON songs(play_count DESC);
CREATE INDEX idx_songs_created_at ON songs(created_at DESC);
CREATE INDEX idx_songs_status ON songs(status);

-- artists表
CREATE INDEX idx_artists_fan_count ON artists(fan_count DESC);
CREATE INDEX idx_artists_name ON artists(name);

-- playlists表
CREATE INDEX idx_playlists_play_count ON playlists(play_count DESC);
CREATE INDEX idx_playlists_created_at ON playlists(created_at DESC);
```

## 第三方登录配置（待开发）

需要添加系统配置表，支持动态配置第三方登录方式：

```sql
CREATE TABLE oauth_config (
    id INT PRIMARY KEY AUTO_INCREMENT,
    provider VARCHAR(50) NOT NULL COMMENT '提供商：wechat, github, google, qq, weibo等',
    app_id VARCHAR(200) NOT NULL COMMENT '应用ID',
    app_secret VARCHAR(500) NOT NULL COMMENT '应用密钥',
    is_enabled TINYINT DEFAULT 0 COMMENT '是否启用：0-否，1-是',
    display_name VARCHAR(100) COMMENT '显示名称',
    icon_url VARCHAR(500) COMMENT '图标URL',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## 注意事项

1. 排行榜支持Redis缓存（可选实现，每小时更新一次）
2. 所有API都支持用户认证（OptionalAuth）
3. 排序参数支持：hot(热门)、new(最新)、name(名称)、rising(飙升)
4. 分页参数：page和pageSize，默认20条/页
5. 字母筛选仅对歌手列表有效
