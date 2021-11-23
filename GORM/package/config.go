package config

var HOST, USER, PASSWORD, DB, PORT string = "localhost", "root", "123456789", "golang", "5432"
var DSN string = "host=" + HOST + " user=" + USER + " password=" + PASSWORD + " dbname=" + DB + " port=" + PORT + " sslmode=disable TimeZone=Asia/Istanbul"
