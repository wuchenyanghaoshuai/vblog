package conf

type Config struct {
	// 这个对象维护整个程序的配置
	Mysql *Mysql `json:"mysql"`
}

type Mysql struct {
	Host     string `json:"host"`
	DB       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}
