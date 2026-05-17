package service

import (
	"context"
	"errors"
	"time"

	"github.com/bailemi/gateway/internal/model"
	"github.com/bailemi/gateway/internal/repository"
	"github.com/bailemi/gateway/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

type AuthService struct {
	userRepo   *repository.UserRepository
	jwtManager *jwt.JWTManager
}

func NewAuthService(userRepo *repository.UserRepository, jwtManager *jwt.JWTManager) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

func (s *AuthService) Register(ctx context.Context, req *model.RegisterRequest) (*model.AuthResponse, error) {
	if req.Email == "" && req.Phone == "" {
		return nil, errors.New("邮箱或手机号至少填写一个")
	}

	exists, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err == nil && exists != nil {
		return nil, errors.New("用户名已存在")
	}

	if req.Email != "" {
		exists, err := s.userRepo.FindByEmail(ctx, req.Email)
		if err == nil && exists != nil {
			return nil, errors.New("邮箱已被注册")
		}
	}

	if req.Phone != "" {
		exists, err := s.userRepo.FindByPhone(ctx, req.Phone)
		if err == nil && exists != nil {
			return nil, errors.New("手机号已被注册")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Status:       1,
		Role:         0,
	}

	if req.Email != "" {
		user.Email = &req.Email
	}
	if req.Phone != "" {
		user.Phone = &req.Phone
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	profile := &model.UserProfile{
		UserID:   user.ID,
		Nickname: &req.Username,
	}
	s.userRepo.CreateProfile(ctx, profile)

	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID, uuid.New().String())
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		UserID:       user.ID,
		Username:     user.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    s.jwtManager.GetAccessTokenTTL(),
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *model.LoginRequest, ip string) (*model.AuthResponse, error) {
	var user *model.User
	var err error

	switch req.LoginType {
	case "email":
		user, err = s.userRepo.FindByEmail(ctx, req.Account)
	case "phone":
		user, err = s.userRepo.FindByPhone(ctx, req.Account)
	case "username":
		user, err = s.userRepo.FindByUsername(ctx, req.Account)
	default:
		return nil, errors.New("无效的登录方式")
	}

	if err != nil {
		return nil, errors.New("用户不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	if user.Status == 0 {
		return nil, errors.New("账号已被禁用")
	}

	s.userRepo.UpdateLoginInfo(ctx, user.ID, ip)

	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID, uuid.New().String())
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		UserID:       user.ID,
		Username:     user.Username,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    s.jwtManager.GetAccessTokenTTL(),
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*model.AuthResponse, error) {
	claims, err := s.jwtManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("refresh token无效或已过期")
	}

	user, err := s.userRepo.FindByID(ctx, claims.UserID)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	if user.Status == 0 {
		return nil, errors.New("账号已被禁用")
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID, uuid.New().String())
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		UserID:       user.ID,
		Username:     user.Username,
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    s.jwtManager.GetAccessTokenTTL(),
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, userID uint64) error {
	return nil
}

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUserByID(ctx context.Context, id uint64) (*model.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	profile, _ := s.userRepo.GetProfile(ctx, id)

	followCount, _ := s.userRepo.GetFollowingCount(ctx, id)
	fanCount, _ := s.userRepo.GetFollowerCount(ctx, id)

	playlistCount := int64(0)
	favoriteCount := int64(0)

	response := &model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		AvatarURL: user.AvatarURL,
		Role:      user.Role,
		Profile: &model.ProfileInfo{
			Nickname: profile.Nickname,
			Bio:      profile.Bio,
			Gender:   profile.Gender,
			Birthday: profile.Birthday,
			Location: profile.Location,
		},
		Stats: &model.UserStats{
			FollowCount:    followCount,
			FanCount:       fanCount,
			PlaylistCount:  playlistCount,
			FavoriteCount:  favoriteCount,
		},
	}

	return response, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, userID uint64, req *model.UpdateProfileRequest) error {
	profile, err := s.userRepo.GetProfile(ctx, userID)
	if err != nil {
		profile = &model.UserProfile{UserID: userID}
	}

	if req.Nickname != nil {
		profile.Nickname = req.Nickname
	}
	if req.Bio != nil {
		profile.Bio = req.Bio
	}
	if req.Gender != nil {
		profile.Gender = req.Gender
	}
	if req.Birthday != nil {
		profile.Birthday = req.Birthday
	}
	if req.Location != nil {
		profile.Location = req.Location
	}

	if profile.ID == 0 {
		return s.userRepo.CreateProfile(ctx, profile)
	}
	return s.userRepo.UpdateProfile(ctx, profile)
}

