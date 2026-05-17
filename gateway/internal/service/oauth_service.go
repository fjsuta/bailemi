package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"bailemi/gateway/internal/model"
	"bailemi/gateway/pkg/database"

	"gorm.io/gorm"
)

type OAuthService struct {
	db            *gorm.DB
	config        *OAuthConfig
}

type OAuthConfig struct {
	Google    ProviderConfig `json:"google"`
	Microsoft ProviderConfig `json:"microsoft"`
	Apple     ProviderConfig `json:"apple"`
	Wechat    WechatConfig   `json:"wechat"`
	QQ        QQConfig       `json:"qq"`
}

type ProviderConfig struct {
	Enabled      bool     `json:"enabled"`
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURI  string   `json:"redirect_uri"`
	AuthURL      string   `json:"auth_url"`
	TokenURL     string   `json:"token_url"`
	UserInfoURL  string   `json:"user_info_url"`
	Scopes       []string `json:"scopes"`
}

type WechatConfig struct {
	ProviderConfig
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type QQConfig struct {
	ProviderConfig
	AppID  string `json:"app_id"`
	AppKey string `json:"app_key"`
}

func NewOAuthService() *OAuthService {
	config := loadOAuthConfig()
	return &OAuthService{
		db:     database.GetDB(),
		config: config,
	}
}

func loadOAuthConfig() *OAuthConfig {
	configPath := os.Getenv("OAUTH_CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/oauth.json"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return getDefaultConfig()
	}

	var config OAuthConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return getDefaultConfig()
	}
	return &config
}

func getDefaultConfig() *OAuthConfig {
	return &OAuthConfig{
		Google: ProviderConfig{
			Enabled:     true,
			AuthURL:     "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL:    "https://oauth2.googleapis.com/token",
			UserInfoURL: "https://www.googleapis.com/oauth2/v2/userinfo",
			Scopes:      []string{"openid", "email", "profile"},
		},
		Microsoft: ProviderConfig{
			Enabled:     true,
			AuthURL:     "https://login.microsoftonline.com/common/oauth2/v2.0/authorize",
			TokenURL:    "https://login.microsoftonline.com/common/oauth2/v2.0/token",
			UserInfoURL: "https://graph.microsoft.com/v1.0/me",
			Scopes:      []string{"openid", "email", "profile", "User.Read"},
		},
		Apple: ProviderConfig{
			Enabled:     true,
			AuthURL:     "https://appleid.apple.com/auth/authorize",
			TokenURL:    "https://appleid.apple.com/auth/token",
			UserInfoURL: "https://appleid.apple.com/auth/userinfo",
			Scopes:      []string{"name", "email"},
		},
		Wechat: WechatConfig{
			ProviderConfig: ProviderConfig{
				Enabled:    true,
				AuthURL:    "https://open.weixin.qq.com/connect/qrconnect",
				TokenURL:   "https://api.weixin.qq.com/sns/oauth2/access_token",
				UserInfoURL: "https://api.weixin.qq.com/sns/userinfo",
			},
		},
		QQ: QQConfig{
			ProviderConfig: ProviderConfig{
				Enabled:    true,
				AuthURL:    "https://graph.qq.com/oauth2.0/authorize",
				TokenURL:   "https://graph.qq.com/oauth2.0/token",
				UserInfoURL: "https://graph.qq.com/oauth2.0/me",
			},
		},
	}
}

func (s *OAuthService) GetPublicConfig() *OAuthConfig {
	return s.config
}

func (s *OAuthService) GetAuthorizationURL(provider string, callback string) (string, error) {
	var authURL string
	var clientID string
	var scopes []string

	switch provider {
	case "google":
		if !s.config.Google.Enabled {
			return "", errors.New("google login is disabled")
		}
		clientID = s.config.Google.ClientID
		scopes = s.config.Google.Scopes
		authURL = s.config.Google.AuthURL
	case "microsoft":
		if !s.config.Microsoft.Enabled {
			return "", errors.New("microsoft login is disabled")
		}
		clientID = s.config.Microsoft.ClientID
		scopes = s.config.Microsoft.Scopes
		authURL = s.config.Microsoft.AuthURL
	case "apple":
		if !s.config.Apple.Enabled {
			return "", errors.New("apple login is disabled")
		}
		clientID = s.config.Apple.ClientID
		scopes = s.config.Apple.Scopes
		authURL = s.config.Apple.AuthURL
	case "wechat":
		if !s.config.Wechat.Enabled {
			return "", errors.New("wechat login is disabled")
		}
		clientID = s.config.Wechat.AppID
		scopes = s.config.Wechat.Scopes
		authURL = s.config.Wechat.AuthURL
	case "qq":
		if !s.config.QQ.Enabled {
			return "", errors.New("qq login is disabled")
		}
		clientID = s.config.QQ.AppID
		scopes = s.config.QQ.Scopes
		authURL = s.config.QQ.AuthURL
	default:
		return "", errors.New("unsupported provider")
	}

	redirectURI := s.getRedirectURI(provider)
	state := fmt.Sprintf("%d", time.Now().UnixNano())

	params := url.Values{}
	params.Set("client_id", clientID)
	params.Set("redirect_uri", redirectURI)
	params.Set("response_type", "code")
	params.Set("scope", strings.Join(scopes, " "))
	params.Set("state", state)

	return authURL + "?" + params.Encode(), nil
}

