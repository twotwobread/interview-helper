package bootstrap

type Config struct {
	Env              Env
	DatabaseInstance DatabaseInstance
}

func NewConfig() Config {
	config := Config{}
	config.Env = NewEnv()
	config.DatabaseInstance = NewDatabaseInstance(&config.Env)
	return config
}

func (config *Config) CloseDBConnection() {
	CloseMongoDBConnection(config.DatabaseInstance.Client)
}
