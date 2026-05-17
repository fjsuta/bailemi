package handler

import (
	"net/http"

	"bailemi/gateway/internal/model"
	"bailemi/gateway/internal/service"
	"bailemi/gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	updateService   *service.UpdateService
	storageService  *service.StorageService
	icpService      *service.ICPService
}

func NewAdminHandler(update *service.UpdateService, storage *service.StorageService, icp *service.ICPService) *AdminHandler {
	return &AdminHandler{
		updateService:  update,
		storageService: storage,
		icpService:     icp,
	}
}

// CheckUpdate 检查更新
func (h *AdminHandler) CheckUpdate(c *gin.Context) {
	hasUpdate, message, err := h.updateService.CheckUpdate()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "UPDATE_CHECK_FAILED", err.Error())
		return
	}

	response.Success(c, gin.H{
		"has_update": hasUpdate,
		"message":    message,
		"version":    h.updateService.GetCurrentVersion(),
	})
}

// DoUpdate 执行更新
func (h *AdminHandler) DoUpdate(c *gin.Context) {
	err := h.updateService.DoUpdate()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "UPDATE_FAILED", err.Error())
		return
	}

	response.Success(c, gin.H{
		"message": "更新成功，请重启服务",
	})
}

// GetUpdateLogs 获取更新日志
func (h *AdminHandler) GetUpdateLogs(c *gin.Context) {
	page := getIntQuery(c, "page", 1)
	pageSize := getIntQuery(c, "page_size", 20)

	logs, total, err := h.updateService.GetUpdateLogs(page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "GET_LOGS_FAILED", err.Error())
		return
	}

	response.SuccessWithPagination(c, logs, page, pageSize, total)
}

// GetICPInfo 获取备案信息
func (h *AdminHandler) GetICPInfo(c *gin.Context) {
	info, err := h.icpService.GetInfo()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "GET_ICP_FAILED", err.Error())
		return
	}

	response.Success(c, info)
}

// SaveICPInfo 保存备案信息
func (h *AdminHandler) SaveICPInfo(c *gin.Context) {
	var info model.ICPInfo
	if err := c.ShouldBindJSON(&info); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_PARAMS", err.Error())
		return
	}

	if err := h.icpService.SaveInfo(&info); err != nil {
		response.Error(c, http.StatusInternalServerError, "SAVE_ICP_FAILED", err.Error())
		return
	}

	response.Success(c, gin.H{
		"message": "保存成功",
	})
}

// GetStorageConfig 获取存储配置
func (h *AdminHandler) GetStorageConfig(c *gin.Context) {
	config := h.storageService.GetConfig()
	response.Success(c, config)
}

// SaveStorageConfig 保存存储配置
func (h *AdminHandler) SaveStorageConfig(c *gin.Context) {
	var config model.StorageConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		response.Error(c, http.StatusBadRequest, "INVALID_PARAMS", err.Error())
		return
	}

	if err := h.storageService.UpdateConfig(&config); err != nil {
		response.Error(c, http.StatusInternalServerError, "SAVE_STORAGE_FAILED", err.Error())
		return
	}

	response.Success(c, gin.H{
		"message": "保存成功",
	})
}