func (s *OAuthService) getRedirectURI(provider string) string {
	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}
	return fmt.Sprintf("%s/api/v1/oauth/%s/callback", baseURL, provider)
}

func (s *OAuthService) HandleCallback(provider string, code string, state string) (*model.OAuthUserInfo, error) {
	switch provider {
	case "google":
		return s.handleGoogleCallback(code)
	case "microsoft":
		return s.handleMicrosoftCallback(code)
	case "apple":
		return s.handleAppleCallback(code)
	case "wechat":
		return s.handleWechatCallback(code)
	case "qq":
		return s.handleQQCallback(code)
	default:
		return nil, errors.New("unsupported provider")
	}
}

func (s *OAuthService) handleGoogleCallback(code string) (*model.OAuthUserInfo, error) {
	resp, err := s.fetchToken(s.config.Google.TokenURL, map[string]string{
		"code":          code,
		"client_id":     s.config.Google.ClientID,
		"client_secret": s.config.Google.ClientSecret,
		"redirect_uri":  s.getRedirectURI("google"),
		"grant_type":    "authorization_code",
	})
	if err != nil {
		return nil, err
	}

	userInfo, err := s.fetchUserInfo(s.config.Google.UserInfoURL, resp["access_token"].(string))
	if err != nil {
		return nil, err
	}

	return &model.OAuthUserInfo{
		Provider: "google",
		ID:       userInfo["id"].(string),
		Email:    userInfo["email"].(string),
		Name:     userInfo["name"].(string),
		Nickname: userInfo["name"].(string),
		Avatar:   userInfo["picture"].(string),
	}, nil
}

func (s *OAuthService) handleMicrosoftCallback(code string) (*model.OAuthUserInfo, error) {
	resp, err := s.fetchToken(s.config.Microsoft.TokenURL, map[string]string{
		"code":          code,
		"client_id":     s.config.Microsoft.ClientID,
		"client_secret": s.config.Microsoft.ClientSecret,
		"redirect_uri":  s.getRedirectURI("microsoft"),
		"grant_type":    "authorization_code",
	})
	if err != nil {
		return nil, err
	}

	userInfo, err := s.fetchUserInfo(s.config.Microsoft.UserInfoURL+"?$select=id,displayName,mail,userPrincipalName,photo", resp["access_token"].(string))
	if err != nil {
		return nil, err
	}

	email := ""
	if mail, ok := userInfo["mail"].(string); ok {
		email = mail
	} else if upn, ok := userInfo["userPrincipalName"].(string); ok {
		email = upn
	}

	return &model.OAuthUserInfo{
		Provider: "microsoft",
		ID:       userInfo["id"].(string),
		Email:    email,
		Name:     userInfo["displayName"].(string),
		Nickname: userInfo["displayName"].(string),
	}, nil
}

func (s *OAuthService) handleAppleCallback(code string) (*model.OAuthUserInfo, error) {
	resp, err := s.fetchToken(s.config.Apple.TokenURL, map[string]string{
		"code":          code,
		"client_id":     s.config.Apple.ClientID,
		"client_secret": s.config.Apple.ClientSecret,
		"redirect_uri":  s.getRedirectURI("apple"),
		"grant_type":    "authorization_code",
	})
	if err != nil {
		return nil, err
	}

	userInfo, err := s.fetchUserInfo(s.config.Apple.UserInfoURL, resp["access_token"].(string))
	if err != nil {
		return nil, err
	}

	return &model.OAuthUserInfo{
		Provider: "apple",
		ID:       userInfo["sub"].(string),
		Email:    userInfo["email"].(string),
		Name:     userInfo["name"].(map[string]interface{}),
		Nickname: userInfo["name"].(map[string]interface{})["firstName"].(string),
	}, nil
}

