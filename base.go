package commbase

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/json-iterator/go"
)

type BaseController struct {
	beego.Controller
}

var ConfigInfo config.Configer
var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	ConfigInfo = beego.AppConfig
}
