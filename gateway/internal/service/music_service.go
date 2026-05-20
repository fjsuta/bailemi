package service

import (
	"context"
	"encoding/json"

	"github.com/bailemi/gateway/internal/model"
	"github.com/bailemi/gateway/internal/repository"
)

type MusicService struct {
	musicRepo  *repository.MusicRepository
	socialRepo *repository.SocialRepository
}

func NewMusicService(musicRepo *repository.MusicRepository, socialRepo *repository.SocialRepository) *MusicService {
	return &MusicService{
		musicRepo:  musicRepo,
		socialRepo: socialRepo,
	}
}

func (s *MusicService) GetSongByID(ctx context.Context, songID uint64, userID *uint64) (*model.SongResponse, error) {
	song, err := s.musicRepo.GetSongByID(ctx, songID)
	if err != nil {
		return nil, err
	}

	var isFavorited bool
	if userID != nil && *userID > 0 {
		isFavorited, _ = s.socialRepo.IsFavorited(ctx, *userID, 1, songID)
	}

	qualityTypes := []string{"standard", "high", "lossless"}
	if song.Bitrate != nil && *song.Bitrate < 320 {
		qualityTypes = []string{"standard"}
	} else if song.Bitrate != nil && *song.Bitrate < 1000 {
		qualityTypes = []string{"standard", "high"}
	}

	artistBrief := &model.ArtistBrief{
		ID:        song.Artist.ID,
		Name:      song.Artist.Name,
		AvatarURL: song.Artist.AvatarURL,
	}

	var albumBrief *model.AlbumBrief
	if song.Album != nil {
		albumBrief = &model.AlbumBrief{
			ID:       song.Album.ID,
			Title:    song.Album.Title,
			CoverURL: song.Album.CoverURL,
		}
	}

	var genreBrief *model.GenreBrief
	if song.Genre != nil {
		genreBrief = &model.GenreBrief{
			ID:   song.Genre.ID,
			Name: song.Genre.Name,
		}
	}

	return &model.SongResponse{
		ID:             song.ID,
		Title:          song.Title,
		Artist:         artistBrief,
		Album:          albumBrief,
		Genre:          genreBrief,
		Duration:       song.Duration,
		IsVIP:          false,
		HasLyric:       song.LyricStatus > 0,
		LyricType:      song.LyricStatus,
		QualityTypes:   qualityTypes,
		PlayCount:      song.PlayCount,
		FavoriteCount:  song.FavoriteCount,
		CommentCount:   song.CommentCount,
		PublishAt:      song.PublishAt,
		IsFavorited:    isFavorited,
	}, nil
}

func (s *MusicService) GetAlbumByID(ctx context.Context, albumID uint64, userID *uint64) (*model.AlbumResponse, error) {
	album, err := s.musicRepo.GetAlbumByID(ctx, albumID)
	if err != nil {
		return nil, err
	}

	artistBrief := &model.ArtistBrief{
		ID:   album.Artist.ID,
		Name: album.Artist.Name,
	}

	var genreBrief *model.GenreBrief
	if album.Genre != nil {
		genreBrief = &model.GenreBrief{
			ID:   album.Genre.ID,
			Name: album.Genre.Name,
		}
	}

	songs, err := s.musicRepo.ListSongs(ctx, 1, 100, nil, &albumID, nil, nil)
	if err != nil {
		return nil, err
	}

	songBriefs := make([]model.SongBrief, len(songs))
	for i, song := range songs {
		songBriefs[i] = model.SongBrief{
			ID:       song.ID,
			Title:    song.Title,
			Duration: song.Duration,
			IsVIP:    false,
		}
	}

	return &model.AlbumResponse{
		ID:          album.ID,
		Title:       album.Title,
		Artist:      artistBrief,
		CoverURL:    album.CoverURL,
		Description: album.Description,
		ReleaseDate: album.ReleaseDate,
		Genre:       genreBrief,
		SongCount:   album.SongCount,
		PlayCount:   album.PlayCount,
		Songs:       songBriefs,
	}, nil
}

