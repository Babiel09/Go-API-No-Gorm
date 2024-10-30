package configs

import "github.com/spf13/viper"

type APIConfigs struct {
	Port string
}

type DBConfigs struct {
	Host     string
	User     string
	DataBase string
	Pass     string
	Port     string
}

type config struct {
	API APIConfigs
	DB  DBConfigs
}

var cnf = new(config)

func init() {
	viper.SetDefault("api.port", "8000")
	viper.SetDefault("database.port", "5433")
	viper.SetDefault("database.host", "localhost")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cnf.API = APIConfigs{
		Port: viper.GetString("api.port"),
	}
	cnf.DB = DBConfigs{
		Port:     viper.GetString("database.port"),
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.user"),
		DataBase: viper.GetString("database.name"),
		Pass:     viper.GetString("database.pass"),
	}
	return nil
}

func GetAPIPort() string {
	return cnf.API.Port
}

func GetDB() DBConfigs {
	return cnf.DB
}
