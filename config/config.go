package config

const MASTER_DSN = "host=localhost user=postgres password=password dbname=database port=55000 " +
	"sslmode=disable TimeZone=Asia/Bangkok"

const SLAVE_DSN = "host=localhost user=postgres password=password dbname=database port=55079 " +
	"sslmode=disable TimeZone=Asia/Bangkok"
