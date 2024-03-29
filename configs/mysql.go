package configs

const DriverName = "mysql"

type DbConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	Database  string
	IsRunning bool
}

var DbMasterList = []DbConfig{
	{
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Pwd:       "tester",
		Database:  "eilieili",
		IsRunning: false,
	},
}

var DbMaster DbConfig = DbMasterList[0]
