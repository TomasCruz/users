package configuration

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func ConfigFromEnvVars(envFile string) (Config, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	minLogLevel, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_MIN_LOG_LEVEL")
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	port, err := readAndCheckIntEnvVar("HEX_TEMPLATE_USERS_WEB_PORT")
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	gRPCPort, err := readAndCheckIntEnvVar("HEX_TEMPLATE_USERS_GRPC_PORT")
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	dbURL, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_DB_URL")
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	dbMigrationPath, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_DB_MIGRATION_PATH")
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	kafkaURL, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_KAFKA_BROKER")
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	createUserTopic, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_KAFKA_TOPIC_CREATE_USER")
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	updateUserTopic, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_KAFKA_TOPIC_UPDATE_USER")
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	deleteUserTopic, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_KAFKA_TOPIC_DELETE_USER")
	if err != nil {
		return Config{}, errors.WithStack(err)
	}

	return Config{
		MinLogLevel:     minLogLevel,
		Port:            port,
		GRPCPort:        gRPCPort,
		DBURL:           dbURL,
		DBMigrationPath: dbMigrationPath,
		KafkaURL:        kafkaURL,
		CreateUserTopic: createUserTopic,
		UpdateUserTopic: updateUserTopic,
		DeleteUserTopic: deleteUserTopic,
	}, nil
}
