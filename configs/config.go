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

var cf = new(config)

func init() {
	viper.SetDefault("api.port", "8000")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.host", "localhost")
}

func Load() error {
	viper.SetConfigName("config") //Defino o nome do arquivo
	viper.SetConfigType("toml")   //Defino o tipo do arquivo
	viper.AddConfigPath(".")      //Aviso ao Viper que eu quero configurar este arquivo.
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cf.API = APIConfigs{
		Port: viper.GetString("api.port"),
	}
	cf.DB = DBConfigs{
		Port:     viper.GetString("database.port"),
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.user"),
		DataBase: viper.GetString("database.name"),
		Pass:     viper.GetString("database.pass"),
	}
	return nil
}

func GetAPIPort() string {
	return cf.API.Port
}

func GetDB() DBConfigs {
	return cf.DB
}
