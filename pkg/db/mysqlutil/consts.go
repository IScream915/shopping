package mysqlutil

const (
	DefaultHost                      = "127.0.0.1"
	DefaultUser                      = "root"
	DefaultPort                      = 3306
	DefaultMaxOpenConns              = 1000
	DefaultMaxIdleConns              = 100
	DefaultMaxLifetime               = 3600
	DefaultTablePrefix               = "t_"
	DefaultLogLevel                  = "info"
	TxDBContextKey      TxContextKey = "TX_CONTEXT_KEY"
)
