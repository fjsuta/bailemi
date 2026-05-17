package handler

import (
	"strconv"

	"github.com/bailemi/gateway/internal/model"
	"github.com/bailemi/gateway/internal/service"
	"github.com/bailemi/gateway/pkg/errors"
	"github.com/bailemi/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

type PlayHandler struct {
	playService  *service.PlayService
	musicService *service.MusicService
}

func NewPlayHandler(playService *service.PlayService, musicService *service.MusicService) *PlayHandler {
	return &PlayHandler{
		playService:  playService,
		musicService: musicService,
	}
}

func (h *PlayHandler) GetPlayURL(c *gin.Context) {
	songID, err := strconv.ParseUint(c.Param("song_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌曲ID")
		return
	}

	quality := c.DefaultQuery("quality", "standard")

	playURL, err := h.playService.GetPlayURL(c.Request.Context(), songID, quality)
	if err != nil {
		errors.Error(c, 404, 40001, "歌曲不存在")
		return
	}

	response.Success(c, playURL)
}

func (h *PlayHandler) ReportPlay(c *gin.Context) {
	userID := c.GetUint64("user_id")

	var req model.PlayReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	err := h.playService.ReportPlay(c.Request.Context(), userID, &req, c.ClientIP())
	if err != nil {
		errors.Error(c, 500, 10000, "上报失败")
		return
	}

	response.Success(c, gin.H{"message": "上报成功"})
}

func (h *PlayHandler) GetPlayHistory(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "page_size", 20)

	histories, total, err := h.playService.GetPlayHistory(c.Request.Context(), userID, page, pageSize)
	if err != nil {
		errors.Error(c, 500, 10000, "获取失败")
		return
	}

	response.SuccessWithPagination(c, histories, page, pageSize, total)
}

func (h *PlayHandler) GetRank(c *gin.Context) {
	rankType := c.Param("type")
	date := c.DefaultQuery("date", "")

	rank, err := h.playService.GetRank(c.Request.Context(), rankType, date, 50)
	if err != nil {
		errors.Error(c, 500, 10000, "获取失败")
		return
	}

	response.Success(c, rank)
}
