package config

import (
	viper "github.com/spf13/viper"
	"os"
)

var Config *Conf

type Conf struct {
	Mysql  map[string]*Mysql `yaml:"mysql"`
	System *System           ` yaml:"system"`
}

type Mysql struct {
	Dialect  string `yaml:"dialect"`
	DbHost   string `yaml:"dbHost"`
	DbPort   string `yaml:"dbPort"`
	DbName   string `yaml:"dbName"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type System struct {
	Domain   string ` yaml:"domain"`
	Version  string `yaml:"version"`
	AppEnv   string `yaml:"appEnv"`
	HttpPort string `yaml:"HttpPort"`
	Host     string `yaml:"Host"`
}

func InitConfig() {
	workDir, _ := os.Getwd()                       // 获取工作目录
	viper.SetConfigName("config")                  // 设置要搜寻的配置文件名称
	viper.SetConfigType("yaml")                    // 设置要搜寻得配置文件类型
	viper.AddConfigPath(workDir + "/config/local") // 设置搜寻配置文件得路径
	viper.AddConfigPath(workDir)                   // 同上
	err := viper.ReadInConfig()                    //读取
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config) // 将信息绑定到变量中
	if err != nil {
		panic(err)
	}
}
