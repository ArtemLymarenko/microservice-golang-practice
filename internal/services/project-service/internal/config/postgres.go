package config

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Dialect  string `yaml:"dialect"`
	Port     int    `yaml:"port"`
	PoolMin  int    `yaml:"poolMin"`
	PoolMax  int    `yaml:"poolMax"`
}

func (p Postgres) GetUser() string     { return p.User }
func (p Postgres) GetPassword() string { return p.Password }
func (p Postgres) GetHost() string     { return p.Host }
func (p Postgres) GetName() string     { return p.Name }
func (p Postgres) GetDialect() string  { return p.Dialect }
func (p Postgres) GetPort() int        { return p.Port }
func (p Postgres) GetPoolMin() int     { return p.PoolMin }
func (p Postgres) GetPoolMax() int     { return p.PoolMax }
