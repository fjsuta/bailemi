package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bailemi/gateway/internal/model"
	"gorm.io/gorm"
)

type SocialRepository struct {
	db *gorm.DB
}

func NewSocialRepository(db *gorm.DB) *SocialRepository {
	return &SocialRepository{db: db}
}

func (r *SocialRepository) CreatePlaylist(ctx context.Context, playlist *model.Playlist) error {
	return r.db.WithContext(ctx).Create(playlist).Error
}

func (r *SocialRepository) GetPlaylistByID(ctx context.Context, id uint64) (*model.Playlist, error) {
	var playlist model.Playlist
	err := r.db.WithContext(ctx).Preload("User").Where("id = ? AND deleted_at IS NULL", id).First(&playlist).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (r *SocialRepository) GetUserPlaylists(ctx context.Context, userID uint64, page, pageSize int) ([]model.Playlist, int64, error) {
	var playlists []model.Playlist
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Playlist{}).Where("user_id = ? AND deleted_at IS NULL", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.WithContext(ctx).Where("user_id = ? AND deleted_at IS NULL", userID).Offset(offset).Limit(pageSize).Order("updated_at DESC").Find(&playlists).Error
	return playlists, total, err
}

func (r *SocialRepository) UpdatePlaylist(ctx context.Context, playlist *model.Playlist) error {
	return r.db.WithContext(ctx).Save(playlist).Error
}

func (r *SocialRepository) DeletePlaylist(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Model(&model.Playlist{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func (r *SocialRepository) AddSongsToPlaylist(ctx context.Context, playlistID uint64, songIDs []uint64) error {
	var playlistSongs []model.PlaylistSong
	maxOrder := 0

	r.db.WithContext(ctx).Model(&model.PlaylistSong{}).Where("playlist_id = ?", playlistID).Select("COALESCE(MAX(sort_order), 0)").Scan(&maxOrder)

	for i, songID := range songIDs {
		playlistSongs = append(playlistSongs, model.PlaylistSong{
			PlaylistID: playlistID,
			SongID:     songID,
			SortOrder:  maxOrder + i + 1,
		})
	}

	err := r.db.WithContext(ctx).Create(&playlistSongs).Error
	if err != nil {
		return err
	}

	return r.db.WithContext(ctx).Model(&model.Playlist{}).Where("id = ?", playlistID).UpdateColumn("song_count", gorm.Expr("song_count + ?", len(songIDs))).Error
}

func (r *SocialRepository) RemoveSongsFromPlaylist(ctx context.Context, playlistID uint64, songIDs []uint64) error {
	result := r.db.WithContext(ctx).Where("playlist_id = ? AND song_id IN ?", playlistID, songIDs).Delete(&model.PlaylistSong{})
	if result.Error != nil {
		return result.Error
	}

	return r.db.WithContext(ctx).Model(&model.Playlist{}).Where("id = ?", playlistID).UpdateColumn("song_count", gorm.Expr("song_count - ?", result.RowsAffected)).Error
}

func (r *SocialRepository) GetPlaylistSongs(ctx context.Context, playlistID uint64) ([]model.PlaylistSong, error) {
	var playlistSongs []model.PlaylistSong
	err := r.db.WithContext(ctx).Where("playlist_id = ?", playlistID).Preload("Song").Preload("Song.Artist").Preload("Song.Album").Order("sort_order ASC").Find(&playlistSongs).Error
	return playlistSongs, err
}

func (r *SocialRepository) UpdatePlaylistSongOrder(ctx context.Context, playlistID uint64, songIDs []uint64) error {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for i, songID := range songIDs {
		if err := tx.Model(&model.PlaylistSong{}).Where("playlist_id = ? AND song_id = ?", playlistID, songID).Update("sort_order", i).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (r *SocialRepository) IncrementPlaylistPlayCount(ctx context.Context, playlistID uint64) error {
	return r.db.WithContext(ctx).Model(&model.Playlist{}).Where("id = ?", playlistID).UpdateColumn("play_count", gorm.Expr("play_count + 1")).Error
}

func (r *SocialRepository) CreateFavorite(ctx context.Context, favorite *model.Favorite) error {
	err := r.db.WithContext(ctx).Create(favorite).Error
	if err != nil {
		return err
	}

	targetType := favorite.TargetType
	targetID := favorite.TargetID

	switch targetType {
	case 1:
		r.db.WithContext(ctx).Model(&model.Song{}).Where("id = ?", targetID).UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1"))
	case 2:
		r.db.WithContext(ctx).Model(&model.Playlist{}).Where("id = ?", targetID).UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1"))
	case 3:
		r.db.WithContext(ctx).Model(&model.Album{}).Where("id = ?", targetID).UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1"))
	}

	return nil
}

func (r *SocialRepository) DeleteFavorite(ctx context.Context, userID uint64, targetType uint8, targetID uint64) error {
	result := r.db.WithContext(ctx).Where("user_id = ? AND target_type = ? AND target_id = ?", userID, targetType, targetID).Delete(&model.Favorite{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return nil
	}

	switch targetType {
	case 1:
		r.db.WithContext(ctx).Model(&model.Song{}).Where("id = ?", targetID).UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1"))
	case 2:
		r.db.WithContext(ctx).Model(&model.Playlist{}).Where("id = ?", targetID).UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1"))
	case 3:
		r.db.WithContext(ctx).Model(&model.Album{}).Where("id = ?", targetID).UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1"))
	}

	return nil
}

func (r *SocialRepository) IsFavorited(ctx context.Context, userID uint64, targetType uint8, targetID uint64) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Favorite{}).Where("user_id = ? AND target_type = ? AND target_id = ?", userID, targetType, targetID).Count(&count).Error
	return count > 0, err
}

func (r *SocialRepository) GetUserFavorites(ctx context.Context, userID uint64, targetType uint8, page, pageSize int) ([]model.Favorite, int64, error) {
	var favorites []model.Favorite
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Favorite{}).Where("user_id = ? AND target_type = ?", userID, targetType)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&favorites).Error
	return favorites, total, err
}

func (r *SocialRepository) CreateComment(ctx context.Context, comment *model.Comment) error {
	err := r.db.WithContext(ctx).Create(comment).Error
	if err != nil {
		return err
	}

	if comment.ParentID == nil {
		switch comment.TargetType {
		case 1:
			r.db.WithContext(ctx).Model(&model.Song{}).Where("id = ?", comment.TargetID).UpdateColumn("comment_count", gorm.Expr("comment_count + 1"))
		}
	} else {
		r.db.WithContext(ctx).Model(&model.Comment{}).Where("id = ?", *comment.ParentID).UpdateColumn("reply_count", gorm.Expr("reply_count + 1"))
	}

	return nil
}

func (r *SocialRepository) GetCommentByID(ctx context.Context, id uint64) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.WithContext(ctx).Preload("User").Where("id = ? AND deleted_at IS NULL", id).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *SocialRepository) GetComments(ctx context.Context, targetType uint8, targetID uint64, page, pageSize int, sort string) ([]model.Comment, int64, error) {
	var comments []model.Comment
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Comment{}).Where("target_type = ? AND target_id = ? AND parent_id IS NULL AND deleted_at IS NULL", targetType, targetID)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if sort == "hot" {
		err = query.Preload("User").Preload("Replies", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User").Where("deleted_at IS NULL").Limit(3).Order("like_count DESC")
		}).Offset(offset).Limit(pageSize).Order("is_pinned DESC, like_count DESC, created_at DESC").Find(&comments).Error
	} else {
		err = query.Preload("User").Preload("Replies", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User").Where("deleted_at IS NULL").Limit(3).Order("created_at ASC")
		}).Offset(offset).Limit(pageSize).Order("is_pinned DESC, created_at DESC").Find(&comments).Error
	}

	return comments, total, err
}

