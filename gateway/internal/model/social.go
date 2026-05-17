package model

import (
	"time"
)

type Playlist struct {
	ID            uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint64     `gorm:"not null;index" json:"user_id"`
	Title         string     `gorm:"size:100;not null" json:"title"`
	Description   *string    `gorm:"size:500" json:"description"`
	CoverURL      *string    `gorm:"size:500" json:"cover_url"`
	Tags          *string    `gorm:"type:json" json:"tags"`
	SongCount     uint       `gorm:"not null;default:0" json:"song_count"`
	PlayCount     uint64     `gorm:"not null;default:0" json:"play_count"`
	FavoriteCount uint       `gorm:"not null;default:0" json:"favorite_count"`
	IsPublic     uint8      `gorm:"not null;default:1" json:"is_public"`
	IsOfficial   uint8      `gorm:"not null;default:0" json:"is_official"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	User         User       `gorm:"foreignKey:UserID" json:"creator,omitempty"`
}

func (Playlist) TableName() string {
	return "playlists"
}

type PlaylistSong struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	PlaylistID  uint64    `gorm:"not null;uniqueIndex:uk_playlist_song;index:idx_playlist_id" json:"playlist_id"`
	SongID      uint64    `gorm:"not null;uniqueIndex:uk_playlist_song;index:idx_song_id" json:"song_id"`
	SortOrder   int       `gorm:"not null;default:0" json:"sort_order"`
	AddedAt     time.Time `gorm:"autoCreateTime" json:"added_at"`
	Playlist    Playlist  `gorm:"foreignKey:PlaylistID" json:"-"`
	Song        Song      `gorm:"foreignKey:SongID" json:"song"`
}

func (PlaylistSong) TableName() string {
	return "playlist_songs"
}

type Favorite struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64    `gorm:"not null;uniqueIndex:uk_user_target;index" json:"user_id"`
	TargetID  uint64    `gorm:"not null;uniqueIndex:uk_user_target" json:"target_id"`
	TargetType uint8    `gorm:"not null;uniqueIndex:uk_user_target;index:idx_target" json:"target_type"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Favorite) TableName() string {
	return "favorites"
}

type Comment struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64     `gorm:"not null;index" json:"user_id"`
	TargetID  uint64     `gorm:"not null;index:idx_target" json:"target_id"`
	TargetType uint8     `gorm:"not null;index:idx_target" json:"target_type"`
	ParentID  *uint64    `gorm:"index:idx_parent_id" json:"parent_id"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	LikeCount uint       `gorm:"not null;default:0;index:idx_like_count" json:"like_count"`
	ReplyCount uint       `gorm:"not null;default:0" json:"reply_count"`
	IsPinned  uint8      `gorm:"not null;default:0" json:"is_pinned"`
	Status    uint8      `gorm:"not null;default:1" json:"status"`
	CreatedAt time.Time  `gorm:"autoCreateTime;index" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	User      User       `gorm:"foreignKey:UserID" json:"user"`
	Replies   []Comment  `gorm:"foreignKey:ParentID" json:"replies,omitempty"`
}

func (Comment) TableName() string {
	return "comments"
}

type Like struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     uint64    `gorm:"not null;uniqueIndex:uk_user_target" json:"user_id"`
	TargetID   uint64    `gorm:"not null;uniqueIndex:uk_user_target;index:idx_target" json:"target_id"`
	TargetType uint8     `gorm:"not null;uniqueIndex:uk_user_target" json:"target_type"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Like) TableName() string {
	return "likes"
}

type Message struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	FromUserID uint64    `gorm:"not null;index" json:"from_user_id"`
	ToUserID   uint64    `gorm:"not null;index" json:"to_user_id"`
	Content    string    `gorm:"type:text;not null" json:"content"`
	IsRead    uint8     `gorm:"not null;default:0" json:"is_read"`
	CreatedAt  time.Time `gorm:"autoCreateTime;index" json:"created_at"`
}

func (Message) TableName() string {
	return "messages"
}

type Report struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	ReporterID uint64    `gorm:"not null;index" json:"reporter_id"`
	TargetID   uint64    `gorm:"not null" json:"target_id"`
	TargetType uint8     `gorm:"not null" json:"target_type"`
	Reason     string    `gorm:"type:text;not null" json:"reason"`
	Status     uint8     `gorm:"not null;default:0" json:"status"`
	HandleNote *string   `gorm:"type:text" json:"handle_note"`
	HandleBy   *uint64   `json:"handle_by"`
	HandleAt   *time.Time `json:"handle_at"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Report) TableName() string {
	return "reports"
}

type PlaylistResponse struct {
	ID            uint64        `json:"id"`
	Title         string        `json:"title"`
	Description   *string       `json:"description,omitempty"`
	CoverURL      *string       `json:"cover_url,omitempty"`
	Tags          []string      `json:"tags,omitempty"`
	Creator       *UserBrief    `json:"creator,omitempty"`
	SongCount     uint          `json:"song_count"`
	PlayCount     uint64        `json:"play_count"`
	FavoriteCount uint          `json:"favorite_count"`
	IsPublic     bool          `json:"is_public"`
	IsFavorited  bool          `json:"is_favorited"`
	Songs         []SongBrief   `json:"songs,omitempty"`
	CreatedAt     time.Time     `json:"created_at"`
}

type UserBrief struct {
	ID        uint64  `json:"id"`
	Username  string  `json:"username"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

type CommentResponse struct {
	ID         uint64            `json:"id"`
	User       *UserBrief         `json:"user"`
	Content    string            `json:"content"`
	LikeCount  uint              `json:"like_count"`
	ReplyCount uint              `json:"reply_count"`
	IsLiked    bool              `json:"is_liked"`
	IsPinned   bool              `json:"is_pinned"`
	Replies    []CommentResponse `json:"replies,omitempty"`
	CreatedAt  time.Time         `json:"created_at"`
}

type CreatePlaylistRequest struct {
	Title       string   `json:"title" binding:"required,min=1,max=100"`
	Description *string  `json:"description" binding:"max=500"`
	Tags        []string `json:"tags"`
	IsPublic   bool     `json:"is_public"`
}

type UpdatePlaylistRequest struct {
	Title       *string  `json:"title" binding:"omitempty,min=1,max=100"`
	Description *string  `json:"description" binding:"omitempty,max=500"`
	Tags        []string `json:"tags"`
	IsPublic   *bool    `json:"is_public"`
}

type PlaylistSongsRequest struct {
	SongIDs []uint64 `json:"song_ids" binding:"required,min=1"`
}

type CommentRequest struct {
	TargetType string `json:"target_type" binding:"required,oneof=song playlist album dynamic"`
	TargetID   uint64 `json:"target_id" binding:"required"`
	ParentID   *uint64 `json:"parent_id"`
	Content    string `json:"content" binding:"required,min=1,max=1000"`
}
