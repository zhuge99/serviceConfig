package modData

type CDataConfig struct {
}

var g_SingleDataConfig *CDataConfig = &CDataConfig{}

func getSingleDataConfig() *CDataConfig {
	return g_SingleDataConfig
}
