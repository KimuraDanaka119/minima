package minima

type Config struct {
	Middleware []Handler
	Logger     bool
	Router     *Router
	ErrorPath  string
	ErrorData  interface{}
}

func NewConfig() *Config {
	return &Config{Logger: false}
}
