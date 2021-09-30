package goenv

import (
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/env"
)

var globalEnv = gosupport.NewGlobalEnvSingleton()

func LoadEnv(envfile string) error {
	err := env.LoadEnv2DataManage(envfile)
	return err
}

func GetString(key string) string {
	return globalEnv.GetString(key)
}

func Get(key string) gosupport.StrTo  {
	return gosupport.StrTo(globalEnv.GetString(key))
}
