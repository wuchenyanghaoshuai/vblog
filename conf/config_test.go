package conf_test

import (
	"github.com/wuchenyanghaoshuai/vblog/conf"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C())
}


func TestLoadConfigFromEnv(t *testing.T){
	err := conf.LoadConfigFromEnv()
	if err !=nil{
		t.Fatal(err)
	}
	t.Log(conf.C())
}
