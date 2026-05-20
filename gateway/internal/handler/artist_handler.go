package handler

import (
	"strconv"

	"github.com/bailemi/gateway/internal/service"
	"github.com/bailemi/gateway/pkg/errors"
	"github.com/bailemi/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

type ArtistHandler struct {
	musicService *service.MusicService
}

func NewArtistHandler(musicService *service.MusicService) *ArtistHandler {
	return &ArtistHandler{musicService: musicService}
}

func (h *ArtistHandler) GetArtist(c *gin.Context) {
	artistID, err := strconv.ParseUint(c.Param("artist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的音乐人ID")
		return
	}

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	artist, err := h.musicService.GetArtistByID(c.Request.Context(), artistID, userID)
	if err != nil {
		errors.Error(c, 404, errors.ErrArtistNotFound.Code, "音乐人不存在")
		return
	}

	response.Success(c, artist)
}

func (h *ArtistHandler) ListArtists(c *gin.Context) {
	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "pageSize", 20)
	sort := c.DefaultQuery("sort", "hot")
	letter := c.Query("letter")

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	artists, total, err := h.musicService.ListArtists(c.Request.Context(), page, pageSize, sort, letter, userID)
	if err != nil {
		errors.Error(c, 500, errors.ErrInternalError.Code, "获取音乐人列表失败")
		return
	}

	response.Success(c, gin.H{
		"list": artists,
		"total": total,
		"page": page,
		"pageSize": pageSize,
	})
}

func (h *ArtistHandler) GetArtistSongs(c *gin.Context) {
	artistID, err := strconv.ParseUint(c.Param("artist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的音乐人ID")
		return
	}

	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "pageSize", 20)

	songs, total, err := h.musicService.GetArtistSongs(c.Request.Context(), artistID, page, pageSize, "")
	if err != nil {
		errors.Error(c, 500, errors.ErrInternalError.Code, "获取音乐人歌曲失败")
		return
	}

	response.Success(c, gin.H{
		"list": songs,
		"total": total,
		"page": page,
		"pageSize": pageSize,
	})
}

func (h *ArtistHandler) GetArtistAlbums(c *gin.Context) {
	artistID, err := strconv.ParseUint(c.Param("artist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的音乐人ID")
		return
	}

	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "pageSize", 20)

	albums, total, err := h.musicService.GetArtistAlbums(c.Request.Context(), artistID, page, pageSize)
	if err != nil {
		errors.Error(c, 500, errors.ErrInternalError.Code, "获取音乐人专辑失败")
		return
	}

	response.Success(c, gin.H{
		"list": albums,
		"total": total,
		"page": page,
		"pageSize": pageSize,
	})
}

func getIntQuery(c *gin.Context, key string, defaultVal int) int {
	val, err := strconv.Atoi(c.Query(key))
	if err != nil || val < 1 {
		return defaultVal
	}
	if val > 100 {
		return 100
	}
	return val
}
