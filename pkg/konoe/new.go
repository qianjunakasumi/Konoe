package konoe

import (
	"github.com/qianjunakasumi/Konoe/internal/app/config"
	"github.com/qianjunakasumi/Konoe/internal/pkg/logger"

	"github.com/Mrs4s/MiraiGo/client"
	"go.uber.org/zap"
)

// Konoe 近江。
type Konoe struct {
	cli *client.QQClient // 客户端
}

// New 返回 Konoe。
func New() (k *Konoe, err error) {

	logger.Init()

	c, err := config.LoadConfig()
	if err != nil {
		return
	}

	k = &Konoe{cli: client.NewClient(int64(c.Number), c.Password)}
	return
}

// login 登录。
func (k *Konoe) login(relogin bool) {

	if k.cli.Online {
		return
	}
	for sta, err := k.cli.Login(); ; sta, err = k.cli.Login() {

		if sta == nil {
			continue
		}
		if (err == nil && sta.Success) || err == client.ErrAlreadyOnline { // 登录成功
			break
		}
		logger.Error("登录失败，需要进一步操作", zap.Error(err))

		switch sta.Error {
		case client.SMSNeededError:
			// TODO 要求输入验证码
		}

		logger.Error("无法登录：" + sta.ErrorMessage)
	}
	if relogin {
		logger.Info("重连成功")
	}

	err := k.cli.ReloadGroupList()
	if err != nil {
		logger.Error("加载群列表失败", zap.Error(err))
		return
	}

	err = k.cli.ReloadFriendList()
	if err != nil {
		logger.Error("加载好友列表失败", zap.Error(err))
		return
	}

	logger.Info("登录成功")
	return
}

// Start 启动。
func (k *Konoe) Start() error {
	k.login(false)
	return k.listen()
}
