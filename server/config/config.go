package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructuer:"zap" json:"zap" yaml:"zap"`
}
