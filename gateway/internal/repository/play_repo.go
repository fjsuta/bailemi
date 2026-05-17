package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/bailemi/gateway/internal/model"
	"gorm.io/gorm"
)

type PlayRepository struct {
	db *gorm.DB
}

func NewPlayRepository(db *gorm.DB) *PlayRepository {
	return &PlayRepository{db: db}
}

func (r *PlayRepository) CreatePlayHistory(ctx context.Context, history *model.PlayHistory) error {
	return r.db.WithContext(ctx).Create(history).Error
}

func (r *PlayRepository) GetUserPlayHistory(ctx context.Context, userID uint64, page, pageSize int) ([]model.PlayHistory, int64, error) {
	var histories []model.PlayHistory
	var total int64

	err := r.db.WithContext(ctx).Model(&model.PlayHistory{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.WithContext(ctx).Where("user_id = ?", userID).Preload("Song").Preload("Song.Artist").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&histories).Error
	return histories, total, err
}

func (r *PlayRepository) GetRecentlyPlayed(ctx context.Context, userID uint64, limit int) ([]model.PlayHistory, error) {
	var histories []model.PlayHistory
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Preload("Song").Preload("Song.Artist").Order("created_at DESC").Limit(limit).Find(&histories).Error
	return histories, err
}

func (r *PlayRepository) SaveSearchHistory(ctx context.Context, userID uint64, keyword string, resultCount uint) error {
	history := &model.SearchHistory{
		UserID:      userID,
		Keyword:     keyword,
		ResultCount: resultCount,
	}
	return r.db.WithContext(ctx).Create(history).Error
}

func (r *PlayRepository) GetSearchHistory(ctx context.Context, userID uint64, limit int) ([]model.SearchHistory, error) {
	var histories []model.SearchHistory
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Limit(limit).Find(&histories).Error
	return histories, err
}

func (r *PlayRepository) DeleteSearchHistory(ctx context.Context, userID uint64, id uint64) error {
	return r.db.WithContext(ctx).Where("id = ? AND user_id = ?", id, userID).Delete(&model.SearchHistory{}).Error
}

func (r *PlayRepository) ClearSearchHistory(ctx context.Context, userID uint64) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&model.SearchHistory{}).Error
}

func (r *PlayRepository) CreateUpload(ctx context.Context, upload *model.Upload) error {
	return r.db.WithContext(ctx).Create(upload).Error
}

func (r *PlayRepository) GetUploadByID(ctx context.Context, id uint64) (*model.Upload, error) {
	var upload model.Upload
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&upload).Error
	if err != nil {
		return nil, err
	}
	return &upload, nil
}

func (r *PlayRepository) GetUserUploads(ctx context.Context, userID uint64, page, pageSize int, status *uint8) ([]model.Upload, int64, error) {
	var uploads []model.Upload
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Upload{}).Where("user_id = ?", userID)

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&uploads).Error
	return uploads, total, err
}

func (r *PlayRepository) UpdateUpload(ctx context.Context, upload *model.Upload) error {
	return r.db.WithContext(ctx).Save(upload).Error
}

func (r *PlayRepository) DeleteUpload(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Upload{}).Error
}

func (r *PlayRepository) GetUploadStats(ctx context.Context, userID uint64) (map[string]interface{}, error) {
	var uploads []model.Upload
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&uploads).Error
	if err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"total_uploads": len(uploads),
		"published":     0,
		"reviewing":     0,
		"rejected":      0,
		"total_plays":   0,
		"total_favorites": 0,
	}

	for _, upload := range uploads {
		switch upload.Status {
		case 1:
			stats["published"] = stats["published"].(int) + 1
		case 2:
			stats["reviewing"] = stats["reviewing"].(int) + 1
		case 3:
			stats["rejected"] = stats["rejected"].(int) + 1
		}
	}

	return stats, nil
}

type RankRepository struct {
	db *gorm.DB
}

func NewRankRepository(db *gorm.DB) *RankRepository {
	return &RankRepository{db: db}
}

