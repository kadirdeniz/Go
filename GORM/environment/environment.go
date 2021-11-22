package environment

var PORT, POSTGRES_DB, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD string = "8000", "golang", "localhost", "5432", "root", "123456789"
var ConnectionString string = "host=" + POSTGRES_HOST + " user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " port=" + POSTGRES_PORT + " sslmode=disable TimeZone=Asia/Istanbul"
