package config

type Config struct {
	Host     string
	Port     string
	User     string
	Pass     string
	DbName   string
	HttpPort string
	SrvKey   string
}

func SetConfig() Config {
	var config Config

	//set configuration here
	config.Host = "localhost"
	config.Port = "3306"
	config.User = "indro"
	config.Pass = "12345"
	config.DbName = "decision_data"
	config.HttpPort = "1234"
	return config
}