func (s *MusicService) GetArtistByID(ctx context.Context, artistID uint64, userID *uint64) (*model.ArtistResponse, error) {
	artist, err := s.musicRepo.GetArtistByID(ctx, artistID)
	if err != nil {
		return nil, err
	}

	var isFollowed bool
	if userID != nil && *userID > 0 && artist.UserID != nil {
		isFollowed, _ = s.socialRepo.IsUserFollowing(ctx, *userID, *artist.UserID)
	}

	hotSongs, _ := s.musicRepo.ListSongs(ctx, 1, 10, &artistID, nil, nil, nil)
	latestAlbums, _ := s.musicRepo.GetAlbumsByArtist(ctx, artistID, 1, 5)

	songBriefs := make([]model.SongBrief, len(hotSongs))
	for i, song := range hotSongs {
		songBriefs[i] = model.SongBrief{
			ID:       song.ID,
			Title:    song.Title,
			Duration: song.Duration,
			IsVIP:    false,
		}
	}

	albumBriefs := make([]model.AlbumBrief, len(latestAlbums))
	for i, album := range latestAlbums {
		albumBriefs[i] = model.AlbumBrief{
			ID:       album.ID,
			Title:    album.Title,
			CoverURL: album.CoverURL,
		}
	}

	return &model.ArtistResponse{
		ID:            artist.ID,
		Name:          artist.Name,
		AvatarURL:     artist.AvatarURL,
		CoverURL:      artist.CoverURL,
		Bio:           artist.Bio,
		Region:        artist.Region,
		IsVerified:    artist.IsVerified == 1,
		FanCount:      artist.FanCount,
		SongCount:     artist.SongCount,
		AlbumCount:    uint(len(latestAlbums)),
		IsFollowed:    isFollowed,
		HotSongs:      songBriefs,
		LatestAlbums:  albumBriefs,
	}, nil
}

func (s *MusicService) ListArtists(ctx context.Context, page, pageSize int, sort, letter string, userID *uint64) ([]model.ArtistBrief, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	artists, total, err := s.musicRepo.ListArtists(ctx, page, pageSize, sort, letter)
	if err != nil {
		return nil, 0, err
	}

	briefs := make([]model.ArtistBrief, len(artists))
	for i, artist := range artists {
		briefs[i] = model.ArtistBrief{
			ID:        artist.ID,
			Name:      artist.Name,
			AvatarURL: artist.AvatarURL,
		}
	}

	return briefs, total, nil
}

func (s *MusicService) GetLyricBySongID(ctx context.Context, songID uint64) (*model.LyricResponse, error) {
	lyric, err := s.musicRepo.GetLyricBySongID(ctx, songID)
	if err != nil {
		return nil, err
	}

	var lyricData model.LyricResponse
	if err := json.Unmarshal([]byte(lyric.Content), &lyricData); err != nil {
		lyricData = model.LyricResponse{
			Type:  "text",
			Lines: []model.LyricLine{},
		}
	}

	return &lyricData, nil
}

func (s *MusicService) GetGenres(ctx context.Context) ([]model.GenreResponse, error) {
	genres, err := s.musicRepo.GetGenres(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]model.GenreResponse, len(genres))
	for i, genre := range genres {
		children, _ := s.musicRepo.GetChildGenres(ctx, genre.ID)
		childBriefs := make([]model.GenreBrief, len(children))
		for j, child := range children {
			childBriefs[j] = model.GenreBrief{
				ID:   child.ID,
				Name: child.Name,
			}
		}

		nameEn := genre.NameEn
		responses[i] = model.GenreResponse{
			ID:       genre.ID,
			Name:     genre.Name,
			NameEn:   nameEn,
			Children: childBriefs,
		}
	}

	return responses, nil
}

func (s *MusicService) GetHotSongs(ctx context.Context, limit int) ([]model.SongBrief, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	songs, err := s.musicRepo.GetHotSongs(ctx, limit)
	if err != nil {
		return nil, err
	}

	briefs := make([]model.SongBrief, len(songs))
	for i, song := range songs {
		briefs[i] = model.SongBrief{
			ID:       song.ID,
			Title:    song.Title,
			Duration: song.Duration,
			IsVIP:    false,
		}
	}

	return briefs, nil
}

func (s *MusicService) GetNewSongs(ctx context.Context, limit int) ([]model.SongBrief, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	songs, err := s.musicRepo.GetNewSongs(ctx, limit)
	if err != nil {
		return nil, err
	}

	briefs := make([]model.SongBrief, len(songs))
	for i, song := range songs {
		briefs[i] = model.SongBrief{
			ID:       song.ID,
			Title:    song.Title,
			Duration: song.Duration,
			IsVIP:    false,
		}
	}

	return briefs, nil
}

func (s *MusicService) IncrementPlayCount(ctx context.Context, songID uint64) error {
	return s.musicRepo.IncrementPlayCount(ctx, songID)
}

func (s *MusicService) GetArtistSongs(ctx context.Context, artistID uint64, page, pageSize int, status string) ([]model.SongBrief, int64, error) {
	var statusValue *uint8
	if status != "" && status != "all" {
		switch status {
		case "published":
			v := uint8(1)
			statusValue = &v
		case "reviewing":
			v := uint8(2)
			statusValue = &v
		case "rejected":
			v := uint8(3)
			statusValue = &v
		}
	}

	songs, total, err := s.musicRepo.GetArtistSongs(ctx, artistID, page, pageSize, statusValue)
	if err != nil {
		return nil, 0, err
	}

	briefs := make([]model.SongBrief, len(songs))
	for i, song := range songs {
		briefs[i] = model.SongBrief{
			ID:       song.ID,
			Title:    song.Title,
			Duration: song.Duration,
			IsVIP:    false,
		}
	}

	return briefs, total, nil
}

