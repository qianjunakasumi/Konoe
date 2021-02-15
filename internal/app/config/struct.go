package config

type (
	// Client 客户端。
	Client struct {
		Number   uint64 `yaml:"number"`   // 号码
		Password string `yaml:"password"` // 密码
	}

	// Device 设备。
	Device struct {
		Display     string `yaml:"display"`     // 显示
		Model       string `yaml:"model"`       // 型号
		Brand       string `yaml:"brand"`       // 制造商
		Board       string `yaml:"board"`       // 主板
		Device      string `yaml:"device"`      // 设备
		Product     string `yaml:"product"`     // 产品
		Fingerprint string `yaml:"fingerprint"` // 编译序列
		Kernel      string `yaml:"kernel"`      // 内核版本
		SIM         string `yaml:"sim"`         // 运营商信息
		MAC         string `yaml:"mac"`         // MAC 地址
		IMEI        string `yaml:"imei"`        // IMEI
	}
)
