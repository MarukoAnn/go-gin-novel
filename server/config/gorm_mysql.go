package config

import "fmt"

type Mysql struct {
	GeneralDB `yaml:":,inline" mapstructure:",squash"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", m.Username, m.Password, m.Host, m.Port, m.Dbname)
}
