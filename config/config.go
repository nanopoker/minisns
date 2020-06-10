package config

var (
	REDIS_HOST           = "localhost:6379"
	REDIS_PASSWORD       = ""
	REDIS_DATABASE       = 0
	REDIS_KEY_DURATION   = 7200
	REDIS_MAXIDLECONNS   = 3
	REDIS_MAXACTIVECONNS = 7000
	REDIS_DIALNETWORK    = "tcp"

	DB_HOST     = "localhost:3306"
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_DATABASE = "minisns"

	HTTP_HOST = "localhost"
	HTTP_PORT = "10001"

	TCP_HOST = "localhost"
	TCP_PORT = "20001"

	COOKIE_DURATION = 7200
)
