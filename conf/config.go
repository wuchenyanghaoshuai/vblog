package conf

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

func DefaultConfig() *Config {
	return &Config{
		Mysql: &Mysql{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "vblog",
			Username: "root",
			Password: "123456",
		},
		App: &App{
			HttpPort: 7080,
			HttpHost: "127.0.0.1",
		},

	}
}

func (c *Config) String() string {
	dj, _ := json.Marshal(c)
	return string(dj)
}

type App struct {
	HttpHost string `json:"http_host" toml:"http_host" env:"HTTP_HOST"`
	HttpPort int `json:"http_port" toml:"http_port" env:"HTTP_PORT"`
}

func (a *App) HttpAddr() string {
	return fmt.Sprintf("%s:%d", a.HttpHost, a.HttpPort)
}

type Config struct {
	// 这个对象维护整个程序的配置
	Mysql *Mysql `json:"mysql"`
	App   *App   `json:"app" toml:"app"`
}

type Mysql struct {
	Host     string `json:"host" toml:"host" env:"MYSQL_HOST"`
	Port     int    `json:"port" toml:"port" env:"MYSQL_PORT"`
	DB       string `json:"database" toml:"database" env:"MYSQL_DB"`
	Username string `json:"username" toml:"username" env:"MYSQL_USERNAME"`
	Password string `json:"password" toml:"password" env:"MYSQL_PASSWORD"`
	lock     sync.Mutex
	//缓存一个对象
	conn *gorm.DB
}

// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m *Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB,
	)
}

// 返回一个数据库链接，返回一个全局单例
func (m *Mysql) GetConn() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.conn == nil {
		conn, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		m.conn = conn
	}
	return m.conn
}
