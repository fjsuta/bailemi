package repository

import (
	"context"
	"encoding/json"

	"github.com/bailemi/gateway/internal/model"
	"gorm.io/gorm"
)

type MusicRepository struct {
	db *gorm.DB
}

func NewMusicRepository(db *gorm.DB) *MusicRepository {
	return &MusicRepository{db: db}
}

func (r *MusicRepository) GetSongByID(ctx context.Context, id uint64) (*model.Song, error) {
	var song model.Song
	err := r.db.WithContext(ctx).Preload("Artist").Preload("Album").Preload("Genre").Where("id = ? AND deleted_at IS NULL", id).First(&song).Error
	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (r *MusicRepository) GetSongsByIDs(ctx context.Context, ids []uint64) ([]model.Song, error) {
	var songs []model.Song
	err := r.db.WithContext(ctx).Preload("Artist").Preload("Album").Where("id IN ? AND deleted_at IS NULL", ids).Find(&songs).Error
	return songs, err
}

func (r *MusicRepository) ListSongs(ctx context.Context, page, pageSize int, artistID *uint64, albumID *uint64, genreID *uint, status *uint8) ([]model.Song, int64, error) {
	var songs []model.Song
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Song{}).Where("deleted_at IS NULL")

	if artistID != nil {
		query = query.Where("artist_id = ?", *artistID)
	}
	if albumID != nil {
		query = query.Where("album_id = ?", *albumID)
	}
	if genreID != nil {
		query = query.Where("genre_id = ?", *genreID)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	} else {
		query = query.Where("status = ?", 1)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Preload("Artist").Preload("Album").Preload("Genre").Offset(offset).Limit(pageSize).Order("play_count DESC, id DESC").Find(&songs).Error
	return songs, total, err
}

func (r *MusicRepository) GetAlbumByID(ctx context.Context, id uint64) (*model.Album, error) {
	var album model.Album
	err := r.db.WithContext(ctx).Preload("Artist").Preload("Genre").Where("id = ? AND deleted_at IS NULL", id).First(&album).Error
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func (r *MusicRepository) GetAlbumsByArtist(ctx context.Context, artistID uint64, page, pageSize int) ([]model.Album, int64, error) {
	var albums []model.Album
	var total int64

	err := r.db.WithContext(ctx).Model(&model.Album{}).Where("artist_id = ? AND deleted_at IS NULL", artistID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = r.db.WithContext(ctx).Where("artist_id = ? AND deleted_at IS NULL", artistID).Offset(offset).Limit(pageSize).Order("release_date DESC").Find(&albums).Error
	return albums, total, err
}

func (r *MusicRepository) GetArtistByID(ctx context.Context, id uint64) (*model.Artist, error) {
	var artist model.Artist
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&artist).Error
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (r *MusicRepository) GetArtistByUserID(ctx context.Context, userID uint64) (*model.Artist, error) {
	var artist model.Artist
	err := r.db.WithContext(ctx).Where("user_id = ? AND deleted_at IS NULL", userID).First(&artist).Error
	if err != nil {
		return nil, err
	}
	return &artist, nil
}

func (r *MusicRepository) GetLyricBySongID(ctx context.Context, songID uint64) (*model.Lyric, error) {
	var lyric model.Lyric
	err := r.db.WithContext(ctx).Where("song_id = ?", songID).First(&lyric).Error
	if err != nil {
		return nil, err
	}
	return &lyric, nil
}

func (r *MusicRepository) GetGenres(ctx context.Context) ([]model.Genre, error) {
	var genres []model.Genre
	err := r.db.WithContext(ctx).Where("parent_id IS NULL").Order("sort_order ASC").Find(&genres).Error
	return genres, err
}

func (r *MusicRepository) GetGenreByID(ctx context.Context, id uint) (*model.Genre, error) {
	var genre model.Genre
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&genre).Error
	if err != nil {
		return nil, err
	}
	return &genre, nil
}

func (r *MusicRepository) GetChildGenres(ctx context.Context, parentID uint) ([]model.Genre, error) {
	var genres []model.Genre
	err := r.db.WithContext(ctx).Where("parent_id = ?", parentID).Order("sort_order ASC").Find(&genres).Error
	return genres, err
}

func (r *MusicRepository) SearchSongs(ctx context.Context, keyword string, page, pageSize int) ([]model.Song, int64, error) {
	var songs []model.Song
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Song{}).Where("deleted_at IS NULL AND status = ?", 1)

	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Preload("Artist").Preload("Album").Preload("Genre").Offset(offset).Limit(pageSize).Order("play_count DESC").Find(&songs).Error
	return songs, total, err
}

func (r *MusicRepository) SearchArtists(ctx context.Context, keyword string, page, pageSize int) ([]model.Artist, int64, error) {
	var artists []model.Artist
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Artist{}).Where("deleted_at IS NULL AND status = ?", 1)

	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("fan_count DESC").Find(&artists).Error
	return artists, total, err
}

