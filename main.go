package main

import (
	"encoding/json"
	"fmt"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goResourceService/global"
)

func main() {
	fmt.Println("程序启动")
	defer fmt.Println("程序退出")
	//=================================================================
	//获取配置
	sysConfig, err := global.GetSysConfig()
	if err != nil {
		fmt.Println("配置文件获取失败.[", err, "]")
		return
	}
	//=================================================================
	//配置输出
	configJson, err := json.Marshal(sysConfig)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(configJson))
	}
	//=================================================================

	fmt.Println(go_tool.Guid())

}
