package go_demo


// 设计函数时,遇到需要传递大量参数的情况，为了保持代码整洁，可以考虑以下两种策略：
// 1.结构体作为参数
// 2.变长参数

type ServiceOptions struct {
	Host string
	Port string
	SSL bool
}

func ConnectService(opts ServiceOptions) {
	// Connect logic...
}

type ServiceConfig struct {
	SSL bool
}

type ServiceOption func(config *ServiceConfig)

func WithSSL(ssl bool) ServiceOption {
	return func(config *ServiceConfig) {
		config.SSL = ssl
	}
}


func ConnectDB(opts ...ServiceOption) {
	cfg := &ServiceConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
}

//
func main () {
	ConnectDB(WithSSL(true), WithSSL(false))
}

