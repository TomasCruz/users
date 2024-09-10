package configuration

import (
	"fmt"
)

// Config holds environment variable values, it's populated on startup
type Config struct {
	Port            string
	DBURL           string
	DBMigrationPath string
	KafkaBroker     string
	CreateUserTopic string
	UpdateUserTopic string
	DeleteUserTopic string
}

func (c Config) String() string {
	portStr := fmt.Sprintf("Port:\t\t%s", c.Port)
	dbUrlStr := fmt.Sprintf("DBURL:\t\t%s", c.DBURL)
	dbMigrationPath := fmt.Sprintf("DBMigrationPath:\t%s", c.DBMigrationPath)
	kafkaBroker := fmt.Sprintf("KafkaBroker:\t\t%s", c.KafkaBroker)
	createUserTopic := fmt.Sprintf("CreateUserTopic:\t\t%s", c.CreateUserTopic)
	updateUserTopic := fmt.Sprintf("UpdateUserTopic:\t\t%s", c.UpdateUserTopic)
	deleteUserTopic := fmt.Sprintf("DeleteUserTopic:\t\t%s", c.DeleteUserTopic)
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n", portStr, dbUrlStr, dbMigrationPath, kafkaBroker, createUserTopic, updateUserTopic, deleteUserTopic)
}
