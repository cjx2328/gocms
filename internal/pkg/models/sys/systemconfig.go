package sys

import (
	"github.com/cjx2328/gocms/internal/pkg/models/basemodel"
	"github.com/cjx2328/gocms/internal/pkg/models/db"

	"github.com/jinzhu/gorm"
)

// 菜单
type Systemconfig struct {
	basemodel.Model
	Status      uint8  `gorm:"column:status;type:tinyint(1);not null;" json:"status" form:"status"`             // 状态(1:启用 2:不启用)
	Title        string `gorm:"column:memo;size:64;" json:"title" form:"title"`                                    // 备注
	Content         string `gorm:"column:url;size:72;" json:"content" form:"content"`                                       // 菜单URL
}

// 表名
func (Systemconfig) TableName() string {
	return TableName("systemconfig")
}

// 添加前
func (m *Systemconfig) BeforeCreate(scope *gorm.Scope) error {

	return nil
}

// 更新前
func (m *Systemconfig) BeforeUpdate(scope *gorm.Scope) error {

	return nil
}

// 获取菜单有权限的操作列表
func (Systemconfig) GetSystemconfigList() (err error) {
	sql := `select * from tb_sys_systemconfig`
	err = db.DB.Raw(sql ).Error
	return
}

// 获取列表
func (Systemconfig) Savesytemconfig(){

}




