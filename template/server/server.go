package server

import "emoney.cn/createcode/template"

func init(){
	template.RegisterTemplate(new(Server))
}


type Server struct{}

func(t *Server) Path() string{
	return "server"
}

func (t *Server) FileName() string{
	return "server.go"
}

func (t *Server) String() string {
	return `package server

import (
	"{project}/global"
	_ "fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/config"
	"github.com/devfeel/middleware/cors"
	"strconv"
)

func StartServer(configPath string) error {
	//初始化DotServer
	global.DotApp = dotweb.Classic()

	appConfig := config.MustInitConfig(configPath + "/dotweb.conf")
	global.DotApp.SetConfig(appConfig)
	global.DotApp.SetDevelopmentMode()
	global.DotApp.UseRequestLog()
	global.DotApp.Use(cors.Middleware(cors.NewConfig().UseDefault()))

	//设置路由
	InitRoute(global.DotApp.HttpServer)

	global.InnerLogger.Debug("dotweb.StartServer => " + strconv.Itoa(appConfig.Server.Port))
	err := global.DotApp.Start()
	if err != nil {
		panic(err)
	}
	return err
}`
}


