package postgresql

import "fmt"

type Configuration struct {
	Port     int
	Host     string
	User     string
	Password string
	Dbname   string
}

var Conf *Configuration = &Configuration{
	Port:     5432,
	Host:     "localhost",
	User:     "root",
	Password: "123456789",
	Dbname:   "golang",
}

var ConnectionString string = fmt.Sprintf("postgres://%v:%v@%v:%d/%v", Conf.User, Conf.Password, Conf.Host, Conf.Port, Conf.Dbname)
