package sys

import (
	"github.com/cjx2328/gocms/internal/pkg/models/basemodel"
	"github.com/jinzhu/gorm"
	"time"
)

type Adminlogs struct {
	basemodel.Model
	Status      uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"`
	Content        string `gorm:"column:content;size:255;" json:"content" form:"content"`
	URL         string `gorm:"column:url;size:72;" json:"url" form:"url"`
	Title        string `gorm:"column:title;size:32;not null;" json:"title" form:"title"`
	Code        string `gorm:"column:code;size:32;not null;" json:"code" form:"code"`
	IP        string `gorm:"column:ip;size:32;not null;" json:"ip" form:"ip"`

}
// 表名
func (Adminlogs) TableName() string {
	return TableName("adminlogs")
}

func (m *Adminlogs) BeforeCreate(scope *gorm.Scope) error{

	m.CreatedAt=time.Now()
	m.UpdatedAt=time.Now()

}

func