func (r *SocialRepository) DeleteComment(ctx context.Context, id uint64) error {
	comment, err := r.GetCommentByID(ctx, id)
	if err != nil {
		return err
	}

	err = r.db.WithContext(ctx).Model(&model.Comment{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	if err != nil {
		return err
	}

	if comment.ParentID == nil {
		switch comment.TargetType {
		case 1:
			r.db.WithContext(ctx).Model(&model.Song{}).Where("id = ?", comment.TargetID).UpdateColumn("comment_count", gorm.Expr("comment_count - 1"))
		}
	} else {
		r.db.WithContext(ctx).Model(&model.Comment{}).Where("id = ?", *comment.ParentID).UpdateColumn("reply_count", gorm.Expr("reply_count - 1"))
	}

	return nil
}

func (r *SocialRepository) CreateLike(ctx context.Context, like *model.Like) error {
	return r.db.WithContext(ctx).Create(like).Error
}

func (r *SocialRepository) DeleteLike(ctx context.Context, userID uint64, targetType uint8, targetID uint64) error {
	return r.db.WithContext(ctx).Where("user_id = ? AND target_type = ? AND target_id = ?", userID, targetType, targetID).Delete(&model.Like{}).Error
}

func (r *SocialRepository) IsLiked(ctx context.Context, userID uint64, targetType uint8, targetID uint64) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Like{}).Where("user_id = ? AND target_type = ? AND target_id = ?", userID, targetType, targetID).Count(&count).Error
	return count > 0, err
}

