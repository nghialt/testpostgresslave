package config

const MASTER_DSN = "host=localhost user=postgres password=password dbname=database port=55005 " +
	"sslmode=disable TimeZone=Asia/Bangkok"

const SLAVE_DSN = "host=localhost user=postgres password=password dbname=database port=55004 " +
	"sslmode=disable TimeZone=Asia/Bangkok"
