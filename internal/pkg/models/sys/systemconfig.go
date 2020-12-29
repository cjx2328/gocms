package sys

import (
	"fmt"
	"github.com/cjx2328/gocms/internal/pkg/models/db"

	"github.com/jinzhu/gorm"
)

// 系统配置数据库表
type Systemconfig struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                    // 主键

	Status      uint8  `gorm:"column:status;type:int(1);not null;" json:"status" form:"status"`             // 状态(1:启用 2:不启用)
	Title        string `gorm:"column:title;size:64;" json:"title" form:"title"`                                    // 备注
	Content         string `gorm:"column:content;size:72;" json:"content" form:"content"`                                       // 菜单URL
}


// 系统配置保存表单model
type Systemconfigsave struct {

	Title        string `gorm:"column:title;size:64;" json:"title" form:"title"`                                    // 备注
	Content         string `gorm:"column:content;size:72;" json:"content" form:"content"`                                       // 菜单URL
}




type Systemconfigjson struct {
	Title string
	Content string
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

// 获取配置列表
func (Systemconfig) GetSystemconfigList(model ,out interface{} ) (err error) {
	//sql := `select * from tb_sys_systemconfig`
	db := db.DB.Model(model)
	return db.Find(out ).Error
}

// 获取列表
func (Systemconfig) Savesytemconfig(  getitemlists []Systemconfig,out interface{})(err error){

	db := db.DB.Model(&Systemconfig{})
	verifydata := &Systemconfig{}
	for _,items := range getitemlists{
		fmt.Println(items)
		wherestring := " title='" + items.Title +"'"
		errcheck := db.Find(verifydata,wherestring).Error
		fmt.Print("find db error\n\r")
		fmt.Println(errcheck)
		if errcheck != nil {
			fmt.Print("check db error\n\r")
			fmt.Println(err)
			err1 :=db.Create(&items)
			fmt.Print("create db error \n\r")
			fmt.Println(err1)
		}else{
			err1 :=db.Where(wherestring).Update(&items)
			fmt.Print("update db error\n\r")
			fmt.Println(err1)
		}
		fmt.Println(verifydata)
	}

	return err

}






