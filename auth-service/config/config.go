package config

type Config struct {
	Port string
	Database Database
}

type Database struct {
	Driver string `mapstructure:"driver"`
	Host string `mapstructure:"host"`
	Post string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName string `mapstrutcture:"db_name"`
	Config string `mapstructure:"config"`
}
