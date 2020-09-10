package global

// Configuration 配置
type Configuration struct {
	// AppName 当前应用的名称
	AppName string
	Port    int
	// 日志记录级别。默认为 warn
	// 可配置级别有： panic fatal error warn info debug trace
	LogLevel string
	// APIPrefix API前缀。请以 / 开头配置此项目
	// 默认值为 /api
	APIPrefix string
	DBConf
	JwtConfig
}

// DBConf 数据库连接配置
type DBConf struct {
	Username, Password, Hostname, Database, Params string
	Debug                                          bool
}

// JwtConfig 生成jwt的配置
type JwtConfig struct {
	Secret string
	Iss    string
	// Exp 过期时间（分钟）。务必写入正整数
	Exp int
}