func (s *MusicService) SearchSongs(keyword string, page, pageSize int) ([]model.SongBrief, int64, error) {
	songs, total, err := s.musicRepo.SearchSongs(context.Background(), keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	briefs := make([]model.SongBrief, len(songs))
	for i, song := range songs {
		briefs[i] = model.SongBrief{
			ID:       song.ID,
			Title:    song.Title,
			Duration: song.Duration,
			IsVIP:    false,
		}
	}

	return briefs, total, nil
}

func (s *MusicService) SearchArtists(keyword string, page, pageSize int) ([]model.ArtistBrief, int64, error) {
	artists, total, err := s.musicRepo.SearchArtists(context.Background(), keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	briefs := make([]model.ArtistBrief, len(artists))
	for i, artist := range artists {
		briefs[i] = model.ArtistBrief{
			ID:        artist.ID,
			Name:      artist.Name,
			AvatarURL: artist.AvatarURL,
		}
	}

	return briefs, total, nil
}

func (s *MusicService) SearchAlbums(keyword string, page, pageSize int) ([]model.AlbumBrief, int64, error) {
	albums, total, err := s.musicRepo.SearchAlbums(context.Background(), keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	briefs := make([]model.AlbumBrief, len(albums))
	for i, album := range albums {
		briefs[i] = model.AlbumBrief{
			ID:       album.ID,
			Title:    album.Title,
			CoverURL: album.CoverURL,
		}
	}

	return briefs, total, nil
}

func (s *MusicService) GetArtistAlbums(ctx context.Context, artistID uint64, page, pageSize int) ([]model.AlbumBrief, int64, error) {
	albums, total, err := s.musicRepo.GetAlbumsByArtist(ctx, artistID, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	briefs := make([]model.AlbumBrief, len(albums))
	for i, album := range albums {
		briefs[i] = model.AlbumBrief{
			ID:       album.ID,
			Title:    album.Title,
			CoverURL: album.CoverURL,
		}
	}

	return briefs, total, nil
}

func (s *MusicService) GetChartSongs(ctx context.Context, chartType string, limit int, userID *uint64) ([]model.SongResponse, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	var songs []model.Song
	var err error

	switch chartType {
	case "hot":
		songs, err = s.musicRepo.GetHotSongs(ctx, limit)
	case "new":
		songs, err = s.musicRepo.GetNewSongs(ctx, limit)
	case "rising":
		songs, err = s.musicRepo.GetRisingSongs(ctx, limit)
	default:
		songs, err = s.musicRepo.GetHotSongs(ctx, limit)
	}

	if err != nil {
		return nil, err
	}

	responses := make([]model.SongResponse, len(songs))
	for i, song := range songs {
		var isFavorited bool
		if userID != nil && *userID > 0 {
			isFavorited, _ = s.socialRepo.IsFavorited(ctx, *userID, 1, song.ID)
		}

		qualityTypes := []string{"standard", "high", "lossless"}
		if song.Bitrate != nil && *song.Bitrate < 320 {
			qualityTypes = []string{"standard"}
		} else if song.Bitrate != nil && *song.Bitrate < 1000 {
			qualityTypes = []string{"standard", "high"}
		}

		var albumBrief *model.AlbumBrief
		if song.Album != nil {
			albumBrief = &model.AlbumBrief{
				ID:       song.Album.ID,
				Title:    song.Album.Title,
				CoverURL: song.Album.CoverURL,
			}
		}

		var genreBrief *model.GenreBrief
		if song.Genre != nil {
			genreBrief = &model.GenreBrief{
				ID:   song.Genre.ID,
				Name: song.Genre.Name,
			}
		}

		responses[i] = model.SongResponse{
			ID:             song.ID,
			Title:          song.Title,
			Artist:         &model.ArtistBrief{ID: song.Artist.ID, Name: song.Artist.Name, AvatarURL: song.Artist.AvatarURL},
			Album:          albumBrief,
			Genre:          genreBrief,
			Duration:       song.Duration,
			IsVIP:          false,
			HasLyric:       song.LyricStatus > 0,
			LyricType:      song.LyricStatus,
			QualityTypes:   qualityTypes,
			PlayCount:      song.PlayCount,
			FavoriteCount:  song.FavoriteCount,
			CommentCount:   song.CommentCount,
			PublishAt:      song.PublishAt,
			IsFavorited:    isFavorited,
		}
	}

	return responses, nil
}
