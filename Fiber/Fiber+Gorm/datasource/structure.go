package datasource

import "fmt"

type DB struct {
	Name     string
	Host     string
	Port     int
	Db       string
	User     string
	Password string
}

func (db *DB) DSNConverter() string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Istanbul", db.Host, db.User, db.Password, db.Db, db.Port)
}

func (db *DB) URIConverter() string {
	return fmt.Sprintf("%v://%v:%v@%v:%v/%v", db.Name, db.User, db.Password, db.Host, db.Port, db.Db)
}
