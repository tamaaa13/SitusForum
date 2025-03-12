package configs

import "github.com/spf13/viper"

var config *Config

type option struct {
	configFolders []string
	configFile    string
	configType    string
}

func Init(opts ...Option) error {
	opt := &option{
		configFolders: getDefaultConfigFolder(),
		configFile:    getDefaultConfigFile(),
		configType:    getDefaultConfigType(),
	}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	for _, configFolders := range opt.configFolders {
		viper.AddConfigPath(configFolders)
	}
	viper.SetConfigName(opt.configFile)
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	// Bind variabel environment
	_ = viper.BindEnv("service.port", "SERVER_PORT")
	_ = viper.BindEnv("service.secretJWT", "JWT_SECRET")
	_ = viper.BindEnv("database.password", "MYSQL_ROOT_PASSWORD")
	_ = viper.BindEnv("database.name", "MYSQL_DATABASE")
	_ = viper.BindEnv("database.port", "DB_PORT")

	config = new(Config)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return err
	}

	// Makesure Service Port memiliki tanda ":" di depan
	port := viper.GetString("service.port")
	if port != "" && port[0] != ':' {
		port = ":" + port
	}
	config.Service.Port = port

	config.Database.DataSourceName = "root:" + viper.GetString("database.password") +
		"@tcp(localhost:" + viper.GetString("database.port") + ")/" +
		viper.GetString("database.name") + "?parseTime=true"

	return nil
}

type Option func(*option)

func getDefaultConfigFolder() []string {
	return []string{"./configs"}
}

func getDefaultConfigFile() string {
	return "config"
}

func getDefaultConfigType() string {
	return "yaml"
}

func WithConfigFolder(configFolders []string) Option {
	return func(opt *option) {
		opt.configFolders = configFolders
	}
}

func WithConfigFile(configFile string) Option {
	return func(opt *option) {
		opt.configFile = configFile
	}
}

func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
