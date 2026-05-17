package model

import (
	"time"
)

type PlayHistory struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint64    `gorm:"not null;index;index:idx_user_created" json:"user_id"`
	SongID        uint64    `gorm:"not null;index" json:"song_id"`
	PlayDuration  uint      `gorm:"not null;default:0" json:"play_duration"`
	PlaySource    *uint8    `json:"play_source"`
	Quality       *string   `gorm:"size:20" json:"quality"`
	ClientInfo    *string   `gorm:"size:200" json:"client_info"`
	IP            *string   `gorm:"size:45" json:"ip"`
	CreatedAt     time.Time `gorm:"autoCreateTime;index:idx_user_created;index" json:"created_at"`
	User          User      `gorm:"foreignKey:UserID" json:"-"`
	Song          Song      `gorm:"foreignKey:SongID" json:"song"`
}

func (PlayHistory) TableName() string {
	return "play_history"
}

type SearchHistory struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint64    `gorm:"not null;index;index:idx_user_created" json:"user_id"`
	Keyword      string    `gorm:"size:200;not null;index:idx_keyword" json:"keyword"`
	ResultCount  uint      `gorm:"not null;default:0" json:"result_count"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (SearchHistory) TableName() string {
	return "search_history"
}

type UserBehaviorLog struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64    `gorm:"not null;index" json:"user_id"`
	Action    string    `gorm:"size:50;not null" json:"action"`
	TargetType string   `gorm:"size:50" json:"target_type"`
	TargetID  *uint64   `json:"target_id"`
	Extra     *string   `gorm:"type:json" json:"extra"`
	IP        *string   `gorm:"size:45" json:"ip"`
	UserAgent *string   `gorm:"size:500" json:"user_agent"`
	CreatedAt time.Time `gorm:"autoCreateTime;index" json:"created_at"`
}

func (UserBehaviorLog) TableName() string {
	return "user_behavior_logs"
}

type Upload struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint64     `gorm:"not null;index" json:"user_id"`
	SongID       *uint64    `gorm:"index" json:"song_id"`
	Title        string     `gorm:"size:200;not null" json:"title"`
	FileName     string     `gorm:"size:500;not null" json:"file_name"`
	FileSize     uint64     `json:"file_size"`
	FileHash     string     `gorm:"size:64;not null" json:"file_hash"`
	StoragePath  string     `gorm:"size:500;not null" json:"storage_path"`
	CCLicense    string     `gorm:"size:20;not null" json:"cc_license"`
	CCAttribution *string   `gorm:"size:500" json:"cc_attribution"`
	IsOriginal  uint8      `gorm:"not null;default:1" json:"is_original"`
	Status      uint8      `gorm:"not null;default:0" json:"status"`
	ErrorMsg    *string    `gorm:"type:text" json:"error_msg"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Upload) TableName() string {
	return "uploads"
}

type CCLicense struct {
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	NameEn               string `json:"name_en"`
	Description          string `json:"description"`
	URL                  string `json:"url"`
	RequiresAttribution  bool   `json:"requires_attribution"`
	AllowsCommercial     bool   `json:"allows_commercial"`
	AllowsDerivatives    bool   `json:"allows_derivatives"`
}

