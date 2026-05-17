package handler

import (
	"github.com/bailemi/gateway/internal/model"
	"github.com/bailemi/gateway/internal/service"
	"github.com/bailemi/gateway/pkg/errors"
	"github.com/bailemi/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
	userService *service.UserService
}

func NewAuthHandler(authService *service.AuthService, userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, errors.ErrInvalidToken.Code, "参数错误: "+err.Error())
		return
	}

	if req.Email == "" && req.Phone == "" {
		errors.Error(c, 400, 20005, "邮箱或手机号至少填写一个")
		return
	}

	if len(req.Password) < 8 {
		errors.Error(c, 400, errors.ErrInvalidPassword.Code, "密码长度至少8位")
		return
	}

	result, err := h.authService.Register(c.Request.Context(), &req)
	if err != nil {
		errors.Error(c, 400, errors.ErrUserExists.Code, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, errors.ErrInvalidCredentials.Code, "参数错误: "+err.Error())
		return
	}

	ip := c.ClientIP()
	result, err := h.authService.Login(c.Request.Context(), &req, ip)
	if err != nil {
		errors.Error(c, 400, errors.ErrInvalidCredentials.Code, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req model.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, errors.ErrInvalidToken.Code, "参数错误")
		return
	}

	result, err := h.authService.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		errors.Error(c, 401, errors.ErrRefreshTokenExpired.Code, err.Error())
		return
	}

	response.Success(c, result)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	userID := c.GetUint64("user_id")
	h.authService.Logout(c.Request.Context(), userID)
	response.Success(c, nil)
}

func (h *AuthHandler) SendCode(c *gin.Context) {
	var req struct {
		Target string `json:"target" binding:"required"`
		Type   string `json:"type" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	response.Success(c, gin.H{"message": "验证码已发送"})
}

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	user, err := h.userService.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		errors.Error(c, 404, errors.ErrUserNotFound.Code, "用户不存在")
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetUint64("user_id")

	var req service.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	err := h.userService.UpdateProfile(c.Request.Context(), userID, &req)
	if err != nil {
		errors.Error(c, 400, 10001, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "更新成功"})
}

func (h *UserHandler) UploadAvatar(c *gin.Context) {
	userID := c.GetUint64("user_id")
	
	file, err := c.FormFile("avatar")
	if err != nil {
		errors.Error(c, 400, errors.ErrInvalidAvatarFormat.Code, "请上传头像文件")
		return
	}

	if file.Size > 5*1024*1024 {
		errors.Error(c, 400, errors.ErrFileTooLarge.Code, "文件大小超过限制")
		return
	}

	response.Success(c, gin.H{
		"avatar_url": "https://cdn.bailemi.com/avatars/" + string(rune(userID)) + ".jpg",
		"message": "上传成功",
	})
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID := c.Param("user_id")
	
	var uid uint64
	if _, err := parseUint64(userID, &uid); err != nil {
		errors.Error(c, 400, 10001, "无效的用户ID")
		return
	}

	user, err := h.userService.GetPublicUserProfile(c.Request.Context(), uid)
	if err != nil {
		errors.Error(c, 404, errors.ErrUserNotFound.Code, "用户不存在")
		return
	}

	response.Success(c, user)
}

func (h *UserHandler) FollowUser(c *gin.Context) {
	currentUserID := c.GetUint64("user_id")
	
	userID := c.Param("user_id")
	var targetID uint64
	if _, err := parseUint64(userID, &targetID); err != nil {
		errors.Error(c, 400, 10001, "无效的用户ID")
		return
	}

	err := h.userService.Follow(c.Request.Context(), currentUserID, targetID)
	if err != nil {
		errors.Error(c, 400, 30003, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "关注成功"})
}

func (h *UserHandler) UnfollowUser(c *gin.Context) {
	currentUserID := c.GetUint64("user_id")
	
	userID := c.Param("user_id")
	var targetID uint64
	if _, err := parseUint64(userID, &targetID); err != nil {
		errors.Error(c, 400, 10001, "无效的用户ID")
		return
	}

	err := h.userService.Unfollow(c.Request.Context(), currentUserID, targetID)
	if err != nil {
		errors.Error(c, 400, 10001, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "取消关注成功"})
}

func (h *UserHandler) GetFollowing(c *gin.Context) {
	userID := c.Param("user_id")
	var uid uint64
	if _, err := parseUint64(userID, &uid); err != nil {
		errors.Error(c, 400, 10001, "无效的用户ID")
		return
	}

	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "page_size", 20)

	users, total, err := h.userService.GetFollowing(c.Request.Context(), uid, page, pageSize)
	if err != nil {
		errors.Error(c, 500, 10000, "获取失败")
		return
	}

	response.SuccessWithPagination(c, users, page, pageSize, total)
}

func (h *UserHandler) GetFollowers(c *gin.Context) {
	userID := c.Param("user_id")
	var uid uint64
	if _, err := parseUint64(userID, &uid); err != nil {
		errors.Error(c, 400, 10001, "无效的用户ID")
		return
	}

	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "page_size", 20)

	users, total, err := h.userService.GetFollowers(c.Request.Context(), uid, page, pageSize)
	if err != nil {
		errors.Error(c, 500, 10000, "获取失败")
		return
	}

	response.SuccessWithPagination(c, users, page, pageSize, total)
}

func (h *UserHandler) DeleteAccount(c *gin.Context) {
	userID := c.GetUint64("user_id")

	var req struct {
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "请输入密码确认")
		return
	}

	err := h.userService.DeleteAccount(c.Request.Context(), userID, req.Password)
	if err != nil {
		errors.Error(c, 400, 10001, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "账号已注销"})
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID := c.GetUint64("user_id")

	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		errors.Error(c, 400, 10001, "参数错误")
		return
	}

	err := h.userService.ChangePassword(c.Request.Context(), userID, req.OldPassword, req.NewPassword)
	if err != nil {
		errors.Error(c, 400, 10001, err.Error())
		return
	}

	response.Success(c, gin.H{"message": "密码修改成功"})
}

func parseUint64(s string, result *uint64) (bool, error) {
	var v uint64
	for _, c := range s {
		if c < '0' || c > '9' {
			return false, nil
		}
		v = v*10 + uint64(c-'0')
	}
	*result = v
	return true, nil
}

func getIntQuery(c *gin.Context, key string, defaultValue int) int {
	if val := c.Query(key); val != "" {
		var v int
		for _, c := range val {
			if c < '0' || c > '9' {
				return defaultValue
			}
			v = v*10 + int(c-'0')
		}
		return v
	}
	return defaultValue
}
