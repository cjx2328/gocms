package sys

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"os"

	"path"
	"strconv"
	"strings"
	"time"

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
	fmt.Println(getlist)

  sendconfig :=  make(map[string]sys.Systemconfig)

  for _,listdata := range getlist{
	  sendconfig[listdata.Title] =  listdata
	  fmt.Println(listdata)
  }

	fmt.Println(&sendconfig)

	common.ResSuccess(c,   sendconfig)

	return
}
///添加和更新2个功能 如果不存在就是添加，如果已经存在就是更新
func (Systemconfig) Save(c *gin.Context){
		modelslist := []sys.Systemconfig{}
		err := c.Bind(&modelslist)
		if err !=nil {
			common.ResErrSrv(c,err)
			return
		}
	fmt.Println(&modelslist)
	fmt.Println(modelslist)
	getlist := []sys.Systemconfig{}
		systemconfigs :=sys.Systemconfig{}
		err =systemconfigs.Savesytemconfig(modelslist,getlist)
		if err !=nil{
			common.ResErrSrv(c,err)
			return
		}
	common.ResSuccessMsg(c )
}

/// 上传图片
func (Systemconfig)  SaveUploadedFile(c *gin.Context ) {

	headers ,err := c.FormFile("file")
	fmt.Println(err)
	if err != nil{
		fmt.Println(err)
	}
	crutime := time.Now().Unix() + rand.Int63n(10)
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime , 10))
	token := fmt.Sprintf("%x", h.Sum(nil))

	//dst :=headers.Filename
	dst := path.Join("./upload", headers.Filename)

	fmt.Println(dst)
	if err := c.SaveUploadedFile(headers,dst) ; err != nil{
		fmt.Println(err)
	}

	fmt.Println("fullFilename =", headers.Filename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(headers.Filename)
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)
	fmt.Println("fileSuffix =", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	fmt.Println("filenameOnly =", filenameOnly)
	newdst :=path.Join("./upload" , token+""+fileSuffix)

	errrename :=os.Rename(dst , newdst)

	if errrename !=nil{
		fmt.Println(errrename)
	}

    common.ResSuccess(c,   newdst)
}


