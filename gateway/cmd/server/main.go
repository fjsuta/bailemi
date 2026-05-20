package main

import (
	"fmt"
	"log"

	"github.com/bailemi/gateway/configs"
	"github.com/bailemi/gateway/internal/handler"
	"github.com/bailemi/gateway/internal/middleware"
	"github.com/bailemi/gateway/internal/repository"
	"github.com/bailemi/gateway/internal/service"
	"github.com/bailemi/gateway/pkg/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := configs.LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	userRepo := repository.NewUserRepository(db)
	musicRepo := repository.NewMusicRepository(db)
	socialRepo := repository.NewSocialRepository(db)
	playRepo := repository.NewPlayRepository(db)

	jwtManager := jwt.NewJWTManager(
		cfg.JWT.Secret,
		cfg.JWT.AccessTokenTTL,
		cfg.JWT.RefreshTokenTTL,
	)

	authService := service.NewAuthService(userRepo, jwtManager)
	userService := service.NewUserService(userRepo)
	musicService := service.NewMusicService(musicRepo, socialRepo)
	socialService := service.NewSocialService(socialRepo, musicRepo)
	playService := service.NewPlayService(playRepo, musicRepo)

	authHandler := handler.NewAuthHandler(authService, userService)
	userHandler := handler.NewUserHandler(userService)
	musicHandler := handler.NewMusicHandler(musicService)
	artistHandler := handler.NewArtistHandler(musicService)
	playlistHandler := handler.NewPlaylistHandler(socialService)
	playHandler := handler.NewPlayHandler(playService, musicService)
	searchHandler := handler.NewSearchHandler(musicService, socialService)
	commentHandler := handler.NewCommentHandler(socialService)
	chartHandler := handler.NewChartHandler(musicService)

	authMiddleware := middleware.NewAuthMiddleware(jwtManager)

	r := gin.Default()

	r.Use(corsMiddleware())
	r.Use(requestLogger())

	api := r.Group("/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
			auth.POST("/logout", authMiddleware.RequireAuth(), authHandler.Logout)
			auth.POST("/send-code", authHandler.SendCode)
		}

		user := api.Group("/user")
		{
			user.GET("/me", authMiddleware.RequireAuth(), userHandler.GetCurrentUser)
			user.PUT("/profile", authMiddleware.RequireAuth(), userHandler.UpdateProfile)
			user.POST("/avatar", authMiddleware.RequireAuth(), userHandler.UploadAvatar)
			user.GET("/:user_id", userHandler.GetUserProfile)
			user.POST("/:user_id/follow", authMiddleware.RequireAuth(), userHandler.FollowUser)
			user.DELETE("/:user_id/follow", authMiddleware.RequireAuth(), userHandler.UnfollowUser)
			user.GET("/:user_id/following", userHandler.GetFollowing)
			user.GET("/:user_id/followers", userHandler.GetFollowers)
			user.GET("/me/playlists", authMiddleware.RequireAuth(), playlistHandler.GetMyPlaylists)
		}

		song := api.Group("/song")
		{
			song.GET("/:song_id", authMiddleware.OptionalAuth(), musicHandler.GetSong)
			song.GET("/:song_id/lyric", musicHandler.GetLyric)
			song.GET("/:song_id/play-url", authMiddleware.OptionalAuth(), playHandler.GetPlayURL)
			song.GET("/hot", musicHandler.GetHotSongs)
			song.GET("/new", musicHandler.GetNewSongs)
		}

		album := api.Group("/album")
		{
			album.GET("/:album_id", authMiddleware.OptionalAuth(), musicHandler.GetAlbum)
		}

		artist := api.Group("/artist")
		{
			artist.GET("", authMiddleware.OptionalAuth(), artistHandler.ListArtists)
			artist.GET("/:artist_id", authMiddleware.OptionalAuth(), artistHandler.GetArtist)
			artist.GET("/:artist_id/songs", artistHandler.GetArtistSongs)
			artist.GET("/:artist_id/albums", artistHandler.GetArtistAlbums)
		}

		genres := api.Group("/genres")
		{
			genres.GET("", musicHandler.GetGenres)
		}

		playlist := api.Group("/playlist")
		{
			playlist.GET("", authMiddleware.OptionalAuth(), playlistHandler.ListPlaylists)
			playlist.POST("", authMiddleware.RequireAuth(), playlistHandler.CreatePlaylist)
			playlist.GET("/:playlist_id", authMiddleware.OptionalAuth(), playlistHandler.GetPlaylist)
			playlist.PUT("/:playlist_id", authMiddleware.RequireAuth(), playlistHandler.UpdatePlaylist)
			playlist.DELETE("/:playlist_id", authMiddleware.RequireAuth(), playlistHandler.DeletePlaylist)
			playlist.POST("/:playlist_id/songs", authMiddleware.RequireAuth(), playlistHandler.AddSongs)
			playlist.DELETE("/:playlist_id/songs", authMiddleware.RequireAuth(), playlistHandler.RemoveSongs)
			playlist.PUT("/:playlist_id/songs/sort", authMiddleware.RequireAuth(), playlistHandler.SortSongs)
			playlist.POST("/:playlist_id/favorite", authMiddleware.RequireAuth(), playlistHandler.FavoritePlaylist)
			playlist.DELETE("/:playlist_id/favorite", authMiddleware.RequireAuth(), playlistHandler.UnfavoritePlaylist)
			playlist.GET("/recommended", playlistHandler.GetRecommendedPlaylists)
		}

		search := api.Group("/search")
		{
			search.GET("", authMiddleware.OptionalAuth(), searchHandler.Search)
			search.GET("/suggest", searchHandler.Suggest)
			search.GET("/hot", searchHandler.GetHotKeywords)
			search.GET("/history", authMiddleware.RequireAuth(), searchHandler.GetSearchHistory)
			search.DELETE("/history", authMiddleware.RequireAuth(), searchHandler.ClearSearchHistory)
		}

		play := api.Group("/play")
		{
			play.POST("/report", authMiddleware.RequireAuth(), playHandler.ReportPlay)
			play.GET("/history", authMiddleware.RequireAuth(), playHandler.GetPlayHistory)
		}

		charts := api.Group("/charts")
		{
			charts.GET("", authMiddleware.OptionalAuth(), chartHandler.GetCharts)
			charts.GET("/hot", authMiddleware.OptionalAuth(), chartHandler.GetHotSongs)
			charts.GET("/new", authMiddleware.OptionalAuth(), chartHandler.GetNewSongs)
			charts.GET("/rising", authMiddleware.OptionalAuth(), chartHandler.GetRisingSongs)
		}

		rank := api.Group("/rank")
		{
			rank.GET("/:type", playHandler.GetRank)
		}

		comment := api.Group("/comment")
		{
			comment.GET("", commentHandler.GetComments)
			comment.POST("", authMiddleware.RequireAuth(), commentHandler.CreateComment)
			comment.DELETE("/:comment_id", authMiddleware.RequireAuth(), commentHandler.DeleteComment)
			comment.POST("/:comment_id/like", authMiddleware.RequireAuth(), commentHandler.LikeComment)
			comment.DELETE("/:comment_id/like", authMiddleware.RequireAuth(), commentHandler.UnlikeComment)
		}

		admin := api.Group("/admin")
		admin.Use(authMiddleware.RequireAuth(), authMiddleware.RequireRole(9))
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "dashboard endpoint"})
			})
			admin.GET("/users", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "users endpoint"})
			})
		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("服务器启动在端口 %s\n", cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
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

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
