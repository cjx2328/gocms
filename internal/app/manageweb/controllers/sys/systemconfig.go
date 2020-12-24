package sys

import (
	"fmt"

	"github.com/cjx2328/gocms/internal/app/manageweb/controllers/common"
	models "github.com/cjx2328/gocms/internal/pkg/models/common"
	"github.com/cjx2328/gocms/internal/pkg/models/sys"


	"github.com/gin-gonic/gin"
)

type Systemconfig struct{}

// 获取系统配置列表
func (Systemconfig) List(c *gin.Context) {

  list:= sys.Systemconfig{}
  getlist := []sys.Systemconfig{}
  list.GetSystemconfigList(&sys.Systemconfig{} , &getlist)
	fmt.Println(&getlist)

  sendconfig :=  make(map[string]*sys.Systemconfig)

  for _,listdata := range getlist{
	  sendconfig[listdata.Title] = &listdata
  }

	fmt.Println(&sendconfig)

	common.ResSuccess(c,   &sendconfig)

	return
}
///添加和更新2个功能 如果不存在就是添加，如果已经存在就是更新
func (Systemconfig) Save(c *gin.Context){
		modelslist :=sys.Systemconfig{}
		err := c.Bind(&modelslist)
		if err !=nil {
			common.ResErrSrv(c,err)
			return
		}

		err = models.Save(modelslist)
		if err !=nil{
			common.ResErrSrv(c,err)
			return
		}
	common.ResSuccessMsg(c )
}



