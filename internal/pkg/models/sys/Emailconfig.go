package sys

type Emailconfig struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id" form:"id"`                    // 主键
	Status      uint8  `gorm:"column:status;type:int(1);not null;" json:"status" form:"status"`             // 状态(1:启用 2:不启用)
	Title        string `gorm:"column:title;size:64;" json:"title" form:"title"`                                    // 备注
	Content         string `gorm:"column:content;size:72;" json:"content" form:"content"`
}
