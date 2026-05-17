
package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("demo-secret-key")

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type Song struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	ArtistID   uint   `json:"artist_id"`
	Album      string `json:"album"`
	AlbumID    uint   `json:"album_id"`
	CoverURL   string `json:"cover_url"`
	Duration   int    `json:"duration"`
	PlayURL    string `json:"play_url"`
	PlayCount  int    `json:"play_count"`
	LikeCount  int    `json:"like_count"`
}

type Album struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	ArtistID  uint   `json:"artist_id"`
	CoverURL  string `json:"cover_url"`
	PlayCount int    `json:"play_count"`
	Songs     []Song `json:"songs,omitempty"`
}

type Artist struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	FansCount int    `json:"fans_count"`
}

type Playlist struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CoverURL    string `json:"cover_url"`
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
	PlayCount   int    `json:"play_count"`
	SongCount   int    `json:"song_count"`
	IsPublic    bool   `json:"is_public"`
	Songs       []Song `json:"songs,omitempty"`
}

type Comment struct {
	ID        uint   `json:"id"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	LikeCount int    `json:"like_count"`
}

var demoUsers = []User{
	{ID: 1, Username: "demo", Nickname: "演示用户", Avatar: "https://picsum.photos/200/200?random=1", Email: "demo@example.com"},
	{ID: 2, Username: "musicfan", Nickname: "音乐爱好者", Avatar: "https://picsum.photos/200/200?random=2", Email: "fan@example.com"},
}

var demoArtists = []Artist{
	{ID: 1, Name: "周杰伦", AvatarURL: "https://picsum.photos/200/200?random=10", FansCount: 10000000},
	{ID: 2, Name: "林俊杰", AvatarURL: "https://picsum.photos/200/200?random=11", FansCount: 8000000},
	{ID: 3, Name: "邓紫棋", AvatarURL: "https://picsum.photos/200/200?random=12", FansCount: 6000000},
	{ID: 4, Name: "陈奕迅", AvatarURL: "https://picsum.photos/200/200?random=13", FansCount: 9000000},
}

var demoAlbums = []Album{
	{ID: 1, Title: "范特西", Artist: "周杰伦", ArtistID: 1, CoverURL: "https://picsum.photos/300/300?random=20", PlayCount: 5000000},
	{ID: 2, Title: "曹操", Artist: "林俊杰", ArtistID: 2, CoverURL: "https://picsum.photos/300/300?random=21", PlayCount: 3000000},
}

var demoSongs = []Song{
	{ID: 1, Title: "双截棍", Artist: "周杰伦", ArtistID: 1, Album: "范特西", AlbumID: 1, CoverURL: "https://picsum.photos/300/300?random=20", Duration: 185, PlayCount: 1000000, LikeCount: 50000},
	{ID: 2, Title: "爱在西元前", Artist: "周杰伦", ArtistID: 1, Album: "范特西", AlbumID: 1, CoverURL: "https://picsum.photos/300/300?random=20", Duration: 245, PlayCount: 800000, LikeCount: 40000},
	{ID: 3, Title: "简单爱", Artist: "周杰伦", ArtistID: 1, Album: "范特西", AlbumID: 1, CoverURL: "https://picsum.photos/300/300?random=20", Duration: 270, PlayCount: 1200000, LikeCount: 60000},
	{ID: 4, Title: "曹操", Artist: "林俊杰", ArtistID: 2, Album: "曹操", AlbumID: 2, CoverURL: "https://picsum.photos/300/300?random=21", Duration: 260, PlayCount: 700000, LikeCount: 35000},
	{ID: 5, Title: "江南", Artist: "林俊杰", ArtistID: 2, Album: "曹操", AlbumID: 2, CoverURL: "https://picsum.photos/300/300?random=21", Duration: 240, PlayCount: 900000, LikeCount: 45000},
	{ID: 6, Title: "光年之外", Artist: "邓紫棋", ArtistID: 3, Album: "光年之外", AlbumID: 3, CoverURL: "https://picsum.photos/300/300?random=22", Duration: 235, PlayCount: 600000, LikeCount: 30000},
	{ID: 7, Title: "十年", Artist: "陈奕迅", ArtistID: 4, Album: "黑白灰", AlbumID: 4, CoverURL: "https://picsum.photos/300/300?random=23", Duration: 205, PlayCount: 1500000, LikeCount: 75000},
	{ID: 8, Title: "浮夸", Artist: "陈奕迅", ArtistID: 4, Album: "认了吧", AlbumID: 5, CoverURL: "https://picsum.photos/300/300?random=24", Duration: 280, PlayCount: 850000, LikeCount: 42000},
}

var demoPlaylists = []Playlist{
	{ID: 1, Title: "华语流行精选", Description: "精选华语流行歌曲", CoverURL: "https://picsum.photos/300/300?random=30", UserID: 1, Username: "演示用户", PlayCount: 10000, SongCount: 50, IsPublic: true},
	{ID: 2, Title: "经典老歌回忆", Description: "经典老歌回忆杀", CoverURL: "https://picsum.photos/300/300?random=31", UserID: 2, Username: "音乐爱好者", PlayCount: 8000, SongCount: 40, IsPublic: true},
	{ID: 3, Title: "我的私人歌单", Description: "我的专属音乐", CoverURL: "https://picsum.photos/300/300?random=32", UserID: 1, Username: "演示用户", PlayCount: 5000, SongCount: 30, IsPublic: false},
}

var demoComments = []Comment{
	{ID: 1, Content: "这首歌太好听了！", UserID: 1, Username: "演示用户", AvatarURL: "https://picsum.photos/200/200?random=1", CreatedAt: "2024-01-15 10:30:00", LikeCount: 120},
	{ID: 2, Content: "百听不厌的经典", UserID: 2, Username: "音乐爱好者", AvatarURL: "https://picsum.photos/200/200?random=2", CreatedAt: "2024-01-14 15:20:00", LikeCount: 85},
}

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	api := r.Group("/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", register)
			auth.POST("/login", login)
			auth.POST("/refresh", refreshToken)
			auth.POST("/logout", logout)
			auth.POST("/send-code", sendCode)
		}

		user := api.Group("/user")
		{
			user.GET("/me", requireAuth(), getCurrentUser)
			user.PUT("/profile", requireAuth(), updateProfile)
			user.POST("/avatar", requireAuth(), uploadAvatar)
			user.GET("/:user_id", getUserProfile)
			user.POST("/:user_id/follow", requireAuth(), followUser)
			user.DELETE("/:user_id/follow", requireAuth(), unfollowUser)
			user.GET("/:user_id/following", getFollowing)
			user.GET("/:user_id/followers", getFollowers)
			user.GET("/me/playlists", requireAuth(), getMyPlaylists)
		}

		song := api.Group("/song")
		{
			song.GET("/:song_id", getSong)
			song.GET("/:song_id/lyric", getLyric)
			song.GET("/:song_id/play-url", getPlayURL)
			song.GET("/hot", getHotSongs)
			song.GET("/new", getNewSongs)
		}

		album := api.Group("/album")
		{
			album.GET("/:album_id", getAlbum)
		}

		artist := api.Group("/artist")
		{
			artist.GET("/:artist_id", getArtist)
		}

		genres := api.Group("/genres")
		{
			genres.GET("", getGenres)
		}

		playlist := api.Group("/playlist")
		{
			playlist.POST("", requireAuth(), createPlaylist)
			playlist.GET("/:playlist_id", getPlaylist)
			playlist.PUT("/:playlist_id", requireAuth(), updatePlaylist)
			playlist.DELETE("/:playlist_id", requireAuth(), deletePlaylist)
			playlist.POST("/:playlist_id/songs", requireAuth(), addSongsToPlaylist)
			playlist.DELETE("/:playlist_id/songs", requireAuth(), removeSongsFromPlaylist)
			playlist.PUT("/:playlist_id/songs/sort", requireAuth(), sortPlaylistSongs)
			playlist.POST("/:playlist_id/favorite", requireAuth(), favoritePlaylist)
			playlist.DELETE("/:playlist_id/favorite", requireAuth(), unfavoritePlaylist)
			playlist.GET("/recommended", getRecommendedPlaylists)
		}

		search := api.Group("/search")
		{
			search.GET("", search)
			search.GET("/suggest", searchSuggest)
			search.GET("/hot", getHotKeywords)
			search.GET("/history", requireAuth(), getSearchHistory)
			search.DELETE("/history", requireAuth(), clearSearchHistory)
		}

		play := api.Group("/play")
		{
			play.POST("/report", requireAuth(), reportPlay)
			play.GET("/history", requireAuth(), getPlayHistory)
		}

		rank := api.Group("/rank")
		{
			rank.GET("/:type", getRank)
		}

		comment := api.Group("/comment")
		{
			comment.GET("", getComments)
			comment.POST("", requireAuth(), createComment)
			comment.DELETE("/:comment_id", requireAuth(), deleteComment)
			comment.POST("/:comment_id/like", requireAuth(), likeComment)
			comment.DELETE("/:comment_id/like", requireAuth(), unlikeComment)
		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("演示 API 服务器启动在端口 8080")
	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未授权"})
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]
		claims := &amp;Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "无效的 token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

func register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "注册成功",
		"data": gin.H{
			"access_token":  generateToken(1, "demo"),
			"refresh_token": generateRefreshToken(1, "demo"),
			"user":          demoUsers[0],
		},
	})
}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"data": gin.H{
			"access_token":  generateToken(1, "demo"),
			"refresh_token": generateRefreshToken(1, "demo"),
			"user":          demoUsers[0],
		},
	})
}

func refreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "刷新成功",
		"data": gin.H{
			"access_token":  generateToken(1, "demo"),
			"refresh_token": generateRefreshToken(1, "demo"),
		},
	})
}

func logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登出成功"})
}

func sendCode(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "验证码已发送"})
}

func getCurrentUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": demoUsers[0]})
}

func updateProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新成功", "data": demoUsers[0]})
}

func uploadAvatar(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "上传成功", "data": gin.H{"avatar_url": "https://picsum.photos/200/200?random=new"}})
}

func getUserProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": demoUsers[0]})
}

func followUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "关注成功"})
}

func unfollowUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "取消关注成功"})
}

func getFollowing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": gin.H{"list": []User{}, "total": 0}})
}

func getFollowers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": gin.H{"list": []User{}, "total": 0}})
}

func getMyPlaylists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": gin.H{"list": demoPlaylists, "total": len(demoPlaylists)}})
}

func getSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("song_id"))
	for _, song := range demoSongs {
		if song.ID == uint(id) {
			song.PlayURL = "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3"
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": song})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "歌曲不存在"})
}

func getLyric(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"data": gin.H{
			"lyric": "[00:00.00]暂无歌词\n[00:05.00]请欣赏音乐",
		},
	})
}

func getPlayURL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"data": gin.H{
			"play_url": "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3",
		},
	})
}

func getHotSongs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": gin.H{"list": demoSongs, "total": len(demoSongs)}})
}

func getNewSongs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": gin.H{"list": demoSongs, "total": len(demoSongs)}})
}

func getAlbum(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("album_id"))
	for _, album := range demoAlbums {
		if album.ID == uint(id) {
			var songs []Song
			for _, song := range demoSongs {
				if song.AlbumID == album.ID {
					songs = append(songs, song)
				}
			}
			album.Songs = songs
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": album})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "专辑不存在"})
}

func getArtist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("artist_id"))
	for _, artist := range demoArtists {
		if artist.ID == uint(id) {
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": artist})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "歌手不存在"})
}

func getGenres(c *gin.Context) {
	genres := []map[string]interface{}{
		{"id": 1, "name": "流行", "cover_url": "https://picsum.photos/200/200?random=g1"},
		{"id": 2, "name": "摇滚", "cover_url": "https://picsum.photos/200/200?random=g2"},
		{"id": 3, "name": "民谣", "cover_url": "https://picsum.photos/200/200?random=g3"},
		{"id": 4, "name": "电子", "cover_url": "https://picsum.photos/200/200?random=g4"},
		{"id": 5, "name": "古典", "cover_url": "https://picsum.photos/200/200?random=g5"},
		{"id": 6, "name": "爵士", "cover_url": "https://picsum.photos/200/200?random=g6"},
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": gin.H{"list": genres, "total": len(genres)}})
}

func createPlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "创建成功", "data": demoPlaylists[0]})
}

func getPlaylist(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("playlist_id"))
	for _, playlist := range demoPlaylists {
		if playlist.ID == uint(id) {
			playlist.Songs = demoSongs
			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": playlist})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": 404, "message": "歌单不存在"})
}

func updatePlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "更新成功"})
}

func deletePlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
}

func addSongsToPlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "添加成功"})
}

func removeSongsFromPlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "移除成功"})
}

func sortPlaylistSongs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "排序成功"})
}

func favoritePlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "收藏成功"})
}

func unfavoritePlaylist(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "取消收藏成功"})
}

func getRecommendedPlaylists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": gin.H{"list": demoPlaylists, "total": len(demoPlaylists)}})
}

func search(c *gin.Context) {
	keyword := c.Query("keyword")
	var songs []Song
	for _, song := range demoSongs {
		if len(keyword) == 0 || contains(song.Title, keyword) || contains(song.Artist, keyword) {
			songs = append(songs, song)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"songs":   gin.H{"list": songs, "total": len(songs)},
			"artists": gin.H{"list": demoArtists, "total": len(demoArtists)},
			"albums":  gin.H{"list": demoAlbums, "total": len(demoAlbums)},
		},
	})
}

func searchSuggest(c *gin.Context) {
	suggestions := []string{"周杰伦", "林俊杰", "邓紫棋", "陈奕迅", "流行音乐"}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": suggestions})
}

func getHotKeywords(c *gin.Context) {
	keywords := []map[string]interface{}{
		{"keyword": "周杰伦", "score": 100000},
		{"keyword": "林俊杰", "score": 80000},
		{"keyword": "邓紫棋", "score": 60000},
		{"keyword": "陈奕迅", "score": 90000},
		{"keyword": "流行音乐", "score": 70000},
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": keywords})
}

func getSearchHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": []string{}})
}

func clearSearchHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "清除成功"})
}

func reportPlay(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "上报成功"})
}

func getPlayHistory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": gin.H{"list": demoSongs[:3], "total": 3}})
}

func getRank(c *gin.Context) {
	rankType := c.Param("type")
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"type":  rankType,
			"list":  demoSongs,
			"total": len(demoSongs),
		},
	})
}

func getComments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": gin.H{"list": demoComments, "total": len(demoComments)}})
}

func createComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "发表成功", "data": demoComments[0]})
}

func deleteComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "删除成功"})
}

func likeComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "点赞成功"})
}

func unlikeComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "取消点赞成功"})
}

func generateToken(userID uint, username string) string {
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &amp;Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

func generateRefreshToken(userID uint, username string) string {
	expirationTime := time.Now().Add(30 * 24 * time.Hour)
	claims := &amp;Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtSecret)
	return tokenString
}

func contains(s, substr string) bool {
	return len(s) &gt;= len(substr) &amp;&amp; (s == substr || len(s) &gt; len(substr) &amp;&amp; (s[:len(substr)] == substr || containsHelper(s, substr)))
}

func containsHelper(s, substr string) bool {
	for i := 1; i &lt;= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

