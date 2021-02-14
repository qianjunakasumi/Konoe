package config

import (
	"github.com/qianjunakasumi/Konoe/internal/pkg/logger"

	"github.com/Mrs4s/MiraiGo/client"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// LoadConfig 加载配置。
func LoadConfig() (c *Client, err error) {
	c, err = GetClientConfig()
	if err != nil {
		return
	}

	err = SetDeviceConfig()
	return
}

// GetClientConfig 获取客户端配置。
func GetClientConfig() (c *Client, err error) {

	f, err := ioutil.ReadFile(".konoe/config.yml")
	if err != nil {
		logger.Error("无法读取客户端配置文件", zap.Error(err))
		return
	}

	err = yaml.Unmarshal(f, &c)
	if err != nil {
		logger.Error("无法解析配置文件或损坏的客户端配置文件", zap.Error(err))
		return
	}

	return
}

// SetDeviceConfig 设置设备配置。
func SetDeviceConfig() (err error) {

	var c *Device

	f, err := ioutil.ReadFile(".konoe/device.yml")
	if err != nil {
		logger.Error("无法读取设备配置文件", zap.Error(err))
		return
	}

	err = yaml.Unmarshal(f, &c)
	if err != nil {
		logger.Error("无法解析配置文件或损坏的设备配置文件", zap.Error(err))
		return
	}

	setDeviceConfig(c)
	return
}

// setDeviceConfig 设置设备配置。
func setDeviceConfig(c *Device) {
	client.SystemDeviceInfo.Protocol = client.AndroidPhone

	client.SystemDeviceInfo.Display = []byte(c.Display)
	client.SystemDeviceInfo.Model = []byte(c.Model)
	client.SystemDeviceInfo.Brand = []byte(c.Brand)
	client.SystemDeviceInfo.Board = []byte(c.Board)
	client.SystemDeviceInfo.Device = []byte(c.Device)
	client.SystemDeviceInfo.Product = []byte(c.Product)

	client.SystemDeviceInfo.FingerPrint = []byte(c.Fingerprint)
	client.SystemDeviceInfo.ProcVersion = []byte(c.Kernel)
	client.SystemDeviceInfo.SimInfo = []byte(c.SIM)
	client.SystemDeviceInfo.MacAddress = []byte(c.MAC)
	client.SystemDeviceInfo.IpAddress = []byte(c.IP)
	client.SystemDeviceInfo.IMEI = c.IMEI
}