func (s *OAuthService) handleWechatCallback(code string) (*model.OAuthUserInfo, error) {
	tokenURL := fmt.Sprintf("%s?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		s.config.Wechat.TokenURL, s.config.Wechat.AppID, s.config.Wechat.AppSecret, code)

	resp, err := s.httpGet(tokenURL)
	if err != nil {
		return nil, err
	}

	accessToken := resp["access_token"].(string)
	openid := resp["openid"].(string)

	userInfoURL := fmt.Sprintf("%s?access_token=%s&openid=%s", s.config.Wechat.UserInfoURL, accessToken, openid)
	userInfo, err := s.httpGet(userInfoURL)
	if err != nil {
		return nil, err
	}

	unionid := ""
	if u, ok := userInfo["unionid"].(string); ok {
		unionid = u
	}

	return &model.OAuthUserInfo{
		Provider: "wechat",
		ID:       openid,
		Email:    "",
		Name:     userInfo["nickname"].(string),
		Nickname: userInfo["nickname"].(string),
		Avatar:   userInfo["headimgurl"].(string),
		UnionID:  unionid,
	}, nil
}

func (s *OAuthService) handleQQCallback(code string) (*model.OAuthUserInfo, error) {
	tokenURL := fmt.Sprintf("%s?grant_type=authorization_code&client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		s.config.QQ.TokenURL, s.config.QQ.AppID, s.config.QQ.AppKey, code, url.QueryEscape(s.getRedirectURI("qq")))

	resp, err := s.httpGet(tokenURL)
	if err != nil {
		return nil, err
	}

	accessToken := resp["access_token"].(string)

	openidURL := fmt.Sprintf("%s?access_token=%s", s.config.QQ.UserInfoURL, accessToken)
	openidResp, err := s.httpGet(openidURL)
	if err != nil {
		return nil, err
	}

	openid := openidResp["openid"].(string)

	userInfoURL := fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s",
		accessToken, s.config.QQ.AppID, openid)
	userInfo, err := s.httpGet(userInfoURL)
	if err != nil {
		return nil, err
	}

	return &model.OAuthUserInfo{
		Provider: "qq",
		ID:       openid,
		Email:    "",
		Name:     userInfo["nickname"].(string),
		Nickname: userInfo["nickname"].(string),
		Avatar:   userInfo["figureurl_qq_2"].(string),
	}, nil
}

func (s *OAuthService) fetchToken(tokenURL string, params map[string]string) (map[string]interface{}, error) {
	formData := url.Values{}
	for k, v := range params {
		formData.Set(k, v)
	}

	body, err := s.httpPost(tokenURL, formData.Encode(), map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	})
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (s *OAuthService) fetchUserInfo(userInfoURL string, accessToken string) (map[string]interface{}, error) {
	return s.httpGet(userInfoURL + "?access_token=" + accessToken)
}

func (s *OAuthService) httpGet(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *OAuthService) httpPost(url string, body string, headers map[string]string) (map[string]interface{}, error) {
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *OAuthService) BindAccount(userID uint64, provider string) error {
	var existing model.OAuthProvider
	if err := s.db.Where("user_id = ? AND provider = ?", userID, provider).First(&existing).Error; err == nil {
		return errors.New("this account is already bound")
	}

	var user model.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return errors.New("user not found")
	}

	oauth := &model.OAuthProvider{
		Provider: provider,
		UserID:   userID,
	}
	return s.db.Create(oauth).Error
}

func (s *OAuthService) UnbindAccount(userID uint64, provider string) error {
	return s.db.Where("user_id = ? AND provider = ?", userID, provider).Delete(&model.OAuthProvider{}).Error
}

func (s *OAuthService) GetUserBindings(userID uint64) ([]model.OAuthProvider, error) {
	var bindings []model.OAuthProvider
	err := s.db.Where("user_id = ?", userID).Find(&bindings).Error
	return bindings, err
}

func (s *OAuthService) FindOrCreateUser(userInfo *model.OAuthUserInfo) (*model.User, error) {
	var oauth model.OAuthProvider
	err := s.db.Where("provider = ? AND provider_id = ?", userInfo.Provider, userInfo.ID).First(&oauth).Error

	if err == nil {
		var user model.User
		if err := s.db.First(&user, oauth.UserID).Error; err != nil {
			return nil, err
		}
		return &user, nil
	}

	var newUser model.User
	newUser.Username = userInfo.Nickname
	newUser.Nickname = userInfo.Nickname
	newUser.AvatarURL = userInfo.Avatar
	newUser.Email = userInfo.Email
	newUser.Status = "active"

	if err := s.db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	oauth = model.OAuthProvider{
		Provider:   userInfo.Provider,
		ProviderID: userInfo.ID,
		UserID:     newUser.ID,
		AvatarURL:  userInfo.Avatar,
		Email:      userInfo.Email,
		Nickname:   userInfo.Nickname,
		UnionID:    userInfo.UnionID,
	}
	if err := s.db.Create(&oauth).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (s *OAuthService) SaveOAuthToken(userID uint64, userInfo *model.OAuthUserInfo, accessToken, refreshToken string, expiresIn int64) error {
	oauth := &model.OAuthProvider{
		Provider:     userInfo.Provider,
		ProviderID:   userInfo.ID,
		UserID:       userID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	if expiresIn > 0 {
		expiresAt := time.Now().Add(time.Duration(expiresIn) * time.Second)
		oauth.ExpiresAt = &expiresAt
	}

	return s.db.Where("provider = ? AND provider_id = ?", userInfo.Provider, userInfo.ID).Assign(oauth).FirstOrCreate(oauth).Error
}
