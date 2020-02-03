package config

import (
	env "github.com/caarlos0/env/v6"
)

type Config struct {
	AppName string `env:"APP_NAME"`

	// MongoDB config
	MongoDBEndpoint         string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName             string `env:"MONGODB_NAME,required"`
	MongoDBStaffTableName   string `env:"MONGODB_STAFF_TABLE_NAME,required"`
	MongoDBCompanyTableName string `env:"MONGODB_COMPANY_TABLE_NAME,required"`

	// Jaeger config
	JaegerAgentHost string `env:"JAEGER_AGENT_HOST,required"`
	JaegerAgentPort string `env:"JAEGER_AGENT_PORT,required"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
