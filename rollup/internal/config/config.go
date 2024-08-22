package config

import (
	"fmt"
	"scroll-tech/common/database"

	"github.com/spf13/viper"
)

// Config load configuration items.
type Config struct {
	L1Config *L1Config        `mapstructure:"l1_config"`
	L2Config *L2Config        `mapstructure:"l2_config"`
	DBConfig *database.Config `mapstructure:"db_config"`
}

// NewConfig returns a new instance of Config.
func NewConfig(file string) (*Config, error) {
	fmt.Printf("Loading config from file: %s\n", file)

	viper.SetConfigFile(file)
	viper.SetConfigType("json")
	viper.SetEnvPrefix("SCROLL_ROLLUP")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return nil, err
	}
	fmt.Println("Successfully read config file")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return nil, err
	}
	fmt.Println("Successfully unmarshaled config")

	fmt.Printf("Config: %+v\n", cfg)
	fmt.Printf("L1 Config: %+v\n", cfg.L1Config)
	fmt.Printf("L2 Config: %+v\n", cfg.L2Config)
	fmt.Printf("DB Config: %+v\n", cfg.DBConfig)
	fmt.Printf("L1 Endpoint: %s\n", cfg.L1Config.Endpoint)
	fmt.Printf("L2 Endpoint: %s\n", cfg.L2Config.Endpoint)
	fmt.Printf("DB DSN: %s\n", cfg.DBConfig.DSN)

	return &cfg, nil
}
