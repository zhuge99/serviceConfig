package modConfig

type CEnvironmentConfig struct {
}

var g_SingleEnvConfig *CEnvironmentConfig = &CEnvironmentConfig{}

func getSingleEnvConfig() *CEnvironmentConfig {
	return g_SingleEnvConfig
}

func (instSelf *CEnvironmentConfig) Initialize() error {

	return nil
}
