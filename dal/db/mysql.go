package db

import (
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	base = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	//获取dbconf
	mysqlConf := getDbConf()
	//构造dsn
	dsn := fmt.Sprintf(base, mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Dbname)
	//连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true}) //关闭默认事务提高性能
	if err != nil {
		panic(err)
	}
	log.Info("mysql连接成功")

}

func getDbConf() mysqlConf {
	viper.SetConfigName("db_conf")
	viper.AddConfigPath("./config")
	if err = viper.ReadInConfig(); err != nil {
		log.Error(errors.WithStack(err))
		panic("viper readInConfig error")
	}
	var dbconf conf
	if err = viper.Unmarshal(&dbconf); err != nil {
		log.Error(errors.WithStack(err))
		panic("viper Unmarshal error")
	}
	return dbconf.Mysql
}

type conf struct {
	Mysql mysqlConf `json:"mysql"`
}
type mysqlConf struct {
	Dbname   string `json:"dbname"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}
