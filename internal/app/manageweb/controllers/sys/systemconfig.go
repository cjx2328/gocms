package sys

import (
	"encoding/json"
	"time"

	"github.com/cjx2328/gocms/internal/app/manageweb/controllers/common"
	models "github.com/cjx2328/gocms/internal/pkg/models/common"
	"github.com/cjx2328/gocms/internal/pkg/models/sys"
	"github.com/cjx2328/gocms/pkg/cache"
	"github.com/cjx2328/gocms/pkg/convert"
	"github.com/cjx2328/gocms/pkg/hash"
	"github.com/cjx2328/gocms/pkg/jwt"
	"github.com/cjx2328/gocms/pkg/logger"
	"github.com/cjx2328/gocms/pkg/util"


	"github.com/gin-gonic/gin"
)

type Systemconfig struct{}

// 获取系统配置列表
func (Systemconfig) List(c *gin.Context) {
  list:= sys.Systemconfig{}
  err :=list.GetSystemconfigList
  if err != nil{
  	common.ResErrSrv(c)
  	return
  }
 common.ResSuccessPage(c,total,&list)


}

// 用户登出
func (Systemconfig) Saves(c *gin.Context) {
	t := c.GetHeader(common.TOKEN_KEY)
	if t == "" {
		common.ResFail(c, "操作失败")
		return
	}
	u,ok:=jwt.ParseToken(t)
	if !ok {
		common.ResFail(c, "操作失败")
		return
	}
	cid:=u["uuid"]
	if cid == "" {
		common.ResFail(c, "操作失败")
		return
	}
	cache.Del([]byte(cid))
	common.ResSuccessMsg(c)
}