func (r *MusicRepository) SearchAlbums(ctx context.Context, keyword string, page, pageSize int) ([]model.Album, int64, error) {
	var albums []model.Album
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Album{}).Where("deleted_at IS NULL AND status = ?", 1)

	if keyword != "" {
		query = query.Where("title LIKE ?", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Preload("Artist").Offset(offset).Limit(pageSize).Order("play_count DESC").Find(&albums).Error
	return albums, total, err
}

func (r *MusicRepository) IncrementPlayCount(ctx context.Context, songID uint64) error {
	return r.db.WithContext(ctx).Model(&model.Song{}).Where("id = ?", songID).UpdateColumn("play_count", gorm.Expr("play_count + 1")).Error
}

func (r *MusicRepository) IncrementCommentCount(ctx context.Context, songID uint64) error {
	return r.db.WithContext(ctx).Model(&model.Song{}).Where("id = ?", songID).UpdateColumn("comment_count", gorm.Expr("comment_count + 1")).Error
}

func (r *MusicRepository) IncrementFavoriteCount(ctx context.Context, songID uint64) error {
	return r.db.WithContext(ctx).Model(&model.Song{}).Where("id = ?", songID).UpdateColumn("favorite_count", gorm.Expr("favorite_count + 1")).Error
}

func (r *MusicRepository) DecrementFavoriteCount(ctx context.Context, songID uint64) error {
	return r.db.WithContext(ctx).Model(&model.Song{}).Where("id = ?", songID).UpdateColumn("favorite_count", gorm.Expr("favorite_count - 1")).Error
}

func (r *MusicRepository) UpdateAlbumSongCount(ctx context.Context, albumID uint64) error {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Song{}).Where("album_id = ? AND deleted_at IS NULL", albumID).Count(&count).Error
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&model.Album{}).Where("id = ?", albumID).Update("song_count", count).Error
}

func (r *MusicRepository) UpdateArtistSongCount(ctx context.Context, artistID uint64) error {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Song{}).Where("artist_id = ? AND deleted_at IS NULL", artistID).Count(&count).Error
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Model(&model.Artist{}).Where("id = ?", artistID).Update("song_count", count).Error
}

func (r *MusicRepository) GetHotSongs(ctx context.Context, limit int) ([]model.Song, error) {
	var songs []model.Song
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND status = ?", 1).Preload("Artist").Preload("Album").Preload("Genre").Order("play_count DESC").Limit(limit).Find(&songs).Error
	return songs, err
}

func (r *MusicRepository) GetNewSongs(ctx context.Context, limit int) ([]model.Song, error) {
	var songs []model.Song
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL AND status = ?", 1).Preload("Artist").Preload("Album").Preload("Genre").Order("created_at DESC").Limit(limit).Find(&songs).Error
	return songs, err
}

func (r *MusicRepository) GetArtistSongs(ctx context.Context, artistID uint64, page, pageSize int, status *uint8) ([]model.Song, int64, error) {
	var songs []model.Song
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Song{}).Where("artist_id = ? AND deleted_at IS NULL", artistID)

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Preload("Artist").Preload("Album").Preload("Genre").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&songs).Error
	return songs, total, err
}

func ConvertTagsToSlice(tagsJSON *string) []string {
	if tagsJSON == nil || *tagsJSON == "" {
		return []string{}
	}
	var tags []string
	json.Unmarshal([]byte(*tagsJSON), &tags)
	return tags
}

func ConvertSliceToTags(tags []string) string {
	if len(tags) == 0 {
		return ""
	}
	data, _ := json.Marshal(tags)
	return string(data)
}
