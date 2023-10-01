package route

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	workdir, _ := os.Getwd()
	logfile, err := os.Create(workdir + "/logs/gin_http.log")
	if err != nil {
		fmt.Printf("配置ginhttp日志失败:%v", err)
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(logfile)
	r := gin.Default()
	return r
}
