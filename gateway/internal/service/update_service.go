package service

import (
	"encoding/json"
	"os"
	"os/exec"
	"time"

	"bailemi/gateway/internal/model"
	"bailemi/gateway/internal/repository"
)

type UpdateService struct {
	repo         *repository.UpdateRepository
	repoURL      string
	checkEnabled bool
}

func NewUpdateService(repo *repository.UpdateRepository) *UpdateService {
	url := os.Getenv("UPDATE_REPO_URL")
	if url == "" {
		url = "https://github.com/fjsuta/bailemi"
	}
	enabled := os.Getenv("UPDATE_CHECK_ENABLE") != "false"

	return &UpdateService{
		repo:         repo,
		repoURL:      url,
		checkEnabled: enabled,
	}
}

// CheckUpdate 检查更新
func (s *UpdateService) CheckUpdate() (bool, string, error) {
	if !s.checkEnabled {
		return false, "更新检查已禁用", nil
	}

	// 执行 git fetch 检查远程更新
	cmd := exec.Command("git", "fetch", "origin")
	cmd.Dir = getProjectRoot()
	err := cmd.Run()
	if err != nil {
		return false, "", err
	}

	// 检查本地和远程是否有差异
	cmd = exec.Command("git", "rev-list", "--count", "HEAD...origin/master")
	cmd.Dir = getProjectRoot()
	output, err := cmd.Output()
	if err != nil {
		return false, "", err
	}

	count := string(output)
	if count != "0" {
		return true, "发现新的更新", nil
	}

	return false, "当前已是最新版本", nil
}

// GetCurrentVersion 获取当前版本
func (s *UpdateService) GetCurrentVersion() string {
	cmd := exec.Command("git", "log", "-1", "--format=%h")
	cmd.Dir = getProjectRoot()
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return string(output)
}

// DoUpdate 执行更新
func (s *UpdateService) DoUpdate() error {
	log := &model.UpdateLog{
		Version:     "latest",
		Description: "自动更新",
		Status:      "running",
	}
	if err := s.repo.CreateLog(log); err != nil {
		return err
	}

	now := time.Now()
	log.StartedAt = &now
	s.repo.UpdateLog(log)

	// 拉取代码
	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Dir = getProjectRoot()
	err := cmd.Run()
	if err != nil {
		log.Status = "failed"
		log.Error = err.Error()
		s.repo.UpdateLog(log)
		return err
	}

	// 成功
	log.Status = "completed"
	completedAt := time.Now()
	log.CompletedAt = &completedAt
	s.repo.UpdateLog(log)

	return nil
}

// GetUpdateLogs 获取更新日志
func (s *UpdateService) GetUpdateLogs(page, pageSize int) ([]model.UpdateLog, int64, error) {
	return s.repo.GetLogs(page, pageSize)
}

func getProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	return dir
}
