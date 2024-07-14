package orchestrator

import "github.com/pkg/errors"

func (d *Docker) Create(info DockerInfo) error {
	return nil
}

func validate(info DockerInfo) error {
	// 验证信息是否有效
	switch info.AccessMode {
	case DockerAccessModeLocalEnv:
	case DockerAccessModeAgent:
	case DockerAccessModeTCP:
	case DockerAccessModeTCPWithTLS:
	default:
		return errors.Errorf("can't")
	}
	return nil
}
