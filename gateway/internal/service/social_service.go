package service

import (
	"context"
	"encoding/json"

	"github.com/bailemi/gateway/internal/model"
	"github.com/bailemi/gateway/internal/repository"
)

type SocialService struct {
	socialRepo *repository.SocialRepository
	musicRepo *repository.MusicRepository
}

func NewSocialService(socialRepo *repository.SocialRepository, musicRepo *repository.MusicRepository) *SocialService {
	return &SocialService{
		socialRepo: socialRepo,
		musicRepo: musicRepo,
	}
}

func (s *SocialService) CreatePlaylist(ctx context.Context, userID uint64, req *model.CreatePlaylistRequest) (*model.PlaylistResponse, error) {
	playlist := &model.Playlist{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		IsPublic:    1,
		IsOfficial:  0,
	}

	if len(req.Tags) > 0 {
		tagsJSON, _ := json.Marshal(req.Tags)
		tagsStr := string(tagsJSON)
		playlist.Tags = &tagsStr
	}

	if req.IsPublic {
		playlist.IsPublic = 1
	} else {
		playlist.IsPublic = 0
	}

	if err := s.socialRepo.CreatePlaylist(ctx, playlist); err != nil {
		return nil, err
	}

	return s.buildPlaylistResponse(ctx, playlist, 0)
}

func (s *SocialService) GetPlaylist(ctx context.Context, playlistID uint64, userID *uint64) (*model.PlaylistResponse, error) {
	playlist, err := s.socialRepo.GetPlaylistByID(ctx, playlistID)
	if err != nil {
		return nil, err
	}

	isFavorited := false
	if userID != nil && *userID > 0 {
		isFavorited, _ = s.socialRepo.IsFavorited(ctx, *userID, 2, playlistID)
	}

	return s.buildPlaylistResponse(ctx, playlist, isFavorited)
}

func (s *SocialService) UpdatePlaylist(ctx context.Context, playlistID, userID uint64, req *model.UpdatePlaylistRequest) error {
	playlist, err := s.socialRepo.GetPlaylistByID(ctx, playlistID)
	if err != nil {
		return err
	}

	if playlist.UserID != userID {
		return err
	}

	if req.Title != nil {
		playlist.Title = *req.Title
	}
	if req.Description != nil {
		playlist.Description = req.Description
	}
	if req.Tags != nil {
		tagsJSON, _ := json.Marshal(req.Tags)
		tagsStr := string(tagsJSON)
		playlist.Tags = &tagsStr
	}
	if req.IsPublic != nil {
		if *req.IsPublic {
			playlist.IsPublic = 1
		} else {
			playlist.IsPublic = 0
		}
	}

	return s.socialRepo.UpdatePlaylist(ctx, playlist)
}

func (s *SocialService) DeletePlaylist(ctx context.Context, playlistID, userID uint64) error {
	playlist, err := s.socialRepo.GetPlaylistByID(ctx, playlistID)
	if err != nil {
		return err
	}

	if playlist.UserID != userID {
		return err
	}

	return s.socialRepo.DeletePlaylist(ctx, playlistID)
}

func (s *SocialService) GetUserPlaylists(ctx context.Context, userID uint64, page, pageSize int) ([]model.PlaylistResponse, int64, error) {
	playlists, total, err := s.socialRepo.GetUserPlaylists(ctx, userID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]model.PlaylistResponse, len(playlists))
	for i, playlist := range playlists {
		resp, _ := s.buildPlaylistResponse(ctx, &playlist, false)
		responses[i] = *resp
	}

	return responses, total, nil
}

func (s *SocialService) AddSongsToPlaylist(ctx context.Context, playlistID, userID uint64, songIDs []uint64) error {
	playlist, err := s.socialRepo.GetPlaylistByID(ctx, playlistID)
	if err != nil {
		return err
	}

	if playlist.UserID != userID {
		return err
	}

	return s.socialRepo.AddSongsToPlaylist(ctx, playlistID, songIDs)
}

func (s *SocialService) RemoveSongsFromPlaylist(ctx context.Context, playlistID, userID uint64, songIDs []uint64) error {
	playlist, err := s.socialRepo.GetPlaylistByID(ctx, playlistID)
	if err != nil {
		return err
	}

	if playlist.UserID != userID {
		return err
	}

	return s.socialRepo.RemoveSongsFromPlaylist(ctx, playlistID, songIDs)
}

func (s *SocialService) SortPlaylistSongs(ctx context.Context, playlistID, userID uint64, songIDs []uint64) error {
	playlist, err := s.socialRepo.GetPlaylistByID(ctx, playlistID)
	if err != nil {
		return err
	}

	if playlist.UserID != userID {
		return err
	}

	return s.socialRepo.UpdatePlaylistSongOrder(ctx, playlistID, songIDs)
}

func (s *SocialService) FavoritePlaylist(ctx context.Context, userID, playlistID uint64) error {
	favorite := &model.Favorite{
		UserID:     userID,
		TargetID:   playlistID,
		TargetType: 2,
	}
	return s.socialRepo.CreateFavorite(ctx, favorite)
}

func (s *SocialService) UnfavoritePlaylist(ctx context.Context, userID, playlistID uint64) error {
	return s.socialRepo.DeleteFavorite(ctx, userID, 2, playlistID)
}

