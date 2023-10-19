package config

import (
	"117web/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB     //声明数据库指针
	Config *viper.Viper //声明配置viper指针
)

func IntiConfig() {
	Config = viper.New() /*为Config申请内存*/

	Config.SetConfigName("config") /*设置配置文件的信息*/
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./config")

	if err := Config.ReadInConfig(); err != nil { /*读取配置文件，出错打印初始化配置错误*/
		fmt.Printf("初始化配置错误")
		return
	}

	fmt.Println("初始化配置完成")
}

func InitDB() {
	dsn := Dsn() /*配置数据库url*/

	var err error /*连接数据库*/
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err.Error())
	}

	if DB == nil { /*判断数据库初始化是否成功*/
		// MySQL未初始化
		fmt.Println("MySQL未初始化")
	} else {
		// MySQL已初始化
		fmt.Println("\nMySQL已初始化:" + dsn + "\n")
	}

	sqlDB, err := DB.DB()

	sqlDB.SetMaxOpenConns(10) // 设置连接池中的最大连接数
	sqlDB.SetMaxIdleConns(5)  // 设置连接池中的最大空闲连接数

}
func Dsn() string { //gorm配置数据库url
	/*  dsn := "user:pass@tcp(host:port)/dbname?dev"*/
	user := Config.GetString("mysql.username") /*基本信息赋值*/
	pass := Config.GetString("mysql.password")
	host := Config.GetString("mysql.host")
	port := Config.GetString("mysql.port")
	dbname := Config.GetString("mysql.dbname")
	dev := Config.GetString("mysql.dev")

	return user + ":" + pass + "@tcp" + "(" + host + ":" + port + ")/" + dbname + "?" + dev
}

func Init() {
	IntiConfig()
	InitDB()
	DB.AutoMigrate(&model.Person{}, &model.User{})
}
