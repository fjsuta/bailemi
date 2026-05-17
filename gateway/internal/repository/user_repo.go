package repository

import (
	"context"

	"github.com/bailemi/gateway/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) FindByID(ctx context.Context, id uint64) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("username = ? AND deleted_at IS NULL", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByPhone(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("phone = ? AND deleted_at IS NULL", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) CreateProfile(ctx context.Context, profile *model.UserProfile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *UserRepository) GetProfile(ctx context.Context, userID uint64) (*model.UserProfile, error) {
	var profile model.UserProfile
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *UserRepository) UpdateProfile(ctx context.Context, profile *model.UserProfile) error {
	return r.db.WithContext(ctx).Save(profile).Error
}

func (r *UserRepository) UpdateLoginInfo(ctx context.Context, userID uint64, ip string) error {
	return r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"last_login_at": gorm.Expr("NOW()"),
		"last_login_ip": ip,
	}).Error
}

func (r *UserRepository) CountByRole(ctx context.Context, role uint8) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("role = ? AND deleted_at IS NULL", role).Count(&count).Error
	return count, err
}

func (r *UserRepository) CountByStatus(ctx context.Context, status uint8) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Where("status = ? AND deleted_at IS NULL", status).Count(&count).Error
	return count, err
}

func (r *UserRepository) List(ctx context.Context, page, pageSize int, role *uint8, status *uint8, keyword string) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	query := r.db.WithContext(ctx).Model(&model.User{}).Where("deleted_at IS NULL")

	if role != nil {
		query = query.Where("role = ?", *role)
	}
	if status != nil {
		query = query.Where("status = ?", *status)
	}
	if keyword != "" {
		query = query.Where("username LIKE ? OR email LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Order("id DESC").Find(&users).Error
	return users, total, err
}

func (r *UserRepository) Follow(ctx context.Context, followerID, followingID uint64) error {
	follow := &model.Follow{
		FollowerID:  followerID,
		FollowingID: followingID,
	}
	return r.db.WithContext(ctx).Create(follow).Error
}

func (r *UserRepository) Unfollow(ctx context.Context, followerID, followingID uint64) error {
	return r.db.WithContext(ctx).Where("follower_id = ? AND following_id = ?", followerID, followingID).Delete(&model.Follow{}).Error
}

func (r *UserRepository) IsFollowing(ctx context.Context, followerID, followingID uint64) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Follow{}).Where("follower_id = ? AND following_id = ?", followerID, followingID).Count(&count).Error
	return count > 0, err
}

func (r *UserRepository) GetFollowingCount(ctx context.Context, userID uint64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Follow{}).Where("follower_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *UserRepository) GetFollowerCount(ctx context.Context, userID uint64) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Follow{}).Where("following_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *UserRepository) GetFollowingList(ctx context.Context, userID uint64, page, pageSize int) ([]model.User, int64, error) {
	var follows []model.Follow
	var total int64
	var users []model.User

	offset := (page - 1) * pageSize

	err := r.db.WithContext(ctx).Model(&model.Follow{}).Where("follower_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("follower_id = ?", userID).Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&follows).Error
	if err != nil {
		return nil, 0, err
	}

	if len(follows) == 0 {
		return users, total, nil
	}

	followingIDs := make([]uint64, len(follows))
	for i, f := range follows {
		followingIDs[i] = f.FollowingID
	}

	err = r.db.WithContext(ctx).Where("id IN ? AND deleted_at IS NULL", followingIDs).Find(&users).Error
	return users, total, err
}

func (r *UserRepository) GetFollowerList(ctx context.Context, userID uint64, page, pageSize int) ([]model.User, int64, error) {
	var follows []model.Follow
	var total int64
	var users []model.User

	offset := (page - 1) * pageSize

	err := r.db.WithContext(ctx).Model(&model.Follow{}).Where("following_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.WithContext(ctx).Where("following_id = ?", userID).Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&follows).Error
	if err != nil {
		return nil, 0, err
	}

	if len(follows) == 0 {
		return users, total, nil
	}

	followerIDs := make([]uint64, len(follows))
	for i, f := range follows {
		followerIDs[i] = f.FollowerID
	}

	err = r.db.WithContext(ctx).Where("id IN ? AND deleted_at IS NULL", followerIDs).Find(&users).Error
	return users, total, err
}
