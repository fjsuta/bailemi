package handler

import (
	"strconv"

	"github.com/bailemi/gateway/internal/model"
	"github.com/bailemi/gateway/internal/service"
	"github.com/bailemi/gateway/pkg/errors"
	"github.com/bailemi/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

type PlaylistHandler struct {
	socialService *service.SocialService
}

func NewPlaylistHandler(socialService *service.SocialService) *PlaylistHandler {
	return &PlaylistHandler{socialService: socialService}
}

func (h *PlaylistHandler) CreatePlaylist(c *gin.Context) {
	userID := c.GetUint64("user_id")

	var req model.CreatePlaylistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	playlist, err := h.socialService.CreatePlaylist(c.Request.Context(), userID, &req)
	if err != nil {
		errors.Error(c, 500, 10000, "创建失败")
		return
	}

	response.Success(c, playlist)
}

func (h *PlaylistHandler) GetPlaylist(c *gin.Context) {
	playlistID, err := strconv.ParseUint(c.Param("playlist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌单ID")
		return
	}

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	playlist, err := h.socialService.GetPlaylist(c.Request.Context(), playlistID, userID)
	if err != nil {
		errors.Error(c, 404, 50001, "歌单不存在")
		return
	}

	response.Success(c, playlist)
}

func (h *PlaylistHandler) UpdatePlaylist(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	playlistID, err := strconv.ParseUint(c.Param("playlist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌单ID")
		return
	}

	var req model.UpdatePlaylistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	err = h.socialService.UpdatePlaylist(c.Request.Context(), playlistID, userID, &req)
	if err != nil {
		errors.Error(c, 403, 50002, "不能修改他人的歌单")
		return
	}

	response.Success(c, gin.H{"message": "更新成功"})
}

func (h *PlaylistHandler) DeletePlaylist(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	playlistID, err := strconv.ParseUint(c.Param("playlist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌单ID")
		return
	}

	err = h.socialService.DeletePlaylist(c.Request.Context(), playlistID, userID)
	if err != nil {
		errors.Error(c, 403, 50002, "不能删除他人的歌单")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

func (h *PlaylistHandler) GetMyPlaylists(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "page_size", 20)

	playlists, total, err := h.socialService.GetUserPlaylists(c.Request.Context(), userID, page, pageSize)
	if err != nil {
		errors.Error(c, 500, 10000, "获取失败")
		return
	}

	response.SuccessWithPagination(c, playlists, page, pageSize, total)
}

func (h *PlaylistHandler) AddSongs(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	playlistID, err := strconv.ParseUint(c.Param("playlist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌单ID")
		return
	}

	var req model.PlaylistSongsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	err = h.socialService.AddSongsToPlaylist(c.Request.Context(), playlistID, userID, req.SongIDs)
	if err != nil {
		errors.Error(c, 403, 50002, "不能修改他人的歌单")
		return
	}

	response.Success(c, gin.H{"message": "添加成功"})
}

func (h *PlaylistHandler) RemoveSongs(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	playlistID, err := strconv.ParseUint(c.Param("playlist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌单ID")
		return
	}

	var req model.PlaylistSongsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	err = h.socialService.RemoveSongsFromPlaylist(c.Request.Context(), playlistID, userID, req.SongIDs)
	if err != nil {
		errors.Error(c, 403, 50002, "不能修改他人的歌单")
		return
	}

	response.Success(c, gin.H{"message": "移除成功"})
}

func (h *PlaylistHandler) SortSongs(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	playlistID, err := strconv.ParseUint(c.Param("playlist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌单ID")
		return
	}

	var req struct {
		SongIDs []uint64 `json:"song_ids" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	err = h.socialService.SortPlaylistSongs(c.Request.Context(), playlistID, userID, req.SongIDs)
	if err != nil {
		errors.Error(c, 403, 50002, "不能修改他人的歌单")
		return
	}

	response.Success(c, gin.H{"message": "排序成功"})
}

func (h *PlaylistHandler) FavoritePlaylist(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	playlistID, err := strconv.ParseUint(c.Param("playlist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌单ID")
		return
	}

	err = h.socialService.FavoritePlaylist(c.Request.Context(), userID, playlistID)
	if err != nil {
		errors.Error(c, 500, 10000, "收藏失败")
		return
	}

	response.Success(c, gin.H{"message": "收藏成功"})
}

func (h *PlaylistHandler) UnfavoritePlaylist(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	playlistID, err := strconv.ParseUint(c.Param("playlist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌单ID")
		return
	}

	err = h.socialService.UnfavoritePlaylist(c.Request.Context(), userID, playlistID)
	if err != nil {
		errors.Error(c, 500, 10000, "取消收藏失败")
		return
	}

	response.Success(c, gin.H{"message": "取消收藏成功"})
}

func (h *PlaylistHandler) GetRecommendedPlaylists(c *gin.Context) {
	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "page_size", 20)

	playlists, total, err := h.socialService.GetRecommendedPlaylists(c.Request.Context(), page, pageSize)
	if err != nil {
		errors.Error(c, 500, 10000, "获取失败")
		return
	}

	response.SuccessWithPagination(c, playlists, page, pageSize, total)
}

func (h *PlaylistHandler) ListPlaylists(c *gin.Context) {
	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "page_size", 20)
	tag := c.Query("tag")
	sort := c.DefaultQuery("sort", "hot")

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	playlists, total, err := h.socialService.ListPlaylists(c.Request.Context(), page, pageSize, tag, sort, userID)
	if err != nil {
		errors.Error(c, 500, 10000, "获取失败")
		return
	}

	response.SuccessWithPagination(c, playlists, page, pageSize, total)
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
