package configs

type RdsConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	IsRunning bool
}

var RdsCacheList = []RdsConfig{
	{
		Host:      "127.0.0.1",
		Port:      6379,
		User:      "root",
		Pwd:       "tester",
		IsRunning: true,
	},
}

var RdsCache RdsConfig = RdsCacheList[0]
