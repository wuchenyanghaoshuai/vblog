package conf_test

import (
	"os"
	"testing"
	"vblog/apps/conf"

	"github.com/go-playground/assert/v2"
)


func TestToYAML(t *testing.T) {
	config := conf.DefaultConfig().ToYAML()
	t.Log(config)
}

func TestLoadFromEnv(t *testing.T) {
	os.Setenv("DATASOURCE_USERNAME","wuchenyang")

	conf.LoadConfigFromEnv()
	t.Log(conf.C().ToYAML())
	//assert 进行断言 对比两个env是否相等
	assert.Equal(t,conf.C().MySQL.Username,"wuchenyang")
}

func TestLoadFromYAML(t *testing.T) {
	
	conf.LoadConfigFromYaml("/Users/wuchenyang/code/gocode/vblog/apps/conf/application.yml")
	t.Log(conf.C().ToYAML())
}