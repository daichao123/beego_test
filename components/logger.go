package components

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/adapter/config"
	"github.com/beego/beego/v2/core/logs"
	"log"
)

func InitLog() (error error) {
	newConfig, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		log.Fatal(err)
		return
	}
	maxLines, lenerr := newConfig.Int64("log:maxlines")
	if lenerr != nil {
		maxLines = 1000
	}
	logConf := make(map[string]interface{})
	logConf["filename"] = newConfig.String("log::log_path")
	port, _ := newConfig.Int("server::listen_port")
	level, _ := newConfig.Int("log::log_level")
	logConf["level"] = level
	logConf["maxlines"] = maxLines
	configStr, err1 := json.Marshal(logConf)
	if err1 != nil {
		fmt.Println("marshal failed,err:", err1)
		return
	}
	fmt.Println("port:", port)
	error1 := logs.SetLogger(logs.AdapterFile, string(configStr))
	if error1 != nil {
		log.Fatal(error1)
	}
	logs.SetLogFuncCall(true)
	return nil
}
