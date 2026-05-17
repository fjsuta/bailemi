package tasks

import (
	"log"
	"time"

	"bailemi/gateway/internal/service"
)

type ScheduledTasks struct {
	updateService *service.UpdateService
}

func NewScheduledTasks(updateService *service.UpdateService) *ScheduledTasks {
	return &ScheduledTasks{
		updateService: updateService,
	}
}

// Start 启动定时任务
func (t *ScheduledTasks) Start() {
	// 每天随机时间检查更新
	go t.runUpdateCheck()
}

func (t *ScheduledTasks) runUpdateCheck() {
	for {
		now := time.Now()
		nextCheck := getNextRandomCheckTime(now)
		duration := nextCheck.Sub(now)
		
		log.Printf("下一次更新检查将在 %s 后执行", duration)
		
		time.Sleep(duration)
		
		log.Println("开始检查更新...")
		hasUpdate, message, err := t.updateService.CheckUpdate()
		if err != nil {
			log.Printf("更新检查失败: %v", err)
		} else if hasUpdate {
			log.Printf("发现新更新: %s", message)
		} else {
			log.Printf(message)
		}
	}
}

func getNextRandomCheckTime(base time.Time) time.Time {
	tomorrow := base.Add(24 * time.Hour)
	
	// 在明天的 00:00 到 24:00 之间随机选择一个时间
	hour := 2 + int(base.UnixNano())%6 // 2-7 点之间
	minute := int(base.UnixNano()) % 60
	second := int(base.UnixNano()) % 60
	
	return time.Date(
		tomorrow.Year(), tomorrow.Month(), tomorrow.Day(),
		hour, minute, second, 0,
		tomorrow.Location(),
	)
}
