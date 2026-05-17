package model

import (
	"time"
)

type User struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string     `gorm:"size:50;not null;uniqueIndex" json:"username"`
	Email        *string    `gorm:"size:100;uniqueIndex" json:"email"`
	Phone        *string    `gorm:"size:20;uniqueIndex" json:"phone"`
	PasswordHash string     `gorm:"size:255;not null" json:"-"`
	AvatarURL    *string    `gorm:"size:500" json:"avatar_url"`
	Role         uint8      `gorm:"not null;default:0" json:"role"`
	Status       uint8      `gorm:"not null;default:1" json:"status"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	LastLoginIP  *string    `gorm:"size:45" json:"last_login_ip"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (User) TableName() string {
	return "users"
}

type UserProfile struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64    `gorm:"uniqueIndex;not null" json:"user_id"`
	Nickname  *string   `gorm:"size:50" json:"nickname"`
	Bio       *string   `gorm:"size:500" json:"bio"`
	Gender    *uint8    `gorm:"default:0" json:"gender"`
	Birthday  *string   `json:"birthday"`
	Location  *string   `gorm:"size:100" json:"location"`
	Website   *string   `gorm:"size:200" json:"website"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
}

func (UserProfile) TableName() string {
	return "user_profiles"
}

type Follow struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	FollowerID  uint64    `gorm:"not null;index:idx_follower;uniqueIndex:uk_pair" json:"follower_id"`
	FollowingID uint64    `gorm:"not null;index:idx_following;uniqueIndex:uk_pair" json:"following_id"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Follow) TableName() string {
	return "follows"
}

type VIPRecord struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64    `gorm:"not null;index" json:"user_id"`
	PlanType  uint8     `gorm:"not null" json:"plan_type"`
	StartTime time.Time `gorm:"not null" json:"start_time"`
	ExpireTime time.Time `gorm:"not null;index" json:"expire_time"`
	OrderID   *uint64   `json:"order_id"`
	IsActive  uint8     `gorm:"not null;default:1;index" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (VIPRecord) TableName() string {
	return "vip_records"
}

type UserStats struct {
	FollowCount    int64 `json:"follow_count"`
	FanCount       int64 `json:"fan_count"`
	PlaylistCount  int64 `json:"playlist_count"`
	FavoriteCount  int64 `json:"favorite_count"`
}

type UserResponse struct {
	ID        uint64      `json:"id"`
	Username  string      `json:"username"`
	Email     *string     `json:"email,omitempty"`
	Phone     *string     `json:"phone,omitempty"`
	AvatarURL *string     `json:"avatar_url,omitempty"`
	Role      uint8       `json:"role"`
	Profile   *ProfileInfo `json:"profile,omitempty"`
	Stats     *UserStats  `json:"stats,omitempty"`
}

type ProfileInfo struct {
	Nickname  *string `json:"nickname,omitempty"`
	Bio       *string `json:"bio,omitempty"`
	Gender    *uint8  `json:"gender,omitempty"`
	Birthday  *string `json:"birthday,omitempty"`
	Location  *string `json:"location,omitempty"`
}

type RegisterRequest struct {
	Username   string `json:"username" binding:"required,min=3,max=50"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password" binding:"required,min=8,max=64"`
	VerifyCode string `json:"verify_code" binding:"required"`
}

type LoginRequest struct {
	LoginType    string `json:"login_type" binding:"required,oneof=email phone username"`
	Account      string `json:"account" binding:"required"`
	Password     string `json:"password" binding:"required"`
	CaptchaToken string `json:"captcha_token"`
}

type AuthResponse struct {
	UserID       uint64 `json:"user_id"`
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
