package app

import (
	"github.com/TomasCruz/users/internal/configuration"
	"github.com/joho/godotenv"
)

func setupFromEnvVars() (configuration.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return configuration.Config{}, err
	}

	port, err := readAndCheckIntEnvVar("HEX_TEMPLATE_USERS_WEB_PORT")
	if err != nil {
		return configuration.Config{}, err
	}

	dbURL, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_DB_URL")
	if err != nil {
		return configuration.Config{}, err
	}

	kafkaBroker, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_KAFKA_BROKER")
	if err != nil {
		return configuration.Config{}, err
	}

	createUserTopic, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_KAFKA_TOPIC_CREATE_USER")
	if err != nil {
		return configuration.Config{}, err
	}

	updateUserTopic, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_KAFKA_TOPIC_UPDATE_USER")
	if err != nil {
		return configuration.Config{}, err
	}

	deleteUserTopic, err := readAndCheckEnvVar("HEX_TEMPLATE_USERS_KAFKA_TOPIC_DELETE_USER")
	if err != nil {
		return configuration.Config{}, err
	}

	return configuration.Config{
		Port:            port,
		DbURL:           dbURL,
		KafkaBroker:     kafkaBroker,
		CreateUserTopic: createUserTopic,
		UpdateUserTopic: updateUserTopic,
		DeleteUserTopic: deleteUserTopic,
	}, nil
}
