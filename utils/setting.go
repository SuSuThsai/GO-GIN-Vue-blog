package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppNode  string
	HttpPort string
	JwtKey   string

	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassWorld string
	DbName      string

	AccessKey string
	SecretKey string
	Bucket    string
	Server    string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadDatabase(file)
	LoadGinIu(file)
}

func LoadServer(file *ini.File) {
	AppNode = file.Section("server").Key("AppNode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("255149331")
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("mysql")
	DbPassWorld = file.Section("database").Key("DbPassWorld").MustString("")
	DbName = file.Section("database").Key("DbName").MustString("root")
}

func LoadGinIu(file *ini.File) {
	AccessKey = file.Section("giniu").Key("AccessKey").String()
	SecretKey = file.Section("giniu").Key("SecretKey").String()
	Bucket = file.Section("giniu").Key("Bucket").String()
	Server = file.Section("giniu").Key("Server").String()
}
