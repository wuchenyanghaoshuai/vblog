package conf

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func DefaultConfig()*Config {
	return &Config{
		Application: &Application{
			Host: "127.0.0.1",
			Port: 8088,
			Domain: "http://127.0.0.1",
		},
		MySQL: &MySQL{
			Host: "127.0.0.1",
			Port: 3307,
			DB: "go15",
			Username: "root",
			Password: "123456",
			Debug: true,
		},
	}
}


type Config struct {
	Application *Application `toml:"app" yaml:"app" json:"app"`
	MySQL *MySQL `toml:"mysql" yaml:"mysql" json:"mysql"`
}


func (a *Application) GinRootRouter() gin.IRouter {
	r := a.GinServer()
	if a.root == nil {
		a.root = r.Group("vblog").Group("api").Group("v1")
	}
	return a.root
}

func (a *Application) GinServer() *gin.Engine {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.server == nil {
		a.server = gin.Default()
	}
	return a.server
}

func (a *Application) Address() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func (a *Application) Start() error {
	r := a.GinServer()
	return r.Run(a.Address())
}
//应用服务
// 比如服务监听端口
type Application struct {
	Host string `toml:"host" yaml:"host" json:"host"`
	Port int `toml:"port" yaml:"port" json:"port"`
	Domain string `toml:"domain" yaml:"domain" json:"domain"`
	server *gin.Engine
	lock sync.Mutex
	root gin.IRouter
}

//数据库
type MySQL struct {
	Host string `toml:"host" yaml:"host" json:"host" env:"DATASOURCE_HOST"`
	Port int `toml:"port" yaml:"port" json:"port" env:"DATASOURCE_PORT"`
	DB string `toml:"database" yaml:"database" json:"database" env:"DATASOURCE_DB"`
	Username string `toml:"username" yaml:"username" json:"username" env:"DATASOURCE_USERNAME"`
	Password string `toml:"password" yaml:"password" json:"password" env:"DATASOURCE_PASSWORD"`
	Debug bool `toml:"debug" yaml:"debug" json:"debug" env:"DATASOURCE_DEBUG"`
	db *gorm.DB
	lock sync.Mutex
}


func (m *MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.DB)
}

func (m *MySQL) GetDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()
	if m.db == nil {
		db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{})
		
		if err != nil {
			panic(err)
		}
		m.db = db
		if m.Debug {
			m.db = m.db.Debug()
		}
	}
	return m.db
}