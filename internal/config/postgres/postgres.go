package postgres

import "fmt"

type Postgres struct {
	User     string `yaml:"user" env-required:"true"`
	Name     string `yaml:"name" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     int    `yaml:"port" env-required:"true"`
}

func (p Postgres) Dsn() string {
	return fmt.Sprintf(
		`host=%s port=%v dbname=%s user=%s password=%s sslmode=disable`,
		p.Host, p.Port, p.Name, p.User, p.Password)
}
