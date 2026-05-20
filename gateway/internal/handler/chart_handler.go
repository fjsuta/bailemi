package handler

import (
	"strconv"

	"github.com/bailemi/gateway/internal/service"
	"github.com/bailemi/gateway/pkg/errors"
	"github.com/bailemi/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

type ChartHandler struct {
	musicService *service.MusicService
}

func NewChartHandler(musicService *service.MusicService) *ChartHandler {
	return &ChartHandler{musicService: musicService}
}

func (h *ChartHandler) GetCharts(c *gin.Context) {
	chartType := c.DefaultQuery("type", "hot")
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	songs, err := h.musicService.GetChartSongs(c.Request.Context(), chartType, limit, userID)
	if err != nil {
		errors.Error(c, 500, errors.ErrInternalError.Code, "获取榜单失败")
		return
	}

	response.Success(c, gin.H{
		"type": chartType,
		"songs": songs,
	})
}

func (h *ChartHandler) GetHotSongs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	songs, err := h.musicService.GetChartSongs(c.Request.Context(), "hot", limit, userID)
	if err != nil {
		errors.Error(c, 500, errors.ErrInternalError.Code, "获取热门歌曲失败")
		return
	}

	response.Success(c, gin.H{
		"type": "hot",
		"songs": songs,
	})
}

func (h *ChartHandler) GetNewSongs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	songs, err := h.musicService.GetChartSongs(c.Request.Context(), "new", limit, userID)
	if err != nil {
		errors.Error(c, 500, errors.ErrInternalError.Code, "获取新歌失败")
		return
	}

	response.Success(c, gin.H{
		"type": "new",
		"songs": songs,
	})
}

func (h *ChartHandler) GetRisingSongs(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	songs, err := h.musicService.GetChartSongs(c.Request.Context(), "rising", limit, userID)
	if err != nil {
		errors.Error(c, 500, errors.ErrInternalError.Code, "获取飙升榜失败")
		return
	}

	response.Success(c, gin.H{
		"type": "rising",
		"songs": songs,
	})
}
