package config

import "github.com/spf13/viper"

type Cfg struct {
	DB DBcfg
}

func LoadConfig() (*Cfg, error) {
	dbConf := loadDBConfig()
	if err := dbConf.validate(); err != nil {
		return nil, err
	}

	return &Cfg{
		DB: dbConf,
	}, nil
}

// loads DB Config
func loadDBConfig() DBcfg {
	viper.SetEnvPrefix("MYSQL")
	viper.BindEnv("USERNAME", "PASS", "DBNAME", "HOST")

	return DBcfg{
		DB_USER:     viper.GetString("USERNAME"),
		DB_PASSWORD: viper.GetString("PASS"),
		DB_NAME:     viper.GetString("DBNAME"),
		DB_HOST:     viper.GetString("HOST"),
	}
}
