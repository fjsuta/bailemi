package handler

import (
	"strconv"

	"github.com/bailemi/gateway/internal/model"
	"github.com/bailemi/gateway/internal/service"
	"github.com/bailemi/gateway/pkg/errors"
	"github.com/bailemi/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	socialService *service.SocialService
}

func NewCommentHandler(socialService *service.SocialService) *CommentHandler {
	return &CommentHandler{socialService: socialService}
}

func (h *CommentHandler) GetComments(c *gin.Context) {
	targetType := c.Query("type")
	targetIDStr := c.Query("target_id")
	
	if targetType == "" || targetIDStr == "" {
		errors.Error(c, 400, 10001, "缺少参数")
		return
	}

	targetID, err := strconv.ParseUint(targetIDStr, 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的目标ID")
		return
	}

	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "page_size", 20)
	sort := c.DefaultQuery("sort", "hot")

	targetTypeValue := uint8(1)
	switch targetType {
	case "song":
		targetTypeValue = 1
	case "playlist":
		targetTypeValue = 2
	case "album":
		targetTypeValue = 3
	case "dynamic":
		targetTypeValue = 4
	}

	comments, total, err := h.socialService.GetComments(c.Request.Context(), targetTypeValue, targetID, page, pageSize, sort)
	if err != nil {
		errors.Error(c, 500, 10000, "获取失败")
		return
	}

	response.SuccessWithPagination(c, comments, page, pageSize, total)
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	userID := c.GetUint64("user_id")

	var req model.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	targetTypeValue := uint8(1)
	switch req.TargetType {
	case "song":
		targetTypeValue = 1
	case "playlist":
		targetTypeValue = 2
	case "album":
		targetTypeValue = 3
	case "dynamic":
		targetTypeValue = 4
	}

	comment, err := h.socialService.CreateComment(c.Request.Context(), userID, targetTypeValue, req.TargetID, req.Content, req.ParentID)
	if err != nil {
		errors.Error(c, 500, 10000, "评论失败")
		return
	}

	response.Success(c, comment)
}

func (h *CommentHandler) DeleteComment(c *gin.Context) {
	userID := c.GetUint64("user_id")
	role := c.GetUint8("role")

	commentID, err := strconv.ParseUint(c.Param("comment_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的评论ID")
		return
	}

	err = h.socialService.DeleteComment(c.Request.Context(), commentID, userID, role)
	if err != nil {
		errors.Error(c, 500, 10000, "删除失败")
		return
	}

	response.Success(c, gin.H{"message": "删除成功"})
}

func (h *CommentHandler) LikeComment(c *gin.Context) {
	userID := c.GetUint64("user_id")

	commentID, err := strconv.ParseUint(c.Param("comment_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的评论ID")
		return
	}

	err = h.socialService.LikeComment(c.Request.Context(), userID, commentID)
	if err != nil {
		errors.Error(c, 500, 10000, "点赞失败")
		return
	}

	response.Success(c, gin.H{"message": "点赞成功"})
}

func (h *CommentHandler) UnlikeComment(c *gin.Context) {
	userID := c.GetUint64("user_id")

	commentID, err := strconv.ParseUint(c.Param("comment_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的评论ID")
		return
	}

	err = h.socialService.UnlikeComment(c.Request.Context(), userID, commentID)
	if err != nil {
		errors.Error(c, 500, 10000, "取消点赞失败")
		return
	}

	response.Success(c, gin.H{"message": "取消点赞成功"})
}
