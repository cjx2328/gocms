package sys

import (
	"github.com/cjx2328/gocms/internal/app/manageweb/controllers/common"
	"github.com/cjx2328/gocms/internal/pkg/models/sys"
	"github.com/gin-gonic/gin"
)

type Aadminlogs struct {

}

func (Aadminlogs) List(c *gin.Context){
	listdata := sys.Adminlogs{}
	 list := listdata.Status
	 common.ResSuccess(c , list)
}



