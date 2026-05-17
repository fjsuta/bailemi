package service

import (
	"context"
	"fmt"
	"time"

	"github.com/bailemi/gateway/internal/model"
	"github.com/bailemi/gateway/internal/repository"
)

type PlayService struct {
	playRepo  *repository.PlayRepository
	musicRepo *repository.MusicRepository
}

func NewPlayService(playRepo *repository.PlayRepository, musicRepo *repository.MusicRepository) *PlayService {
	return &PlayService{
		playRepo:  playRepo,
		musicRepo: musicRepo,
	}
}

func (s *PlayService) GetPlayURL(ctx context.Context, songID uint64, quality string) (*model.PlayURLResponse, error) {
	song, err := s.musicRepo.GetSongByID(ctx, songID)
	if err != nil {
		return nil, err
	}

	qualityMap := map[string]struct {
		bitrate  uint
		format   string
	}{
		"standard":  {128, "mp3"},
		"high":      {320, "mp3"},
		"lossless":  {1000, "flac"},
	}

	q, ok := qualityMap[quality]
	if !ok {
		q = qualityMap["standard"]
	}

	fileSize := uint64(0)
	if song.FileSize != nil {
		fileSize = *song.FileSize
	}

	sign := fmt.Sprintf("%d_%s_%d", songID, quality, time.Now().Unix())
	
	return &model.PlayURLResponse{
		URL:      fmt.Sprintf("https://cdn.bailemi.com/stream/%d_%s.%s?sign=%s", songID, quality, q.format, sign),
		Quality:  quality,
		Bitrate:  q.bitrate,
		Format:   q.format,
		FileSize: fileSize,
		ExpireIn: 3600,
	}, nil
}

func (s *PlayService) ReportPlay(ctx context.Context, userID uint64, req *model.PlayReportRequest, clientIP string) error {
	var sourceType *uint8
	if req.PlaySource != "" {
		sourceMap := map[string]uint8{
			"search":  1,
			"recommend": 2,
			"playlist": 3,
			"album":   4,
			"fm":      5,
		}
		if v, ok := sourceMap[req.PlaySource]; ok {
			sourceType = &v
		}
	}

	history := &model.PlayHistory{
		UserID:        userID,
		SongID:        req.SongID,
		PlayDuration:  req.PlayDuration,
		PlaySource:    sourceType,
		Quality:       &req.Quality,
		IP:            &clientIP,
	}

	if err := s.playRepo.CreatePlayHistory(ctx, history); err != nil {
		return err
	}

	if req.Completed || req.PlayDuration > req.TotalDuration/2 {
		s.musicRepo.IncrementPlayCount(ctx, req.SongID)
	}

	return nil
}

func (s *PlayService) GetPlayHistory(ctx context.Context, userID uint64, page, pageSize int) ([]model.PlayHistoryResponse, int64, error) {
	histories, total, err := s.playRepo.GetUserPlayHistory(ctx, userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]model.PlayHistoryResponse, len(histories))
	for i, h := range histories {
		responses[i] = model.PlayHistoryResponse{
			Song: model.SongBrief{
				ID:       h.Song.ID,
				Title:    h.Song.Title,
				Duration: h.Song.Duration,
				IsVIP:    false,
			},
			PlayedAt:     h.CreatedAt,
			PlayDuration: h.PlayDuration,
			Quality:      "",
		}
		if h.Quality != nil {
			responses[i].Quality = *h.Quality
		}
	}

	return responses, total, nil
}

func (s *PlayService) GetRank(ctx context.Context, rankType string, date string, limit int) (*model.RankResponse, error) {
	if limit <= 0 {
		limit = 50
	}
	if limit > 100 {
		limit = 100
	}

	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	var songs []model.Song
	var err error

	switch rankType {
	case "daily_play", "weekly_play":
		songs, err = s.musicRepo.GetHotSongs(ctx, limit)
	case "new_songs":
		songs, err = s.musicRepo.GetNewSongs(ctx, limit)
	default:
		songs, err = s.musicRepo.GetHotSongs(ctx, limit)
	}

	if err != nil {
		return nil, err
	}

	items := make([]model.RankItem, len(songs))
	for i, song := range songs {
		items[i] = model.RankItem{
			Rank: uint64(i + 1),
			Song: model.SongBrief{
				ID:       song.ID,
				Title:    song.Title,
				Duration: song.Duration,
				IsVIP:    false,
			},
			PlayCount: song.PlayCount,
			Change:    0,
		}
	}

	return &model.RankResponse{
		Type:      rankType,
		Date:      date,
		UpdatedAt: time.Now(),
		Items:     items,
	}, nil
}

func (s *PlayService) SaveSearchHistory(ctx context.Context, userID uint64, keyword string, resultCount uint) error {
	return s.playRepo.SaveSearchHistory(ctx, userID, keyword, resultCount)
}

func (s *PlayService) GetSearchHistory(ctx context.Context, userID uint64, limit int) ([]model.SearchHistory, error) {
	return s.playRepo.GetSearchHistory(ctx, userID, limit)
}

func (s *PlayService) ClearSearchHistory(ctx context.Context, userID uint64) error {
	return s.playRepo.ClearSearchHistory(ctx, userID)
}
