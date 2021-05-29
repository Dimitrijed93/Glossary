package util

const (
	PORT                     = ":3030"
	CONNECTION_STRING_DOCKER = "host=glossaryDB user=postgres password=dimit dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Belgrade"
	CONNECTION_STRING_LOCAL  = "host=localhost user=postgres password=dimit dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Belgrade"
	CONNECTION_STRING_K8     = "host=192.168.56.1 user=postgres password=dimit dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Belgrade"
)
