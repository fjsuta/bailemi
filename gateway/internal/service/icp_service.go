package service

import (
	"bailemi/gateway/internal/model"
	"bailemi/gateway/internal/repository"
)

type ICPService struct {
	repo *repository.ICPRepository
}

func NewICPService(repo *repository.ICPRepository) *ICPService {
	return &ICPService{repo: repo}
}

func (s *ICPService) GetInfo() (*model.ICPInfo, error) {
	return s.repo.GetInfo()
}

func (s *ICPService) SaveInfo(info *model.ICPInfo) error {
	return s.repo.SaveInfo(info)
}
