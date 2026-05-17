package handler

import (
	"net/http"

	"bailemi/gateway/internal/model"
	"bailemi/gateway/internal/service"
	"bailemi/gateway/pkg/errors"
	"bailemi/gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type OAuthHandler struct {
	oauthService *service.OAuthService
}

func NewOAuthHandler(oauthService *service.OAuthService) *OAuthHandler {
	return &OAuthHandler{
		oauthService: oauthService,
	}
}

// GetOAuthConfig 获取OAuth配置（公开接口）
func (h *OAuthHandler) GetOAuthConfig(c *gin.Context) {
	config := h.oauthService.GetPublicConfig()
	
	enabledProviders := make([]string, 0)
	if config.Google.Enabled {
		enabledProviders = append(enabledProviders, "google")
	}
	if config.Microsoft.Enabled {
		enabledProviders = append(enabledProviders, "microsoft")
	}
	if config.Apple.Enabled {
		enabledProviders = append(enabledProviders, "apple")
	}
	if config.Wechat.Enabled {
		enabledProviders = append(enabledProviders, "wechat")
	}
	if config.QQ.Enabled {
		enabledProviders = append(enabledProviders, "qq")
	}

	response.Success(c, gin.H{
		"enabled_providers": enabledProviders,
	})
}

// Authorize 授权跳转
func (h *OAuthHandler) Authorize(c *gin.Context) {
	provider := c.Param("provider")
	callback := c.Query("callback")
	
	authURL, err := h.oauthService.GetAuthorizationURL(provider, callback)
	if err != nil {
		errors.Error(c, http.StatusBadRequest, errors.ErrInvalidRequest.Code, errors.ErrInvalidRequest.Message)
		return
	}
	
	c.Redirect(http.StatusTemporaryRedirect, authURL)
}

// Callback OAuth回调
func (h *OAuthHandler) Callback(c *gin.Context) {
	provider := c.Param("provider")
	code := c.Query("code")
	state := c.Query("state")

	userInfo, err := h.oauthService.HandleCallback(provider, code, state)
	if err != nil {
		errors.Error(c, http.StatusUnauthorized, errors.ErrInvalidToken.Code, err.Error())
		return
	}

	response.Success(c, userInfo)
}

// Bind 绑定账号
func (h *OAuthHandler) Bind(c *gin.Context) {
	var req model.OAuthBindRequest
	if err := c.ShouldBind(&req); err != nil {
		errors.Error(c, http.StatusBadRequest, errors.ErrInvalidRequest.Code, errors.ErrInvalidRequest.Message)
		return
	}

	userID := c.GetUint64("user_id")
	if userID == 0 {
		errors.Error(c, http.StatusUnauthorized, errors.ErrUnauthorized.Code, errors.ErrUnauthorized.Message)
		return
	}

	err := h.oauthService.BindAccount(userID, req.Provider)
	if err != nil {
		errors.Error(c, http.StatusBadRequest, errors.ErrInvalidRequest.Code, err.Error())
		return
	}

	response.Success(c, gin.H{
		"message": "绑定成功",
	})
}

// Unbind 解除绑定
func (h *OAuthHandler) Unbind(c *gin.Context) {
	provider := c.Param("provider")
	userID := c.GetUint64("user_id")

	err := h.oauthService.UnbindAccount(userID, provider)
	if err != nil {
		errors.Error(c, http.StatusBadRequest, errors.ErrInvalidRequest.Code, err.Error())
		return
	}

	response.Success(c, gin.H{
		"message": "解除绑定成功",
	})
}

// GetBindings 获取已绑定的账号
func (h *OAuthHandler) GetBindings(c *gin.Context) {
	userID := c.GetUint64("user_id")

	bindings, err := h.oauthService.GetUserBindings(userID)
	if err != nil {
		errors.Error(c, http.StatusInternalServerError, errors.ErrInternalServerError.Code, err.Error())
		return
	}

	response.Success(c, bindings)
}
