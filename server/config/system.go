package config

type System struct {
	ENV    string `mapstructure:"env" json:"env" yaml:"env"`
	DbType string `mapstructure:"db_type" json:"db_type" yaml:"db_type"`
}
