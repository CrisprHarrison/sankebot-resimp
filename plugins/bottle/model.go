package bottle

import (
	"time"

	"github.com/RicheyJang/PaimengBot/manager"
	log "github.com/sirupsen/logrus"
)

type DriftingBottleModel struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	FromID    int64     `gorm:"index"`
	Content   string    `gorm:"type:text"`
	ImagePath string    `gorm:"type:text"` // 新增字段存储图片路径
	CreatedAt time.Time `gorm:"index"`
}

func init() {
	err := manager.GetDB().AutoMigrate(&DriftingBottleModel{})
	if err != nil {
		log.Errorf("[SQL] Bottle DriftingBottleModel初始化失败，err: %v", err)
	}
}
