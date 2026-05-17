package model

import (
	"time"
)

type Artist struct {
	ID         uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     *uint64    `gorm:"uniqueIndex" json:"user_id"`
	Name       string     `gorm:"size:100;not null;index" json:"name"`
	AvatarURL  *string    `gorm:"size:500" json:"avatar_url"`
	CoverURL   *string    `gorm:"size:500" json:"cover_url"`
	Bio        *string    `json:"bio"`
	Region     *string    `gorm:"size:50;index" json:"region"`
	IsVerified uint8      `gorm:"not null;default:0;index" json:"is_verified"`
	Status     uint8      `gorm:"not null;default:1" json:"status"`
	FanCount   uint       `gorm:"not null;default:0" json:"fan_count"`
	SongCount  uint       `gorm:"not null;default:0" json:"song_count"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

func (Artist) TableName() string {
	return "artists"
}

type Album struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	ArtistID     uint64     `gorm:"not null;index" json:"artist_id"`
	Title        string     `gorm:"size:200;not null" json:"title"`
	CoverURL     *string    `gorm:"size:500" json:"cover_url"`
	Description  *string    `json:"description"`
	ReleaseDate  *string    `json:"release_date"`
	GenreID      *uint      `gorm:"index" json:"genre_id"`
	SongCount    uint       `gorm:"not null;default:0" json:"song_count"`
	PlayCount    uint64     `gorm:"not null;default:0" json:"play_count"`
	Status       uint8      `gorm:"not null;default:1" json:"status"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	Artist       Artist     `gorm:"foreignKey:ArtistID" json:"artist"`
	Genre        Genre      `gorm:"foreignKey:GenreID" json:"genre,omitempty"`
}

func (Album) TableName() string {
	return "albums"
}

type Song struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Title           string     `gorm:"size:200;not null;index:idx_title" json:"title"`
	ArtistID        uint64     `gorm:"not null;index" json:"artist_id"`
	AlbumID         *uint64    `gorm:"index" json:"album_id"`
	UploaderID      uint64     `gorm:"not null;index" json:"uploader_id"`
	GenreID         *uint      `gorm:"index" json:"genre_id"`
	Duration        uint       `gorm:"not null;default:0" json:"duration"`
	CoverURL        *string    `gorm:"size:500" json:"cover_url"`
	LyricStatus     uint8      `gorm:"not null;default:0" json:"lyric_status"`
	ISRC           *string    `gorm:"size:20" json:"isrc"`
	Bitrate         *uint      `json:"bitrate"`
	FileSize        *uint64    `json:"file_size"`
	FileHash        *string    `gorm:"size:64" json:"file_hash"`
	StoragePath     string     `gorm:"size:500;not null" json:"storage_path"`
	CCLicense       *string    `gorm:"size:20;index" json:"cc_license"`
	CCAttribution   *string    `gorm:"size:500" json:"cc_attribution"`
	IsOriginal     uint8      `gorm:"not null;default:1" json:"is_original"`
	PlayCount       uint64     `gorm:"not null;default:0;index:idx_play_count" json:"play_count"`
	FavoriteCount   uint       `gorm:"not null;default:0" json:"favorite_count"`
	CommentCount    uint       `gorm:"not null;default:0" json:"comment_count"`
	ShareCount      uint       `gorm:"not null;default:0" json:"share_count"`
	Status          uint8      `gorm:"not null;default:2;index" json:"status"`
	PublishAt       *time.Time `gorm:"index" json:"publish_at"`
	CreatedAt       time.Time  `gorm:"autoCreateTime;index" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
	Artist          Artist     `gorm:"foreignKey:ArtistID" json:"artist"`
	Album           *Album     `gorm:"foreignKey:AlbumID" json:"album,omitempty"`
	Genre           *Genre     `gorm:"foreignKey:GenreID" json:"genre,omitempty"`
	Lyric           *Lyric     `gorm:"foreignKey:ID" json:"lyric,omitempty"`
}

func (Song) TableName() string {
	return "songs"
}

type Lyric struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	SongID       uint64    `gorm:"uniqueIndex;not null" json:"song_id"`
	Content      string    `gorm:"LONGTEXT;not null" json:"content"`
	ContentType  uint8     `gorm:"not null;default:1" json:"content_type"`
	Translation  *string   `gorm:"LONGTEXT" json:"translation"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Lyric) TableName() string {
	return "lyrics"
}

