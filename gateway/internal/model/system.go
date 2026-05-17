package model

import "time"

// SystemConfig 系统配置表
type SystemConfig struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	Key       string    `json:"key" gorm:"uniqueIndex;not null"`
	Value     string    `json:"value" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ICPInfo 公安备案信息表
type ICPInfo struct {
	ID             uint64    `json:"id" gorm:"primaryKey"`
	ICPNumber      string    `json:"icp_number" gorm:"size:50"`
	PSBRecord      string    `json:"psb_record" gorm:"size:50"`
	Domain         string    `json:"domain" gorm:"size:255"`
	CompanyName    string    `json:"company_name" gorm:"size:255"`
	LegalPerson    string    `json:"legal_person" gorm:"size:100"`
	Contact        string    `json:"contact" gorm:"size:100"`
	ContactPhone   string    `json:"contact_phone" gorm:"size:20"`
	Address        string    `json:"address" gorm:"size:500"`
	ICPImageURL    string    `json:"icp_image_url" gorm:"size:500"`
	PSBImageURL    string    `json:"psb_image_url" gorm:"size:500"`
	Status         string    `json:"status" gorm:"size:20;default:'pending'"`
	Remark         string    `json:"remark" gorm:"type:text"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// UpdateLog 更新日志表
type UpdateLog struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	Version     string    `json:"version" gorm:"size:50"`
	Description string    `json:"description" gorm:"type:text"`
	Status      string    `json:"status" gorm:"size:20;default:'pending'"`
	StartedAt   *time.Time `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at"`
	Error       string    `json:"error" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
}

// StorageConfig 存储配置
type StorageConfig struct {
	Type         string        `json:"type"`
	Local        LocalStorage  `json:"local"`
	AliyunOSS    AliyunConfig  `json:"aliyun_oss"`
	TencentCOS   TencentConfig `json:"tencent_cos"`
	Qiniu        QiniuConfig   `json:"qiniu"`
	Upyun        UpyunConfig   `json:"upyun"`
	HuaweiOBS    HuaweiConfig  `json:"huawei_obs"`
	BaiduBOS     BaiduConfig   `json:"baidu_bos"`
}

type LocalStorage struct {
	Path string `json:"path"`
}

type AliyunConfig struct {
	Enabled         bool   `json:"enabled"`
	AccessKeyID     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Bucket          string `json:"bucket"`
	Region          string `json:"region"`
	Endpoint        string `json:"endpoint"`
}

type TencentConfig struct {
	Enabled    bool   `json:"enabled"`
	SecretID   string `json:"secret_id"`
	SecretKey  string `json:"secret_key"`
	Bucket     string `json:"bucket"`
	Region     string `json:"region"`
}

type QiniuConfig struct {
	Enabled    bool   `json:"enabled"`
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
	Bucket     string `json:"bucket"`
	Domain     string `json:"domain"`
}

type UpyunConfig struct {
	Enabled    bool   `json:"enabled"`
	Bucket     string `json:"bucket"`
	Operator   string `json:"operator"`
	Password   string `json:"password"`
	Domain     string `json:"domain"`
}

type HuaweiConfig struct {
	Enabled      bool   `json:"enabled"`
	AccessKey    string `json:"access_key"`
	SecretKey    string `json:"secret_key"`
	Bucket       string `json:"bucket"`
	Region       string `json:"region"`
	Endpoint     string `json:"endpoint"`
}

type BaiduConfig struct {
	Enabled      bool   `json:"enabled"`
	AccessKey    string `json:"access_key"`
	SecretKey    string `json:"secret_key"`
	Bucket       string `json:"bucket"`
	Region       string `json:"region"`
	Endpoint     string `json:"endpoint"`
}
