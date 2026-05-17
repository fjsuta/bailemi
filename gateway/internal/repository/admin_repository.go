package repository

import (
	"bailemi/gateway/internal/model"

	"gorm.io/gorm"
)

type UpdateRepository struct {
	db *gorm.DB
}

func NewUpdateRepository(db *gorm.DB) *UpdateRepository {
	return &UpdateRepository{db: db}
}

func (r *UpdateRepository) CreateLog(log *model.UpdateLog) error {
	return r.db.Create(log).Error
}

func (r *UpdateRepository) UpdateLog(log *model.UpdateLog) error {
	return r.db.Save(log).Error
}

func (r *UpdateRepository) GetLogs(page, pageSize int) ([]model.UpdateLog, int64, error) {
	var logs []model.UpdateLog
	var total int64
	
	offset := (page - 1) * pageSize
	
	err := r.db.Model(&model.UpdateLog{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	
	err = r.db.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&logs).Error
	return logs, total, err
}

type ICPRepository struct {
	db *gorm.DB
}

func NewICPRepository(db *gorm.DB) *ICPRepository {
	return &ICPRepository{db: db}
}

func (r *ICPRepository) GetInfo() (*model.ICPInfo, error) {
	var info model.ICPInfo
	err := r.db.FirstOrCreate(&info, model.ICPInfo{}).Error
	return &info, err
}

func (r *ICPRepository) SaveInfo(info *model.ICPInfo) error {
	if info.ID == 0 {
		return r.db.Create(info).Error
	}
	return r.db.Save(info).Error
}
