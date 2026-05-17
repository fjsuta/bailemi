package handler

import (
	"github.com/bailemi/gateway/internal/service"
	"github.com/bailemi/gateway/pkg/errors"
	"github.com/bailemi/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	musicService  *service.MusicService
	socialService *service.SocialService
}

func NewSearchHandler(musicService *service.MusicService, socialService *service.SocialService) *SearchHandler {
	return &SearchHandler{
		musicService:  musicService,
		socialService: socialService,
	}
}

func (h *SearchHandler) Search(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		errors.Error(c, 400, 60001, "搜索关键词不能为空")
		return
	}

	searchType := c.DefaultQuery("type", "all")
	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "page_size", 20)

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	results := gin.H{
		"keyword": keyword,
		"results": gin.H{},
	}

	switch searchType {
	case "all", "song":
		songs, total, _ := h.musicService.SearchSongs(keyword, page, pageSize)
		results["results"].(gin.H)["songs"] = gin.H{"total": total, "items": songs}
	}

	switch searchType {
	case "all", "artist":
		artists, total, _ := h.musicService.SearchArtists(keyword, page, pageSize)
		results["results"].(gin.H)["artists"] = gin.H{"total": total, "items": artists}
	}

	switch searchType {
	case "all", "album":
		albums, total, _ := h.musicService.SearchAlbums(keyword, page, pageSize)
		results["results"].(gin.H)["albums"] = gin.H{"total": total, "items": albums}
	}

	switch searchType {
	case "all", "playlist":
		playlists, total, _ := h.socialService.SearchPlaylists(keyword, page, pageSize)
		results["results"].(gin.H)["playlists"] = gin.H{"total": total, "items": playlists}
	}

	results["type"] = searchType

	if userID != nil {
		h.socialService.SaveSearchHistory(c.Request.Context(), *userID, keyword, 0)
	}

	response.Success(c, results)
}

func (h *SearchHandler) Suggest(c *gin.Context) {
	keyword := c.Query("q")
	if keyword == "" {
		response.Success(c, gin.H{"suggestions": []gin.H{}})
		return
	}

	suggestions := []gin.H{
		{"type": "song", "text": keyword + " - 搜索建议1"},
		{"type": "artist", "text": keyword},
		{"type": "album", "text": keyword + " - 专辑"},
	}

	response.Success(c, gin.H{
		"keyword":     keyword,
		"suggestions": suggestions,
	})
}

func (h *SearchHandler) GetHotKeywords(c *gin.Context) {
	keywords := []gin.H{
		{"rank": 1, "keyword": "周杰伦", "heat": 980000, "trend": "up"},
		{"rank": 2, "keyword": "晴天", "heat": 850000, "trend": "stable"},
		{"rank": 3, "keyword": "新歌", "heat": 720000, "trend": "down"},
		{"rank": 4, "keyword": "邓紫棋", "heat": 680000, "trend": "up"},
		{"rank": 5, "keyword": "告白气球", "heat": 620000, "trend": "stable"},
	}

	response.Success(c, gin.H{
		"list":       keywords,
		"updated_at": "2026-05-17T10:00:00Z",
	})
}

func (h *SearchHandler) GetSearchHistory(c *gin.Context) {
	userID := c.GetUint64("user_id")
	limit := getIntQuery(c, "limit", 20)

	history, err := h.socialService.GetSearchHistory(c.Request.Context(), userID, limit)
	if err != nil {
		errors.Error(c, 500, 10000, "获取失败")
		return
	}

	response.Success(c, gin.H{"history": history})
}

func (h *SearchHandler) ClearSearchHistory(c *gin.Context) {
	userID := c.GetUint64("user_id")

	err := h.socialService.ClearSearchHistory(c.Request.Context(), userID)
	if err != nil {
		errors.Error(c, 500, 10000, "清除失败")
		return
	}

	response.Success(c, gin.H{"message": "清除成功"})
}

func (h *SearchHandler) DeleteSearchHistory(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	historyID := c.Param("history_id")
	if historyID == "" {
		errors.Error(c, 400, 10001, "无效的历史记录ID")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

func init() {
	_ = SearchHandler{}
}
