package conf

type Server struct {
	Base Base `json "base"`
	Mysql Mysql `json "mysql"`
}