type Fingerprint struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	SongID        uint64    `gorm:"uniqueIndex;not null" json:"song_id"`
	HashData      []byte    `gorm:"MEDIUMBLOB;not null" json:"-"`
	FeatureCount  uint      `gorm:"not null" json:"feature_count"`
	Duration      float64   `gorm:"not null" json:"duration"`
	SchemaVersion uint8     `gorm:"not null;default:1" json:"schema_version"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (Fingerprint) TableName() string {
	return "fingerprints"
}

type Genre struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `gorm:"size:50;not null;uniqueIndex" json:"name"`
	NameEn     *string   `gorm:"size:50" json:"name_en"`
	ParentID   *uint     `gorm:"index" json:"parent_id"`
	SortOrder  int       `gorm:"not null;default:0" json:"sort_order"`
	IconURL    *string   `gorm:"size:500" json:"icon_url"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Genre) TableName() string {
	return "genres"
}

type SongResponse struct {
	ID             uint64       `json:"id"`
	Title          string       `json:"title"`
	Artist         *ArtistBrief `json:"artist"`
	Album          *AlbumBrief  `json:"album,omitempty"`
	Genre          *GenreBrief  `json:"genre,omitempty"`
	Duration       uint         `json:"duration"`
	IsVIP         bool         `json:"is_vip"`
	HasLyric      bool         `json:"has_lyric"`
	LyricType     uint8        `json:"lyric_type"`
	QualityTypes  []string     `json:"quality_types"`
	PlayCount     uint64       `json:"play_count"`
	FavoriteCount uint         `json:"favorite_count"`
	CommentCount  uint         `json:"comment_count"`
	PublishAt     *time.Time   `json:"publish_at"`
	IsFavorited   bool         `json:"is_favorited"`
}

type ArtistBrief struct {
	ID       uint64  `json:"id"`
	Name     string  `json:"name"`
	AvatarURL *string `json:"avatar_url,omitempty"`
}

type AlbumBrief struct {
	ID       uint64  `json:"id"`
	Title    string  `json:"title"`
	CoverURL *string `json:"cover_url,omitempty"`
}

type GenreBrief struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type LyricResponse struct {
	Type  string        `json:"type"`
	Lines []LyricLine   `json:"lines"`
}

type LyricLine struct {
	Time       float64       `json:"time"`
	Text       string        `json:"text"`
	Translation *string      `json:"translation,omitempty"`
	Words      []LyricWord   `json:"words,omitempty"`
}

type LyricWord struct {
	Text  string  `json:"text"`
	Start float64 `json:"start"`
	End   float64 `json:"end"`
}

type ArtistResponse struct {
	ID         uint64        `json:"id"`
	Name       string        `json:"name"`
	AvatarURL  *string       `json:"avatar_url"`
	CoverURL   *string       `json:"cover_url"`
	Bio        *string       `json:"bio"`
	Region     *string       `json:"region"`
	IsVerified bool          `json:"is_verified"`
	FanCount   uint          `json:"fan_count"`
	SongCount  uint          `json:"song_count"`
	AlbumCount uint          `json:"album_count"`
	IsFollowed bool          `json:"is_followed"`
	HotSongs   []SongBrief   `json:"hot_songs,omitempty"`
	LatestAlbums []AlbumBrief `json:"latest_albums,omitempty"`
}

type SongBrief struct {
	ID       uint64  `json:"id"`
	Title    string  `json:"title"`
	Duration uint    `json:"duration"`
	IsVIP   bool    `json:"is_vip"`
}

type AlbumResponse struct {
	ID           uint64        `json:"id"`
	Title        string        `json:"title"`
	Artist       *ArtistBrief  `json:"artist"`
	CoverURL     *string       `json:"cover_url"`
	Description  *string       `json:"description"`
	ReleaseDate  *string       `json:"release_date"`
	Genre        *GenreBrief   `json:"genre"`
	SongCount    uint          `json:"song_count"`
	PlayCount    uint64        `json:"play_count"`
	Songs        []SongBrief   `json:"songs,omitempty"`
}

type GenreResponse struct {
	ID       uint            `json:"id"`
	Name     string           `json:"name"`
	NameEn   *string          `json:"name_en"`
	Children []GenreBrief     `json:"children,omitempty"`
}
