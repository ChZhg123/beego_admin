package conf

type Mysql struct {
	Host string `json "host"`
	Port int `json "port"`
	DataBase string `json "database"`
	Username string `json "username"`
	Password string `json "password"`
}