func (s *UserService) Follow(ctx context.Context, followerID, followingID uint64) error {
	if followerID == followingID {
		return errors.New("不能关注自己")
	}

	exists, err := s.userRepo.FindByID(ctx, followingID)
	if err != nil {
		return errors.New("用户不存在")
	}

	isFollowing, _ := s.userRepo.IsFollowing(ctx, followerID, followingID)
	if isFollowing {
		return errors.New("已经关注该用户")
	}

	return s.userRepo.Follow(ctx, followerID, followingID)
}

func (s *UserService) Unfollow(ctx context.Context, followerID, followingID uint64) error {
	return s.userRepo.Unfollow(ctx, followerID, followingID)
}

func (s *UserService) GetFollowing(ctx context.Context, userID uint64, page, pageSize int) ([]model.UserResponse, int64, error) {
	users, total, err := s.userRepo.GetFollowingList(ctx, userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]model.UserResponse, len(users))
	for i, user := range users {
		responses[i] = model.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			AvatarURL: user.AvatarURL,
			Role:      user.Role,
		}
	}

	return responses, total, nil
}

func (s *UserService) GetFollowers(ctx context.Context, userID uint64, page, pageSize int) ([]model.UserResponse, int64, error) {
	users, total, err := s.userRepo.GetFollowerList(ctx, userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]model.UserResponse, len(users))
	for i, user := range users {
		responses[i] = model.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			AvatarURL: user.AvatarURL,
			Role:      user.Role,
		}
	}

	return responses, total, nil
}

func (s *UserService) IsFollowing(ctx context.Context, followerID, followingID uint64) (bool, error) {
	return s.userRepo.IsFollowing(ctx, followerID, followingID)
}

func MaskPhone(phone string) string {
	if len(phone) < 7 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-4:]
}

func MaskEmail(email string) string {
	parts := splitEmail(email)
	if len(parts) != 2 {
		return email
	}
	name := parts[0]
	domain := parts[1]
	if len(name) <= 2 {
		return name + "***@" + domain
	}
	return name[:2] + "***@" + domain
}

func splitEmail(email string) []string {
	for i := len(email) - 1; i >= 0; i-- {
		if email[i] == '@' {
			return []string{email[:i], email[i+1:]}
		}
	}
	return []string{}
}

type UpdateProfileRequest struct {
	Nickname *string `json:"nickname"`
	Bio      *string `json:"bio"`
	Gender   *uint8  `json:"gender"`
	Birthday *string `json:"birthday"`
	Location *string `json:"location"`
}

func (s *UserService) GetPublicUserProfile(ctx context.Context, userID uint64) (*model.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	profile, _ := s.userRepo.GetProfile(ctx, userID)
	followCount, _ := s.userRepo.GetFollowingCount(ctx, userID)
	fanCount, _ := s.userRepo.GetFollowerCount(ctx, userID)

	maskedPhone := ""
	if user.Phone != nil {
		maskedPhone = MaskPhone(*user.Phone)
	}

	maskedEmail := ""
	if user.Email != nil {
		maskedEmail = MaskEmail(*user.Email)
	}

	return &model.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     &maskedEmail,
		Phone:     &maskedPhone,
		AvatarURL: user.AvatarURL,
		Role:      user.Role,
		Profile: &model.ProfileInfo{
			Nickname: profile.Nickname,
			Bio:      profile.Bio,
			Gender:   profile.Gender,
			Birthday: profile.Birthday,
			Location: profile.Location,
		},
		Stats: &model.UserStats{
			FollowCount: followCount,
			FanCount:    fanCount,
		},
	}, nil
}
