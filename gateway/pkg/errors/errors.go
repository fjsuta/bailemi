package errors

import "fmt"

type AppError struct {
	Code    int
	Message string
	HTTPStatus int
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code int, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

func (e *AppError) WithMessage(msg string) *AppError {
	return &AppError{
		Code:       e.Code,
		Message:    msg,
		HTTPStatus: e.HTTPStatus,
	}
}

func FromCode(code int) *AppError {
	if err, ok := ErrorCodes[code]; ok {
		return err
	}
	return ErrSystem
}

var ErrorCodes = map[int]*AppError{
	10000: ErrSystem,
	10001: ErrServiceUnavailable,
	10002: ErrRateLimit,
	10003: ErrMethodNotAllowed,
	10004: ErrNotFound,

	20001: ErrUnauthorized,
	20002: ErrInvalidToken,
	20003: ErrRefreshTokenExpired,
	20004: ErrInvalidCredentials,
	20005: ErrInvalidVerifyCode,
	20006: ErrUserExists,
	20007: ErrAccountDisabled,
	20008: ErrInvalidPassword,

	30001: ErrUserNotFound,
	30002: ErrNoPermission,
	30003: ErrCannotFollowSelf,
	30004: ErrAlreadyFollowed,
	30005: ErrInvalidAvatarFormat,
	30006: ErrFileTooLarge,

	40001: ErrSongNotFound,
	40002: ErrAlbumNotFound,
	40003: ErrArtistNotFound,
	40004: ErrSongUnavailable,
	40005: ErrLyricNotFound,

	50001: ErrPlaylistNotFound,
	50002: ErrCannotModifyOthersPlaylist,

	60001: ErrInvalidKeyword,

	70001: ErrCommentNotFound,

	80001: ErrRegionRestricted,
	80002: ErrPlayURLExpired,

	100001: ErrUnsupportedAudioFormat,
	100002: ErrUploadFileTooLarge,
	100003: ErrInvalidCCLicense,
	100004: ErrNoPermissionModify,
	100005: ErrCannotModifyStatus,
	100006: ErrFingerprintFailed,
}

var (
	ErrSystem              = New(10000, "系统内部错误", 500)
	ErrServiceUnavailable = New(10001, "服务暂时不可用", 503)
	ErrRateLimit          = New(10002, "请求过于频繁", 429)
	ErrMethodNotAllowed   = New(10003, "请求方法不支持", 405)
	ErrNotFound           = New(10004, "接口不存在", 404)

	ErrUnauthorized       = New(20001, "未登录或Token已过期", 401)
	ErrInvalidToken       = New(20002, "Token无效", 401)
	ErrRefreshTokenExpired = New(20003, "Refresh Token已过期", 401)
	ErrInvalidCredentials = New(20004, "用户名或密码错误", 400)
	ErrInvalidVerifyCode  = New(20005, "验证码错误或已过期", 400)
	ErrUserExists         = New(20006, "用户名/邮箱/手机号已注册", 409)
	ErrAccountDisabled    = New(20007, "账号已被禁用", 403)
	ErrInvalidPassword    = New(20008, "密码格式不符合要求", 400)

	ErrUserNotFound       = New(30001, "用户不存在", 404)
	ErrNoPermission       = New(30002, "无权访问该用户信息", 403)
	ErrCannotFollowSelf   = New(30003, "不能关注自己", 400)
	ErrAlreadyFollowed    = New(30004, "已经关注该用户", 409)
	ErrInvalidAvatarFormat = New(30005, "头像文件格式不支持", 400)
	ErrFileTooLarge       = New(30006, "文件大小超过限制", 400)

	ErrSongNotFound       = New(40001, "歌曲不存在", 404)
	ErrAlbumNotFound      = New(40002, "专辑不存在", 404)
	ErrArtistNotFound     = New(40003, "音乐人不存在", 404)
	ErrSongUnavailable    = New(40004, "歌曲已下架", 403)
	ErrLyricNotFound      = New(40005, "歌词不存在", 404)

	ErrPlaylistNotFound   = New(50001, "歌单不存在", 404)
	ErrCannotModifyOthersPlaylist = New(50002, "不能修改他人的歌单", 403)

	ErrInvalidKeyword     = New(60001, "搜索关键词无效", 400)

	ErrCommentNotFound   = New(70001, "评论不存在", 404)

	ErrRegionRestricted  = New(80001, "地区限制，暂不可播放", 403)
	ErrPlayURLExpired    = New(80002, "播放URL已过期", 410)

	ErrUnsupportedAudioFormat = New(100001, "不支持的音频格式", 400)
	ErrUploadFileTooLarge     = New(100002, "文件大小超过限制", 400)
	ErrInvalidCCLicense       = New(100003, "无效的CC协议类型", 400)
	ErrNoPermissionModify     = New(100004, "无权修改该作品", 403)
	ErrCannotModifyStatus     = New(100005, "作品状态不允许修改", 403)
	ErrFingerprintFailed      = New(100006, "音频指纹提取失败", 500)
)

func ParamError(format string, args ...interface{}) *AppError {
	return New(10001, fmt.Sprintf(format, args...), 400)
}
