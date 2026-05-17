package service

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"bailemi/gateway/internal/model"
)

type StorageService struct {
	config *model.StorageConfig
}

func NewStorageService(configPath string) (*StorageService, error) {
	config, err := loadStorageConfig(configPath)
	if err != nil {
		config = getDefaultStorageConfig()
	}
	return &StorageService{config: config}, nil
}

func (s *StorageService) UploadFile(file multipart.File, filename string) (string, error) {
	switch s.config.Type {
	case "local":
		return s.uploadLocal(file, filename)
	case "aliyun_oss":
		return s.uploadAliyunOSS(file, filename)
	case "tencent_cos":
		return s.uploadTencentCOS(file, filename)
	case "qiniu":
		return s.uploadQiniu(file, filename)
	case "upyun":
		return s.uploadUpyun(file, filename)
	case "huawei_obs":
		return s.uploadHuaweiOBS(file, filename)
	case "baidu_bos":
		return s.uploadBaiduBOS(file, filename)
	default:
		return s.uploadLocal(file, filename)
	}
}

func (s *StorageService) uploadLocal(file multipart.File, filename string) (string, error) {
	uploadPath := s.config.Local.Path
	if uploadPath == "" {
		uploadPath = "./uploads"
	}
	
	if err := os.MkdirAll(uploadPath, 0755); err != nil {
		return "", err
	}

	dst, err := os.Create(filepath.Join(uploadPath, filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return "/uploads/" + filename, nil
}

func (s *StorageService) uploadAliyunOSS(file multipart.File, filename string) (string, error) {
	// 阿里云OSS上传实现（占位）
	return "", nil
}

func (s *StorageService) uploadTencentCOS(file multipart.File, filename string) (string, error) {
	// 腾讯云COS上传实现（占位）
	return "", nil
}

func (s *StorageService) uploadQiniu(file multipart.File, filename string) (string, error) {
	// 七牛云上传实现（占位）
	return "", nil
}

func (s *StorageService) uploadUpyun(file multipart.File, filename string) (string, error) {
	// 又拍云上传实现（占位）
	return "", nil
}

func (s *StorageService) uploadHuaweiOBS(file multipart.File, filename string) (string, error) {
	// 华为云OBS上传实现（占位）
	return "", nil
}

func (s *StorageService) uploadBaiduBOS(file multipart.File, filename string) (string, error) {
	// 百度云BOS上传实现（占位）
	return "", nil
}

func (s *StorageService) GetConfig() *model.StorageConfig {
	return s.config
}

func (s *StorageService) UpdateConfig(config *model.StorageConfig) error {
	s.config = config
	return saveStorageConfig("./config/storage.json", config)
}

func loadStorageConfig(path string) (*model.StorageConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config model.StorageConfig
	err = json.Unmarshal(data, &config)
	return &config, err
}

func saveStorageConfig(path string, config *model.StorageConfig) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func getDefaultStorageConfig() *model.StorageConfig {
	return &model.StorageConfig{
		Type: "local",
		Local: model.LocalStorage{
			Path: "./uploads",
		},
	}
}