func (r *RankRepository) GetDailyRank(ctx context.Context, date string, limit int) ([]model.Song, error) {
	var songs []model.Song
	startOfDay, _ := time.Parse("2006-01-02", date)
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := r.db.WithContext(ctx).Model(&model.Song{}).
		Where("deleted_at IS NULL AND status = ?", 1).
		Preload("Artist").
		Preload("Album").
		Order("play_count DESC").
		Limit(limit).
		Find(&songs).Error

	return songs, err
}

func (r *RankRepository) GetNewSongsRank(ctx context.Context, date string, limit int) ([]model.Song, error) {
	targetDate, _ := time.Parse("2006-01-02", date)
	weekAgo := targetDate.AddDate(0, 0, -7)

	var songs []model.Song
	err := r.db.WithContext(ctx).Model(&model.Song{}).
		Where("deleted_at IS NULL AND status = ? AND created_at >= ?", 1, weekAgo).
		Preload("Artist").
		Preload("Album").
		Order("created_at DESC").
		Limit(limit).
		Find(&songs).Error

	return songs, err
}

func (r *PlayRepository) CountTotalUsers(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("deleted_at IS NULL").Count(&count).Error
	return count, err
}

func (r *PlayRepository) CountTotalSongs(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Song{}).Where("deleted_at IS NULL AND status = ?", 1).Count(&count).Error
	return count, err
}

func (r *PlayRepository) CountVIPUsers(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("role >= ? AND deleted_at IS NULL", 1).Count(&count).Error
	return count, err
}

func (r *PlayRepository) GetTodayNewUsers(ctx context.Context) (int64, error) {
	var count int64
	startOfDay := time.Now().Truncate(24 * time.Hour)
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("created_at >= ? AND deleted_at IS NULL", startOfDay).Count(&count).Error
	return count, err
}

func (r *PlayRepository) GetTodayPlayCount(ctx context.Context) (int64, error) {
	var count int64
	startOfDay := time.Now().Truncate(24 * time.Hour)
	err := r.db.WithContext(ctx).Model(&model.PlayHistory{}).Where("created_at >= ?", startOfDay).Count(&count).Error
	return count, err
}

func (r *PlayRepository) GetDAU(ctx context.Context) (int64, error) {
	var count int64
	startOfDay := time.Now().Truncate(24 * time.Hour)
	err := r.db.WithContext(ctx).Model(&model.PlayHistory{}).Where("created_at >= ?", startOfDay).Distinct("user_id").Count(&count).Error
	return count, err
}

func (r *PlayRepository) GetMAU(ctx context.Context) (int64, error) {
	var count int64
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)
	err := r.db.WithContext(ctx).Model(&model.PlayHistory{}).Where("created_at >= ?", thirtyDaysAgo).Distinct("user_id").Count(&count).Error
	return count, err
}

func (r *PlayRepository) GetPlayCountByDateRange(ctx context.Context, startDate, endDate string) (map[string]int64, error) {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	var results []struct {
		Date  string
		Count int64
	}

	err := r.db.WithContext(ctx).Model(&model.PlayHistory{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("created_at >= ? AND created_at < ?", start, end.AddDate(0, 0, 1)).
		Group("DATE(created_at)").
		Order("date").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	counts := make(map[string]int64)
	for _, r := range results {
		counts[r.Date] = r.Count
	}

	return counts, nil
}

func GetDateString(date time.Time) string {
	return date.Format("2006-01-02")
}

func ParseDateString(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func (r *PlayRepository) GetDailyUsersByRange(ctx context.Context, startDate, endDate string) ([]int64, error) {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	var results []struct {
		Date  string
		Count int64
	}

	err := r.db.WithContext(ctx).Model(&model.PlayHistory{}).
		Select("DATE(created_at) as date, COUNT(DISTINCT user_id) as count").
		Where("created_at >= ? AND created_at < ?", start, end.AddDate(0, 0, 1)).
		Group("DATE(created_at)").
		Order("date").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	counts := make([]int64, len(results))
	for i, r := range results {
		counts[i], _ = strconv.ParseInt(r.Date, 10, 64)
	}

	return counts, nil
}
