package sys

import (
	"fmt"

	"github.com/cjx2328/gocms/internal/app/manageweb/controllers/common"

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

func (Systemconfig) Save(c *gin.Context){
		models :=sys.Systemconfig{}
		err := c.Bind(&models)
		if err !=nil {
			common.ResErrSrv(c,err)
			return
		}
	common.ResSuccess(c , &models)
}



