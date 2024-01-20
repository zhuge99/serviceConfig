package modConfig

import (
	"os"
	"serviceConfig/modUserCustomize"
)

type CEnvironmentConfig struct {
}

var g_SingleEnvConfig *CEnvironmentConfig = &CEnvironmentConfig{}

func getSingleEnvConfig() *CEnvironmentConfig {
	return g_SingleEnvConfig
}

func (instSelf *CEnvironmentConfig) Initialize() error {
	appid := os.Getenv(modUserCustomize.DIGCD_Key_AppID)
	return nil
}
