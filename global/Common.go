package global

import "github.com/BurntSushi/toml"

func GetSysConfig() (sysConfig SysConfig,err error){
	_,err = toml.DecodeFile("config.toml",&sysConfig)
	return
}
