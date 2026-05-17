package model

import (
	"time"
)

// OAuthProvider 第三方登录提供商
type OAuthProvider struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	Provider    string    `json:"provider" gorm:"uniqueIndex;size:50;not null"` // google, microsoft, apple, wechat, qq
	ProviderID  string    `json:"provider_id" gorm:"size:255;not null"`         // 第三方用户ID
	UserID      uint64    `json:"user_id" gorm:"index;not null"`
	User        *User      `json:"-" gorm:"foreignKey:UserID"`
	AccessToken string    `json:"-" gorm:"type:text"`  // 访问令牌
	RefreshToken string   `json:"-" gorm:"type:text"` // 刷新令牌
	ExpiresAt   *time.Time `json:"expires_at"`
	AvatarURL   string    `json:"avatar_url" gorm:"size:500"`
	Email       string    `json:"email" gorm:"size:100"`
	Nickname    string    `json:"nickname" gorm:"size:100"`
	UnionID     string    `json:"union_id" gorm:"size:255"` // 微信UnionID
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// OAuthConfig OAuth配置
type OAuthConfig struct {
	Google    ProviderConfig `json:"google"`
	Microsoft ProviderConfig `json:"microsoft"`
	Apple    ProviderConfig `json:"apple"`
	Wechat   WechatConfig  `json:"wechat"`
	QQ       QQConfig      `json:"qq"`
}

// ProviderConfig OAuth提供商配置
type ProviderConfig struct {
	Enabled     bool   `json:"enabled"`
	ClientID    string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI string `json:"redirect_uri"`
	AuthURL    string `json:"auth_url"`
	TokenURL   string `json:"token_url"`
	UserInfoURL string `json:"user_info_url"`
	Scopes     []string `json:"scopes"`
}

// WechatConfig 微信配置
type WechatConfig struct {
	ProviderConfig
	AppID       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
}

// QQConfig QQ配置
type QQConfig struct {
	ProviderConfig
	AppID  string `json:"app_id"`
	AppKey string `json:"app_key"`
}

// OAuthCallbackRequest OAuth回调请求
type OAuthCallbackRequest struct {
	Code  string `json:"code" form:"code"`
	State string `json:"state" form:"state"`
}

// OAuthBindRequest 绑定请求
type OAuthBindRequest struct {
	Provider string `json:"provider" form:"provider" binding:"required"`
}

// OAuthConfigResponse OAuth配置响应
type OAuthConfigResponse struct {
	EnabledProviders []string `json:"enabled_providers"`
	Config          OAuthConfig `json:"config"`
}

// OAuthUserInfo OAuth用户信息
type OAuthUserInfo struct {
	Provider string `json:"provider"`
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	UnionID  string `json:"union_id,omitempty"`
}
