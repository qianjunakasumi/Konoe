package config

// Client 客户端
type Client struct {
	Number   uint64 `yaml:"number"`   // 号码
	Password string `yaml:"password"` // 密码
}
