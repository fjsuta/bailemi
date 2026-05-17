package handler

import (
	"strconv"

	"github.com/bailemi/gateway/internal/service"
	"github.com/bailemi/gateway/pkg/errors"
	"github.com/bailemi/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

type MusicHandler struct {
	musicService *service.MusicService
}

func NewMusicHandler(musicService *service.MusicService) *MusicHandler {
	return &MusicHandler{musicService: musicService}
}

func (h *MusicHandler) GetSong(c *gin.Context) {
	songID, err := strconv.ParseUint(c.Param("song_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌曲ID")
		return
	}

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	song, err := h.musicService.GetSongByID(c.Request.Context(), songID, userID)
	if err != nil {
		errors.Error(c, 404, errors.ErrSongNotFound.Code, "歌曲不存在")
		return
	}

	response.Success(c, song)
}

func (h *MusicHandler) GetLyric(c *gin.Context) {
	songID, err := strconv.ParseUint(c.Param("song_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的歌曲ID")
		return
	}

	lyric, err := h.musicService.GetLyricBySongID(c.Request.Context(), songID)
	if err != nil {
		errors.Error(c, 404, errors.ErrLyricNotFound.Code, "歌词不存在")
		return
	}

	response.Success(c, lyric)
}

func (h *MusicHandler) GetAlbum(c *gin.Context) {
	albumID, err := strconv.ParseUint(c.Param("album_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的专辑ID")
		return
	}

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	album, err := h.musicService.GetAlbumByID(c.Request.Context(), albumID, userID)
	if err != nil {
		errors.Error(c, 404, errors.ErrAlbumNotFound.Code, "专辑不存在")
		return
	}

	response.Success(c, album)
}

func (h *MusicHandler) GetArtist(c *gin.Context) {
	artistID, err := strconv.ParseUint(c.Param("artist_id"), 10, 64)
	if err != nil {
		errors.Error(c, 400, 10001, "无效的音乐人ID")
		return
	}

	var userID *uint64
	if uid, exists := c.Get("user_id"); exists {
		id := uid.(uint64)
		userID = &id
	}

	artist, err := h.musicService.GetArtistByID(c.Request.Context(), artistID, userID)
	if err != nil {
		errors.Error(c, 404, errors.ErrArtistNotFound.Code, "音乐人不存在")
		return
	}

	response.Success(c, artist)
}

func (h *MusicHandler) GetGenres(c *gin.Context) {
	genres, err := h.musicService.GetGenres(c.Request.Context())
	if err != nil {
		errors.Error(c, 500, 10000, "获取流派列表失败")
		return
	}

	response.Success(c, genres)
}

func (h *MusicHandler) GetHotSongs(c *gin.Context) {
	limit := getIntQuery(c, "limit", 20)
	
	songs, err := h.musicService.GetHotSongs(c.Request.Context(), limit)
	if err != nil {
		errors.Error(c, 500, 10000, "获取热门歌曲失败")
		return
	}

	response.Success(c, songs)
}

func (h *MusicHandler) GetNewSongs(c *gin.Context) {
	limit := getIntQuery(c, "limit", 20)
	
	songs, err := h.musicService.GetNewSongs(c.Request.Context(), limit)
	if err != nil {
		errors.Error(c, 500, 10000, "获取新歌曲失败")
		return
	}

	response.Success(c, songs)
}