func (s *SocialService) GetRecommendedPlaylists(ctx context.Context, page, pageSize int) ([]model.PlaylistResponse, int64, error) {
	playlists, total, err := s.socialRepo.GetPublicPlaylists(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]model.PlaylistResponse, len(playlists))
	for i, playlist := range playlists {
		resp, _ := s.buildPlaylistResponse(ctx, &playlist, false)
		responses[i] = *resp
	}

	return responses, total, nil
}

func (s *SocialService) buildPlaylistResponse(ctx context.Context, playlist *model.Playlist, isFavorited bool) (*model.PlaylistResponse, error) {
	playlistSongs, _ := s.socialRepo.GetPlaylistSongs(ctx, playlist.ID)

	user := model.UserBrief{
		ID:       playlist.User.ID,
		Username: playlist.User.Username,
		AvatarURL: playlist.User.AvatarURL,
	}

	var tags []string
	if playlist.Tags != nil && *playlist.Tags != "" {
		json.Unmarshal([]byte(*playlist.Tags), &tags)
	}

	songs := make([]model.SongBrief, 0, len(playlistSongs))
	for _, ps := range playlistSongs {
		songs = append(songs, model.SongBrief{
			ID:       ps.Song.ID,
			Title:    ps.Song.Title,
			Duration: ps.Song.Duration,
			IsVIP:    false,
		})
	}

	return &model.PlaylistResponse{
		ID:            playlist.ID,
		Title:         playlist.Title,
		Description:   playlist.Description,
		CoverURL:      playlist.CoverURL,
		Tags:          tags,
		Creator:       &user,
		SongCount:     playlist.SongCount,
		PlayCount:     playlist.PlayCount,
		FavoriteCount: playlist.FavoriteCount,
		IsPublic:     playlist.IsPublic == 1,
		IsFavorited:  isFavorited,
		Songs:         songs,
		CreatedAt:     playlist.CreatedAt,
	}, nil
}

func (s *SocialService) SearchPlaylists(keyword string, page, pageSize int) ([]model.PlaylistResponse, int64, error) {
	ctx := context.Background()
	playlists, total, err := s.socialRepo.SearchPlaylists(ctx, keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]model.PlaylistResponse, len(playlists))
	for i, playlist := range playlists {
		resp, _ := s.buildPlaylistResponse(ctx, &playlist, false)
		responses[i] = *resp
	}

	return responses, total, nil
}

func (s *SocialService) GetComments(ctx context.Context, targetType uint8, targetID uint64, page, pageSize int, sort string) ([]model.CommentResponse, int64, error) {
	comments, total, err := s.socialRepo.GetComments(ctx, targetType, targetID, page, pageSize, sort)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]model.CommentResponse, len(comments))
	for i, comment := range comments {
		responses[i] = model.CommentResponse{
			ID:        comment.ID,
			User: &model.UserBrief{
				ID:        comment.User.ID,
				Username:  comment.User.Username,
				AvatarURL: comment.User.AvatarURL,
			},
			Content:   comment.Content,
			LikeCount: comment.LikeCount,
			ReplyCount: comment.ReplyCount,
			IsPinned:  comment.IsPinned == 1,
			CreatedAt: comment.CreatedAt,
		}
	}

	return responses, total, nil
}

func (s *SocialService) CreateComment(ctx context.Context, userID uint64, targetType uint8, targetID uint64, content string, parentID *uint64) (*model.CommentResponse, error) {
	comment := &model.Comment{
		UserID:     userID,
		TargetID:   targetID,
		TargetType: targetType,
		ParentID:   parentID,
		Content:    content,
		Status:     1,
	}

	if err := s.socialRepo.CreateComment(ctx, comment); err != nil {
		return nil, err
	}

	return &model.CommentResponse{
		ID:         comment.ID,
		Content:    content,
		LikeCount:  0,
		ReplyCount: 0,
		CreatedAt: comment.CreatedAt,
	}, nil
}

func (s *SocialService) DeleteComment(ctx context.Context, commentID, userID uint64, role uint8) error {
	return s.socialRepo.DeleteComment(ctx, commentID)
}

func (s *SocialService) LikeComment(ctx context.Context, userID, commentID uint64) error {
	like := &model.Like{
		UserID:     userID,
		TargetID:   commentID,
		TargetType: 1,
	}
	if err := s.socialRepo.CreateLike(ctx, like); err != nil {
		return err
	}
	return s.socialRepo.IncrementLikeCount(ctx, commentID)
}

func (s *SocialService) UnlikeComment(ctx context.Context, userID, commentID uint64) error {
	if err := s.socialRepo.DeleteLike(ctx, userID, 1, commentID); err != nil {
		return err
	}
	return s.socialRepo.DecrementLikeCount(ctx, commentID)
}

func (s *SocialService) SaveSearchHistory(ctx context.Context, userID uint64, keyword string, resultCount uint) error {
	return s.socialRepo.SaveSearchHistory(ctx, userID, keyword, resultCount)
}

func (s *SocialService) GetSearchHistory(ctx context.Context, userID uint64, limit int) ([]model.SearchHistory, error) {
	return s.socialRepo.GetSearchHistory(ctx, userID, limit)
}

func (s *SocialService) ClearSearchHistory(ctx context.Context, userID uint64) error {
	return s.socialRepo.ClearSearchHistory(ctx, userID)
}

func (s *SocialService) GetPublicPlaylists(ctx context.Context, page, pageSize int) ([]model.Playlist, int64, error) {
	return s.socialRepo.GetPublicPlaylists(ctx, page, pageSize)
}

type SearchHistory = model.SearchHistory