var CCLicenses = []CCLicense{
	{
		Code:                 "cc0",
		Name:                 "CC0 1.0 通用",
		NameEn:               "CC0 1.0 Universal",
		Description:          "公有领域 dedication — 放弃所有版权，任何人可自由使用",
		URL:                  "https://creativecommons.org/publicdomain/zero/1.0/",
		RequiresAttribution:  false,
		AllowsCommercial:     true,
		AllowsDerivatives:    true,
	},
	{
		Code:                 "cc-by",
		Name:                 "署名 4.0 国际",
		NameEn:               "Attribution 4.0 International",
		Description:          "允许商用和改编，但必须标明原作者",
		URL:                  "https://creativecommons.org/licenses/by/4.0/",
		RequiresAttribution:  true,
		AllowsCommercial:    true,
		AllowsDerivatives:    true,
	},
	{
		Code:                 "cc-by-sa",
		Name:                 "署名-相同方式共享 4.0 国际",
		NameEn:               "Attribution-ShareAlike 4.0 International",
		Description:          "允许商用和改编，但必须标明原作者且衍生作品需使用相同协议",
		URL:                  "https://creativecommons.org/licenses/by-sa/4.0/",
		RequiresAttribution:  true,
		AllowsCommercial:     true,
		AllowsDerivatives:    true,
	},
	{
		Code:                 "cc-by-nc",
		Name:                 "署名-非商业性使用 4.0 国际",
		NameEn:               "Attribution-NonCommercial 4.0 International",
		Description:          "允许改编，但必须标明原作者且禁止商业用途",
		URL:                  "https://creativecommons.org/licenses/by-nc/4.0/",
		RequiresAttribution:  true,
		AllowsCommercial:    false,
		AllowsDerivatives:    true,
	},
	{
		Code:                 "cc-by-nd",
		Name:                 "署名-禁止演绎 4.0 国际",
		NameEn:               "Attribution-NoDerivatives 4.0 International",
		Description:          "允许商用，但必须标明原作者且禁止修改作品",
		URL:                  "https://creativecommons.org/licenses/by-nd/4.0/",
		RequiresAttribution:  true,
		AllowsCommercial:    true,
		AllowsDerivatives:    false,
	},
	{
		Code:                 "cc-by-nc-sa",
		Name:                 "署名-非商业性-相同方式共享 4.0 国际",
		NameEn:               "Attribution-NonCommercial-ShareAlike 4.0 International",
		Description:          "允许改编，但必须标明原作者，禁止商业用途，衍生作品需相同协议",
		URL:                  "https://creativecommons.org/licenses/by-nc-sa/4.0/",
		RequiresAttribution:  true,
		AllowsCommercial:    false,
		AllowsDerivatives:    true,
	},
	{
		Code:                 "cc-by-nc-nd",
		Name:                 "署名-非商业性-禁止演绎 4.0 国际",
		NameEn:               "Attribution-NonCommercial-NoDerivatives 4.0 International",
		Description:          "允许商用但必须标明原作者，禁止商业用途和修改",
		URL:                  "https://creativecommons.org/licenses/by-nc-nd/4.0/",
		RequiresAttribution:  true,
		AllowsCommercial:    false,
		AllowsDerivatives:    false,
	},
}

type PlayURLResponse struct {
	URL       string `json:"url"`
	Quality   string `json:"quality"`
	Bitrate   uint   `json:"bitrate"`
	Format    string `json:"format"`
	FileSize  uint64 `json:"file_size"`
	ExpireIn  int    `json:"expire_in"`
}

type PlayReportRequest struct {
	SongID       uint64 `json:"song_id" binding:"required"`
	PlayDuration uint   `json:"play_duration"`
	TotalDuration uint  `json:"total_duration"`
	PlaySource   string `json:"play_source"`
	SourceID     *uint64 `json:"source_id"`
	Quality      string `json:"quality"`
	Completed    bool   `json:"completed"`
}

type PlayHistoryResponse struct {
	Song         SongBrief `json:"song"`
	PlayedAt     time.Time `json:"played_at"`
	PlayDuration uint      `json:"play_duration"`
	Quality      string    `json:"quality"`
}

type RankType string

const (
	RankDailyPlay  RankType = "daily_play"
	RankWeeklyPlay RankType = "weekly_play"
	RankNewSongs   RankType = "new_songs"
	RankSurging    RankType = "surging"
	RankHotComment RankType = "hot_comment"
)

type RankItem struct {
	Rank      uint64   `json:"rank"`
	Song      SongBrief `json:"song"`
	PlayCount uint64   `json:"play_count"`
	Change    int      `json:"change"`
}

type RankResponse struct {
	Type      string     `json:"type"`
	Date      string     `json:"date"`
	UpdatedAt time.Time  `json:"updated_at"`
	Items     []RankItem `json:"items"`
}
