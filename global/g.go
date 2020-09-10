package global

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 全局变量声明
var (
	// Router 全局路由实例
	Router *gin.Engine
	// Vreader 配置文件读取组件(viper.Viper)
	Vreader *viper.Viper
	Conf    Configuration

	AppPath string
	MyDB    *gorm.DB

	MyLog *logrus.Logger
	// Enforcer casbin的权限认证器
	Enforcer *casbin.Enforcer
)

// ------------------ 全局变量初始化方法 -----------------------

// InitViper 初始化配置文件读取工具
func InitViper() Configuration {

	Vreader = viper.New()
	// 一些配置的默认值
	Vreader.SetDefault("Port", "8080")
	Vreader.SetDefault("LogLevel", "warn")
	Vreader.SetDefault("APIPrefix", "/api")
	Vreader.SetConfigName("app")
	Vreader.SetConfigType("toml")
	Vreader.AddConfigPath(AppPath)
	if err := Vreader.ReadInConfig(); err != nil {
		fmt.Println(err)
		panic("读取配置文件失败！")
	}
	err := Vreader.Unmarshal(&Conf)
	if err != nil {
		fmt.Println(err)
		panic("不能把配置文件内容序列化到结构体中！")
	}
	return Conf
}

// InitLog 初始化日志组件
func InitLog(conf Configuration) *logrus.Logger {
	MyLog = logrus.New()
	lv, err := logrus.ParseLevel(Conf.LogLevel)
	if err != nil {
		fmt.Printf("设置日志记录级别失败，使用默认 warn 级别。%s \n", err)
		MyLog.SetLevel(logrus.WarnLevel)
	} else {
		MyLog.SetLevel(lv)
	}
	MyLog.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	}
	MyLog.SetOutput(os.Stdout)
	MyLog.Debugln("日志组件初始化成功！当前日志级别为 " + MyLog.Level.String())
	return MyLog
}

// InitGin 初始化 gin 配置
func InitGin() *gin.Engine {
	Router = gin.Default()
	return Router
}

// InitDB 初始化数据库连接
func InitDB() *gorm.DB {
	dbConf := Conf.DBConf
	// s should like "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		dbConf.Username, dbConf.Password,
		dbConf.Hostname, dbConf.Database,
		dbConf.Params)
	sqlLogger := logger.New(
		// MyLog, // logrus 组件
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 关闭多彩日志，特别是输出到其他日志组件时更要关闭
		},
	)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		Logger: sqlLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名单数
			TablePrefix:   "",   // 表前缀
		},
	})
	if err != nil {
		// gorm v2 不推荐使用db.Close()方法了
		// 其方法调用隐藏的更深了。
		// 暂时不Close()看看吧
		// _db, _ := MyDB.DB()
		// _db.Close()
		fmt.Printf("数据库连接失败！连接配置为： %s, 错误为： %s \n", s, err)
		panic("数据库连接失败！")
	}
	// 是否开启debug模式
	if dbConf.Debug {
		MyDB = db.Debug()
	} else {
		MyDB = db
	}
	MyLog.Debugln("数据库连接成功！")

	return MyDB
}

// InitCasbin 初始化 Casbin
func InitCasbin() *casbin.Enforcer {
	a, err := gormadapter.NewAdapterByDB(MyDB)
	if err != nil {
		MyLog.Fatalln("casbin 适配器初始化失败！ ", err)
	}
	e, err := casbin.NewEnforcer(AppPath+"/models.conf", a)
	if err != nil {
		MyLog.Fatalln("casbin Enforcer 初始化失败！ ", err)
	} else {
		MyLog.Debug(" Casbin 初始化成功！")
	}
	Enforcer = e
	return e
}