func (r *SocialRepository) IncrementLikeCount(ctx context.Context, commentID uint64) error {
	return r.db.WithContext(ctx).Model(&model.Comment{}).Where("id = ?", commentID).UpdateColumn("like_count", gorm.Expr("like_count + 1")).Error
}

func (r *SocialRepository) DecrementLikeCount(ctx context.Context, commentID uint64) error {
	return r.db.WithContext(ctx).Model(&model.Comment{}).Where("id = ?", commentID).UpdateColumn("like_count", gorm.Expr("like_count - 1")).Error
}

func ConvertTagsJSON(tagsJSON *string) []string {
	if tagsJSON == nil || *tagsJSON == "" {
		return []string{}
	}
	var tags []string
	json.Unmarshal([]byte(*tagsJSON), &tags)
	return tags
}

func (r *SocialRepository) IsUserFollowing(ctx context.Context, userID, targetUserID uint64) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Follow{}).Where("follower_id = ? AND following_id = ?", userID, targetUserID).Count(&count).Error
	return count > 0, err
}

func (r *SocialRepository) GetUserPlaylistsSimple(ctx context.Context, userID uint64) ([]model.Playlist, error) {
	var playlists []model.Playlist
	err := r.db.WithContext(ctx).Where("user_id = ? AND deleted_at IS NULL", userID).Order("updated_at DESC").Find(&playlists).Error
	return playlists, err
}

func (r *SocialRepository) SearchPlaylists(ctx context.Context, keyword string, page, pageSize int) ([]model.Playlist, int64, error) {
	var playlists []model.Playlist
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Playlist{}).Where("is_public = 1 AND deleted_at IS NULL")

	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Preload("User").Offset(offset).Limit(pageSize).Order("favorite_count DESC").Find(&playlists).Error
	return playlists, total, err
}

func (r *SocialRepository) GetPublicPlaylists(ctx context.Context, page, pageSize int) ([]model.Playlist, int64, error) {
	var playlists []model.Playlist
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Playlist{}).Where("is_public = 1 AND deleted_at IS NULL").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.WithContext(ctx).Where("is_public = 1 AND deleted_at IS NULL").Preload("User").Offset(offset).Limit(pageSize).Order("play_count DESC").Find(&playlists).Error
	return playlists, total, err
}
