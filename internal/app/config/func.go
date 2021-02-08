package config

import (
	"github.com/qianjunakasumi/Konoe/internal/pkg/logger"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// GetClientConfig 获取配置
func GetClientConfig() (conf *Client, err error) {

	f, err := ioutil.ReadFile(".konoe/config.yml")
	if err != nil {
		logger.Error("无法读取配置文件", zap.Error(err))
		return nil, err
	}

	err = yaml.Unmarshal(f, conf)
	if err != nil {
		logger.Error("无法解析配置文件或损坏的配置文件", zap.Error(err))
		return nil, err
	}

	return
}
