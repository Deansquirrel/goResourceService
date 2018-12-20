package global

type SysConfig struct {
	DbConfig dbConfig `toml:"DbConfig"`
}

type dbConfig struct {
	Server   string `toml:"server"`
	Port     int    `toml:"port"`
	DbName   string `toml:"db"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}
