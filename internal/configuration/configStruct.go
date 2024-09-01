package configuration

// Config holds environment variable values, it's populated on startup
type Config struct {
	Port            string
	DbURL           string
	DbMigrationPath string
	KafkaBroker     string
	CreateUserTopic string
	UpdateUserTopic string
	DeleteUserTopic string
}
