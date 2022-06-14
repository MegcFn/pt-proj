package util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
)

var _db *gorm.DB

type Conf struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	DBname   string `yaml:"DBname"`
	Timeout  string `yaml:"timeout"`
}

func init() {
	//dsn configuration
	conf := new(Conf)
	var err error
	conf.getConf()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DBname, conf.Timeout)
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败, error=" + err.Error())
	}
	sqlDB, _ := _db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)
}

func GetDB() *gorm.DB {
	return _db
}

// The function for reading yamlfile
func (c *Conf) getConf() {
	//Read filepath
	yamlFile, err := ioutil.ReadFile("config/DBconfig.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, c)
}
