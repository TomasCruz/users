package configuration

import (
	"strings"
)

type Config struct {
	MinLogLevel     string
	Port            string
	DBURL           string
	DBMigrationPath string
	KafkaURL        string
	CreateUserTopic string
	UpdateUserTopic string
	DeleteUserTopic string
}

func (c Config) String() string {
	var sb strings.Builder

	sb.WriteString("MinLogLevel:\t\t")
	sb.WriteString(c.MinLogLevel)
	sb.WriteRune('\n')

	sb.WriteString("Port:\t\t\t")
	sb.WriteString(c.Port)
	sb.WriteRune('\n')

	sb.WriteString("DBURL:\t\t\t")
	sb.WriteString(c.DBURL)
	sb.WriteRune('\n')

	sb.WriteString("DBMigrationPath:\t")
	sb.WriteString(c.DBMigrationPath)
	sb.WriteRune('\n')

	sb.WriteString("KafkaURL:\t\t")
	sb.WriteString(c.KafkaURL)
	sb.WriteRune('\n')

	sb.WriteString("CreateUserTopic:\t")
	sb.WriteString(c.CreateUserTopic)
	sb.WriteRune('\n')

	sb.WriteString("UpdateUserTopic:\t")
	sb.WriteString(c.UpdateUserTopic)
	sb.WriteRune('\n')

	sb.WriteString("DeleteUserTopic:\t")
	sb.WriteString(c.DeleteUserTopic)
	sb.WriteRune('\n')

	return sb.String()
}
